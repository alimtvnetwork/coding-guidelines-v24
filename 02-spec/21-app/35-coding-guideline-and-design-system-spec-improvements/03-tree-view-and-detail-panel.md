# Generic Tree View and Detail Panel

**Updated:** 28-Oct-26
**Evidence:** `antigravity-manager/src/components/instances/PromptTreeView/`
**Rule:** describe roles and tokens. Do not require that product's class names.

## Layout

```text
+------------------+-------------------------------+
| tree pane        | detail pane (flex-1)          |
| default 340px    | header, tabs, body            |
| min 280 max 760  |                               |
+------------------+-------------------------------+
```

Collapsed tree is a 40px rail labeled with a show action. It is not `display: none` and it is not width 0. Persist the width. Clamp it on read.

The shell can fill a dialog at about 90vh. The detail surface is the neutral canvas. The tree surface is a step above the canvas.

## Density (one size)

| Token | Value | Role |
| :--- | :--- | :--- |
| `--tree-font-size` | 12px | Row label |
| `--tree-row-pad-y` | 2px | Vertical padding |
| `--tree-row-pad-x` | 6px | Horizontal padding |
| `--tree-row-gap` | 2px | Gap between siblings |
| `--tree-indent` | 10px | Extra inset per depth |
| `--tree-rule` | 1px | Left rule on nested groups |
| `--tree-glyph-size` | 10px to 11px | `├──` and `└──` in a mono face |
| `--tree-chevron` | 14px root, 12px branch | Icon swap, not rotation |
| `--tree-type-icon` | 14px | Kind mark |
| `--tree-pane` | 340px | Default width |
| `--tree-pane-min` | 280px | Drag floor |
| `--tree-pane-max` | 760px | Drag ceiling |
| `--tree-rail` | 40px | Collapsed width |

Do not add a comfortable density unless the user asks. Do not give the row a tall fixed height.

## Color by role

Hover is the same at every depth: a light slate wash, or a dark zinc wash at about 60 percent. Hover is not the selected color.

| Role | Selected treatment | Notes |
| :--- | :--- | :--- |
| Root | Blue wash, dark text on light, cyan-tinted text on dark | The container, not the record |
| Branch | Violet wash and a violet hairline | The selectable record |
| Branch while a leaf inside it is active | Quieter violet, about 15 percent | Shows context without stealing the leaf |
| Leaf | Sky wash plus a 2px sky left rule | The item the panel renders |
| Nested agent | Solid violet, light text | The only row that may take a 1px focus ring |

Ordinary rows set no focus outline. A running mark is a small success dot (`#1af18d` in the reference). Do not paint the whole row green.

Reference chrome tokens that the rows do **not** use for hover:

- `--ui-surface-hover` `#2a1c3f`
- `--ui-surface-pressed` `#33244c`

Those belong to chrome. Copying them onto rows is the usual miss.

## Interaction

1. A chevron exists only when the node has children.
2. Chevron and row are separate targets. The chevron does not select.
3. Swap `expanded` and `collapsed` icons. Do not rotate a single icon.
4. Arrow keys move across expanded rows and wrap.
5. Outside a single-root scope, cap a branch at 50 children and say that the list is capped.
6. Selecting a branch loads the detail for that branch's current leaf: the running leaf if one exists, otherwise the last leaf.
7. The detail pane never stays on the previously selected record.

## Detail pane

Show, in order:

1. Identity: name, kind, sequence, elapsed state.
2. Primary actions: copy, send, enqueue. Do not hide them in another dialog.
3. Tabs: preview, raw, edit. Add instruction versus results when the record has both.
4. Body: full text, or a preview of about 200 words plus a word count and a byte size.
5. Result block when present: latest output, step count, running or finished.

### Empty states

| Condition | UI |
| :--- | :--- |
| Tree has no rows | Skeleton. No sentence. |
| Tree has rows and a root is selected | Root name, one sentence telling the user to select a leaf, and the primary action. |
| Nothing is selected | One sentence: select a leaf from the hierarchy to inspect it. |

## Mapping recipe

When an AI builds this in another product:

1. Rename root, branch, leaf, and nested agent to that product's nouns.
2. Keep the numbers in the token table unless the user asks for a different density.
3. Keep four selection treatments. Collapsing them into one color fails the check.
4. Bind the pane to the leaf.
5. Do not import `antigravity-manager` components. Copy the behavior.

## WP Exam motion that may sit beside a tree

Use these only for choice lists and admin rails, not for tree rows:

- Choice hover moves on X by 6px in 220ms with `cubic-bezier(0.16, 1, 0.3, 1)`.
- A rail that opens from 64px to 256px animates `width` for 300ms on that same family of curves.
- Theme swatches select with a 2px ring shadow so the box does not jump.

Source: `wp-exam/src/styles/theme.less` and `wp-exam/src/styles/theme-transition.css`.
