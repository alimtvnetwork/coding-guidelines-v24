import glob
import re

files = glob.glob("01-prompts/**/*.md", recursive=True)
count = 0

for file in files:
    with open(file, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Replace (slashCommand:X) with (slashCommand;X)
    new_content, num_subs = re.subn(r'\(slashCommand:([^)]+)\)', r'(slashCommand;\1)', content)
    
    if num_subs > 0:
        with open(file, 'w', encoding='utf-8') as f:
            f.write(new_content)
        count += 1

print(f"Updated {count} files.")
