# Spec Audits

> **/goal** Keep audit results next to the specs they judge, and use each failed row as the next edit.
> **/learn** An audit file is a work list. It is not a score to store and leave.

**Version:** 4.3.0
**Updated:** 2026-10-02
**Status:** Active
**Ambiguity:** Older notes under `.ai-memory/audits/` are not this module.

---

## Where audits go

Spec audits live in this folder, `02-spec/25-spec-audits/`. They do not live under `.ai-memory/audits/`.

An audit improves the spec it names. A failed row is fixed in that spec, then the audit row is marked done. A passing score that hides a contradiction is a failed audit.

---

## Inventory

| # | File | What it judges | Next use |
|---|---|---|---|
| 01 | [01-v3-v4-v5-prompt-benchmark-audit.md](./01-v3-v4-v5-prompt-benchmark-audit.md) | Execute prompts V3, V4, V5 | Prompt edits only. Do not treat its self-score as measured. |
| 02 | [02-v3-v4-v5-v6-prompt-benchmark-audit.md](./02-v3-v4-v5-v6-prompt-benchmark-audit.md) | Execute prompts through V6 | Same rule as 01. |
| 03 | [03-design-spec-source-ledger/readme.md](./03-design-spec-source-ledger/readme.md) | Numbers in `02-spec/07-design-system/` files 24–38 | Verify each unverified row against the source, or mark the spec line `Not specified in source. Do not invent.` |

---

## Rule

Do not copy a client, product, or repository name into an audit. Cite the spec path and the symbol.
