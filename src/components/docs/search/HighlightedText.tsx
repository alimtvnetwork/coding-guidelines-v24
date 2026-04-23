import { fuzzyMatch } from "./fuzzyMatch";

interface HighlightedTextProps {
  text: string;
  query: string;
  highlightIndices?: number[];
}

function isQueryEmpty(query: string): boolean {
  return query.trim().length === 0;
}

function extendOrFlush(
  ranges: [number, number][],
  current: { start: number; end: number },
  next: number
): void {
  const isContiguous = next === current.end + 1;

  if (isContiguous) {
    current.end = next;

    return;
  }

  ranges.push([current.start, current.end]);
  current.start = next;
  current.end = next;
}

function buildHighlightRanges(indices: number[]): [number, number][] {
  if (indices.length === 0) {
    return [];
  }

  const sorted = [...indices].sort((a, b) => a - b);
  const ranges: [number, number][] = [];
  const cursor = { start: sorted[0], end: sorted[0] };

  for (let i = 1; i < sorted.length; i++) {
    extendOrFlush(ranges, cursor, sorted[i]);
  }

  ranges.push([cursor.start, cursor.end]);

  return ranges;
}

function renderHighlightedRanges(text: string, ranges: [number, number][]): React.ReactNode[] {
  const parts: React.ReactNode[] = [];
  const cursor = { lastEnd: 0 };

  for (const [start, end] of ranges) {
    if (start > cursor.lastEnd) {
      parts.push(<span key={`t-${cursor.lastEnd}`}>{text.slice(cursor.lastEnd, start)}</span>);
    }

    parts.push(
      <mark key={`h-${start}`} className="bg-primary/25 text-foreground rounded-sm px-0.5">
        {text.slice(start, end + 1)}
      </mark>
    );
    cursor.lastEnd = end + 1;
  }

  if (cursor.lastEnd < text.length) {
    parts.push(<span key={`t-${cursor.lastEnd}`}>{text.slice(cursor.lastEnd)}</span>);
  }

  return parts;
}

export function HighlightedText({ text, query, highlightIndices }: HighlightedTextProps) {
  if (isQueryEmpty(query)) {
    return <>{text}</>;
  }

  const indices = highlightIndices ?? fuzzyMatch(text, query).matchedIndices;

  if (indices.length === 0) {
    return <>{text}</>;
  }

  const ranges = buildHighlightRanges(indices);
  const parts = renderHighlightedRanges(text, ranges);

  return <>{parts}</>;
}
