#!/bin/bash

# Generate Swagger Documentation Script
# This script regenerates the Swagger documentation for the RPG API

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}🚀 Generating Swagger Documentation${NC}"
echo -e "${CYAN}===================================${NC}"

# Add Go bin to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Check if swag is installed
if ! command -v swag &> /dev/null; then
    echo -e "${YELLOW}Installing swag tool...${NC}"
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate Swagger documentation
echo -e "${YELLOW}Generating Swagger docs...${NC}"
swag init -g cmd/api/main.go

echo -e "${GREEN}✅ Swagger documentation generated successfully!${NC}"
echo -e "${CYAN}📁 Files created:${NC}"
echo -e "  ├─ docs/docs.go"
echo -e "  ├─ docs/swagger.json"
echo -e "  └─ docs/swagger.yaml"

echo -e "${CYAN}🌐 Access Swagger UI at: http://localhost:8080/swagger/index.html${NC}"
echo -e "${CYAN}📄 Access Swagger JSON at: http://localhost:8080/swagger/doc.json${NC}"

echo -e "${GREEN}✨ Done!${NC}"
