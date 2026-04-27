import { useState, useMemo } from "react";
import type { SpecNode } from "@/types/spec";
import specData from "@/data/specTree.json";
import { SpecEntryType } from "@/constants/enums";
import { treeDiagDebug, treeDiagWarn, TreeDiagCategory } from "@/lib/treeDiagnostics";

const flattenFiles = (nodes: SpecNode[]): SpecNode[] => {
  const result: SpecNode[] = [];
  for (const node of nodes) {

    if (node.type === SpecEntryType.File) result.push(node);

    if (node.children) result.push(...flattenFiles(node.children));
  }

  return result;
};

function summarizeTree(tree: SpecNode[]) {
  const topLevel = tree.map((n) => ({ name: n.name, type: n.type, childCount: n.children?.length ?? 0 }));
  const folderCount = tree.filter((n) => n.type === SpecEntryType.Folder).length;
  const fileCount = tree.filter((n) => n.type === SpecEntryType.File).length;

  return { topLevel, folderCount, fileCount, totalRoots: tree.length };
}

export function useSpecData() {
  const tree = specData.specTree as SpecNode[];
  const allFiles = useMemo(() => flattenFiles(tree), [tree]);

  const summary = useMemo(() => summarizeTree(tree), [tree]);

  if (summary.totalRoots === 0) {
    treeDiagWarn(TreeDiagCategory.UseSpecData, "specTree.json produced ZERO root nodes — spec sidebar will be empty.", summary);
  } else {
    treeDiagDebug(TreeDiagCategory.UseSpecData, "Loaded specTree.json", { ...summary, allFilesCount: allFiles.length });
  }

  return { tree, allFiles };
}

function isQueryBlank(query: string): boolean {
  return query.trim().length === 0;
}

export function useSpecSearch(allFiles: SpecNode[], query: string) {
  return useMemo(() => {

    if (isQueryBlank(query)) {
      return [];
    }

    const q = query.toLowerCase();

    return allFiles
      .filter(
        (f) =>
          f.name.toLowerCase().includes(q) ||
          (f.content && f.content.toLowerCase().includes(q))
      )
      .slice(0, 20);
  }, [allFiles, query]);
}
