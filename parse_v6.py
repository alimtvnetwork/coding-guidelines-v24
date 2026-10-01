import re
with open("01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md", "r", encoding="utf-8") as f:
    text = f.read()

print("Sections:")
for match in re.finditer(r'^(#+ .*)$', text, re.MULTILINE):
    print(match.group(1))
