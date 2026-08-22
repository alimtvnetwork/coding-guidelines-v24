find . -name "*.md" -not -path "*/node_modules/*" | grep -E "[A-Z]" > files.txt; while read f; do lower="${f,,}"; git mv "$f" "$lower" 2>/dev/null || mv "$f" "$lower"; done < files.txt; rm files.txt
