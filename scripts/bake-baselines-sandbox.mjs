#!/usr/bin/env node
/**
 * scripts/bake-baselines-sandbox.mjs
 *
 * Wrapper for `bunx playwright test slides-app/tests/visual.spec.ts
 * --update-snapshots` that resolves `LD_LIBRARY_PATH` from `/nix/store`
 * so Chromium's shared libraries (glib, libX11, GTK adjacents) load
 * inside the Lovable sandbox. On developer machines and CI runners
 * the system loader already exposes those libraries, so this wrapper
 * refuses to run unless `/nix/store` exists AND `libglib-2.0.so.0`
 * cannot already be resolved by the default loader path. That guard
 * prevents the wrapper from silently masking a real "chromium
 * missing" failure off-sandbox.
 *
 * Recipe reference: slides-app/docs/visual-baseline-sandbox-recipe.md
 * Root cause it addresses: v5.117 shipped a strict guard while zero
 * baselines existed because `slides:bake-baselines` errored out on
 * "libglib-2.0.so.0: cannot open shared object file" inside the
 * sandbox and nobody had captured the fix.
 */
import { existsSync, readdirSync, statSync } from "node:fs";
import { execFileSync, spawnSync } from "node:child_process";
import { join } from "node:path";

const NIX_STORE = "/nix/store";

function isSandbox() {
  if (!existsSync(NIX_STORE)) return false;
  // Off-sandbox linux boxes may have /nix/store but also have glib
  // on the loader path — in that case defer to the plain bake script.
  const probe = spawnSync("ldconfig", ["-p"], { encoding: "utf8" });
  if (probe.status !== 0) return true; // no ldconfig, assume sandbox
  return !/libglib-2\.0\.so\.0/.test(probe.stdout);
}

function findFirst(soname) {
  // Depth-limited scan of /nix/store/*/lib for a single soname.
  // Bounded fan-out: only reads first-level lib/ dirs, no recursion.
  const entries = readdirSync(NIX_STORE);
  for (const e of entries) {
    const libDir = join(NIX_STORE, e, "lib");
    try {
      if (!statSync(libDir).isDirectory()) continue;
    } catch { continue; }
    const candidate = join(libDir, soname);
    if (existsSync(candidate)) return libDir;
  }
  return null;
}

function collectLibDirs() {
  const entries = readdirSync(NIX_STORE);
  const dirs = [];
  for (const e of entries) {
    const libDir = join(NIX_STORE, e, "lib");
    try {
      if (statSync(libDir).isDirectory()) dirs.push(libDir);
    } catch { /* skip */ }
  }
  return dirs;
}

function main() {
  if (!isSandbox()) {
    console.error("[bake:sandbox] refusing to run: this looks like a non-sandbox host.");
    console.error("[bake:sandbox] Use `npm run slides:bake-baselines` instead.");
    console.error("[bake:sandbox] (This guard prevents silently hiding real chromium failures.)");
    process.exit(2);
  }

  const glibDir = findFirst("libglib-2.0.so.0");
  const x11Dir = findFirst("libX11.so.6");
  if (!glibDir || !x11Dir) {
    console.error(`[bake:sandbox] ERROR: could not resolve required libs in ${NIX_STORE}`);
    console.error(`[bake:sandbox]   libglib-2.0.so.0 -> ${glibDir ?? "NOT FOUND"}`);
    console.error(`[bake:sandbox]   libX11.so.6      -> ${x11Dir ?? "NOT FOUND"}`);
    console.error("[bake:sandbox] See slides-app/docs/visual-baseline-sandbox-recipe.md");
    process.exit(3);
  }

  const nixLibs = collectLibDirs().join(":");
  const ldPath = [glibDir, x11Dir, nixLibs, process.env.LD_LIBRARY_PATH || ""].filter(Boolean).join(":");

  console.log(`[bake:sandbox] glib   -> ${glibDir}`);
  console.log(`[bake:sandbox] libX11 -> ${x11Dir}`);
  console.log(`[bake:sandbox] /nix/store/*/lib entries: ${collectLibDirs().length}`);
  console.log(`[bake:sandbox] launching Playwright (--update-snapshots, workers=1)`);

  const result = spawnSync(
    "bunx",
    ["playwright", "test", "tests/visual.spec.ts", "--update-snapshots", "--workers=1", "--reporter=list"],
    { cwd: "slides-app", stdio: "inherit", env: { ...process.env, LD_LIBRARY_PATH: ldPath } },
  );

  if (result.status !== 0) {
    console.error(`[bake:sandbox] Playwright exited with ${result.status}`);
    process.exit(result.status ?? 1);
  }

  console.log("[bake:sandbox] Bake complete. Verify with:");
  console.log("  node scripts/validate-visual-baselines.mjs --strict --list");
}

main();
