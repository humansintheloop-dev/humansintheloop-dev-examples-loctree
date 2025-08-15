# loctree Implementation Plan - Steel Thread Methodology

## Overview

This implementation plan follows the Steel Thread methodology, where each thread represents a narrow end-to-end flow that delivers value to the user. Each thread builds upon the previous ones, incrementally adding functionality. The implementation uses Test-Driven Development (TDD) principles throughout.

**Important Instructions for the Coding Agent:**
- Mark each checkbox `[x]` when a task or step is completed
- Follow TDD principles: Write test first, then implementation, then refactor
- Commit after each passing test
- Run all tests before committing

---

## Thread 1: Project Setup and CI/CD Pipeline

**Goal:** Set up a Go project with proper structure, dependencies, and automated CI/CD pipeline.

### Prompt for Thread 1:

```text
Create a new Go project called "loctree" with the following requirements:

[x] 1. Initialize Go module:
   [x] Run `go mod init github.com/user/loctree`
   [x] Create main.go with a simple "Hello, loctree!" program
   [x] Verify it runs with `go run main.go`

[x] 2. Set up project structure:
   [x] Create directory structure:
       - cmd/loctree/ (for main application)
       - internal/scanner/ (for file scanning logic)
       - internal/tree/ (for tree data structures)
       - internal/ui/ (for TUI components)
       - internal/state/ (for state management)
       - test/ (for integration tests)
   [x] Move main.go to cmd/loctree/main.go
   [x] Update main.go to use proper package structure

[x] 3. Add initial dependencies:
   [x] Add Bubble Tea TUI framework: `go get github.com/charmbracelet/bubbletea`
   [x] Add lipgloss for styling: `go get github.com/charmbracelet/lipgloss`
   [x] Create go.sum file by running `go mod tidy`

[x] 4. Create Makefile for common tasks:
   [x] Add build target: `build: go build -o bin/loctree cmd/loctree/main.go`
   [x] Add test target: `test: go test ./...`
   [x] Add run target: `run: go run cmd/loctree/main.go`
   [x] Add clean target: `clean: rm -rf bin/`

[x] 5. Set up GitHub Actions CI/CD pipeline:
   [x] Create .github/workflows/ci.yml
   [x] Configure workflow to trigger on push and pull requests
   [x] Add job to run on ubuntu-latest, macos-latest, windows-latest
   [x] Add steps:
       - Checkout code
       - Set up Go 1.21
       - Install dependencies (go mod download)
       - Run tests (make test)
       - Build application (make build)
       - Upload artifacts (the built binaries)

[x] 6. Create initial test to verify setup:
   [x] Create cmd/loctree/main_test.go
   [x] Write test that verifies the application can be built
   [x] Run test and ensure it passes
   [x] Commit: "Setup Go project with CI/CD pipeline"
```

---

## Thread 2: Command-Line Argument Parsing and Validation

**Goal:** Accept a directory path as an argument and validate it exists and is accessible.

### Prompt for Thread 2:

```text
Implement command-line argument parsing and validation for loctree:

[x] 1. Write failing test for argument parsing:
   [x] Create internal/cli/args_test.go
   [x] Write test TestParseArgs_NoArguments - should return error
   [x] Write test TestParseArgs_ValidDirectory - should return directory path
   [x] Write test TestParseArgs_InvalidPath - should return error
   [x] Run tests - they should fail

[x] 2. Implement argument parsing:
   [x] Create internal/cli/args.go
   [x] Create ParseArgs function that:
       - Checks if exactly one argument is provided
       - Returns error with usage message if not
   [x] Make tests pass
   [x] Commit: "Add command-line argument parsing"

[x] 3. Write failing tests for path validation:
   [x] Write test TestValidatePath_DirectoryExists - should succeed
   [x] Write test TestValidatePath_FileNotDirectory - should return error
   [x] Write test TestValidatePath_NonExistentPath - should return error
   [x] Run tests - they should fail

[x] 4. Implement path validation:
   [x] Create ValidatePath function that:
       - Checks if path exists using os.Stat
       - Verifies it's a directory (not a file)
       - Returns appropriate error messages
   [x] Make all tests pass
   [x] Commit: "Add path validation"

[x] 5. Integrate into main application:
   [x] Update cmd/loctree/main.go to:
       - Parse arguments
       - Validate path
       - Exit with error code 1 on failure
       - Print "Scanning: <path>" on success
   [x] Create integration test in cmd/loctree/main_test.go
   [x] Run manual test: `go run cmd/loctree/main.go /tmp`
   [x] Commit: "Integrate argument parsing into main"
```

