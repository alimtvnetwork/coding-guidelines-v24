import re
import glob

with open("01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md", "r", encoding="utf-8") as f:
    v6_content = f.read()

gate_match = re.search(r'(### \U0001F6A8 MANDATORY SUBAGENT SPAWNING GATE.*?)(?=---)', v6_content, re.DOTALL)
new_gate = gate_match.group(1).strip()

phase_1_match = re.search(r'(## 6\. Phase 1: Planning Step.*?)(?=---)', v6_content, re.DOTALL)
new_phase_1 = phase_1_match.group(1).strip()

phase_2_match = re.search(r'(## 7\. Phase 2: Execution Step.*?)(?=---)', v6_content, re.DOTALL)
new_phase_2 = phase_2_match.group(1).strip()

turn_yielding_match = re.search(r'(### 7\.3 Turn-Yielding & Verification Protocol.*?)(?=---)', v6_content, re.DOTALL)
# It's already in new_phase_2 because it's a subsection of 7, wait, no. 
# new_phase_2 goes until the first ---. Let's see if there is a --- before 7.3.
# No, there isn't. So new_phase_2 contains 7.1, 7.2, and 7.3.

cg_files = glob.glob("01-prompts/15-cg-execute/*.md")

count = 0
for file in cg_files:
    if "readme.md" in file.lower():
        continue
    with open(file, "r", encoding="utf-8") as f:
        content = f.read()
        
    # Replace gate
    content = re.sub(r'### \U0001F6A8 MANDATORY SUBAGENT SPAWNING GATE.*?---', new_gate + "\n\n---", content, flags=re.DOTALL)
    
    # Replace phase 1
    content = re.sub(r'## 6\. Phase 1B: Spec, Plan & Lean Subtasks.*?---', new_phase_1 + "\n\n---", content, flags=re.DOTALL)
    
    # Replace phase 2 (including turn-yielding)
    content = re.sub(r'## 7\. Phase 2: Mandatory Worker Waves & Coding Guidelines Enforcement.*?---', new_phase_2 + "\n\n---", content, flags=re.DOTALL)
    
    with open(file, "w", encoding="utf-8") as f:
        f.write(content)
    count += 1

print(f"Updated {count} CG files with new V6 blocks.")
