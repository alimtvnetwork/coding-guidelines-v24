// Wrapper that re-exports the upstream Playwright config factory and fixture
// under neutral names. This keeps the rest of the project free of references
// to the upstream package's branding.
export { createLovableConfig as createPlaywrightConfig } from "lovable-agent-playwright-config/config";
export { test, expect } from "lovable-agent-playwright-config/fixture";
