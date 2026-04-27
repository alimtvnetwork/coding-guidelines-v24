# Gap Analysis — Are the Installer Specs "Blind-AI Ready"?

> **Question:** Given only the specifications in this repository (no chat,
> no clarifying questions, no access to a maintainer), could a fresh AI
> agent regenerate `linters-cicd/install.sh` and `linters-cicd/install.ps1`
> faithfully — i.e. byte-for-byte equivalent in *behavior*, not in source?
>
> **Verdict:** **Partially.** The conceptual contract is strong, but
> several concrete details required to produce a working installer are
> implicit, scattered, or contradicted across spec files. A blind AI would
> produce a script that *looks* conformant and passes a casual review,
> but would diverge from the shipped scripts in measurable ways.
>
> **Date:** 2026-04-24
> **Scope reviewed:**
> - `spec/14-update/27-generic-installer-behavior.md` (canonical)
> - `spec/15-distribution-and-runner/01-install-contract.md`
> - `spec/12-cicd-pipeline-workflows/04-install-script-generation.md`
> - Reference implementations: `linters-cicd/install.sh`, `linters-cicd/install.ps1`
> - Public documentation: `linters-cicd/README.md`

---

## 1. What a blind AI would get right (no gap)

These items are unambiguous in the specs and the shipped scripts match
them. A blind AI would reproduce them correctly.

| Capability | Source of truth | Confidence |
|---|---|---|
| Two-mode dispatch (PINNED vs IMPLICIT) at script start | spec 27 §3 | High |
| Exit-code table 0–5 with code `4` semantics dependent on verification flag | spec 27 §8 + README §"Installer exit codes" | High |
| Banner format with `mode / repo / version / source` lines | spec 27 §7 | High |
| `--version` / `-Version` and `--no-verify` / `-NoVerify` flag pair | spec 27 §4, §9 + README | High |
| Reject `http://`, require HTTPS | spec 27 §9 | High |
| Loud warning when verification is disabled | spec 27 §9 ("loud warning") | High |
| Cleanup of temp dir on success and failure (`trap` / `try–finally`) | spec 15 §"Cleanup contract" | High |
| Help/usage available **before** any network call | (only in the script comments — NOT in the spec) | **Gap, see §3.1** |

---

## 2. What a blind AI would get *wrong by default*

These are the items where the spec is silent, ambiguous, or
self-contradictory. A blind AI would have to guess — and would almost
certainly guess differently than the shipped script.

### 2.1 Conflicting "default install target"

| Source | What it says |
|---|---|
| `spec/15-distribution-and-runner/01-install-contract.md` §"What gets installed (default)" | Installs **four folders** (`spec/`, `linters/`, `linter-scripts/`, `linters-cicd/`) into the **current working directory**. |
| `linters-cicd/install.sh` & `install.ps1` | Installs **only `linters-cicd/`** (one folder), default destination `./linters-cicd`. |
| `spec/14-update/27-generic-installer-behavior.md` | Silent on what gets installed — only describes *how* it's fetched. |

**Impact:** A blind AI reading spec 15 would build a four-folder
installer with `--folders` / `-Folders` overrides, `--prompt` /
`-Prompt`, `--force`, `--dry-run`, `--list-versions`, and
`--list-folders`. None of those flags exist in the shipped scripts.
Conversely, an AI reading spec 27 alone would not know what folders to
extract at all.

**Severity:** 🔴 High — these are two different products described as
one.

### 2.2 Conflicting one-liner URLs

- spec 15: `…/raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v18/main/install.sh`
  (i.e. fetched from the **branch tip**)
- README & shipped script: `…/releases/latest/download/install.sh`
  (i.e. fetched from a **release asset**)

The release-asset URL is clearly the intended channel (it's the only one
that supports checksums), but spec 15 still advertises the `main`-branch
URL. A blind AI might wire up neither, both, or the wrong one.

### 2.3 Asset-name resolution is not specified

The shipped Bash installer goes through this dance:

1. Try the un-versioned name `coding-guidelines-linters.zip`.
2. On failure, hit `…/releases/latest` to resolve the tag.
3. Retry with `coding-guidelines-linters-<TAG>.zip`.

None of this is in any spec. A blind AI would either:
- assume a single asset name (and fail when GitHub serves the
  versioned one), or
- always hit the API first (slower, hits rate limits faster).

**Severity:** 🟠 Medium — works fine until releases stop publishing the
short-named alias.

### 2.4 V → V+N discovery is in the spec but **not** in the shipped script

Spec 27 §6 mandates V→V+N parallel discovery with a default
`LOOKAHEAD = 20`. Neither `install.sh` nor `install.ps1` performs any
discovery — they hard-code the repo `alimtvnetwork/coding-guidelines-v18`.

