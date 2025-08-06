package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SwaggerComment struct {
	Summary     string
	Description string
	Tags        string
	Accept      string
	Produce     string
	Security    string
	Params      []Param
	Success     []Response
	Failure     []Response
	Router      string
}

type Param struct {
	Name     string
	In       string
	Required bool
	Type     string
	Desc     string
}

type Response struct {
	Code        string
	Description string
	Type        string
}

type FunctionInfo struct {
	Name        string
	Receiver    string
	HTTPMethod  string
	Path        string
	Params      []Param
	Returns     []string
	HasAuth     bool
	RequestBody string
	LineNumber  int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run auto_swagger.go <handler-directory>")
		fmt.Println("Example: go run auto_swagger.go internal/handlers")
		os.Exit(1)
	}

	handlerDir := os.Args[1]
	if err := processHandlerDirectory(handlerDir); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Swagger comments added successfully!")
	fmt.Println("🔄 Now generating Swagger documentation...")

	if err := generateSwaggerDocs(); err != nil {
		fmt.Printf("Error generating Swagger docs: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("🎉 Swagger documentation generated successfully!")
}

func processHandlerDirectory(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return fmt.Errorf("failed to find Go files: %v", err)
	}

	for _, file := range files {
		if err := processHandlerFile(file); err != nil {
			fmt.Printf("Warning: failed to process %s: %v\n", file, err)
		}
	}

	return nil
}

func processHandlerFile(filename string) error {
	fmt.Printf("Processing: %s\n", filename)

	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, content, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file: %v", err)
	}

	handlers := findHandlersWithoutSwaggerComments(node, fset, content)

	if len(handlers) > 0 {
		newContent := addSwaggerCommentsToFile(content, handlers, fset)
		if err := os.WriteFile(filename, newContent, 0644); err != nil {
			return fmt.Errorf("failed to write file: %v", err)
		}
		fmt.Printf("  ✅ Added Swagger comments for %d functions\n", len(handlers))
	} else {
		fmt.Printf("  ℹ️  All functions already have Swagger comments\n")
	}

	return nil
}

func findHandlersWithoutSwaggerComments(node *ast.File, fset *token.FileSet, content []byte) []FunctionInfo {
	var handlers []FunctionInfo

	ast.Inspect(node, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			if isHandlerFunction(funcDecl) {
				if !hasSwaggerComments(funcDecl, content) {
					handler := extractFunctionInfo(funcDecl, fset)
					handlers = append(handlers, handler)
				}
			}
		}
		return true
	})

	return handlers
}

func isHandlerFunction(funcDecl *ast.FuncDecl) bool {
	if funcDecl.Type.Params != nil && len(funcDecl.Type.Params.List) > 0 {
		for _, param := range funcDecl.Type.Params.List {
			if sel, ok := param.Type.(*ast.SelectorExpr); ok {
				if ident, ok := sel.X.(*ast.Ident); ok {
					if ident.Name == "gin" && sel.Sel.Name == "Context" {
						return true
					}
				}
			}
		}
	}
	return false
}

func hasSwaggerComments(funcDecl *ast.FuncDecl, content []byte) bool {
	if funcDecl.Doc != nil {
		for _, comment := range funcDecl.Doc.List {
			if strings.Contains(comment.Text, "godoc") {
				return true
			}
		}
	}
	return false
}

func extractFunctionInfo(funcDecl *ast.FuncDecl, fset *token.FileSet) FunctionInfo {
	handler := FunctionInfo{
		Name: funcDecl.Name.Name,
	}

	if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
		if star, ok := funcDecl.Recv.List[0].Type.(*ast.StarExpr); ok {
			if sel, ok := star.X.(*ast.SelectorExpr); ok {
				handler.Receiver = sel.Sel.Name
			}
		}
	}

	handler.LineNumber = fset.Position(funcDecl.Pos()).Line

	handler.HTTPMethod, handler.Path = inferHTTPMethodAndPath(handler.Name, handler.Receiver)

	handler.HasAuth = hasAuthentication(handler.Name)

	handler.RequestBody = inferRequestBodyType(handler.Name, handler.Receiver)

	return handler
}

