import os
import re
import sys
from concurrent.futures import ThreadPoolExecutor

# Regex to detect absolute paths and file:/// URIs
def get_repo_patterns():
    root = os.path.abspath('.')
    root_fwd = root.replace('\\', '/')
    root_drive = root_fwd.split(':')[0]
    root_path = root_fwd.split(':')[1] if ':' in root_fwd else root_fwd
    
    p1 = re.compile(rf'(?i)file:///{root_drive}:{re.escape(root_path)}[/\\]+')
    p2 = re.compile(rf'(?i){root_drive}:{re.escape(root_path.replace("/", "\\"))}[/\\]+')
    return p1, p2

# Generic file:/// URIs
GENERIC_FILE_URI = re.compile(r'(?i)file:///(?:[a-z]:[/\\]+|[/\\]+)')
# Generic Win paths (e.g. /Users/Admin). We use a group to capture the path part so we can fix its slashes.
GENERIC_WIN_PATH = re.compile(r'(?i)(?<![a-zA-Z0-9])[a-z]:\\([a-zA-Z0-9_][^ \n\r\t"\'<>()]*)')

EXCLUDE_DIRS = {'.git', 'node_modules', 'dist', 'build', '.ci-out', 'tmp', '.agent'}
EXCLUDE_EXTS = {'.png', '.jpg', '.jpeg', '.zip', '.tar', '.gz', '.woff', '.woff2', '.exe', '.dll', '.bin'}

def fix_file(filepath):
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
    except Exception:
        return False
        
    original = content
    p1, p2 = get_repo_patterns()
    
    # 1. Strip exact repo paths
    content = p1.sub('', content)
    content = p2.sub('', content)
    
    # 2. Generic file:/// URIs -> /
    content = GENERIC_FILE_URI.sub('/', content)
    
    # 3. Generic C:\ paths -> /path
    def win_replacer(match):
        path_part = match.group(1)
        # Ensure we only replace backslashes with forward slashes for the path part
        return '/' + path_part.replace('\\', '/')
        
    content = GENERIC_WIN_PATH.sub(win_replacer, content)

    if content != original:
        with open(filepath, 'w', encoding='utf-8', newline='') as f:
            f.write(content)
        return True
    return False

def main():
    print("Running 04-relative-path-fixer.py...")
    all_files = []
    for root, dirs, files in os.walk('.'):
        dirs[:] = [d for d in dirs if d not in EXCLUDE_DIRS]
        for f in files:
            ext = os.path.splitext(f)[1].lower()
            if ext not in EXCLUDE_EXTS:
                all_files.append(os.path.join(root, f))

    fixed_count = 0
    with ThreadPoolExecutor(max_workers=3) as executor:
        results = executor.map(fix_file, all_files)
        for r in results:
            if r:
                fixed_count += 1

    print(f"✅ Fixed absolute paths in {fixed_count} files.")
    sys.exit(0)

if __name__ == "__main__":
    main()
