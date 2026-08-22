import { useSyncExternalStore } from "react";

/**
 * Per-block review progress. Learners tick Symptom / Rule / Action on every
 * guideline slide; state persists in localStorage and drives the TOC bar.
 */

export enum ProgressBlockType {
  Symptom = "symptom",
  Rule = "rule",
  Action = "action"
}

export const PROGRESS_BLOCKS: readonly ProgressBlockType[] = [ProgressBlockType.Symptom, ProgressBlockType.Rule, ProgressBlockType.Action] as const;

const STORAGE_KEY = "slides-sra-progress-v1";
const EVENT_NAME = "slides-sra-progress-change";

type ProgressMap = Record<string, Partial<Record<ProgressBlockType, boolean>>>;

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

export function getBlock(slideId: string, block: ProgressBlockType): boolean {
  const all = readAll();

  return Boolean(all[slideId]?.[block]);
}

export function setBlock(slideId: string, block: ProgressBlockType, value: boolean): void {
  const all = readAll();
  const current = all[slideId] ?? {};
  const nextEntry = { ...current, [block]: value };
  writeAll({ ...all, [slideId]: nextEntry });
}

export function toggleBlock(slideId: string, block: ProgressBlockType): void {
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
  reviewed: Record<ProgressBlockType, boolean>;
  completedCount: number;
  isComplete: boolean;
} {
  const all = useProgressSnapshot();
  const entry = all[slideId] ?? {};
  const reviewed: Record<ProgressBlockType, boolean> = {
    [ProgressBlockType.Symptom]: Boolean(entry[ProgressBlockType.Symptom]),
    [ProgressBlockType.Rule]: Boolean(entry[ProgressBlockType.Rule]),
    [ProgressBlockType.Action]: Boolean(entry[ProgressBlockType.Action]),
  };
  const completedCount = PROGRESS_BLOCKS.filter((b) => reviewed[b]).length;

  return { reviewed, completedCount, isComplete: completedCount === PROGRESS_BLOCKS.length };
}
