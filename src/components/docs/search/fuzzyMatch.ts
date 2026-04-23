export interface FuzzyMatchResult {
  score: number;
  matchedIndices: number[];
}

const NO_MATCH: FuzzyMatchResult = { score: 0, matchedIndices: [] };

const SCORE_EXACT = 100;
const SCORE_WORD_START = 15;
const SCORE_CONSECUTIVE = 10;
const SCORE_BASE_CHAR = 5;
const PENALTY_GAP = -2;
const MAX_GAP_PENALTY_DISTANCE = 10;
const RATIO_BONUS_MULTIPLIER = 20;
const FUZZY_PREVIEW_LENGTH = 500;

const DEFAULT_CONTEXT_CHARS = 60;
const DEFAULT_MAX_SNIPPETS = 3;

const WORD_BOUNDARY_CHARS = new Set([" ", "-", "_", "/", "."]);

function isWordBoundary(text: string, index: number): boolean {
  if (index === 0) {
    return true;
  }

  return WORD_BOUNDARY_CHARS.has(text[index - 1]);
}

function buildExactMatch(text: string, query: string, exactIdx: number): FuzzyMatchResult {
  const indices = Array.from({ length: query.length }, (_, i) => exactIdx + i);
  const boundaryBonus = isWordBoundary(text, exactIdx) ? SCORE_WORD_START : 0;

  return { score: SCORE_EXACT + boundaryBonus, matchedIndices: indices };
}

function scoreCharacterMatch(text: string, i: number, lastMatchIdx: number): number {
  const boundaryBonus = isWordBoundary(text, i) ? SCORE_WORD_START : 0;
  const isConsecutive = i === lastMatchIdx + 1;

  if (isConsecutive) {
    return SCORE_BASE_CHAR + boundaryBonus + SCORE_CONSECUTIVE;
  }

  const hasPriorMatch = lastMatchIdx >= 0;
  const gapPenalty = hasPriorMatch
    ? PENALTY_GAP * Math.min(i - lastMatchIdx, MAX_GAP_PENALTY_DISTANCE)
    : 0;

  return SCORE_BASE_CHAR + boundaryBonus + gapPenalty;
}

interface FuzzyState {
  matchedIndices: number[];
  queryIdx: number;
  score: number;
  lastMatchIdx: number;
}

function runFuzzyScan(text: string, textLower: string, queryLower: string): FuzzyState {
  const state: FuzzyState = { matchedIndices: [], queryIdx: 0, score: 0, lastMatchIdx: -2 };

  for (let i = 0; i < text.length && state.queryIdx < queryLower.length; i++) {
    if (textLower[i] !== queryLower[state.queryIdx]) {
      continue;
    }

    state.matchedIndices.push(i);
    state.score += scoreCharacterMatch(text, i, state.lastMatchIdx);
    state.lastMatchIdx = i;
    state.queryIdx++;
  }

  return state;
}

export function fuzzyMatch(text: string, query: string): FuzzyMatchResult {
  if (!query) {
    return NO_MATCH;
  }

  const textLower = text.toLowerCase();
  const queryLower = query.toLowerCase();

  const exactIdx = textLower.indexOf(queryLower);

  if (exactIdx !== -1) {
    return buildExactMatch(text, query, exactIdx);
  }

  const state = runFuzzyScan(text, textLower, queryLower);

  if (state.queryIdx < queryLower.length) {
    return NO_MATCH;
  }

  const ratioBonus = Math.round((query.length / text.length) * RATIO_BONUS_MULTIPLIER);

  return { score: state.score + ratioBonus, matchedIndices: state.matchedIndices };
}

export interface FuzzyContentSnippets {
  snippets: string[];
  snippetHighlights: number[][];
}

export interface FuzzyContentOptions {
  contextChars?: number;
  maxSnippets?: number;
}

function appendExactSnippet(
  content: string,
  query: string,
  idx: number,
  contextChars: number,
  acc: FuzzyContentSnippets
): void {
  const start = Math.max(0, idx - contextChars);
  const end = Math.min(content.length, idx + query.length + contextChars);
  const raw = content.slice(start, end).replace(/\n/g, " ");
  const prefix = start > 0 ? "…" : "";
  const suffix = end < content.length ? "…" : "";
  acc.snippets.push(`${prefix}${raw}${suffix}`);

  const highlightStart = idx - start + prefix.length;
  const highlights = Array.from({ length: query.length }, (_, i) => highlightStart + i);
  acc.snippetHighlights.push(highlights);
}

function collectExactSnippets(
  content: string,
  contentLower: string,
  query: string,
  queryLower: string,
  contextChars: number,
  maxSnippets: number
): FuzzyContentSnippets {
  const acc: FuzzyContentSnippets = { snippets: [], snippetHighlights: [] };
  const cursor = { searchFrom: 0 };

  while (acc.snippets.length < maxSnippets) {
    const idx = contentLower.indexOf(queryLower, cursor.searchFrom);

    if (idx === -1) {
      break;
    }

    appendExactSnippet(content, query, idx, contextChars, acc);
    cursor.searchFrom = idx + query.length;
  }

  return acc;
}

interface PreviewWindow {
  start: number;
  end: number;
  prefix: string;
  suffix: string;
}

function buildPreviewWindow(preview: string, result: FuzzyMatchResult, contextChars: number): PreviewWindow {
  const firstMatch = result.matchedIndices[0];
  const lastMatch = result.matchedIndices[result.matchedIndices.length - 1];
  const start = Math.max(0, firstMatch - contextChars);
  const end = Math.min(preview.length, lastMatch + contextChars);

  return {
    start,
    end,
    prefix: start > 0 ? "…" : "",
    suffix: end < preview.length ? "…" : "",
  };
}

function isResultEmpty(result: FuzzyMatchResult): boolean {
  return result.score <= 0 || result.matchedIndices.length === 0;
}

function appendFuzzyPreviewSnippet(
  preview: string,
  query: string,
  contextChars: number,
  acc: FuzzyContentSnippets
): void {
  const result = fuzzyMatch(preview, query);

  if (isResultEmpty(result)) {
    return;
  }

  const window = buildPreviewWindow(preview, result, contextChars);
  const raw = preview.slice(window.start, window.end).replace(/\n/g, " ");
  acc.snippets.push(`${window.prefix}${raw}${window.suffix}`);

  const highlights = result.matchedIndices
    .filter((i) => i >= window.start && i < window.end)
    .map((i) => i - window.start + window.prefix.length);
  acc.snippetHighlights.push(highlights);
}

export function fuzzyMatchContent(
  content: string,
  query: string,
  options: FuzzyContentOptions = {}
): FuzzyContentSnippets {
  const contextChars = options.contextChars ?? DEFAULT_CONTEXT_CHARS;
  const maxSnippets = options.maxSnippets ?? DEFAULT_MAX_SNIPPETS;
  const contentLower = content.toLowerCase();
  const queryLower = query.toLowerCase();

  const acc = collectExactSnippets(content, contentLower, query, queryLower, contextChars, maxSnippets);

  if (acc.snippets.length > 0) {
    return acc;
  }

  const preview = content.slice(0, FUZZY_PREVIEW_LENGTH);
  appendFuzzyPreviewSnippet(preview, query, contextChars, acc);

  return acc;
}
