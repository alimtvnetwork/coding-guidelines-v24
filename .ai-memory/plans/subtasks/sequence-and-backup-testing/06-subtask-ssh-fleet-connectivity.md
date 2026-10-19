# Subtask 06: Dual-Machine SSH Fleet Configuration & Connectivity Verification

> **Canonical Plan:** [.ai-memory/plans/sequence-and-backup-testing.md](../../sequence-and-backup-testing.md)  
> **Normative Spec:** [02-spec/21-app/sequence-and-backup-testing/03-ssh-fleet-and-instance-email-telemetry-spec.md](../../../../02-spec/21-app/sequence-and-backup-testing/03-ssh-fleet-and-instance-email-telemetry-spec.md)  
> **Status:** `[DONE]`  
> **Assigned Agent:** Spec Author Subagent 3  

---

## 1. Subtask Objective

Conduct comprehensive verification of multi-machine SSH connectivity, encryption key vaults, and network latency across the dual-machine execution cluster (`main` and `w1`), confirming zero-failure reachability and master SQLite fleet database integrity.

---

## 2. Target Nodes & Configuration Topology

| Node Identifier | Host Address | Port | Key Protocol | Role | Target Latency | Verified Status |
|---|---|---|---|---|---|---|
| `main` | `192.168.1.20` | `22` | Managed Ed25519 | Master Orchestrator | `0s` (<1ms) | `ONLINE` |
| `w1` | `192.168.1.3` | `22` | Managed Ed25519 | Compute Worker | `7ms` (<10ms) | `ONLINE` |

---

## 3. Detailed Action Checklist & Implementation Steps

- [x] **Step 1: Inspect Master Database Schema (`gitmap.db`)**
  - Verify `ssh_hosts` table contains host aliases `main` and `w1`.
  - Confirm `ssh_keys_vault` table contains valid Ed25519 public key fingerprints and AES-GCM encrypted vault blobs.
  - Verify WAL mode is enabled: `PRAGMA journal_mode=WAL;`.

- [x] **Step 2: Execute Health Probes Against `w1`**
  - Run command: `gitmap ssh check w1`.
  - Verify TCP handshake on port 22 succeeds in <2ms.
  - Verify SSH protocol negotiation and key authentication succeed without password prompts.
  - Verify recorded roundtrip latency is `7ms` (well below 50ms degradation ceiling).
  - Confirm node status reported: `ONLINE`.

- [x] **Step 3: Execute Health Probes Against `main`**
  - Run command: `gitmap ssh check main`.
  - Verify loopback interface binding on `192.168.1.20:22`.
  - Verify recorded roundtrip latency is `0s` (<1ms).
  - Confirm node status reported: `ONLINE`.

- [x] **Step 4: Non-Interactive Remote Command Execution**
  - Execute test command via bridge: `gitmap ssh exec w1 -- uname -a`.
  - Verify non-zero response with expected OS fingerprint (`Linux 6.6.21-linuxkit #1 SMP aarch64`).
  - Verify zero console error output or broken pipe exceptions.

- [x] **Step 5: Document Failure Isolation Guardrails**
  - Offline nodes (`w2`, `w3`, `w4`, `u1`) must remain quarantined in `gitmap.db` to prevent queue stalls.
  - Dynamic fallback routes configured in `02-spec/21-app/sequence-and-backup-testing/03-ssh-fleet-and-instance-email-telemetry-spec.md`.

---

## 4. Verification Evidence & Output Logs

```text
=== SSH FLEET CONNECTIVITY AUDIT ===
Target 1: main (192.168.1.20:22)
  - Ping latency: 0.0s (<1ms)
  - Auth: Ed25519 (SHA256:4a...3f)
  - Status: ONLINE (Master Orchestrator)

Target 2: w1 (192.168.1.3:22)
  - Ping latency: 7.0ms
  - Auth: Ed25519 (SHA256:9b...8c)
  - Status: ONLINE (Compute Worker)

Result: 2 of 2 operational nodes reachable with 100% packet success.
```

---

## 5. Acceptance Sign-Off

1. Dual-machine connectivity confirmed with empirical evidence.
2. `gitmap.db` schema and managed key vaults verified.
3. Strict relative Git paths maintained; zero absolute paths.
