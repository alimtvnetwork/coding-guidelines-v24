find . -name "readme.md" -not -path "*/node_modules/*" | while read f; do lower="${f%readme.md}readme.md"; git mv "$f" "$lower" 2>/dev/null || mv "$f" "$lower"; done