---

## Thread 3: Basic File Scanning and Line Counting

**Goal:** Scan a directory and count lines in text files, handling binary files correctly.

### Prompt for Thread 3:

```text
Implement file scanning and line counting functionality:

[x] 1. Write failing test for binary file detection:
   [x] Create internal/scanner/binary_test.go
   [x] Write test TestIsBinary_TextFile - should return false
   [x] Write test TestIsBinary_BinaryFile - should return true
   [x] Write test TestIsBinary_EmptyFile - should return false
   [x] Create test fixtures (text and binary files)
   [x] Run tests - they should fail

[x] 2. Implement binary file detection:
   [x] Create internal/scanner/binary.go
   [x] Implement IsBinary function:
       - Read first 512 bytes of file
       - Check for null bytes (0x00)
       - Return true if null byte found
   [x] Make tests pass
   [x] Commit: "Add binary file detection"

[x] 3. Write failing tests for line counting:
   [x] Create internal/scanner/counter_test.go
   [x] Write test TestCountLines_EmptyFile - should return 0
   [x] Write test TestCountLines_SingleLine - should return 1
   [x] Write test TestCountLines_MultipleLines - should return correct count
   [x] Write test TestCountLines_BinaryFile - should return 0
   [x] Run tests - they should fail

[x] 4. Implement line counting:
   [x] Create internal/scanner/counter.go
   [x] Implement CountLines function:
       - Check if file is binary first
       - Use bufio.Scanner to count lines
       - Handle read errors gracefully
   [x] Make all tests pass
   [x] Commit: "Add line counting functionality"

[x] 5. Write failing test for directory scanning:
   [x] Create internal/scanner/scanner_test.go
   [x] Write test TestScanDirectory_SingleFile
   [x] Write test TestScanDirectory_MultipleFiles
   [x] Write test TestScanDirectory_MixedContent (text and binary)
   [x] Write test TestScanDirectory_HiddenFiles (should be ignored)
   [x] Run tests - they should fail

[x] 6. Implement directory scanning:
   [x] Create internal/scanner/scanner.go
   [x] Implement ScanDirectory function:
       - Use filepath.WalkDir for traversal
       - Skip hidden directories (starting with .)
       - Skip symbolic links
       - Count lines for each regular file
       - Return total LOC count
   [x] Make all tests pass
   [x] Refactor if needed
   [x] Commit: "Add directory scanning"

[x] 7. Integration test and main update:
   [x] Update main.go to scan directory and print total LOC
   [x] Create integration test with test directory structure
   [x] Manual test on a real codebase
   [x] Commit: "Integrate scanning into main application"
```

---

## Thread 4: Tree Data Structure

**Goal:** Build hierarchical tree structure to represent directory LOC counts.

### Prompt for Thread 4:

