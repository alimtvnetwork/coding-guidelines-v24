/**
 * Command Palette for the Docs Viewer.
 *
 * Opens with Ctrl/Cmd+J. The top-pinned action jumps directly to
 * `spec/00-overview.md` resolved from the flat file list, so it works even
 * when the sidebar tree is stale or collapsed.
 */
import { useEffect, useCallback, useMemo } from "react";
import { BookOpen, FileText, Hash } from "lucide-react";
import {
  CommandDialog,
  CommandInput,
  CommandList,
  CommandEmpty,
  CommandGroup,
  CommandItem,
  CommandShortcut,
  CommandSeparator,
} from "@/components/ui/command";
import type { SpecNode } from "@/types/spec";
import { KeyboardKeyType } from "@/constants/enums";
import { findSpecOverviewFile, SPEC_OVERVIEW_PATH } from "./specOverviewJump";

const MAX_FILE_RESULTS = 30;
const SPEC_ROOT_PREFIX = "spec/";

interface CommandPaletteProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  allFiles: SpecNode[];
  onSelect: (node: SpecNode) => void;
}

function isPaletteShortcut(e: KeyboardEvent): boolean {
  const hasModifier = e.metaKey || e.ctrlKey;

  return hasModifier && e.key.toLowerCase() === KeyboardKeyType.J;
}

/** Global Ctrl+J / Cmd+J listener that opens the command palette. */
export function useCommandPaletteShortcut(onOpen: () => void): void {
  useEffect(() => {
    const handler = (e: KeyboardEvent) => {
      if (!isPaletteShortcut(e)) return;

      e.preventDefault();
      onOpen();
    };

    window.addEventListener("keydown", handler);

    return () => window.removeEventListener("keydown", handler);
  }, [onOpen]);
}

function getTopLevelSpecFolders(allFiles: SpecNode[]): string[] {
  const folders = new Set<string>();

  for (const file of allFiles) {
    if (!file.path.startsWith(SPEC_ROOT_PREFIX)) continue;

    const segments = file.path.split("/");

    if (segments.length < 3) continue;

    folders.add(segments[1]);
  }

  return Array.from(folders).sort();
}

function findFolderOverview(allFiles: SpecNode[], folder: string): SpecNode | null {
  const targetPath = `${SPEC_ROOT_PREFIX}${folder}/00-overview.md`;

  return allFiles.find((f) => f.path === targetPath) ?? null;
}

interface QuickJumpProps {
  overviewFile: SpecNode | null;
  onPick: (node: SpecNode) => void;
}

function QuickJumpGroup({ overviewFile, onPick }: QuickJumpProps) {
  return (
    <CommandGroup heading="Quick Jump">
      <CommandItem
        value="open-spec-overview spec/00-overview.md root index master"
        disabled={!overviewFile}
        onSelect={() => overviewFile && onPick(overviewFile)}
        className="gap-2"
      >
        <BookOpen className="h-4 w-4 text-primary" />
        <div className="flex flex-col">
          <span className="font-semibold">Open Spec Overview</span>
          <span className="text-xs text-muted-foreground">{SPEC_OVERVIEW_PATH}</span>
        </div>
        <CommandShortcut>↵</CommandShortcut>
      </CommandItem>
    </CommandGroup>
  );
}

interface FolderOverviewsProps {
  allFiles: SpecNode[];
  onPick: (node: SpecNode) => void;
}

function FolderOverviewsGroup({ allFiles, onPick }: FolderOverviewsProps) {
  const folders = useMemo(() => getTopLevelSpecFolders(allFiles), [allFiles]);

  if (folders.length === 0) return null;

  return (
    <CommandGroup heading="Spec Folders">
      {folders.map((folder) => {
        const overview = findFolderOverview(allFiles, folder);

        if (!overview) return null;

        return (
          <CommandItem
            key={folder}
            value={`folder ${folder} ${overview.path}`}
            onSelect={() => onPick(overview)}
            className="gap-2"
          >
            <Hash className="h-4 w-4 text-muted-foreground" />
            <span>{folder}</span>
            <span className="ml-auto text-xs text-muted-foreground">overview</span>
          </CommandItem>
        );
      })}
    </CommandGroup>
  );
}

interface AllFilesProps {
  allFiles: SpecNode[];
  onPick: (node: SpecNode) => void;
}

function AllFilesGroup({ allFiles, onPick }: AllFilesProps) {
  const limited = useMemo(() => allFiles.slice(0, MAX_FILE_RESULTS), [allFiles]);

  return (
    <CommandGroup heading="All Files">
      {limited.map((file) => (
        <CommandItem
          key={file.path}
          value={`file ${file.path} ${file.name}`}
          onSelect={() => onPick(file)}
          className="gap-2"
        >
          <FileText className="h-4 w-4 text-muted-foreground" />
          <span className="truncate">{file.name}</span>
          <span className="ml-auto text-xs text-muted-foreground truncate max-w-[260px]">{file.path}</span>
        </CommandItem>
      ))}
    </CommandGroup>
  );
}

export function CommandPalette({ open, onOpenChange, allFiles, onSelect }: CommandPaletteProps) {
  const overviewFile = useMemo(() => findSpecOverviewFile(allFiles), [allFiles]);

  const handlePick = useCallback(
    (node: SpecNode) => {
      onSelect(node);
      onOpenChange(false);
    },
    [onSelect, onOpenChange],
  );

  return (
    <CommandDialog open={open} onOpenChange={onOpenChange}>
      <CommandInput placeholder="Jump to a spec file… (try 'overview')" />
      <CommandList>
        <CommandEmpty>No matching files.</CommandEmpty>
        <QuickJumpGroup overviewFile={overviewFile} onPick={handlePick} />
        <CommandSeparator />
        <FolderOverviewsGroup allFiles={allFiles} onPick={handlePick} />
        <CommandSeparator />
        <AllFilesGroup allFiles={allFiles} onPick={handlePick} />
      </CommandList>
    </CommandDialog>
  );
}
