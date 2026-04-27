/**
 * GitHub Sync Banner.
 *
 * Shows a clear visual confirmation that the local docs view is in sync with
 * the connected GitHub repository, including:
 *  - The git branch + short SHA the build was generated from (baked at build time).
 *  - The "last updated" date from `version.json`.
 *  - The local "last loaded" timestamp for the current session.
 *
 * Lovable's GitHub integration is bidirectional and automatic, so any commit
 * that reaches the connected repo is reflected in the next build. This banner
 * surfaces that state instead of leaving the user guessing.
 */
import { useEffect, useMemo, useState, useCallback } from "react";
import { CheckCircle2, GitBranch, X, RefreshCw } from "lucide-react";
import versionData from "../../../version.json";
import { isClosed, isOpen } from "@/constants/boolFlags";

const DISMISS_STORAGE_KEY = "lovable.github-sync-banner.dismissed-sha";
const RELOAD_DELAY_MS = 100;

interface VersionGit {
  sha?: string;
  shortSha?: string;
  branch?: string;
}

interface VersionFile {
  version?: string;
  updated?: string;
  git?: VersionGit;
}

const version = versionData as VersionFile;

function safeReadDismissedSha(): string | null {
  try {
    return window.localStorage.getItem(DISMISS_STORAGE_KEY);
  } catch {
    return null;
  }
}

function safeWriteDismissedSha(sha: string): void {
  try {
    window.localStorage.setItem(DISMISS_STORAGE_KEY, sha);
  } catch {
    // Ignore storage failures (private mode, etc.)
  }
}

function formatLoadedAt(date: Date): string {
  return date.toLocaleString(undefined, {
    year: "numeric",
    month: "short",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

function useLoadedAt(): string {
  const date = useMemo(() => new Date(), []);

  return useMemo(() => formatLoadedAt(date), [date]);
}

function useDismissalState(currentSha: string | undefined): { isDismissed: boolean; dismiss: () => void } {
  const [isDismissed, setIsDismissed] = useState(isClosed);

  useEffect(() => {
    if (!currentSha) return;

    const stored = safeReadDismissedSha();
    setIsDismissed(stored === currentSha ? isOpen : isClosed);
  }, [currentSha]);

  const dismiss = useCallback(() => {
    if (currentSha) safeWriteDismissedSha(currentSha);

    setIsDismissed(isOpen);
  }, [currentSha]);

  return { isDismissed, dismiss };
}

interface SyncDetailsProps {
  branch: string;
  shortSha: string;
  updated: string;
  loadedAt: string;
}

function SyncDetails({ branch, shortSha, updated, loadedAt }: SyncDetailsProps) {
  return (
    <div className="flex flex-wrap items-center gap-x-4 gap-y-1 text-xs text-muted-foreground min-w-0">
      <span className="inline-flex items-center gap-1.5">
        <GitBranch className="h-3.5 w-3.5" />
        <span className="font-mono text-foreground truncate max-w-[260px]">{branch}</span>
        <span className="text-border">·</span>
        <span className="font-mono text-foreground">{shortSha}</span>
      </span>
      <span className="inline-flex items-center gap-1.5">
        <span className="text-muted-foreground">Repo updated:</span>
        <span className="text-foreground font-medium">{updated}</span>
      </span>
      <span className="inline-flex items-center gap-1.5">
        <span className="text-muted-foreground">Loaded:</span>
        <span className="text-foreground font-medium">{loadedAt}</span>
      </span>
    </div>
  );
}

function BannerActions({ onReload, onDismiss }: { onReload: () => void; onDismiss: () => void }) {
  return (
    <div className="flex items-center gap-1 shrink-0">
      <button
        type="button"
        onClick={onReload}
        title="Reload to pick up the latest synced changes"
        className="inline-flex items-center gap-1.5 rounded-md border border-border bg-background/60 px-2 py-1 text-xs text-foreground hover:bg-muted transition-colors"
      >
        <RefreshCw className="h-3.5 w-3.5" />
        <span className="hidden sm:inline">Reload</span>
      </button>
      <button
        type="button"
        onClick={onDismiss}
        title="Dismiss until the next sync"
        className="inline-flex items-center justify-center h-7 w-7 rounded-md text-muted-foreground hover:text-foreground hover:bg-muted transition-colors"
      >
        <X className="h-3.5 w-3.5" />
      </button>
    </div>
  );
}

export function GithubSyncBanner() {
  const branch = version.git?.branch ?? "unknown";
  const shortSha = version.git?.shortSha ?? version.git?.sha?.slice(0, 7) ?? "unknown";
  const updated = version.updated ?? "unknown";
  const loadedAt = useLoadedAt();
  const { isDismissed, dismiss } = useDismissalState(version.git?.sha);

  const handleReload = useCallback(() => {
    setTimeout(() => window.location.reload(), RELOAD_DELAY_MS);
  }, []);

  if (isDismissed) return null;

  return (
    <div
      role="status"
      aria-live="polite"
      className="flex items-center gap-3 border-b border-border bg-success/10 px-4 py-2 text-sm shrink-0"
    >
      <CheckCircle2 className="h-4 w-4 text-success shrink-0" aria-hidden="true" />
      <div className="flex flex-col sm:flex-row sm:items-center gap-x-4 gap-y-0.5 min-w-0 flex-1">
        <span className="font-medium text-foreground shrink-0">Synced with GitHub</span>
        <SyncDetails branch={branch} shortSha={shortSha} updated={updated} loadedAt={loadedAt} />
      </div>
      <BannerActions onReload={handleReload} onDismiss={dismiss} />
    </div>
  );
}