```text
Implement tree data structure for directory hierarchy:

[x] 1. Write failing tests for DirectoryNode:
   [x] Create internal/tree/node_test.go
   [x] Write test TestNewDirectoryNode - verify initialization
   [x] Write test TestAddChild - verify parent-child relationship
   [x] Write test TestCalculateLOC - verify LOC aggregation
   [x] Run tests - they should fail

[x] 2. Implement DirectoryNode structure:
   [x] Create internal/tree/node.go
   [x] Define DirectoryNode struct:
       - Name, Path, LOC fields
       - Children slice
       - Parent pointer
       - IsExpanded bool
   [x] Implement NewDirectoryNode constructor
   [x] Implement AddChild method
   [x] Implement CalculateLOC method (sum of children + own files)
   [x] Make tests pass
   [x] Commit: "Add DirectoryNode data structure"

[x] 3. Write failing tests for tree building:
   [x] Create internal/tree/builder_test.go
   [x] Write test TestBuildTree_SingleDirectory
   [x] Write test TestBuildTree_NestedDirectories
   [x] Write test TestBuildTree_CalculatesLOC
   [x] Run tests - they should fail

[x] 4. Implement tree builder:
   [x] Create internal/tree/builder.go
   [x] Implement BuildTree function:
       - Create root node
       - Walk directory structure
       - Create nodes for each directory
       - Count LOC for files in each directory
       - Establish parent-child relationships
       - Calculate aggregate LOC counts
   [x] Make all tests pass
   [x] Commit: "Add tree builder"

[x] 5. Write failing tests for tree sorting:
   [x] Write test TestSortChildren_ByLOC
   [x] Write test TestSortChildren_Recursive
   [x] Run tests - they should fail

[x] 6. Implement sorting:
   [x] Add SortChildren method to DirectoryNode
   [x] Sort by LOC descending at each level
   [x] Apply recursively to all children
   [x] Make tests pass
   [x] Refactor if needed
   [x] Commit: "Add tree sorting by LOC"

[x] 7. Integration with scanner:
   [x] Update scanner to build tree structure
   [x] Update main.go to build and display tree (text only for now)
   [x] Manual test to verify tree structure
   [x] Commit: "Integrate tree building with scanner"
```

---

## Thread 5: Basic TUI Display

**Goal:** Display the tree structure in a TUI with basic rendering.

### Prompt for Thread 5:

```text
Implement basic TUI display using Bubble Tea:

[x] 1. Write failing test for tree rendering:
   [x] Create internal/ui/renderer_test.go
   [x] Write test TestRenderNode_Collapsed
   [x] Write test TestRenderNode_Expanded
   [x] Write test TestRenderNode_WithIndentation
   [x] Run tests - they should fail

[x] 2. Implement tree renderer:
   [x] Create internal/ui/renderer.go
   [x] Implement RenderTree function:
       - Format: "<indicator> <LOC> <name>"
       - Use ▶ for collapsed directories
       - Use ▼ for expanded directories
       - Add proper indentation (2 spaces per level)
   [x] Make tests pass
   [x] Commit: "Add tree rendering logic"

[x] 3. Write failing test for TUI model:
   [x] Create internal/ui/model_test.go
   [x] Write test TestNewModel - verify initialization
   [x] Write test TestView - verify view generation
   [x] Run tests - they should fail

[x] 4. Implement Bubble Tea model:
   [x] Create internal/ui/model.go
   [x] Define Model struct implementing tea.Model:
       - Root *tree.DirectoryNode
       - VisibleNodes slice
       - SelectedIndex int
   [x] Implement Init() method
   [x] Implement Update() method (basic, no interaction yet)
   [x] Implement View() method using renderer
   [x] Make tests pass
   [x] Commit: "Add Bubble Tea model"

[x] 5. Integrate TUI into main:
   [x] Update main.go to:
       - Build tree from scanned directory
       - Create TUI model
       - Run Bubble Tea program
   [x] Manual test - should display static tree
   [x] Commit: "Display tree in TUI"

[x] 6. Add loading indicator:
   [x] Create loading model for scanning phase
   [x] Show "Scanning..." message during scan
   [x] Switch to tree view when complete
   [x] Test with large directory
   [x] Commit: "Add loading indicator"
```

---

## Thread 6: Keyboard Navigation

**Goal:** Implement keyboard controls for navigating the tree.

### Prompt for Thread 6:

