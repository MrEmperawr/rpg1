#!/bin/bash

# Codebase Statistics Generator
# Generates fun stats about the authors of the codebase

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo -e "${RED}Error: Not in a git repository${NC}"
    exit 1
fi

echo -e "${CYAN}🚀 RPG Codebase Statistics Generator${NC}"
echo -e "${CYAN}=====================================${NC}"
echo ""

REPO_NAME=$(basename -s .git $(git config --get remote.origin.url 2>/dev/null) || echo "rpg1")
echo -e "${BLUE}📊 Repository: ${REPO_NAME}${NC}"
echo ""

echo -e "${YELLOW}📈 Lines of Code Analysis${NC}"
echo -e "${YELLOW}========================${NC}"

GO_LINES=$(find . -name "*.go" -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" -exec wc -l {} + | tail -1 | awk '{print $1}')
GO_FILES=$(find . -name "*.go" -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" | wc -l | tr -d ' ')

SQL_LINES=$(find . -name "*.sql" -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" -exec wc -l {} + 2>/dev/null | tail -1 | awk '{print $1}' || echo "0")
SQL_FILES=$(find . -name "*.sql" -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" | wc -l | tr -d ' ')

SH_LINES=$(find . -name "*.sh" -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" -exec wc -l {} + 2>/dev/null | tail -1 | awk '{print $1}' || echo "0")
SH_FILES=$(find . -name "*.sh" -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" | wc -l | tr -d ' ')

MD_LINES=$(find . -name "*.md" -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" -exec wc -l {} + 2>/dev/null | tail -1 | awk '{print $1}' || echo "0")
MD_FILES=$(find . -name "*.md" -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" | wc -l | tr -d ' ')

TOTAL_LINES=$((GO_LINES + SQL_LINES + SH_LINES + MD_LINES))
TOTAL_FILES=$((GO_FILES + SQL_FILES + SH_FILES + MD_FILES))

echo -e "${GREEN}📁 Total Files: ${TOTAL_FILES}${NC}"
echo -e "${GREEN}📝 Total Lines: ${TOTAL_LINES}${NC}"
echo -e "  ├─ Go files: ${GO_FILES} files, ${GO_LINES} lines"
echo -e "  ├─ SQL files: ${SQL_FILES} files, ${SQL_LINES} lines"
echo -e "  ├─ Shell scripts: ${SH_FILES} files, ${SH_LINES} lines"
echo -e "  └─ Documentation: ${MD_FILES} files, ${MD_LINES} lines"
echo ""

echo -e "${YELLOW}👥 Author Statistics${NC}"
echo -e "${YELLOW}===================${NC}"

echo -e "${BLUE}📊 Commit Count by Author:${NC}"
git shortlog -sn --all | while read commits author; do
    echo -e "  ${GREEN}${commits}${NC} commits by ${CYAN}${author}${NC}"
done
echo ""

echo -e "${BLUE}📈 Lines Added/Deleted by Author:${NC}"
git log --format='%aN' --all | sort | uniq | while read author; do
    LINES_ADDED=$(git log --author="$author" --pretty=tformat: --numstat | awk '{add+=$1} END {print add}' | tr -d ' ')
    LINES_DELETED=$(git log --author="$author" --pretty=tformat: --numstat | awk '{del+=$2} END {print del}' | tr -d ' ')
    
    LINES_ADDED=${LINES_ADDED:-0}
    LINES_DELETED=${LINES_DELETED:-0}
    
    NET_LINES=$((LINES_ADDED - LINES_DELETED))
    
    if [ $NET_LINES -gt 0 ]; then
        NET_COLOR=$GREEN
        NET_SIGN="+"
    elif [ $NET_LINES -lt 0 ]; then
        NET_COLOR=$RED
        NET_SIGN=""
    else
        NET_COLOR=$YELLOW
        NET_SIGN=""
    fi
    
    echo -e "  ${CYAN}${author}${NC}:"
    echo -e "    ├─ Added: ${GREEN}+${LINES_ADDED}${NC} lines"
    echo -e "    ├─ Deleted: ${RED}-${LINES_DELETED}${NC} lines"
    echo -e "    └─ Net: ${NET_COLOR}${NET_SIGN}${NET_LINES}${NC} lines"
done
echo ""

echo -e "${BLUE}📅 Most Active Day of the Week:${NC}"
git log --format='%ad' --date=format:'%A' --all | sort | uniq -c | sort -nr | head -1 | while read count day; do
    echo -e "  ${GREEN}${day}${NC} with ${CYAN}${count}${NC} commits"
done
echo ""

echo -e "${BLUE}🕐 Most Active Hour:${NC}"
git log --format='%ad' --date=format:'%H' --all | sort | uniq -c | sort -nr | head -1 | while read count hour; do
    echo -e "  ${GREEN}${hour}:00${NC} with ${CYAN}${count}${NC} commits"
done
echo ""

echo -e "${BLUE}📅 Repository Age:${NC}"
FIRST_COMMIT=$(git log --reverse --format='%ad' --date=short | head -1)
LAST_COMMIT=$(git log --format='%ad' --date=short | head -1)
echo -e "  First commit: ${GREEN}${FIRST_COMMIT}${NC}"
echo -e "  Last commit: ${GREEN}${LAST_COMMIT}${NC}"

if [ ! -z "$FIRST_COMMIT" ] && [ ! -z "$LAST_COMMIT" ]; then
    if [[ "$OSTYPE" == "darwin"* ]]; then
        DAYS=$(( ($(date -j -f "%Y-%m-%d" "$LAST_COMMIT" +%s) - $(date -j -f "%Y-%m-%d" "$FIRST_COMMIT" +%s)) / 86400 ))
    else
        DAYS=$(( ($(date -d "$LAST_COMMIT" +%s) - $(date -d "$FIRST_COMMIT" +%s)) / 86400 ))
    fi
    echo -e "  Repository age: ${CYAN}${DAYS} days${NC}"
fi
echo ""

echo -e "${YELLOW}📁 File Type Breakdown${NC}"
echo -e "${YELLOW}=====================${NC}"

echo -e "${BLUE}File extensions:${NC}"
find . -type f -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" | grep -E '\.[a-zA-Z0-9]+$' | sed 's/.*\.//' | sort | uniq -c | sort -nr | head -10 | while read count ext; do
    echo -e "  ${GREEN}${ext}${NC}: ${CYAN}${count}${NC} files"
done
echo ""

# Largest files
echo -e "${BLUE}📏 Largest Files:${NC}"
find . -type f -not -path "./vendor/*" -not -path "./node_modules/*" -not -path "./.git/*" -not -path "./.gitignore" | xargs wc -l 2>/dev/null | sort -nr | head -6 | tail -5 | while read lines filename; do
    if [ "$lines" != "total" ]; then
        echo -e "  ${GREEN}${filename}${NC}: ${CYAN}${lines}${NC} lines"
    fi
done
echo ""

echo -e "${YELLOW}🕒 Recent Activity (Last 7 Days)${NC}"
echo -e "${YELLOW}===============================${NC}"

RECENT_COMMITS=$(git log --since="7 days ago" --oneline | wc -l | tr -d ' ')
RECENT_AUTHORS=$(git log --since="7 days ago" --format='%aN' | sort | uniq | wc -l | tr -d ' ')

echo -e "${GREEN}${RECENT_COMMITS}${NC} commits in the last 7 days"
echo -e "${GREEN}${RECENT_AUTHORS}${NC} authors active in the last 7 days"
echo ""

echo -e "${YELLOW}🎉 Fun Facts${NC}"
echo -e "${YELLOW}===========${NC}"

LONGEST_COMMIT=$(git log --format='%s' --all | awk '{print length, $0}' | sort -nr | head -1 | cut -d' ' -f2-)
echo -e "${BLUE}Longest commit message:${NC}"
echo -e "  ${CYAN}${LONGEST_COMMIT}${NC}"
echo ""

MOST_FILES_COMMIT=$(git log --format='%h %s' --all | while read hash message; do
    files=$(git show --name-only --format="" $hash | wc -l | tr -d ' ')
    echo "$files $hash $message"
done | sort -nr | head -1 | cut -d' ' -f2-)
echo -e "${BLUE}Commit with most files changed:${NC}"
echo -e "  ${CYAN}${MOST_FILES_COMMIT}${NC}"
echo ""

echo -e "${PURPLE}🎯 Repository Summary${NC}"
echo -e "${PURPLE}===================${NC}"
echo -e "📊 Total lines of code: ${GREEN}${TOTAL_LINES}${NC}"
echo -e "👥 Total authors: ${GREEN}$(git log --format='%aN' --all | sort | uniq | wc -l | tr -d ' ')${NC}"
echo -e "📝 Total commits: ${GREEN}$(git rev-list --count --all)${NC}"
echo -e "📁 Total files: ${GREEN}${TOTAL_FILES}${NC}"
echo ""

echo -e "${CYAN}✨ Statistics generation complete! ✨${NC}"
