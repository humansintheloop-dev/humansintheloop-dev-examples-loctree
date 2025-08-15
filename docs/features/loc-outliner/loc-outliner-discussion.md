# LOC Outliner Specification Discussion

## Q1: What specific file types should be counted for LOC?
**Answer:** For now all files

## Q2: How should the application handle binary files?
**Answer:** Treat as 0 LOC

## Q3: What navigation controls should be used?
**Answer:** No mouse. Up and down arrow for next/previous directory. Space to expand/contract. Q to quit

## Q4: Should LOC count include empty lines and comments or have filtering options?
**Answer:** Keep it simple (total line count)

## Q5: Should expanded directories show individual files or only subdirectories?
**Answer:** Only subdirs for now

## Q6: How should the app handle symbolic links and hidden directories?
**Answer:** Ignore symlinks and hidden directories

## Q7: Should there be visual indicators for expanded/collapsed state?
**Answer:** Arrows (▶ for collapsed, ▼ for expanded)

## Q8: How should the currently selected directory be highlighted?
**Answer:** Different color

## Q9: Should empty directories be shown or hidden?
**Answer:** Just show up as 0 LOC

## Q10: Should the app scan all LOC upfront or lazy-load?
**Answer:** To show the total for the root directory it needs to scan everything

## Q11: Should there be a loading indicator during initial scan?
**Answer:** Loading indicator

## Q12: How should directories be sorted?
**Answer:** Sorted - highest LOC first

## Q13: What should happen with invalid paths or file paths?
**Answer:** Error and exit

## Q14: Should there be a maximum depth limit for expanding directories?
**Answer:** No limit

## Q15: What should the executable be named?
**Answer:** loctree

---

## Final Specification Summary

### Application: loctree
A GoLang TUI application that displays lines of code (LOC) counts in an expandable/collapsible directory tree format.

### Core Functionality
- **Input**: Takes a directory path as command-line argument
- **Scanning**: Scans all files upfront to calculate total LOC with loading indicator
- **Display**: Shows LOC count followed by directory name, sorted by highest LOC first
- **Navigation**: Keyboard-only navigation with arrow keys and space bar
- **Tree Structure**: Expandable/collapsible directories with visual indicators (▶/▼)

### Technical Details
- **File Handling**:
  - Count all files (no filtering by extension)
  - Binary files treated as 0 LOC
  - Simple line counting (includes all lines: blank, comments, code)
  - Ignore symbolic links and hidden directories (starting with .)

- **UI Elements**:
  - Indented subdirectories for hierarchy
  - Only show subdirectories (no individual files)
  - Empty directories show as 0 LOC
  - Currently selected directory highlighted with different color
  - Loading indicator during initial scan

- **Controls**:
  - Up/Down arrows: Navigate between directories
  - Space: Expand/collapse current directory
  - Q: Quit application
  - No mouse support

- **Behavior**:
  - No depth limit for expansion
  - Error and exit on invalid/file paths
  - Directories sorted by LOC count (highest first) at all levels