func inferHTTPMethodAndPath(funcName, receiver string) (method, path string) {
	switch {
	case strings.HasPrefix(funcName, "Create"):
		method = "POST"
		path = fmt.Sprintf("/api/%s", strings.ToLower(receiver))
	case strings.HasPrefix(funcName, "Get"):
		if strings.Contains(funcName, "ByID") || strings.Contains(funcName, "By") {
			method = "GET"
			path = fmt.Sprintf("/api/%s/{id}", strings.ToLower(receiver))
		} else {
			method = "GET"
			path = fmt.Sprintf("/api/%s", strings.ToLower(receiver))
		}
	case strings.HasPrefix(funcName, "Update"):
		method = "PUT"
		path = fmt.Sprintf("/api/%s/{id}", strings.ToLower(receiver))
	case strings.HasPrefix(funcName, "Delete"):
		method = "DELETE"
		path = fmt.Sprintf("/api/%s/{id}", strings.ToLower(receiver))
	case strings.HasPrefix(funcName, "Search"):
		method = "GET"
		path = fmt.Sprintf("/api/%s/search", strings.ToLower(receiver))
	case strings.HasPrefix(funcName, "Login"):
		method = "POST"
		path = "/api/auth/login"
	case strings.HasPrefix(funcName, "Register"):
		method = "POST"
		path = "/api/auth/register"
	case strings.HasPrefix(funcName, "Logout"):
		method = "POST"
		path = "/api/auth/logout"
	case strings.HasPrefix(funcName, "Refresh"):
		method = "POST"
		path = "/api/auth/refresh"
	default:
		method = "GET"
		path = fmt.Sprintf("/api/%s", strings.ToLower(receiver))
	}

	return method, path
}

func hasAuthentication(funcName string) bool {
	authFuncs := []string{"Create", "Update", "Delete", "GetCurrent", "Logout", "Set", "Add", "Remove"}
	for _, authFunc := range authFuncs {
		if strings.Contains(funcName, authFunc) {
			return true
		}
	}
	return false
}

func inferRequestBodyType(funcName, receiver string) string {
	switch {
	case strings.HasPrefix(funcName, "Create"):
		return fmt.Sprintf("models.%s", cases.Title(language.English).String(receiver))
	case strings.HasPrefix(funcName, "Update"):
		return fmt.Sprintf("models.%s", cases.Title(language.English).String(receiver))
	case strings.HasPrefix(funcName, "Login"):
		return "LoginRequest"
	case strings.HasPrefix(funcName, "Register"):
		return "RegisterRequest"
	case strings.HasPrefix(funcName, "Set"):
		if strings.Contains(funcName, "Attribute") {
			return "game.CharacterAttribute"
		} else if strings.Contains(funcName, "Skill") {
			return "game.CharacterSkill"
		} else if strings.Contains(funcName, "Quality") {
			return "game.CharacterQuality"
		} else if strings.Contains(funcName, "Equipment") {
			return "game.CharacterEquipment"
		} else if strings.Contains(funcName, "Stats") {
			return "game.CharacterDerivedStats"
		}
		return "interface{}"
	case strings.HasPrefix(funcName, "Add"):
		if strings.Contains(funcName, "Equipment") {
			return "game.CharacterEquipment"
		} else if strings.Contains(funcName, "Specialty") {
			return "AddSkillSpecialtyRequest"
		}
		return "interface{}"
	default:
		return ""
	}
}

func addSwaggerCommentsToFile(content []byte, handlers []FunctionInfo, fset *token.FileSet) []byte {
	lines := strings.Split(string(content), "\n")

	for i := len(handlers) - 1; i >= 0; i-- {
		handler := handlers[i]
		comment := generateSwaggerComment(handler)
		commentLines := strings.Split(comment, "\n")

		insertIndex := handler.LineNumber - 1
		if insertIndex < 0 {
			insertIndex = 0
		}

		newLines := make([]string, 0, len(lines)+len(commentLines))
		newLines = append(newLines, lines[:insertIndex]...)
		newLines = append(newLines, commentLines...)
		newLines = append(newLines, lines[insertIndex:]...)

		lines = newLines
	}

	return []byte(strings.Join(lines, "\n"))
}

