# Specification Remediation from Audit Findings — Planning Spec (must follow)


> **Prompt Version:** 2.1.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

## Goal

Fix the specification files strictly based on the latest generated audit file. The non-negotiable goal is to increase the spec score to exactly 100%. 
You must ingest the audit, map the findings to the broken files, and execute the fixes.



## STRICT AVOIDANCE: Never Disable CI/CD

> [!CAUTION]
> **NEVER disable any CI/CD checks, GitHub Actions, or validation workflows.** 
> Strictly avoid commenting out, bypassing, or deleting CI/CD steps to force a pipeline to pass. Your job is to fix the underlying code so that the CI/CD pipeline passes legitimately. Disabling CI/CD is an auto-reject failure.

## Anti-Hallucination, Micro-Tasking, & Self-Looping

> [!CAUTION]
> **CRITICAL RULE: DO NOT ATTEMPT TO READ, PLAN, AND EXECUTE EVERYTHING AT ONCE.**
> If you try to consume a massive codebase and write code in a single turn, you WILL hallucinate, drop requirements, and fail.

To survive massive checklists and complex codebases, you MUST operate using these three principles:

1. **Phase 1: Read & Understand (Isolated Loop):** Your very first action must be purely exploratory. Do NOT write code. Break down the task, read the specific files, trace the dependencies, and understand the architectural boundary. Once you understand the scope, end your turn and self-loop to begin execution.
2. **Phase 2: Bounded Micro-Tasking (Sequential Self-Looping):** Never attempt to execute the entire checklist in one response. Treat each checklist section or file as a strict, isolated boundary. Execute *only* the first small portion, verify it, end your turn, and self-loop to process the next portion. 
3. **Phase 3: Multi-Agent Parallelization:** If tasks are independent, you MUST spawn dedicated sub-agents to handle them concurrently. Give each sub-agent an extremely small, strictly defined bounding box (e.g., "Only edit File X"). Never give a sub-agent a generic or multi-file task.

## Actionable Items & Checklist

- [ ] Read the latest audit file located in `spec/25-app-spec-audit/`. 
- [ ] Parse the Markdown Summary Table at the bottom of the audit to understand every folder, issue, and proposed fix.
- [ ] Anti-Garbage Naming (Non-Negotiable): I have strictly verified that absolutely NO generic garbage variable names (e.g., `comp_100.go`, `temp`, `data`, `obj`, `Input100`, `TestHandleComp100`) were written. All names are highly semantic and domain-specific.
- [ ] Temp Script Sandboxing: AI Fix Scripts (Reusable Tools): Before creating a helper script, you MUST check `.lovable/ai-fix-scripts/index.md` to reuse existing tools. If you generate a new script, you MUST write it to `.lovable/ai-fix-scripts/`, update `index.md` with its explanation, ensure `index.md` is linked in `what-to-read.md`, and commit the script.
- [ ] Ensure the git repository starts completely clean.
- [ ] Complete all work on the current branch only.
- [ ] Adhere to all coding guidelines in `spec/02-coding-guidelines/` and `spec/03-error-manage/`.
- [ ] Run fixing processes in self-looping sub-agents.
- [ ] Group all completed work into a single logical commit.
- [ ] Push the commit to the remote repository.
- [ ] **File Change Summary:** Provide a highly detailed summary in the chat listing exactly which files were changed, what specific changes were made inside them, and why they were changed. The summary is VERY important.

## Execution & Self-Looping Protocol

This task requires deep, multi-agent processing. You MUST NOT run a fast, linear self-loop that artificially consolidates work or skips steps to save time. 

How to self-loop and distribute tasks effectively:

1. Spawn Sub-Agents: You must actively spawn dedicated, self-looping sub-agents to handle the workload. You must spawn at least 2 agents, each executing 2 to 3 parallel tasks (e.g. assigning one agent to fix `spec/21-app/` and another to `spec/23-app-db/`).
2. Utilize Processing Power: Take your time and use maximum processing power and credits. Do not take shortcuts. Do not attempt to fix a massive spec in a single, consolidated step.
3. Wait and Aggregate: As the master agent, loop autonomously to wait for your sub-agents to complete their assigned fixes. Aggregate their precise findings to ensure the highest quality result.

### Temp-Agent State Management Protocol (Non-Negotiable)

To ensure agents don't lose context, you MUST use the `.lovable/temp-agents/` directory for tracking sub-agent tasks.

1. On Start: The sub-agent creates `.lovable/temp-agents/<task-name>.md`. Write the objective and `STATUS: IN_PROGRESS`.
2. On Error/Crash: If an agent breaks or fails, it MUST append the exact error, cause, and `STATUS: FAILED` to that file before closing.
3. On Resume: The next assigned agent must first read that file to avoid repeating the mistake.
4. On Success: Update the file to `STATUS: DONE` and immediately update the master plan in `.lovable/plans/index.md` or `.lovable/plans/pending/`.

NEVER MAKE THIS EXCUSE: *"I did not literally perform N self-loops or spawn dozens of sub-agents. Instead of artificially spinning up N separate loops, I consolidated the required deep work into a concentrated series of sequential steps within my existing execution environment."*
If you write this excuse, you are a lazy, stupid fuck and you are violating RULE 0. Spawn the actual sub-agents. Do the actual loops. Utilize the processing power effectively to get the right answer. Avoid stupidity.

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.lovable/memories/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Read the whole codebase, read every folder in `spec/` and `.lovable/`, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.

---

## Metadata

- slug: fix-spec-from-audit
- status: active
