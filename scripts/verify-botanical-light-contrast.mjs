#!/usr/bin/env node
/**
 * WCAG 2.x contrast check for botanical-light tokens (file 42 §3).
 * Exit 0 when all required pairs pass AA normal text (4.5:1).
 */

const MIN_RATIO = 4.5;

function hslToRgb(h, s, l) {
  s /= 100;
  l /= 100;
  const c = (1 - Math.abs(2 * l - 1)) * s;
  const x = c * (1 - Math.abs(((h / 60) % 2) - 1));
  const m = l - c / 2;
  let r = 0;
  let g = 0;
  let b = 0;

  if (h < 60) {
    r = c;
    g = x;
  } else if (h < 120) {
    r = x;
    g = c;
  } else if (h < 180) {
    g = c;
    b = x;
  } else if (h < 240) {
    g = x;
    b = c;
  } else if (h < 300) {
    r = x;
    b = c;
  } else {
    r = c;
    b = x;
  }

  return [Math.round((r + m) * 255), Math.round((g + m) * 255), Math.round((b + m) * 255)];
}

function linearize(channel) {
  const v = channel / 255;

  if (v <= 0.03928) {
    return v / 12.92;
  }

  return ((v + 0.055) / 1.055) ** 2.4;
}

function relativeLuminance(rgb) {
  const [r, g, b] = rgb.map(linearize);

  return 0.2126 * r + 0.7152 * g + 0.0722 * b;
}

function contrastRatio(foregroundRgb, backgroundRgb) {
  const l1 = relativeLuminance(foregroundRgb);
  const l2 = relativeLuminance(backgroundRgb);
  const lighter = Math.max(l1, l2);
  const darker = Math.min(l1, l2);

  return (lighter + 0.05) / (darker + 0.05);
}

const tokens = {
  background: [140, 18, 97],
  foreground: [160, 22, 10],
  primary: [142, 70, 30],
  mutedForeground: [160, 9, 42],
  primaryForeground: [0, 0, 100],
  card: [0, 0, 100],
};

const bg = hslToRgb(...tokens.background);
const card = hslToRgb(...tokens.card);
const primary = hslToRgb(...tokens.primary);

const pairs = [
  ["foreground on background", hslToRgb(...tokens.foreground), bg],
  ["foreground on card", hslToRgb(...tokens.foreground), card],
  ["muted-foreground on background", hslToRgb(...tokens.mutedForeground), bg],
  ["primary on background (links/icons only)", primary, bg],
  ["primary-foreground on primary (filled buttons)", hslToRgb(...tokens.primaryForeground), primary],
];

let hasFailure = false;

for (const [label, fg, background] of pairs) {
  const ratio = contrastRatio(fg, background);
  const pass = ratio >= MIN_RATIO;

  if (!pass) {
    hasFailure = true;
  }

  const status = pass ? "PASS" : "FAIL";

  process.stdout.write(`${status} ${label}: ${ratio.toFixed(2)}:1 (min ${MIN_RATIO}:1)\n`);
}

if (hasFailure) {
  process.exit(1);
}

process.stdout.write("All botanical-light contrast pairs passed.\n");
