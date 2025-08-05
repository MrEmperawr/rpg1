# RPG1 API Server

This is the API server for the RPG1 tabletop roleplaying game system.

## Quick Start

### 1. Setup Environment

Copy the example environment file and configure your database:

```bash
cp config/database.env.example config/database.env
# Edit config/database.env with your database credentials
```

### 2. Run Database Migrations and Seeding

First, run the database migrations and seed the database with initial data:

```bash
# Run migrations and seed database
go run cmd/seed/main.go
```

### 3. Start the API Server

```bash
# Start the API server
go run cmd/api/main.go
```

The server will start on `:8080` by default.

## Available Commands

### API Server
```bash
go run cmd/api/main.go
```
Starts the main API server with all SRD endpoints.

### Database Seeding
```bash
go run cmd/seed/main.go
```
Runs database migrations and seeds the database with all initial data (SRD entries, content, attributes, skills, etc.).

## API Endpoints

### Health Check
- `GET /health` - Check if the API and database are running

### SRD (System Reference Document) APIs
- `GET /api/srd/entries` - Get all SRD entries
- `GET /api/srd/entries/:id` - Get SRD entry by ID
- `GET /api/srd/entries/category/:category` - Get SRD entries by category
- `GET /api/srd/categories` - Get all SRD categories
- `GET /api/srd/search?q=<query>` - Search SRD entries
- `GET /api/srd/content` - Get all SRD content
- `GET /api/srd/content/:title` - Get SRD content by title

## Database

The application uses PostgreSQL with GORM as the ORM. All database operations are graceful and handle prepared statement errors and duplicate key constraints automatically.

## Development

### Project Structure
```
api1/
├── cmd/
│   ├── api/          # Main API server
│   └── seed/         # Database seeding command
├── internal/
│   ├── config/       # Configuration management
│   ├── database/     # Database connection and migrations
│   ├── features/     # Domain models
│   ├── handlers/     # HTTP request handlers
│   ├── repository/   # Data access layer
│   └── routes/       # API route definitions
└── config/           # Environment configuration
```

### Adding New Data

To add new SRD entries or content:

1. Add the data to the appropriate seed files in `internal/database/seeds/`
2. Run the seeding command: `go run cmd/seed/main.go`
3. The new data will be available through the API endpoints

### Graceful Error Handling

All seeders are designed to handle:
- Prepared statement errors (common with Supabase)
- Duplicate key constraint violations
- Missing dependencies

The system will log warnings and continue processing rather than failing completely. 