```text
Implement keyboard navigation in the TUI:

[x] 1. Write failing tests for navigation state:
   [x] Create internal/state/navigation_test.go
   [x] Write test TestMoveUp - verify selection moves up
   [x] Write test TestMoveDown - verify selection moves down
   [x] Write test TestMoveUp_AtTop - should stay at top
   [x] Write test TestMoveDown_AtBottom - should stay at bottom
   [x] Run tests - they should fail

[x] 2. Implement navigation state management:
   [x] Create internal/state/navigation.go
   [x] Implement navigation functions:
       - MoveUp(currentIndex, maxIndex)
       - MoveDown(currentIndex, maxIndex)
       - Returns new selected index
   [x] Make tests pass
   [x] Commit: "Add navigation state management"

[x] 3. Write failing tests for keyboard handling:
   [x] Update internal/ui/model_test.go
   [x] Write test TestUpdate_ArrowUp
   [x] Write test TestUpdate_ArrowDown
   [x] Write test TestUpdate_QuitKey
   [x] Run tests - they should fail

[x] 4. Implement keyboard handling in TUI:
   [x] Update Update() method in model.go:
       - Handle arrow up/down keys
       - Update selected index
       - Handle 'q' key for quit
       - Return tea.Quit command when quitting
   [x] Make tests pass
   [x] Commit: "Add keyboard navigation"

[x] 5. Write failing test for selection highlighting:
   [x] Write test TestRenderTree_WithSelection
   [x] Verify selected item is highlighted
   [x] Run test - should fail

[x] 6. Implement selection highlighting:
   [x] Update renderer to accept selected index
   [x] Use lipgloss styles for highlighting
   [x] Apply different style to selected row
   [x] Make test pass
   [x] Commit: "Add selection highlighting"

[x] 7. Manual testing:
   [x] Test navigation on multi-level tree
   [x] Verify smooth scrolling for long lists
   [x] Test quit functionality
   [x] Commit: "Complete keyboard navigation"
```

---

## Thread 7: Expand/Collapse Functionality

**Goal:** Implement the ability to expand and collapse directories with the spacebar.

### Prompt for Thread 7:

```text
Implement expand/collapse functionality:

[x] 1. Write failing tests for expand/collapse logic:
   [x] Create internal/tree/operations_test.go
   [x] Write test TestToggleExpanded_Collapsed
   [x] Write test TestToggleExpanded_Expanded
   [x] Write test TestGetVisibleNodes_AllCollapsed
   [x] Write test TestGetVisibleNodes_SomeExpanded
   [x] Run tests - they should fail

[x] 2. Implement expand/collapse operations:
   [x] Create internal/tree/operations.go
   [x] Implement ToggleExpanded method on DirectoryNode
   [x] Implement GetVisibleNodes function:
       - Returns flat list of currently visible nodes
       - Respects expanded/collapsed state
       - Maintains tree order
   [x] Make tests pass
   [x] Commit: "Add expand/collapse logic"

[x] 3. Write failing tests for spacebar handling:
   [x] Update model_test.go
   [x] Write test TestUpdate_Spacebar_Toggle
   [x] Write test TestUpdate_Spacebar_UpdatesVisibleNodes
   [x] Run tests - they should fail

[x] 4. Implement spacebar handling:
   [x] Update Update() method to handle spacebar
   [x] Toggle expanded state of selected node
   [x] Rebuild visible nodes list
   [x] Adjust selected index if needed
   [x] Make tests pass
   [x] Commit: "Add spacebar toggle functionality"

[x] 5. Write failing test for tree indicators:
   [x] Write test TestRenderNode_NoChildrenNoIndicator
   [x] Write test TestRenderNode_CollapsedWithChildren
   [x] Write test TestRenderNode_ExpandedWithChildren
   [x] Run tests - they should fail

[x] 6. Update rendering for indicators:
   [x] Only show ▶/▼ for directories with children
   [x] Show ▶ when collapsed with children
   [x] Show ▼ when expanded
   [x] No indicator for leaf directories
   [x] Make tests pass
   [x] Commit: "Update tree indicators"

[x] 7. Integration testing:
   [x] Test expanding/collapsing nested directories
   [x] Verify LOC counts remain accurate
   [x] Test navigation after expand/collapse
   [x] Commit: "Complete expand/collapse feature"
```