This is a **conformance gap in the implementation**, not in the spec.
A blind AI following spec 27 would actually produce a *more* compliant
installer than the one currently shipped.

**Acceptance criterion 5 in spec 27 §10 currently fails for both shipped
scripts.**

### 2.5 Implicit-mode source-order step 3 (main-branch fallback)

Spec 27 §5.1 requires: latest release → V+N discovery → `main` branch
tarball, with a visible "unstable" warning. The shipped scripts skip
straight from "latest release" to a hard failure.

**Acceptance criteria 1 and 4 in spec 27 §10 also currently fail.**

### 2.6 Banner field `source:` does not match its allowed values

Spec 27 §7 says:

```
source: release-asset | tag-tarball | main-branch | local-archive
```

The shipped `install.sh` prints:

```
source:  release-asset (latest)
source:  release-asset ($VERSION)
```

The parenthetical suffix is undocumented. A linter that grep'd the
banner against the spec's enum would flag both shipped scripts.

### 2.7 Anti-requirements not enforced in code

Spec 15 §"Anti-requirements" says installers MUST NOT require Python.
The Bash installer falls back to `python3 -c …` for both tag resolution
(line 86) and zip extraction (line 116). A spec-strict blind AI would
omit the Python fallback and instead require `unzip` + `jq`/`grep`.

### 2.8 Checksum-file behavior on 404 is under-specified

- Spec 27 §9: "verify [checksums] by default; allow `--no-verify` only
  with a loud warning."
- Shipped script: if `checksums.txt` is missing, prints a yellow
  warning and **continues** the install with exit 0.

The spec does not say whether a missing checksum file is a hard error
or a soft warning. A blind AI could justifiably implement *either* and
still claim conformance. The shipped behavior leans permissive; a
security-paranoid AI would lean strict and break the user's first run.

### 2.9 PowerShell `--help` handling is undocumented

Both spec 27 and spec 15 describe `--help` for Bash but say nothing
about how PowerShell should accept the bash-style `--help` token (PS
parameter binding does not natively allow it). The shipped `install.ps1`
contains ~10 lines of special-case code (`$MyInvocation.UnboundArguments`
inspection) to make `pwsh ./install.ps1 --help` work. A blind AI would
ship `-Help` only and frustrate cross-platform users.

### 2.10 "Strip outer folder" post-extraction step

Both shipped scripts contain a step to detect and flatten a nested
`linters-cicd/` folder inside the extracted zip:

```bash
if [ -d "$DEST/linters-cicd" ]; then
    cp -R "$DEST/linters-cicd/." "$DEST/"
    rm -rf "$DEST/linters-cicd"
fi
```

This depends on the **release packaging format** (whether the zip is
flat or nested). That format is not documented in any spec. A blind AI
would not know to add this normalisation step, and the user would end
up with `./linters-cicd/linters-cicd/run-all.sh`.

### 2.11 "Next steps" footer is undocumented

Both scripts print:

```
Next steps:
   bash $DEST/run-all.sh --path . --format text
```

This is a UX convention with no spec backing. A blind AI would produce
a script that succeeds silently — technically conformant, practically
worse.

---

## 3. Documentation-quality gaps

### 3.1 Help is required to work offline — but only the script says so

The shipped scripts deliberately handle `-h` / `--help` / `-Help`
**before any network call**, so users with no internet (or behind
restrictive proxies) can still read usage. This is implemented and
commented in `install.ps1`, but it is **not** in the spec.

A blind AI would happily put help inside the main `try { … }` block
after `Invoke-WebRequest`, and `install.ps1 --help` on an air-gapped
machine would crash instead of printing usage.

### 3.2 The exit-code matrix lives in the README, not the spec

