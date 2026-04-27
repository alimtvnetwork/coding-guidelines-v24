/**
 * File-tree diagnostics.
 *
 * Structured logging helper for understanding why the in-app spec sidebar
 * (DocsSidebar / SpecTreeNav / useSpecData) is or isn't rendering what the
 * user expects. Lightweight, zero-cost when disabled.
 *
 * Enable via either:
 *   - URL param: ?diag=tree   (persists for the session)
 *   - localStorage["debug.tree"] = "1"
 *
 * Disable by removing the localStorage key or visiting any page without
 * ?diag=tree (the URL param re-arms it whenever present).
 *
 * The ring buffer is exposed at `window.__treeDiagnostics` for ad-hoc
 * inspection from DevTools.
 */

const STORAGE_KEY = "debug.tree";
const URL_FLAG = "diag";
const URL_FLAG_VALUE = "tree";
const RING_BUFFER_SIZE = 200;
const GLOBAL_KEY = "__treeDiagnostics" as const;

export const TreeDiagCategory = {
  UseSpecData: "useSpecData",
  DocsSidebar: "DocsSidebar",
  SpecTreeNav: "SpecTreeNav",
  Refresh: "Refresh",
} as const;

export type TreeDiagCategoryType = (typeof TreeDiagCategory)[keyof typeof TreeDiagCategory];

export const TreeDiagLevel = {
  Debug: "debug",
  Info: "info",
  Warn: "warn",
} as const;

export type TreeDiagLevelType = (typeof TreeDiagLevel)[keyof typeof TreeDiagLevel];

export interface TreeDiagEntry {
  ts: string;
  level: TreeDiagLevelType;
  category: TreeDiagCategoryType;
  message: string;
  data?: unknown;
}

interface TreeDiagWindow {
  enabled: boolean;
  entries: TreeDiagEntry[];
  copyLogs: () => string;
  clear: () => void;
}

function readUrlFlag(): boolean {
  if (typeof window === "undefined") return false;

  const params = new URLSearchParams(window.location.search);

  return params.get(URL_FLAG) === URL_FLAG_VALUE;
}

function readStorageFlag(): boolean {
  if (typeof window === "undefined") return false;

  try {
    return window.localStorage.getItem(STORAGE_KEY) === "1";
  } catch {
    return false;
  }
}

function persistEnabled(): void {
  if (typeof window === "undefined") return;

  try {
    window.localStorage.setItem(STORAGE_KEY, "1");
  } catch {
    // ignore
  }
}

function computeEnabled(): boolean {
  const fromUrl = readUrlFlag();

  if (fromUrl) {
    persistEnabled();

    return true;
  }

  return readStorageFlag();
}

const ENABLED = computeEnabled();
const buffer: TreeDiagEntry[] = [];

function pushEntry(entry: TreeDiagEntry): void {
  buffer.push(entry);

  if (buffer.length > RING_BUFFER_SIZE) buffer.shift();
}

function emitConsole(entry: TreeDiagEntry): void {
  const prefix = `[tree:${entry.category}]`;

  if (entry.level === TreeDiagLevel.Warn) {
    console.warn(prefix, entry.message, entry.data ?? "");

    return;
  }

  if (entry.level === TreeDiagLevel.Info) {
    console.info(prefix, entry.message, entry.data ?? "");

    return;
  }

  console.debug(prefix, entry.message, entry.data ?? "");
}

function record(level: TreeDiagLevelType, category: TreeDiagCategoryType, message: string, data?: unknown): void {
  if (!ENABLED) return;

  const entry: TreeDiagEntry = {
    ts: new Date().toISOString(),
    level,
    category,
    message,
    data,
  };

  pushEntry(entry);
  emitConsole(entry);
}

export function treeDiagDebug(category: TreeDiagCategoryType, message: string, data?: unknown): void {
  record(TreeDiagLevel.Debug, category, message, data);
}

export function treeDiagInfo(category: TreeDiagCategoryType, message: string, data?: unknown): void {
  record(TreeDiagLevel.Info, category, message, data);
}

export function treeDiagWarn(category: TreeDiagCategoryType, message: string, data?: unknown): void {
  record(TreeDiagLevel.Warn, category, message, data);
}

export function isTreeDiagEnabled(): boolean {
  return ENABLED;
}

export function getTreeDiagEntries(): readonly TreeDiagEntry[] {
  return buffer;
}

export function clearTreeDiagEntries(): void {
  buffer.length = 0;
}

function exportLogsAsString(): string {
  return JSON.stringify(buffer, null, 2);
}

function installGlobalHandle(): void {
  if (typeof window === "undefined") return;

  const handle: TreeDiagWindow = {
    enabled: ENABLED,
    entries: buffer,
    copyLogs: exportLogsAsString,
    clear: clearTreeDiagEntries,
  };

  (window as unknown as Record<string, TreeDiagWindow>)[GLOBAL_KEY] = handle;
}

installGlobalHandle();
