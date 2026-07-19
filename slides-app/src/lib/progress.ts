import { useSyncExternalStore } from "react";

/**
 * Per-block review progress. Learners tick Symptom / Rule / Action on every
 * guideline slide; state persists in localStorage and drives the TOC bar.
 */

export type ProgressBlock = "symptom" | "rule" | "action";
export const PROGRESS_BLOCKS: readonly ProgressBlock[] = ["symptom", "rule", "action"] as const;

const STORAGE_KEY = "slides-sra-progress-v1";
const EVENT_NAME = "slides-sra-progress-change";

type ProgressMap = Record<string, Partial<Record<ProgressBlock, boolean>>>;

function isBrowser(): boolean {
  return typeof window !== "undefined";
}

function readAll(): ProgressMap {
  if (!isBrowser()) return {};
  const raw = window.localStorage.getItem(STORAGE_KEY);
  if (!raw) return {};
  try {
    const parsed = JSON.parse(raw);
    if (parsed && typeof parsed === "object") return parsed as ProgressMap;
  } catch {
    // Corrupt JSON: wipe so future writes start clean.
    window.localStorage.removeItem(STORAGE_KEY);
  }
  return {};
}

function writeAll(next: ProgressMap): void {
  if (!isBrowser()) return;
  window.localStorage.setItem(STORAGE_KEY, JSON.stringify(next));
  window.dispatchEvent(new Event(EVENT_NAME));
}

export function getBlock(slideId: string, block: ProgressBlock): boolean {
  const all = readAll();
  return Boolean(all[slideId]?.[block]);
}

export function setBlock(slideId: string, block: ProgressBlock, value: boolean): void {
  const all = readAll();
  const current = all[slideId] ?? {};
  const nextEntry = { ...current, [block]: value };
  writeAll({ ...all, [slideId]: nextEntry });
}

export function toggleBlock(slideId: string, block: ProgressBlock): void {
  setBlock(slideId, block, !getBlock(slideId, block));
}

export function resetAllProgress(): void {
  writeAll({});
}

function subscribe(callback: () => void): () => void {
  if (!isBrowser()) return () => {};
  const handler = () => callback();
  window.addEventListener(EVENT_NAME, handler);
  window.addEventListener("storage", handler);
  return () => {
    window.removeEventListener(EVENT_NAME, handler);
    window.removeEventListener("storage", handler);
  };
}

function getSnapshot(): string {
  if (!isBrowser()) return "{}";
  return window.localStorage.getItem(STORAGE_KEY) ?? "{}";
}

function getServerSnapshot(): string {
  return "{}";
}

export function useProgressSnapshot(): ProgressMap {
  const raw = useSyncExternalStore(subscribe, getSnapshot, getServerSnapshot);
  try {
    return JSON.parse(raw) as ProgressMap;
  } catch {
    return {};
  }
}

export function useSlideProgress(slideId: string): {
  reviewed: Record<ProgressBlock, boolean>;
  completedCount: number;
  isComplete: boolean;
} {
  const all = useProgressSnapshot();
  const entry = all[slideId] ?? {};
  const reviewed: Record<ProgressBlock, boolean> = {
    symptom: Boolean(entry.symptom),
    rule: Boolean(entry.rule),
    action: Boolean(entry.action),
  };
  const completedCount = PROGRESS_BLOCKS.filter((b) => reviewed[b]).length;
  return { reviewed, completedCount, isComplete: completedCount === PROGRESS_BLOCKS.length };
}
