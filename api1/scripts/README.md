# Scripts Directory

This directory contains utility scripts for the RPG API project.

## 📁 Available Scripts

### 🔄 `auto_swagger.sh` - Auto Swagger Documentation Generator

**What it does:**
- Automatically analyzes your Go handler functions
- Generates Swagger documentation comments based on function signatures
- Infers HTTP methods, paths, parameters, and response types
- Adds the comments directly to your handler files
- Generates the complete Swagger documentation

**Usage:**
```bash
# From the api1 directory
./scripts/auto_swagger.sh
```

**What it analyzes:**
- Function names (Create, Get, Update, Delete, Search, etc.)
- Receiver types (CharacterHandlers, UserHandlers, etc.)
- Function parameters and return types
- Authentication requirements

**Generated comments example:**
```go
// CreateCharacter godoc
// @Summary Create a new character
// @Description Create a new character for the authenticated user
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param character body models.Character true "Character data"
// @Success 201 {object} models.Character
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/characters [post]
func (h *CharacterHandlers) CreateCharacter(c *gin.Context) {
    // Your code here
}
```

### 🔧 `generate_swagger.sh` - Manual Swagger Documentation Generator

**What it does:**
- Regenerates Swagger documentation from existing comments
- Useful when you've manually updated Swagger comments

**Usage:**
```bash
# From the api1 directory
./scripts/generate_swagger.sh
```

### 📊 `codebase_stats.sh` - Codebase Statistics

**What it does:**
- Generates comprehensive statistics about your codebase
- Shows lines of code, author contributions, and activity patterns

**Usage:**
```bash
# From the api1 directory
./scripts/codebase_stats.sh
```

## 🚀 Quick Start

1. **Generate Swagger documentation for all handlers:**
   ```bash
   ./scripts/auto_swagger.sh
   ```

2. **View your API documentation:**
   - Start your server: `go run cmd/api/main.go`
   - Open: http://localhost:8080/swagger/index.html

3. **Regenerate docs after manual changes:**
   ```bash
   ./scripts/generate_swagger.sh
   ```

## 🎯 How the Auto-Swagger Works

### Function Name Analysis
The program analyzes function names to determine:

| Function Prefix | HTTP Method | Path Pattern | Example |
|----------------|-------------|--------------|---------|
| `Create` | POST | `/api/{resource}` | `/api/characters` |
| `Get` | GET | `/api/{resource}` | `/api/characters` |
| `Get` + `ByID` | GET | `/api/{resource}/{id}` | `/api/characters/{id}` |
| `Update` | PUT | `/api/{resource}/{id}` | `/api/characters/{id}` |
| `Delete` | DELETE | `/api/{resource}/{id}` | `/api/characters/{id}` |
| `Search` | GET | `/api/{resource}/search` | `/api/characters/search` |
| `Login` | POST | `/api/auth/login` | `/api/auth/login` |
| `Register` | POST | `/api/auth/register` | `/api/auth/register` |

### Authentication Detection
Functions with these patterns automatically get `@Security BearerAuth`:
- `Create*`, `Update*`, `Delete*`
- `GetCurrent*`, `Logout*`
- `Set*`, `Add*`, `Remove*`

### Parameter Inference
- **Path parameters**: Automatically added for `{id}` patterns
- **Query parameters**: Added for search functions
- **Request body**: Added for POST/PUT operations
- **Response types**: Inferred from function context

## 🔧 Customization

### Adding New Function Patterns
Edit `scripts/auto_swagger.go` and modify the `inferHTTPMethodAndPath` function:

```go
case strings.HasPrefix(funcName, "YourNewPattern"):
    method = "POST"
    path = fmt.Sprintf("/api/%s/your-path", strings.ToLower(receiver))
```

### Custom Response Types
Modify the `inferRequestBodyType` function to add new type mappings:

```go
case strings.HasPrefix(funcName, "YourFunction"):
    return "YourCustomType"
```

## 📋 Requirements

- Go 1.21+
- `swag` tool (automatically installed by the scripts)
- Handler functions must use `gin.Context` as the first parameter

## 🎉 Benefits

1. **Time Saving**: No more manual Swagger comment writing
2. **Consistency**: All endpoints follow the same documentation pattern
3. **Maintenance**: Easy to update when you add new endpoints
4. **Professional**: Generates production-ready API documentation
5. **Interactive**: Test your API directly from the Swagger UI

## 🚨 Notes

- The program only adds comments to functions that don't already have `godoc` comments
- Existing comments are preserved
- Generated comments follow Swagger/OpenAPI 2.0 specification
- All generated files are in the `docs/` directory
