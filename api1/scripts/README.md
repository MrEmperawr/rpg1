# Scripts Directory

This directory contains utility scripts for the RPG API project.

## Available Scripts

### `codebase_stats.sh`

A comprehensive codebase statistics generator that provides fun insights about the project and its contributors.

#### Features

- **Lines of Code Analysis**: Counts lines in Go, SQL, Shell, and Markdown files
- **Author Statistics**: Shows commit counts and lines added/deleted per author
- **Activity Patterns**: Most active day of the week and hour
- **Repository Age**: Time since first commit
- **File Type Breakdown**: Distribution of file extensions
- **Largest Files**: Top 5 largest files by line count
- **Recent Activity**: Commits and authors active in the last 7 days
- **Fun Facts**: Longest commit message, commit with most files changed

#### Usage

```bash
# Make sure you're in the api1 directory
cd api1

# Run the statistics script
./scripts/codebase_stats.sh
```

#### Requirements

- Git repository
- Bash shell
- Unix-like system (macOS/Linux)

#### Example Output

```
🚀 RPG Codebase Statistics Generator
=====================================

📊 Repository: rpg1

📈 Lines of Code Analysis
========================
📁 Total Files: 55
📝 Total Lines: 18002
  ├─ Go files: 47 files, 17361 lines
  ├─ SQL files: 4 files, 56 lines
  ├─ Shell scripts: 1 files, 200 lines
  └─ Documentation: 3 files, 385 lines

👥 Author Statistics
===================
📊 Commit Count by Author:
  19 commits by Marcus Kalmár
  1 commits by Marcus M. Kalmár

📈 Lines Added/Deleted by Author:
  Marcus Kalmár:
    ├─ Added: +38664 lines
    ├─ Deleted: -2246 lines
    └─ Net: +36418 lines

🎯 Repository Summary
===================
📊 Total lines of code: 18002
👥 Total authors: 2
📝 Total commits: 20
📁 Total files: 55

✨ Statistics generation complete! ✨
```

#### What It Excludes

The script automatically excludes:
- `vendor/` directory (Go dependencies)
- `node_modules/` directory (Node.js dependencies)
- `.git/` directory (Git metadata)
- Generated files and binaries

#### Customization

You can modify the script to:
- Add more file types to analyze
- Change the time periods for recent activity
- Add more statistics or fun facts
- Customize the color scheme

## Adding New Scripts

When adding new scripts to this directory:

1. Make them executable: `chmod +x scripts/your_script.sh`
2. Add documentation to this README
3. Follow the same naming convention and style
4. Include error handling and helpful output

## Contributing

Feel free to enhance these scripts or add new ones! Some ideas:
- Database migration helpers
- Environment setup scripts
- Testing utilities
- Deployment helpers
- Code quality checks
