/**
 * Floating diagnostics panel for the in-app spec sidebar / file tree.
 *
 * Visible only when diagnostics are enabled (URL `?diag=tree` or
 * localStorage `debug.tree=1`). Designed to be minimally intrusive — bottom
 * right, semi-transparent, collapsible.
 */
import { useEffect, useState, useCallback } from "react";
import { Bug, Copy, X, RefreshCw } from "lucide-react";
import { isTreeDiagEnabled, getTreeDiagEntries, clearTreeDiagEntries, type TreeDiagEntry } from "@/lib/treeDiagnostics";

const POLL_MS = 500;

interface DiagnosticsSummary {
  total: number;
  warnings: number;
  lastTs: string | null;
  lastMessage: string | null;
}

function summarize(entries: readonly TreeDiagEntry[]): DiagnosticsSummary {
  const warnings = entries.filter((e) => e.level === "warn").length;
  const last = entries[entries.length - 1] ?? null;

  return {
    total: entries.length,
    warnings,
    lastTs: last?.ts ?? null,
    lastMessage: last ? `[${last.category}] ${last.message}` : null,
  };
}

function useDiagnosticsSnapshot(): DiagnosticsSummary {
  const [summary, setSummary] = useState<DiagnosticsSummary>(() => summarize(getTreeDiagEntries()));

  useEffect(() => {
    const interval = window.setInterval(() => {
      setSummary(summarize(getTreeDiagEntries()));
    }, POLL_MS);

    return () => window.clearInterval(interval);
  }, []);

  return summary;
}

function copyToClipboard(text: string): void {
  if (!navigator.clipboard) return;

  void navigator.clipboard.writeText(text);
}

function PanelHeader({ onClose }: { onClose: () => void }) {
  return (
    <div className="flex items-center justify-between px-3 py-2 border-b border-border bg-muted/40">
      <div className="inline-flex items-center gap-1.5 text-xs font-semibold text-foreground">
        <Bug className="h-3.5 w-3.5 text-primary" />
        Tree Diagnostics
      </div>
      <button
        type="button"
        onClick={onClose}
        title="Hide panel (logging stays on)"
        className="text-muted-foreground hover:text-foreground"
      >
        <X className="h-3.5 w-3.5" />
      </button>
    </div>
  );
}

function PanelBody({ summary }: { summary: DiagnosticsSummary }) {
  return (
    <div className="px-3 py-2 text-xs space-y-1">
      <div className="flex items-center gap-3">
        <span className="text-muted-foreground">Entries:</span>
        <span className="font-mono text-foreground">{summary.total}</span>
        <span className="text-muted-foreground">Warnings:</span>
        <span className={`font-mono ${summary.warnings > 0 ? "text-destructive" : "text-foreground"}`}>{summary.warnings}</span>
      </div>
      {summary.lastMessage && (
        <div className="text-muted-foreground truncate" title={summary.lastMessage}>
          Last: <span className="text-foreground">{summary.lastMessage}</span>
        </div>
      )}
    </div>
  );
}

function PanelActions() {
  const handleCopy = useCallback(() => {
    const text = JSON.stringify(getTreeDiagEntries(), null, 2);
    copyToClipboard(text);
  }, []);

  const handleClear = useCallback(() => {
    clearTreeDiagEntries();
  }, []);

  return (
    <div className="flex items-center gap-1 border-t border-border bg-muted/20 px-2 py-1.5">
      <button
        type="button"
        onClick={handleCopy}
        title="Copy all tree-diagnostic entries to clipboard"
        className="inline-flex items-center gap-1 rounded-md px-2 py-1 text-xs text-foreground hover:bg-muted"
      >
        <Copy className="h-3 w-3" /> Copy logs
      </button>
      <button
        type="button"
        onClick={handleClear}
        title="Clear ring buffer"
        className="inline-flex items-center gap-1 rounded-md px-2 py-1 text-xs text-foreground hover:bg-muted"
      >
        <RefreshCw className="h-3 w-3" /> Clear
      </button>
    </div>
  );
}

export function TreeDiagnosticsPanel() {
  const [hidden, setHidden] = useState(false);
  const summary = useDiagnosticsSnapshot();

  if (!isTreeDiagEnabled() || hidden) return null;

  return (
    <div
      role="status"
      aria-live="polite"
      className="fixed bottom-4 right-4 z-50 w-72 rounded-lg border border-border bg-card/95 backdrop-blur shadow-lg overflow-hidden"
    >
      <PanelHeader onClose={() => setHidden(true)} />
      <PanelBody summary={summary} />
      <PanelActions />
    </div>
  );
}
