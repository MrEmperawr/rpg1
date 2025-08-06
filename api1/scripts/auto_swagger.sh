#!/bin/bash

# Auto Swagger Documentation Generator
# This script automatically generates Swagger comments for handler functions
# and then generates the complete Swagger documentation

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}🚀 Auto Swagger Documentation Generator${NC}"
echo -e "${CYAN}=====================================${NC}"

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo -e "${YELLOW}Error: Please run this script from the api1 directory${NC}"
    exit 1
fi

# Check if swag is installed
if ! command -v swag &> /dev/null; then
    echo -e "${YELLOW}Installing swag tool...${NC}"
    export PATH=$PATH:$(go env GOPATH)/bin
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Run the auto-swagger program
echo -e "${YELLOW}Generating Swagger comments for handlers...${NC}"
go run scripts/auto_swagger.go internal/handlers

echo -e "${GREEN}✅ Auto Swagger documentation generation complete!${NC}"
echo -e "${CYAN}🌐 Access Swagger UI at: http://localhost:8080/swagger/index.html${NC}"
echo -e "${CYAN}📄 Access Swagger JSON at: http://localhost:8080/swagger/doc.json${NC}"

echo -e "${GREEN}✨ Done!${NC}"
