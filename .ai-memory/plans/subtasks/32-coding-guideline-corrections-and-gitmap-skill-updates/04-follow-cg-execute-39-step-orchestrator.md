# Subtask 04: Follow CG Execute 39-Step Orchestrator

- **Parent Plan:** [.ai-memory/plans/pending/32-coding-guideline-corrections-and-gitmap-skill-updates.md](.ai-memory/plans/pending/32-coding-guideline-corrections-and-gitmap-skill-updates.md)
- **Target Files:**
  - `.agents/skills/follow-cg-execute/skill.md`
  - `.cursor/skills/follow-cg-execute/skill.md`
- **Status:** Completed

---

## Action Items

1. In `.agents/skills/follow-cg-execute/skill.md` and `.cursor/skills/follow-cg-execute/skill.md`, author the complete 39-step execution sequence.
2. Clearly enumerate all 39 steps with their respective prompt path and target coding guideline focus.
3. Establish the operational loop protocol:
   - Automated via `next` iterations.
   - Steps 1 through 38: fast-forwarding mode (zero build, zero test, bounded 5–8 file refactor passes, atomic local commits).
   - Step 39: terminal verification (full builds, E2E tests, 100% green pass, minor bump release ceremony).
4. Synchronize across both `.agents/skills/` and `.cursor/skills/` in both repositories.