func generateSwaggerComment(handler FunctionInfo) string {
	comment := SwaggerComment{
		Summary:     generateSummary(handler.Name),
		Description: generateDescription(handler.Name),
		Tags:        strings.ToLower(handler.Receiver),
		Accept:      "json",
		Produce:     "json",
		Router:      fmt.Sprintf("%s [%s]", handler.Path, strings.ToLower(handler.HTTPMethod)),
	}

	if handler.HasAuth {
		comment.Security = "BearerAuth"
	}

	comment.Params = generateParams(handler)

	comment.Success = generateSuccessResponses(handler)
	comment.Failure = generateFailureResponses(handler)

	return formatSwaggerComment(comment)
}

func generateSummary(funcName string) string {
	switch {
	case strings.HasPrefix(funcName, "Create"):
		return fmt.Sprintf("Create a new %s", strings.ToLower(strings.TrimPrefix(funcName, "Create")))
	case strings.HasPrefix(funcName, "Get"):
		return fmt.Sprintf("Get %s", strings.ToLower(strings.TrimPrefix(funcName, "Get")))
	case strings.HasPrefix(funcName, "Update"):
		return fmt.Sprintf("Update %s", strings.ToLower(strings.TrimPrefix(funcName, "Update")))
	case strings.HasPrefix(funcName, "Delete"):
		return fmt.Sprintf("Delete %s", strings.ToLower(strings.TrimPrefix(funcName, "Delete")))
	case strings.HasPrefix(funcName, "Search"):
		return fmt.Sprintf("Search %s", strings.ToLower(strings.TrimPrefix(funcName, "Search")))
	case strings.HasPrefix(funcName, "Set"):
		return fmt.Sprintf("Set %s", strings.ToLower(strings.TrimPrefix(funcName, "Set")))
	case strings.HasPrefix(funcName, "Add"):
		return fmt.Sprintf("Add %s", strings.ToLower(strings.TrimPrefix(funcName, "Add")))
	case strings.HasPrefix(funcName, "Remove"):
		return fmt.Sprintf("Remove %s", strings.ToLower(strings.TrimPrefix(funcName, "Remove")))
	default:
		return funcName
	}
}

func generateDescription(funcName string) string {
	switch {
	case strings.HasPrefix(funcName, "Create"):
		return fmt.Sprintf("Create a new %s", strings.ToLower(strings.TrimPrefix(funcName, "Create")))
	case strings.HasPrefix(funcName, "Get"):
		return fmt.Sprintf("Retrieve %s", strings.ToLower(strings.TrimPrefix(funcName, "Get")))
	case strings.HasPrefix(funcName, "Update"):
		return fmt.Sprintf("Update an existing %s", strings.ToLower(strings.TrimPrefix(funcName, "Update")))
	case strings.HasPrefix(funcName, "Delete"):
		return fmt.Sprintf("Delete %s", strings.ToLower(strings.TrimPrefix(funcName, "Delete")))
	case strings.HasPrefix(funcName, "Search"):
		return fmt.Sprintf("Search %s", strings.ToLower(strings.TrimPrefix(funcName, "Search")))
	case strings.HasPrefix(funcName, "Set"):
		return fmt.Sprintf("Set or update %s", strings.ToLower(strings.TrimPrefix(funcName, "Set")))
	case strings.HasPrefix(funcName, "Add"):
		return fmt.Sprintf("Add %s", strings.ToLower(strings.TrimPrefix(funcName, "Add")))
	case strings.HasPrefix(funcName, "Remove"):
		return fmt.Sprintf("Remove %s", strings.ToLower(strings.TrimPrefix(funcName, "Remove")))
	default:
		return funcName
	}
}

