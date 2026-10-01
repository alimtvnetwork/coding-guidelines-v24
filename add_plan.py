import glob
import re

files = glob.glob("01-prompts/14-execute/*.md") + glob.glob("01-prompts/15-cg-execute/*.md")
count = 0

plan_text = "\n\n[/plan](slashCommand;plan) Execute thorough step-by-step planning in the repository before execution. Ensure all deliverables, architecture boundaries, and requirements are clearly defined in the audit ledger and subtask plans before dispatching worker waves."

for file in files:
    with open(file, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Check if already added
    if "[/plan]" not in content:
        # Find the end of the [/learn] block
        match = re.search(r'\[/learn\]\(slashCommand;learn\).*?(?=\n\n)', content, re.DOTALL)
        if match:
            new_content = content[:match.end()] + plan_text + content[match.end():]
            with open(file, 'w', encoding='utf-8') as f:
                f.write(new_content)
            count += 1

print(f"Added plan mode to {count} files.")
