# RPG Backend Project Structure

This project is organized into feature-based modules for better maintainability and separation of concerns.

## Directory Structure

```
api1/
├── cmd/
│   ├── api/           # Main API server
│   └── test-db/       # Database testing utility
├── config/            # Configuration files
├── internal/
│   ├── config/        # Configuration management
│   ├── database/      # Database connection and migrations
│   ├── handlers/      # HTTP request handlers
│   ├── models/        # Shared models (User, Campaign, Character, Species)
│   └── features/      # Feature-based modules
│       ├── auth/      # Authentication and user management
│       ├── game/      # Game mechanics and character management
│       └── srd/       # System Reference Document and rules
├── game_documentation/ # RPG rule documentation
└── supabase/          # Supabase integration
```

## Feature Modules

### Auth Feature (`internal/features/auth/`)
- **Purpose**: User authentication and management
- **Models**: User (aliased from shared models)
- **Future**: Authentication handlers, JWT management, user sessions

### Game Feature (`internal/features/game/`)
- **Purpose**: Core game mechanics and character management
- **Models**:
  - CharacterAttribute: Character's attribute values
  - CharacterSkill: Character's skill values
  - CharacterSkillSpecialty: Character's skill specialties
  - CharacterQuality: Character's positive/negative qualities
  - CharacterEquipment: Character's equipment
  - CharacterDerivedStats: Calculated character stats
- **Future**: Character creation logic, experience spending, derived stat calculations

### SRD Feature (`internal/features/srd/`)
- **Purpose**: System Reference Document and game rules management
- **Models**:
  - SRDEntry: Text-based rule entries
  - Attribute: Base attribute definitions (Brawn, Agility, etc.)
  - Skill: Base skill definitions (Athletics, Melee, etc.)
  - SkillSpecialty: Skill specialty definitions
  - Quality: Positive/negative quality definitions
  - Equipment: Equipment definitions
  - CharacterCreationRule: Structured character creation rules
  - RuleVersion: Version control for rule changes
- **Future**: Rule validation, CMS interface, rule versioning

## Shared Models (`internal/models/`)

Core models shared across features:
- **User**: System users
- **Campaign**: Game sessions/campaigns
- **Character**: Player characters
- **Species**: Character races/species

## Key Design Decisions

1. **Feature Separation**: Each feature is self-contained with its own models and logic
2. **Shared Models**: Core entities (User, Character, etc.) are shared to avoid circular imports
3. **Type Aliases**: Feature modules use type aliases to reference shared models
4. **SRD as CMS**: Game rules are stored as structured data for programmatic access
5. **Versioning**: Rules can be versioned and tracked for changes

## Database Schema

The database is organized into logical groups:
- **Core Tables**: user, campaign, character, species
- **SRD Tables**: srd_entry, attribute, skill, skill_specialty, quality, equipment, character_creation_rule, rule_version
- **Game Tables**: character_attribute, character_skill, character_skill_specialty, character_quality, character_equipment, character_derived_stats

## Future Enhancements

1. **Repository Pattern**: Add repository interfaces for each feature
2. **Service Layer**: Add business logic services for each feature
3. **API Routes**: Organize HTTP handlers by feature
4. **Validation**: Add input validation for each feature
5. **Testing**: Add unit tests for each feature module 