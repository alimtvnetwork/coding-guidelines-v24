# Subtask 03: Fileutil Readme Documentation & AI Skill Guide

## Objective
Update `04-code/golang/pkg/fileutil/readme.md` to document the renamed operation files, `FilePathOps`, the creator patterns, and format a comprehensive AI Skill guide for agentic codebases.

## Target Files
- `04-code/golang/pkg/fileutil/readme.md`

## Implementation Details
1. Document the 1:1 struct-to-filename mapping:
   - `appendOps` in `append_ops.go`
   - `openOps` in `open_ops.go`
   - `createOps` in `create_ops.go`
   - `readOps` in `read_ops.go`
   - `writeOps` in `write_ops.go`
   - `FilePathOps` in `file_path_ops.go`
2. Document `FilePathOps` bound usage:
   - Creating via `File.Target(path)` or `File.At(workDir, relPath)`.
   - Immutable chaining (`WithWorkDir`, `WithRelPath`, `Join`).
   - Pre-flight guarantees (`EnsureParentDir`, `EnsureFile`).
   - Bound I/O (`ReadString()`, `WriteString()`, `AppendLines()`).
3. Add a dedicated section titled:
   `## AI Agent Skill: fileutil Operational Playbook`
   Outlining decision rules, patterns, anti-patterns, and step-by-step guidelines so that an AI reading the README can adopt it as an autonomous skill.