func generateParams(handler FunctionInfo) []Param {
	var params []Param

	if strings.Contains(handler.Path, "{id}") {
		params = append(params, Param{
			Name:     "id",
			In:       "path",
			Required: true,
			Type:     "string",
			Desc:     fmt.Sprintf("%s ID", cases.Title(language.English).String(handler.Receiver)),
		})
	}

	if strings.Contains(handler.Name, "Search") {
		params = append(params, Param{
			Name:     "q",
			In:       "query",
			Required: true,
			Type:     "string",
			Desc:     "Search query",
		})
	}

	if handler.RequestBody != "" && (handler.HTTPMethod == "POST" || handler.HTTPMethod == "PUT") {
		paramName := strings.ToLower(handler.Receiver)
		if strings.Contains(handler.Name, "Login") {
			paramName = "credentials"
		} else if strings.Contains(handler.Name, "Register") {
			paramName = "user"
		} else if strings.Contains(handler.Name, "Set") {
			paramName = strings.ToLower(strings.TrimPrefix(handler.Name, "Set"))
		} else if strings.Contains(handler.Name, "Add") {
			paramName = strings.ToLower(strings.TrimPrefix(handler.Name, "Add"))
		}

		params = append(params, Param{
			Name:     paramName,
			In:       "body",
			Required: true,
			Type:     handler.RequestBody,
			Desc:     fmt.Sprintf("%s data", cases.Title(language.English).String(handler.Receiver)),
		})
	}

	return params
}

func generateSuccessResponses(handler FunctionInfo) []Response {
	var responses []Response

	switch handler.HTTPMethod {
	case "POST":
		responses = append(responses, Response{
			Code:        "201",
			Description: "Created",
			Type:        handler.RequestBody,
		})
	case "GET":
		if strings.Contains(handler.Path, "{id}") {
			responses = append(responses, Response{
				Code:        "200",
				Description: "OK",
				Type:        handler.RequestBody,
			})
		} else {
			responses = append(responses, Response{
				Code:        "200",
				Description: "OK",
				Type:        fmt.Sprintf("array %s", handler.RequestBody),
			})
		}
	case "PUT":
		responses = append(responses, Response{
			Code:        "200",
			Description: "OK",
			Type:        handler.RequestBody,
		})
	case "DELETE":
		responses = append(responses, Response{
			Code:        "200",
			Description: "OK",
			Type:        "DeleteResponse",
		})
	}

	return responses
}

func generateFailureResponses(handler FunctionInfo) []Response {
	var responses []Response

	responses = append(responses, Response{
		Code:        "400",
		Description: "Bad Request",
		Type:        "ErrorResponse",
	})

	if strings.Contains(handler.Path, "{id}") {
		responses = append(responses, Response{
			Code:        "404",
			Description: "Not Found",
			Type:        "ErrorResponse",
		})
	}

	if handler.HasAuth {
		responses = append(responses, Response{
			Code:        "401",
			Description: "Unauthorized",
			Type:        "ErrorResponse",
		})
	}

	responses = append(responses, Response{
		Code:        "500",
		Description: "Internal Server Error",
		Type:        "ErrorResponse",
	})

	return responses
}

func formatSwaggerComment(comment SwaggerComment) string {
	var lines []string

	lines = append(lines, fmt.Sprintf("// %s godoc", comment.Summary))

	lines = append(lines, fmt.Sprintf("// @Summary %s", comment.Summary))

	lines = append(lines, fmt.Sprintf("// @Description %s", comment.Description))

	lines = append(lines, fmt.Sprintf("// @Tags %s", comment.Tags))

	lines = append(lines, "// @Accept json")
	lines = append(lines, "// @Produce json")

	if comment.Security != "" {
		lines = append(lines, fmt.Sprintf("// @Security %s", comment.Security))
	}

	for _, param := range comment.Params {
		required := "false"
		if param.Required {
			required = "true"
		}
		lines = append(lines, fmt.Sprintf("// @Param %s %s %s %s \"%s\"",
			param.Name, param.In, param.Type, required, param.Desc))
	}

	for _, resp := range comment.Success {
		lines = append(lines, fmt.Sprintf("// @Success %s {%s} %s",
			resp.Code, resp.Type, resp.Description))
	}

	for _, resp := range comment.Failure {
		lines = append(lines, fmt.Sprintf("// @Failure %s {%s} %s",
			resp.Code, resp.Type, resp.Description))
	}

	lines = append(lines, fmt.Sprintf("// @Router %s", comment.Router))

	return strings.Join(lines, "\n")
}

func generateSwaggerDocs() error {
	cmd := exec.Command("swag", "init", "-g", "cmd/api/main.go")
	cmd.Dir = "."
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
