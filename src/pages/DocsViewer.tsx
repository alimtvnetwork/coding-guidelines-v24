import { useState, useCallback } from "react";
import { useSpecData } from "@/hooks/useSpecData";
import type { SpecNode } from "@/types/spec";
import { SidebarProvider } from "@/components/ui/sidebar";
import { DocsContent } from "@/pages/DocsViewerComponents";
import { useDeepLinkFile, useFileSelection } from "@/pages/DocsViewerHelpers";
import { SearchDialog, useSearchShortcut } from "@/components/docs/SearchDialog";
import { CommandPalette, useCommandPaletteShortcut } from "@/components/docs/CommandPalette";
import { findSpecOverviewFile } from "@/components/docs/specOverviewJump";
import { GithubSyncBanner } from "@/components/docs/GithubSyncBanner";
import { isOpen, isClosed } from "@/constants/boolFlags";

function useDocsViewerState(allFiles: SpecNode[]) {
  const [activeFile, setActiveFile] = useState<SpecNode | null>(null);
  const [searchQuery, setSearchQuery] = useState("");
  const [searchOpen, setSearchOpen] = useState(isClosed);
  const [paletteOpen, setPaletteOpen] = useState(isClosed);

  useDeepLinkFile(allFiles, setActiveFile);
  const handleSelect = useFileSelection(setActiveFile, setSearchQuery);
  const openSearch = useCallback(() => setSearchOpen(isOpen), []);
  const openPalette = useCallback(() => setPaletteOpen(isOpen), []);
  useSearchShortcut(openSearch);
  useCommandPaletteShortcut(openPalette);

  const handleSearchSelect = useCallback((node: SpecNode) => {
    handleSelect(node);
    setSearchOpen(isClosed);
  }, [handleSelect]);

  const handlePaletteSelect = useCallback((node: SpecNode) => {
    handleSelect(node);
    setPaletteOpen(isClosed);
  }, [handleSelect]);

  const jumpToOverview = useCallback(() => {
    const overview = findSpecOverviewFile(allFiles);

    if (overview) {
      handleSelect(overview);

      return;
    }

    setPaletteOpen(isOpen);
  }, [allFiles, handleSelect]);

  return {
    activeFile, searchQuery, setSearchQuery,
    searchOpen, setSearchOpen, handleSelect, handleSearchSelect, openSearch,
    paletteOpen, setPaletteOpen, handlePaletteSelect, openPalette, jumpToOverview,
  };
}

function DocsViewerLayout({ tree, allFiles, state }: { tree: SpecNode[]; allFiles: SpecNode[]; state: ReturnType<typeof useDocsViewerState> }) {
  return (
    <div className="min-h-screen flex flex-col w-full">
      <GithubSyncBanner />
      <div className="flex-1 flex w-full">
        <DocsContent activeFile={state.activeFile} allFiles={allFiles} tree={tree} onSelect={state.handleSelect} searchQuery={state.searchQuery} setSearchQuery={state.setSearchQuery} onSearchOpen={state.openSearch} onJumpToOverview={state.jumpToOverview} />
      </div>
      <DocsViewerFooter />
    </div>
  );
}

function DocsViewerFooter() {
  return (
    <footer className="border-t border-border py-4 text-center text-sm text-muted-foreground space-y-1">
      <p>
        Built by{" "}
        <a href="https://alimkarim.com/" target="_blank" rel="noopener noreferrer" className="text-primary hover:underline font-medium">Md. Alim Ul Karim</a>
        {" "}— Chief Software Engineer,{" "}
        <a href="https://riseup-asia.com" target="_blank" rel="noopener noreferrer" className="text-primary hover:underline font-medium">Riseup Asia LLC</a>
      </p>
    </footer>
  );
}

export default function DocsViewer() {
  const { tree, allFiles } = useSpecData();
  const state = useDocsViewerState(allFiles);

  return (
    <SidebarProvider>
      <DocsViewerLayout tree={tree} allFiles={allFiles} state={state} />
      <SearchDialog open={state.searchOpen} onOpenChange={state.setSearchOpen} allFiles={allFiles} onSelect={state.handleSearchSelect} />
      <CommandPalette open={state.paletteOpen} onOpenChange={state.setPaletteOpen} allFiles={allFiles} onSelect={state.handlePaletteSelect} />
    </SidebarProvider>
  );
}
