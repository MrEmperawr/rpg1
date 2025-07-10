# GORM Database Setup Guide

This project uses GORM (Go Object Relational Mapper) to manage database operations and migrations with Supabase.

## What's Set Up

### 1. **Models** (`internal/models/models.go`)
- `User`: User accounts with email, first_name, last_name, and display_name
- `Species`: Character species/races with stats (Human, Elf, Dwarf)
- `Character`: Player characters with concept, vice, virtue, and species relationship

### 2. **Database Connection** (`internal/database/database.go`)
- GORM connection to PostgreSQL (Supabase)
- Automatic migration and seeding on startup
- Connection pooling and error handling

### 3. **Migrations** (`internal/database/migrations.go`)
- Automatic schema migration using `AutoMigrate()`
- Database seeding with initial species (Human, Elf, Dwarf)
- Safe to run multiple times

### 4. **Repository Pattern** (`internal/repository/`)
- Clean separation of database operations
- Example: `UserRepository` with CRUD operations

## Database Schema

### Tables
- **`user`**: User accounts (UUID primary keys)
- **`species`**: Character species with stat ranges
- **`character`**: Player characters linked to users and species

### Key Features
- **UUID Primary Keys**: All tables use UUIDs with `gen_random_uuid()` default
- **Foreign Key Relationships**: Characters link to Users and Species
- **Timestamps**: Automatic `created_at` and optional `updated_at` fields

## How to Connect to Supabase

### 1. Get Your Supabase Connection String
1. Go to your Supabase project dashboard
2. Navigate to Settings → Database
3. Copy the "Connection string" (URI format)
4. It should look like: `postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres`

### 2. Set Environment Variable
Create a `.env` file in the `api1` directory:
```env
DATABASE_URL=postgresql://postgres:sSwvqM#13MwYyTRDfCv4@db.gigstzftwvfyusbkaopm.supabase.co:5432/postgres
PORT=8080
ENV=development
```

### 3. Test the Connection
```bash
cd api1
go run cmd/test-db/main.go
```

### 4. Run the Application
```bash
cd api1
go run cmd/api/main.go
```

The application will:
- Connect to your Supabase database
- Verify existing tables are accessible
- Seed the database with initial species if needed
- Start the API server

## How Migrations Work

### Automatic Migration
When you start the application, GORM automatically:
1. Verifies existing tables match the model definitions
2. Adds new columns to existing tables if needed
3. Updates column types if they've changed
4. Maintains existing data

### Adding New Models
1. Add your model to `internal/models/models.go`
2. Add it to the migration list in `internal/database/migrations.go`
3. Restart the application

Example:
```go
// Add to models.go
type Quest struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Title       string    `gorm:"not null"`
    Description string
    CreatedAt   time.Time `gorm:"not null;default:now()"`
    UpdatedAt   *time.Time
}

// Add to migrations.go
err := db.AutoMigrate(
    &models.User{},
    &models.Species{},
    &models.Character{},
    &models.Quest{}, // Add this line
)
```

## Using the Repository Pattern

```go
// Create a repository
userRepo := repository.NewUserRepository()

// Create a user
user := &models.User{
    Email:       "player@example.com",
    FirstName:   "John",
    LastName:    "Doe",
    DisplayName: "Player1",
}
err := userRepo.Create(user)

// Get user by email
user, err := userRepo.GetByEmail("player@example.com")

// Get all users
users, err := userRepo.GetAll()
```

## GORM Features You Can Use

### Tags for Model Definition
```go
type User struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Email       string    `gorm:"not null"`
    FirstName   string    `gorm:"column:first_name"`
    CreatedAt   time.Time `gorm:"not null;default:now()"`
    UpdatedAt   *time.Time `gorm:"column:updated_at"`
}
```

### Relationships
```go
// One-to-Many
type User struct {
    Characters []Character `gorm:"foreignKey:UserID"`
}

// Many-to-One
type Character struct {
    User    User    `gorm:"foreignKey:UserID"`
    Species Species `gorm:"foreignKey:SpeciesID"`
}
```

### Querying
```go
// Find with conditions
db.Where("name = ?", "Human").Find(&species)

// Preload relationships
db.Preload("User").Preload("Species").Find(&characters)

// Raw SQL
db.Raw("SELECT * FROM user WHERE email = ?", email).Scan(&user)
```

## Testing the Connection

Use the test script to verify your database connection:

```bash
go run cmd/test-db/main.go
```

This will:
- Test the database connection
- Verify all tables are accessible
- Show record counts for each table

## Best Practices

1. **Always use transactions** for operations that modify multiple tables
2. **Use the repository pattern** to keep database logic organized
3. **Add indexes** for frequently queried fields
4. **Use UUIDs** for primary keys (already configured)
5. **Validate data** before saving to the database

## Troubleshooting

### Connection Issues
- Check your Supabase connection string
- Ensure your IP is whitelisted in Supabase
- Verify your database password is correct

### Migration Issues
- Check GORM logs for specific error messages
- Ensure your model structs match the actual database schema
- Verify foreign key relationships are correct

### Performance Issues
- Use `Preload()` to avoid N+1 queries
- Add database indexes for frequently queried fields
- Use `Select()` to limit returned fields 