The detailed table showing how exit code `4` behaves with verification
ON vs OFF is in `linters-cicd/README.md`. Spec 27 §8 only lists the
codes; spec 15 §"Exit codes" lists 0/1/2 (and is *inconsistent* with
spec 27's 0–5 list).

A blind AI would have to choose between two contradictory exit-code
contracts (spec 15 vs spec 27). The README clarifies the intent, but
the README is not a spec.

### 3.3 Cross-spec contradictions

| Topic | spec 15 says | spec 27 says | Shipped script |
|---|---|---|---|
| Default install target | 4 folders into CWD | (silent) | 1 folder (`./linters-cicd`) |
| Exit codes | 0, 1, 2 only | 0–5 | 0, 1, 2, 3, 4 |
| Folder selection flags | `--folders` / `-Folders` | (silent) | not implemented |
| Interactive merge | `--prompt` / `-Prompt` | (silent) | not implemented |
| Listings | `--list-versions`, `--list-folders` | (silent) | not implemented |
| Discovery | (silent) | V→V+N with LOOKAHEAD=20 mandatory | not implemented |
| Main fallback | (silent) | mandatory with warning | not implemented |

**These two specs describe two different products and both claim to be
the contract for `install.sh`/`install.ps1`.** This is the single
biggest blind-AI-readiness blocker.

---

## 4. "Could a blind AI rebuild the shipped script?" — verdict matrix

| Aspect | Blind AI from spec 27 | Blind AI from spec 15 | Truth (shipped) |
|---|---|---|---|
| Mode dispatch | ✅ correct | ⚠️ inferred | ✅ |
| Repo / asset URL | ⚠️ generic placeholder | ❌ wrong (raw.githubusercontent) | ✅ release-asset |
| Folders installed | ❓ not specified | ❌ four folders | ✅ one folder |
| Checksum verification | ✅ correct | ⚠️ partial | ✅ |
| Exit code 4 semantics | ✅ correct | ❌ doesn't list code 4 | ✅ |
| V→V+N discovery | ✅ implements (over-conformant) | ❌ unaware | ❌ missing |
| Main-branch fallback | ✅ implements (over-conformant) | ❌ unaware | ❌ missing |
| `-h` / `--help` offline-safe | ⚠️ implements but may be online | ⚠️ same | ✅ explicit |
| PS `--help` token compat | ❌ missing | ❌ missing | ✅ explicit |
| Asset name fallback dance | ❌ missing | ❌ missing | ✅ explicit |
| "Strip outer folder" step | ❌ missing | ❌ missing | ✅ explicit |
| `Next steps:` footer | ❌ missing | ❌ missing | ✅ explicit |

A blind AI given **spec 27 alone** would build a *more spec-compliant*
installer than ships today, but would not match the shipped UX (no
asset-name fallback, no nested-folder flatten, no helpful footer).

A blind AI given **spec 15 alone** would build a substantially
*different product* — a multi-folder bootstrapper, not a single-pack
installer.

A blind AI given **both specs** would be confused and ask clarifying
questions — exactly what the user wants to avoid.

---

## 5. Recommended remediation (prioritised)

### P0 — Resolve the spec 15 vs spec 27 conflict

Either:

- **(a)** Mark `spec/15-distribution-and-runner/01-install-contract.md`
  as describing a *different* installer (e.g. a future
  `bootstrap.sh`), and explicitly state that
  `linters-cicd/install.{sh,ps1}` follow spec 27, **or**
- **(b)** Retire spec 15's install-contract document and consolidate
  into spec 27 with a clear "single-folder pack" appendix.

Until this is done, no AI — blind or otherwise — can implement the
"correct" installer without asking which spec wins.

### P1 — Document the implementation-specific bits in spec 27

Add normative subsections to spec 27:

1. **Asset naming & fallback order** (un-versioned alias → API tag
   resolution → versioned name).
2. **Post-extraction normalisation** (flatten single top-level
   directory if the release zip is nested).
3. **Help must be reachable without network access** (move from script
   comment to spec MUST).
4. **PowerShell `--help` token compatibility** (pattern for handling
   bash-style `--help` in PS).
5. **Behavior on missing `checksums.txt`** (declare hard error vs
   soft warning).
6. **Recommended "Next steps:" footer pattern** (or explicitly mark
   as optional UX).

### P2 — Bring shipped scripts to spec 27 conformance

Either implement V→V+N discovery and main-branch fallback in the
shipped scripts (currently failing acceptance criteria 1, 4, and 5),
or amend spec 27 §10 to make those criteria conditional on installer
family ("generic" vs "single-pack").

### P3 — Move the exit-code matrix into the spec

Replicate the README's "Verification ON vs OFF" exit-code table inside
spec 27 §8. The README can then link back instead of being the de-facto
source of truth.

---

## 6. Bottom line

> **Is the current spec set "blind-AI ready"?**
> **No, not yet.** A blind AI would produce a *plausible* installer
> from spec 27, but it would diverge from the shipped script in 6+
> behaviour-visible ways and would directly contradict spec 15.

Closing the P0 conflict and moving the items in §3.1, §3.2, and §2.3
into spec 27 would push the readiness score from "partially" to
"yes, with one clarifying assumption" (which assumption is exactly the
tradeoff between conformance to spec 27 §6 V→V+N discovery and the
current ship-it-simple shape of `linters-cicd/install.sh`).

---

*Gap analysis — 2026-04-24 — generated from a static read of specs 12, 14, 15 and the shipped `linters-cicd/install.{sh,ps1}` and `README.md`.*
