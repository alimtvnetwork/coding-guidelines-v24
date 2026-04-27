/**
 * Resolves the canonical spec overview file (`spec/00-overview.md`) directly
 * from the flat `allFiles` list. This intentionally bypasses the sidebar tree
 * so the user can jump to the overview even when the tree is stale, collapsed,
 * or visually out-of-sync.
 */
import type { SpecNode } from "@/types/spec";
import { SpecEntryType } from "@/constants/enums";

export const SPEC_OVERVIEW_PATH = "spec/00-overview.md";
const OVERVIEW_FILENAME = "00-overview.md";

function isFile(node: SpecNode): boolean {
  return node.type === SpecEntryType.File;
}

function matchesExactPath(node: SpecNode): boolean {
  return isFile(node) && node.path === SPEC_OVERVIEW_PATH;
}

function matchesOverviewByName(node: SpecNode): boolean {
  if (!isFile(node)) return false;

  const path = node.path;

  return path.startsWith("spec/") && path.endsWith(OVERVIEW_FILENAME) && path.split("/").length === 2;
}

/**
 * Find the spec overview file regardless of tree expansion state.
 * Falls back to a name-based match if the exact path is not present.
 */
export function findSpecOverviewFile(allFiles: SpecNode[]): SpecNode | null {
  const exact = allFiles.find(matchesExactPath);

  if (exact) return exact;

  const byName = allFiles.find(matchesOverviewByName);

  return byName ?? null;
}
