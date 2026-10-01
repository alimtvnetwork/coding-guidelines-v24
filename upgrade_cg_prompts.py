import os
import glob
import re

banned_prefixes = [
    "## Banned Operations Checklist",
    "## Final Step Git Commit",
    "## Pre-Reply / Loop Checklist",
    "## Continuous 2-Phase Self-Loop",
    "## Strictly Avoid",
    "## AI Fix Scripts Memory",
    "## Non-Negotiable Coding Guidelines",
    "## Strict In-Repository Execution",
    "## Metadata",
    "## MUST FOLLOW NON-NEGOTIABLE",
    "## Phase 1:",
    "## Phase 2:",
    "## Phase 3:",
    "## Phase 1B:",
    "## Ruthless Orchestration",
    "## Anti-Hallucination",
    "## STRICT AVOIDANCE",
    "## Phase 1 Violation Ledger Format",
    "## The Unified Master Pipeline",
    "## Execution & Build Policy",
    "## Task Consolidation",
    "## Verification Checklist",
    "## Phase 1A:",
    "## 2-Agent Parallel Orchestration",
    "## \U0001F6A8 High Priority Instructions Below",
    "## End-of-Turn Verification",
    "## Standardized N-Step Self-Loop Architecture"
]

def is_banned(heading):
    heading = re.sub(r'^## \d+\.\s*', '## ', heading.strip())
    for bp in banned_prefixes:
        if heading.startswith(bp):
            return True
    return False

v6_path = "01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md"
with open(v6_path, "r", encoding="utf-8") as f:
    v6_content = f.read()

param_block_match = re.search(r'(```text\nN = 300.*?```)', v6_content, re.DOTALL)
param_block = param_block_match.group(1)

mandatory_gate_match = re.search(r'(### \U0001F6A8 MANDATORY SUBAGENT SPAWNING GATE.*?)(?=---)', v6_content, re.DOTALL)
mandatory_gate = mandatory_gate_match.group(1).strip()

sections_1_to_9_match = re.search(r'(## The Unified Master Pipeline.*?)(?=## 10\. Targeted)', v6_content, re.DOTALL)
sections_1_to_9 = sections_1_to_9_match.group(1).strip()

sections_11_to_end_match = re.search(r'(## 11\. Discovery Toolchain.*)', v6_content, re.DOTALL)
sections_11_to_end = sections_11_to_end_match.group(1).strip()

cg_files = glob.glob("01-prompts/15-cg-execute/*.md")
for cg_file in cg_files:
    if "readme.md" in cg_file.lower() or "cg-execute-in-below-steps" in cg_file.lower():
        continue
    
    with open(cg_file, "r", encoding="utf-8") as f:
        cg_text = f.read()
    
    goal_match = re.search(r'\[/goal\]\(slashCommand:goal\)\s*(.*?)(?=until 100%|with strict|Spawn autonomous)', cg_text, re.IGNORECASE)
    cg_goal = goal_match.group(1).strip() if goal_match else "Autonomously orchestrate and execute the task"
    if cg_goal.endswith(": FIRST showcase"):
        cg_goal = cg_goal.split(": FIRST showcase")[0].strip()
        
    metadata_match = re.search(r'(## Metadata\n\n- slug:.*?status:.*?\n)', cg_text, re.DOTALL)
    metadata = metadata_match.group(1).strip() if metadata_match else ""
    
    # Extract blocks
    blocks = []
    current_heading = None
    current_content = []
    
    for line in cg_text.splitlines():
        if line.startswith("## "):
            if current_heading:
                if not is_banned(current_heading):
                    blocks.append((current_heading, "\n".join(current_content).strip()))
            current_heading = line
            current_content = []
        else:
            if current_heading:
                current_content.append(line)
                
    if current_heading and not is_banned(current_heading):
        blocks.append((current_heading, "\n".join(current_content).strip()))
        
    # Build new text
    new_text = f"{param_block}\n\n"
    new_text += f"[/goal](slashCommand:goal) {cg_goal}: FIRST showcase and list out the given task in visible chat during Turn 1, capture it verbatim, plan it in the repo, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution without calling `invoke_subagent` is an auto-reject failure) in disjoint file boxes using GitMap high-speed commands as primary, prove every single claim with concrete evidence, enforce coding guidelines to 100%, and finish with one atomic GitMap commit that holds strictly this task's files.\n\n"
    new_text += "[/learn](slashCommand:learn) Enforce the Top-Instruction Priority Mandate: whatever directives, custom requirements, checklists, or user instructions are provided ABOVE this prompt outrank everything below. Turn 1 MUST showcase the given task list in visible chat before any background execution. Each rule is stated once (R1 to R16) and cited by ID. Progress lives in the ledger and in `.ai-memory/plans/`, never only in chat.\n\n"
    new_text += f"{mandatory_gate}\n\n---\n\n"
    new_text += f"{sections_1_to_9}\n\n---\n\n"
    new_text += f"## 10. Targeted Verification Checks\n\n"
    new_text += "Confirm scripts exist via harmless workspace call before invoking (R4). Run on changed files/folders only:\n\n"
    new_text += "- **Coding Guidelines & Boolean Linter:** `python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only --ext <.ext>`\n"
    new_text += "- **Relative Path Linter:** `python linter-scripts/check-relative-paths.py`\n"
    new_text += "- **Prompts & Spec Index Linter:** `python linter-scripts/check-prompts-loaded.py`\n"
    new_text += "- **Markdown Link & Doc Path Linter:** `python 03-ai-scripts/22-doc-path-linter.py <folder>`\n"
    new_text += "- **Sequence Integrity Linter:** `python linter-scripts/check-sequence-integrity.py`\n"
    new_text += "- **Forbidden Strings Check:** `python linter-scripts/check-forbidden-strings.py`\n\n---\n\n"
    new_text += f"{sections_11_to_end}\n\n---\n\n"
    
    # Specific payload blocks
    for heading, content in blocks:
        # Avoid double-adding Metadata if it slipped through
        if "Metadata" in heading:
            continue
        new_text += f"{heading}\n\n{content}\n\n---\n\n"
        
    if metadata:
        new_text += f"{metadata}\n"
        
    with open(cg_file, "w", encoding="utf-8") as f:
        f.write(new_text.strip() + "\n")

print(f"Properly upgraded {len(cg_files)} CG execute prompts.")
