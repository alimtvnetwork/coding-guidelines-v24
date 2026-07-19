# Visual baseline bake recipe (Lovable sandbox)

`npm run slides:bake-baselines` invokes Playwright, which needs a
system Chromium plus its shared libraries. In the Lovable sandbox those
libraries are only reachable from `/nix/store/*/lib`; nothing on the
default loader path exposes them, so an unmodified `bunx playwright
test --update-snapshots` dies with:

```
error while loading shared libraries: libglib-2.0.so.0:
cannot open shared object file: No such file or directory
```

This is why the v5.117 strict guard shipped with zero baselines: the
bake failed silently and everyone assumed the sandbox couldn't run
Chromium at all. It can, given the right `LD_LIBRARY_PATH`.

## Recipe (Lovable sandbox only)

Reason for existence: local developer machines and CI runners already
have these libs on the loader path, so this ONLY applies when baking
inside the Lovable sandbox.

1. Resolve the two nix paths that hold `libglib-2.0.so.0` and
   `libX11.so.6` (paths are content-addressed, so the store hash
   changes per rebuild — always resolve, never hardcode):

   ```bash
   GLIB=$(dirname "$(readlink -f "$(nix-store -q --references \
     $(command -v glib-compile-schemas 2>/dev/null || true) 2>/dev/null | head -1)")" 2>/dev/null || true)
   # Fallback: brute-scan the store for the two sonames we need.
   GLIB=$(dirname "$(find /nix/store -maxdepth 4 -name libglib-2.0.so.0 2>/dev/null | head -1)")
   X11=$(dirname "$(find /nix/store -maxdepth 4 -name libX11.so.6 2>/dev/null | head -1)")
   ```

2. Widen `LD_LIBRARY_PATH` to include every `/nix/store/*/lib`
   directory. Chromium pulls in ~40 transitive shared libraries (GTK,
   ATK, NSS, dbus, cups, ...); listing them by hand is fragile and
   breaks on every nix garbage collection.

   ```bash
   NIX_LIBS=$(printf '%s\n' /nix/store/*/lib | paste -sd:)
   export LD_LIBRARY_PATH="$GLIB:$X11:$NIX_LIBS:${LD_LIBRARY_PATH:-}"
   ```

3. Bake all baselines. Runtime is ~3.4 min for 70 slides on one
   worker; scales linearly.

   ```bash
   cd slides-app
   bunx playwright test tests/visual.spec.ts --update-snapshots --workers=1
   ```

4. Verify from repo root:

   ```bash
   node scripts/validate-visual-baselines.mjs --strict --list
   # Expect: deck=<N> baselines=<N> missing=0 strict=true
   ```

## Why not automate this in `slides:bake-baselines`?

Because the fallback (`find /nix/store ...`) is sandbox-specific and
would silently mask real "chromium missing" errors on developer
machines and CI. A future `slides:bake-baselines:sandbox` script can
wrap the recipe above without polluting the default entry point.

## Related

- `scripts/validate-visual-baselines.mjs` — the strict guard.
- `.husky/pre-push` — blocks pushes on missing baselines.
- `.github/workflows/slides-visual.yml` — CI-side gate.