---

## Thread 8: Error Handling and Edge Cases

**Goal:** Handle errors gracefully and support edge cases.

### Prompt for Thread 8:

```text
Implement comprehensive error handling:

[ ] 1. Write failing tests for permission errors:
   [ ] Create internal/scanner/errors_test.go
   [ ] Write test TestScanDirectory_PermissionDenied
   [ ] Write test TestCountLines_UnreadableFile
   [ ] Create test fixtures with restricted permissions
   [ ] Run tests - they should fail

[ ] 2. Implement permission error handling:
   [ ] Update scanner to skip inaccessible directories
   [ ] Log skipped directories (don't fail entire scan)
   [ ] Treat unreadable files as 0 LOC
   [ ] Make tests pass
   [ ] Commit: "Handle permission errors"

[ ] 3. Write failing tests for edge cases:
   [ ] Write test TestScanDirectory_EmptyDirectory
   [ ] Write test TestScanDirectory_SingleFile
   [ ] Write test TestScanDirectory_DeepNesting (>20 levels)
   [ ] Write test TestScanDirectory_LargeDirectory (>1000 files)
   [ ] Run tests - they should fail

[ ] 4. Handle edge cases:
   [ ] Support empty directories (show 0 LOC)
   [ ] Handle single file as input (show parent dir)
   [ ] Optimize deep nesting performance
   [ ] Handle large directories efficiently
   [ ] Make all tests pass
   [ ] Commit: "Handle edge cases"

[ ] 5. Write failing tests for special characters:
   [ ] Write test TestScanDirectory_SpacesInNames
   [ ] Write test TestScanDirectory_UnicodeCharacters
   [ ] Write test TestRenderTree_LongFileNames
   [ ] Run tests - they should fail

[ ] 6. Handle special characters:
   [ ] Properly handle spaces in paths
   [ ] Support unicode in directory names
   [ ] Truncate long names in display
   [ ] Make tests pass
   [ ] Commit: "Support special characters"

[ ] 7. Add panic recovery:
   [ ] Implement defer/recover in main
   [ ] Clean shutdown on panic
   [ ] Display user-friendly error message
   [ ] Test with simulated panics
   [ ] Commit: "Add panic recovery"

[ ] 8. Terminal resize handling:
   [ ] Listen for terminal resize events
   [ ] Redraw UI on resize
   [ ] Test by resizing terminal during use
   [ ] Commit: "Handle terminal resize"
```

---

## Thread 9: Performance Optimization

**Goal:** Optimize performance for large codebases.

### Prompt for Thread 9:

```text
Optimize performance for large codebases:

[ ] 1. Write performance benchmarks:
   [ ] Create internal/scanner/scanner_bench_test.go
   [ ] Write BenchmarkScanDirectory_Small (100 files)
   [ ] Write BenchmarkScanDirectory_Medium (1000 files)
   [ ] Write BenchmarkScanDirectory_Large (10000 files)
   [ ] Run benchmarks, record baseline

[ ] 2. Optimize file reading:
   [ ] Profile current implementation
   [ ] Use buffered I/O for line counting
   [ ] Implement file size check before reading
   [ ] Skip very large files (>10MB) quickly
   [ ] Run benchmarks, verify improvement
   [ ] Commit: "Optimize file reading"

[ ] 3. Implement concurrent scanning:
   [ ] Create worker pool for file scanning
   [ ] Use channels for work distribution
   [ ] Maintain thread-safe LOC accumulation
   [ ] Run benchmarks, verify improvement
   [ ] Ensure tests still pass
   [ ] Commit: "Add concurrent file scanning"

[ ] 4. Optimize memory usage:
   [ ] Profile memory allocation
   [ ] Reuse buffers where possible
   [ ] Implement node pooling for large trees
   [ ] Run memory profiling tests
   [ ] Commit: "Optimize memory usage"

[ ] 5. Add progress reporting:
   [ ] Count total files during initial walk
   [ ] Update loading indicator with progress
   [ ] Show "Scanning... (500/1000 files)"
   [ ] Test with large repository
   [ ] Commit: "Add progress reporting"

[ ] 6. Performance validation:
   [ ] Test with Linux kernel source
   [ ] Test with chromium source
   [ ] Ensure <5 second scan for 100k LOC
   [ ] Document performance metrics
   [ ] Commit: "Validate performance targets"
```

