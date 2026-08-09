import { createContext, useContext, useEffect, useState } from "react";

export enum ThemeType {
  Light = "light",
  Dark = "dark"
}

interface ThemeContextType {
  theme: ThemeType;
  toggleTheme: () => void;
}

const ThemeContext = createContext<ThemeContextType>({
  theme: ThemeType.Light,
  toggleTheme: () => {},
});

const STORAGE_KEY = "docs-theme";

function getInitialTheme(): ThemeType {
  if (typeof window === "undefined") return ThemeType.Light;

  const stored = localStorage.getItem(STORAGE_KEY) as ThemeType | null;

  if (stored) return stored;

  const isDarkPreferred = window.matchMedia("(prefers-color-scheme: dark)").matches;

  return isDarkPreferred ? ThemeType.Dark : ThemeType.Light;
}

function applyThemeToDocument(theme: ThemeType): void {
  const root = document.documentElement;
  root.classList.remove(ThemeType.Light, ThemeType.Dark);
  root.classList.add(theme);
  localStorage.setItem(STORAGE_KEY, theme);
}

export function ThemeProvider({ children }: { children: React.ReactNode }) {
  const [theme, setTheme] = useState<ThemeType>(getInitialTheme);

  useEffect(() => {
    applyThemeToDocument(theme);
  }, [theme]);

  const toggleTheme = () => setTheme((prev) => (prev === ThemeType.Light ? ThemeType.Dark : ThemeType.Light));

  return (
    <ThemeContext.Provider value={{ theme, toggleTheme }}>
      {children}
    </ThemeContext.Provider>
  );
}

export const useTheme = () => useContext(ThemeContext);
