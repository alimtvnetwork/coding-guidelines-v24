/**
 * Named boolean flag constants.
 *
 * Per CODE-RED-024: bare `true`/`false` may not be passed as a positional
 * argument. Use these named constants when calling React `useState`,
 * setter functions, or any other API that takes a boolean.
 */

export const isFlagOn: boolean = true;
export const isFlagOff: boolean = false;

export const isOpen: boolean = true;
export const isClosed: boolean = false;

export const isVisible: boolean = true;
export const isHidden: boolean = false;

export const isActive: boolean = true;
export const isInactive: boolean = false;

export const isCopied: boolean = true;
export const isCopyReset: boolean = false;

export const isLoading: boolean = true;
export const isIdle: boolean = false;

export const isFullscreenOn: boolean = true;
export const isFullscreenOff: boolean = false;

export const isDragging: boolean = true;
export const isDragIdle: boolean = false;

export const isCollapsed: boolean = true;
export const isExpanded: boolean = false;

export const isOverflowing: boolean = true;
export const isOverflowFit: boolean = false;