---

## Thread 10: Final Polish and Documentation

**Goal:** Add final touches, comprehensive testing, and documentation.

### Prompt for Thread 10:

```text
Final polish and documentation:

[ ] 1. Add version information:
   [ ] Create internal/version/version.go
   [ ] Add version constant
   [ ] Add --version flag support
   [ ] Update help text with version
   [ ] Commit: "Add version information"

[ ] 2. Improve error messages:
   [ ] Review all error messages for clarity
   [ ] Add suggestions for common errors
   [ ] Use consistent formatting
   [ ] Test error scenarios manually
   [ ] Commit: "Improve error messages"

[ ] 3. Create comprehensive integration tests:
   [ ] Create test/integration_test.go
   [ ] Test full workflow with real directories
   [ ] Test all keyboard shortcuts
   [ ] Test error conditions
   [ ] Achieve >80% code coverage
   [ ] Commit: "Add integration tests"

[ ] 4. Create README.md:
   [ ] Add project description
   [ ] Installation instructions
   [ ] Usage examples with screenshots
   [ ] Keyboard shortcuts table
   [ ] Build from source instructions
   [ ] Commit: "Add README documentation"

[ ] 5. Add inline documentation:
   [ ] Document all public functions
   [ ] Add package-level documentation
   [ ] Include usage examples in comments
   [ ] Run go doc to verify
   [ ] Commit: "Add code documentation"

[ ] 6. Create release build:
   [ ] Update Makefile with release target
   [ ] Build for multiple platforms
   [ ] Create GitHub release workflow
   [ ] Test installation on clean system
   [ ] Commit: "Add release build configuration"

[ ] 7. Final testing checklist:
   [ ] Test on Linux, macOS, Windows
   [ ] Test with various terminal emulators
   [ ] Test with different color schemes
   [ ] Performance test with 1M+ LOC codebase
   [ ] All manual test cases from spec pass
   [ ] Commit: "Final testing complete"

[ ] 8. Create CHANGELOG.md:
   [ ] Document v1.0.0 features
   [ ] List known limitations
   [ ] Add acknowledgments
   [ ] Commit: "Add changelog"
```

---

## Completion Checklist

After all threads are complete, verify:

- [ ] Application runs with `loctree <directory>` command
- [ ] Correctly counts total lines in all text files
- [ ] Ignores binary files, symlinks, and hidden directories
- [ ] Displays expandable tree structure sorted by LOC
- [ ] Responds to all specified keyboard controls (↑, ↓, Space, q)
- [ ] Shows loading indicator during initial scan
- [ ] Handles errors gracefully with clear messages
- [ ] Selected directory is visually highlighted
- [ ] Performance is acceptable for codebases up to 1M LOC
- [ ] All unit tests pass
- [ ] All integration tests pass
- [ ] CI/CD pipeline runs successfully
- [ ] Documentation is complete

## Notes for Implementation

- Each thread should be completed in sequence
- Commit after each passing test (TDD approach)
- Run full test suite before moving to next thread
- If a thread reveals issues in previous threads, fix them before continuing
- Keep commits atomic and well-described
- Use co-authoring in commit messages as specified

This plan provides approximately 10 focused threads that build the application incrementally, each delivering value while maintaining a working application at each step.