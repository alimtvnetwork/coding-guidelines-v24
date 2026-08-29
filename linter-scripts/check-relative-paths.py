import os
import re
import sys
from concurrent.futures import ThreadPoolExecutor

# Regex to detect absolute paths and file:/// URIs
# - file:/// URIs
# - Windows drive paths: C:\..., D:\...
# Require at least one alphanumeric character after the slash to avoid matching a:\\ n in JSON
ABSOLUTE_PATH_REGEX = re.compile(
    r'(?i)(?:file:///[a-z]:[/\\]+|file:///[/\\]+|(?<![a-zA-Z0-9])(?:[a-z]:\\[a-zA-Z0-9_][^ \n\r\t"\'<>()]*))'
)

EXCLUDE_DIRS = {'.git', 'node_modules', 'dist', 'build', '.ci-out', 'tmp', '.agent'}
EXCLUDE_EXTS = {'.png', '.jpg', '.jpeg', '.zip', '.tar', '.gz', '.woff', '.woff2', '.exe', '.dll', '.bin'}

def scan_file(filepath):
    violations = []
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            for line_no, line in enumerate(f, 1):
                # find all matches
                for match in ABSOLUTE_PATH_REGEX.finditer(line):
                    violations.append((filepath, line_no, match.group(0)))
    except Exception:
        pass
    return violations

def main():
    print("Running check-relative-paths.py...")
    all_files = []
    for root, dirs, files in os.walk('.'):
        dirs[:] = [d for d in dirs if d not in EXCLUDE_DIRS]
        for f in files:
            ext = os.path.splitext(f)[1].lower()
            if ext not in EXCLUDE_EXTS:
                all_files.append(os.path.join(root, f))

    all_violations = []
    with ThreadPoolExecutor(max_workers=3) as executor:
        results = executor.map(scan_file, all_files)
        for res in results:
            if res:
                all_violations.extend(res)

    if all_violations:
        print(f"❌ FOUND {len(all_violations)} ABSOLUTE PATH VIOLATIONS:")
        for v in all_violations:
            print(f"  {v[0]}:{v[1]} -> {v[2]}")
        sys.exit(1)
    else:
        print("✅ No absolute paths found. All paths are relative.")
        sys.exit(0)

if __name__ == "__main__":
    main()
