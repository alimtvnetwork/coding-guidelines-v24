#!/usr/bin/env node
/**
 * scripts/branch-protection-diff.mjs
 *
 * Read-only diff of LIVE GitHub branch-protection required_status_checks
 * for `main` against `.github/branch-protection.expected.json`.
 *
 * Root cause it closes: the `visual` + `smoke` promotion has sat in the
 * backlog for 4 releases because operators had no cheap way to verify
 * "is live state already in sync with expected.json?" without opening
 * the GitHub UI. Post-promotion verification in
 * `.lovable/procedures/branch-protection.md` step 4 also required
 * eyeballing a raw JSON array. This makes both a one-command check.
 *
 * Usage:
 *   npm run branch-protection:diff                       # auto-detect repo
 *   npm run branch-protection:diff -- --repo owner/name  # override
 *
 * Exit codes:
 *   0 = live state matches expected.required exactly
 *   1 = drift (missing or extra contexts on live); prints diff
 *   2 = `gh` CLI missing or not authenticated
 *   3 = repo could not be auto-detected and --repo not provided
 *
 * Never mutates anything. Safe to run without admin.
 */
import { readFileSync, existsSync, mkdirSync, writeFileSync } from "node:fs";
import { spawnSync } from "node:child_process";
import { join } from "node:path";

const EXPECTED_FILE = ".github/branch-protection.expected.json";

function loadExpected() {
  if (!existsSync(EXPECTED_FILE)) {
    console.error(`Missing ${EXPECTED_FILE}`);
    process.exit(1);
  }
  return JSON.parse(readFileSync(EXPECTED_FILE, "utf8"));
}

function parseArgs(argv) {
  const args = { repo: null, reportDir: null };
  for (let i = 0; i < argv.length; i += 1) {
    if (argv[i] === "--repo" && argv[i + 1]) {
      args.repo = argv[i + 1];
      i += 1;
    } else if (argv[i] === "--report" && argv[i + 1]) {
      args.reportDir = argv[i + 1];
      i += 1;
    }
  }
  return args;
}


function detectRepo() {
  const r = spawnSync("gh", ["repo", "view", "--json", "nameWithOwner", "-q", ".nameWithOwner"], {
    encoding: "utf8",
  });
  if (r.status !== 0) return null;
  return r.stdout.trim() || null;
}

function ghAvailable() {
  const r = spawnSync("gh", ["auth", "status"], { encoding: "utf8" });
  return r.status === 0;
}

function fetchLiveContexts(repo) {
  const r = spawnSync(
    "gh",
    [
      "api",
      `repos/${repo}/branches/main/protection/required_status_checks`,
      "--jq",
      ".contexts",
    ],
    { encoding: "utf8" },
  );
  if (r.status !== 0) {
    console.error("gh api failed:");
    console.error(r.stderr);
    process.exit(1);
  }
  return JSON.parse(r.stdout);
}

function diff(live, expected) {
  const liveSet = new Set(live);
  const expSet = new Set(expected);
  const missingOnLive = expected.filter((n) => !liveSet.has(n));
  const extraOnLive = live.filter((n) => !expSet.has(n));
  return { missingOnLive, extraOnLive };
}

function buildReport({ repo, expected, live, missingOnLive, extraOnLive }) {

  const unchanged = expected.filter((n) => live.includes(n));
  const status = !missingOnLive.length && !extraOnLive.length ? "in-sync" : "drift";
  const json = {
    generatedAt: new Date().toISOString(),
    repo,
    status,
    counts: {
      expected: expected.length,
      live: live.length,
      missingOnLive: missingOnLive.length,
      extraOnLive: extraOnLive.length,
      unchanged: unchanged.length,
    },
    expected,
    live,
    missingOnLive,
    extraOnLive,
    unchanged,
  };
  const bullet = (arr) => (arr.length ? arr.map((n) => `- \`${n}\``).join("\n") : "_none_");
  const md = [
    `# Branch Protection Drift Report`,
    ``,
    `- Repo: \`${repo}\``,
    `- Generated: ${json.generatedAt}`,
    `- Status: **${status.toUpperCase()}**`,
    ``,
    `## Counts`,
    ``,
    `| Metric | Count |`,
    `| --- | ---: |`,
    `| Expected required | ${expected.length} |`,
    `| Live required | ${live.length} |`,
    `| Missing on live (drift) | ${missingOnLive.length} |`,
    `| Extra on live (drift) | ${extraOnLive.length} |`,
    `| Unchanged (in sync) | ${unchanged.length} |`,
    ``,
    `## Missing on live (expected but not required)`,
    ``,
    bullet(missingOnLive),
    ``,
    `## Extra on live (required but not expected)`,
    ``,
    bullet(extraOnLive),
    ``,
    `## Unchanged`,
    ``,
    bullet(unchanged),
    ``,
  ].join("\n");
  return { json, md };
}

function writeReport(dir, report) {
  mkdirSync(dir, { recursive: true });
  writeFileSync(join(dir, "drift.json"), JSON.stringify(report.json, null, 2) + "\n");
  writeFileSync(join(dir, "drift.md"), report.md);
}

function main() {
  if (!ghAvailable()) {
    console.error("gh CLI not found or not authenticated. Run `gh auth login`.");
    process.exit(2);
  }
  const args = parseArgs(process.argv.slice(2));
  const repo = args.repo || detectRepo();
  if (!repo) {
    console.error("Repo not detected. Pass --repo owner/name.");
    process.exit(3);
  }
  const expected = loadExpected().required || [];
  const live = fetchLiveContexts(repo);
  const { missingOnLive, extraOnLive } = diff(live, expected);
  const report = buildReport({ repo, expected, live, missingOnLive, extraOnLive });

  console.log(`Repo: ${repo}`);
  console.log(`Expected required (${expected.length}): ${expected.join(", ") || "(none)"}`);
  console.log(`Live required     (${live.length}): ${live.join(", ") || "(none)"}`);

  const inSync = !missingOnLive.length && !extraOnLive.length;
  if (inSync) {
    console.log("");
    console.log("OK, live branch protection matches expected.json exactly.");
  } else {
    console.log("");
    console.log("DRIFT:");
    for (const n of missingOnLive) console.log(`  - missing on live (expected but not required): ${n}`);
    for (const n of extraOnLive) console.log(`  + extra on live  (required but not in expected):   ${n}`);
  }

  if (args.reportDir) {
    writeReport(args.reportDir, report);
    console.log("");
    console.log(`Report written: ${join(args.reportDir, "drift.json")}, ${join(args.reportDir, "drift.md")}`);
  }

  if (!inSync) process.exit(1);
}

export { diff, buildReport };

import { fileURLToPath } from "node:url";
if (process.argv[1] === fileURLToPath(import.meta.url)) main();


export { diff, buildReport };

