# loctree - Lines of Code Tree Viewer Specification

## 1. Executive Summary

`loctree` is a terminal user interface (TUI) application written in Go that provides an interactive, hierarchical view of lines of code (LOC) counts for a given directory structure. Users can navigate through directories using keyboard controls, expanding and collapsing folders to explore code distribution across their project.

## 2. Functional Requirements

### 2.1 Core Features
- Display total LOC count for the root directory
- Show LOC counts for all subdirectories in a tree structure
- Interactive navigation with keyboard controls
- Expandable/collapsible directory nodes
- Real-time directory expansion without re-scanning

### 2.2 Command Line Interface
- **Usage**: `loctree <directory_path>`
- **Input Validation**:
  - Verify the provided path exists
  - Verify the path points to a directory (not a file)
  - Display clear error message and exit with non-zero code on failure

### 2.3 User Controls
| Key | Action |
|-----|--------|
| ↑ (Up Arrow) | Navigate to previous directory in the tree |
| ↓ (Down Arrow) | Navigate to next directory in the tree |
| Space | Toggle expand/collapse state of current directory |
| Q/q | Quit the application |

### 2.4 Display Requirements
- **Format**: `<LOC_COUNT> <DIRECTORY_NAME>`
- **Example**: `1234 src`
- **Visual Indicators**:
  - `▶` prefix for collapsed directories with subdirectories
  - `▼` prefix for expanded directories
  - No indicator for directories without subdirectories
- **Highlighting**: Currently selected directory displayed in different color
- **Indentation**: Each nesting level indented by 2 spaces
- **Sorting**: All directories sorted by LOC count (descending) at each level

## 3. Technical Requirements

### 3.1 File Processing
- **Included Files**: All regular files in the directory tree
- **Binary Detection**: Treat binary files as 0 LOC
- **Line Counting**: Simple newline counting (include all lines: code, comments, blank)
- **Exclusions**:
  - Hidden directories (starting with `.`)
  - Symbolic links (both file and directory symlinks)
  - Special files (devices, pipes, sockets)

### 3.2 Performance Requirements
- **Initial Scan**: Complete directory traversal on startup
- **Loading Indicator**: Display progress during initial scan
- **Memory Efficiency**: Store pre-calculated LOC counts to avoid re-scanning
- **Responsiveness**: Immediate UI response to keyboard input after initial load

### 3.3 Data Structures

```go
type DirectoryNode struct {
    Name        string
    Path        string
    LOC         int
    Children    []*DirectoryNode
    IsExpanded  bool
    Parent      *DirectoryNode
}

type TreeState struct {
    Root            *DirectoryNode
    CurrentNode     *DirectoryNode
    VisibleNodes    []*DirectoryNode
    SelectedIndex   int
}
```

## 4. Architecture Design

### 4.1 Component Overview
1. **File Scanner**: Traverses directory tree and counts lines
2. **Tree Builder**: Constructs hierarchical data structure
3. **UI Renderer**: Manages TUI display and updates
4. **Input Handler**: Processes keyboard events
5. **State Manager**: Maintains application state

### 4.2 Recommended Libraries
- **TUI Framework**: `github.com/charmbracelet/bubbletea` or `github.com/gdamore/tcell`
- **File Walking**: Go standard library `filepath.Walk` or `filepath.WalkDir`
- **Binary Detection**: Check for null bytes in first chunk of file

### 4.3 Processing Flow
1. Parse command-line arguments
2. Validate input directory
3. Display loading indicator
4. Scan directory tree recursively
5. Build hierarchical data structure
6. Sort directories by LOC count
7. Initialize TUI
8. Render initial view
9. Enter event loop for user interaction

## 5. Implementation Details

### 5.1 Binary File Detection
```go
func isBinary(filePath string) bool {
    file, err := os.Open(filePath)
    if err != nil {
        return true // Treat unreadable files as binary
    }
    defer file.Close()
    
    buffer := make([]byte, 512)
    n, _ := file.Read(buffer)
    
    for i := 0; i < n; i++ {
        if buffer[i] == 0 {
            return true
        }
    }
    return false
}
```

### 5.2 Line Counting
```go
func countLines(filePath string) int {
    if isBinary(filePath) {
        return 0
    }
    
    file, err := os.Open(filePath)
    if err != nil {
        return 0
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lines := 0
    for scanner.Scan() {
        lines++
    }
    return lines
}
```

### 5.3 Directory Filtering
```go
func shouldSkipDir(name string, info os.FileInfo) bool {
    // Skip hidden directories
    if strings.HasPrefix(name, ".") {
        return true
    }
    
    // Skip symbolic links
    if info.Mode()&os.ModeSymlink != 0 {
        return true
    }
    
    return false
}
```

## 6. Error Handling

### 6.1 Startup Errors
- **Invalid Path**: "Error: Path does not exist: <path>"
- **Not a Directory**: "Error: Path is not a directory: <path>"
- **No Arguments**: "Usage: loctree <directory_path>"
- **Permission Denied**: "Error: Permission denied accessing: <path>"

### 6.2 Runtime Errors
- **File Access Errors**: Skip file, treat as 0 LOC
- **Directory Access Errors**: Skip directory and subdirectories
- **Terminal Resize**: Redraw UI to fit new dimensions
- **Panic Recovery**: Graceful shutdown with error message

## 7. Testing Plan

### 7.1 Unit Tests
- Binary file detection algorithm
- Line counting for various file types
- Directory filtering logic
- Tree sorting algorithm
- State management operations

### 7.2 Integration Tests
- Full directory scanning with mock filesystem
- UI navigation state transitions
- Expand/collapse operations
- Edge cases (empty directories, single file, deep nesting)

### 7.3 Manual Test Cases
1. **Empty Directory**: Should show 0 LOC
2. **Mixed Content**: Directory with code, binaries, and hidden files
3. **Deep Nesting**: Directory structure > 10 levels deep
4. **Large Codebase**: Test with repositories > 100,000 LOC
5. **Special Characters**: Directories with spaces, unicode in names
6. **Permission Issues**: Directories with restricted access
7. **Symbolic Links**: Verify symlinks are properly ignored

### 7.4 Performance Testing
- Measure scan time for various repository sizes
- Memory usage monitoring during operation
- UI responsiveness under load

## 8. Future Enhancements (Out of Scope)

These features are not required for initial implementation but noted for future consideration:

- File-level display toggle
- Filter by file extensions
- Export results to file
- Search functionality
- Code/comment/blank line separation
- Configuration file support
- Mouse support
- Horizontal scrolling for long names
- Watch mode for real-time updates
- Parallel scanning for performance

## 9. Acceptance Criteria

The implementation is complete when:

1. Application runs with `loctree <directory>` command
2. Correctly counts total lines in all text files
3. Ignores binary files, symlinks, and hidden directories
4. Displays expandable tree structure sorted by LOC
5. Responds to all specified keyboard controls
6. Shows loading indicator during initial scan
7. Handles errors gracefully with clear messages
8. Selected directory is visually highlighted
9. Performance is acceptable for codebases up to 1M LOC
10. All manual test cases pass successfully

## 10. Development Milestones

### Phase 1: Core Functionality (2-3 days)
- Set up Go project structure
- Implement file scanning and LOC counting
- Build directory tree data structure
- Add sorting logic

### Phase 2: TUI Implementation (2-3 days)
- Integrate TUI library
- Implement tree rendering
- Add keyboard navigation
- Implement expand/collapse logic

### Phase 3: Polish & Testing (1-2 days)
- Add loading indicator
- Implement error handling
- Write unit tests
- Perform manual testing
- Fix bugs and optimize performance

### Phase 4: Documentation (1 day)
- Write README with usage instructions
- Add inline code documentation
- Create build and installation instructions

## Appendix A: Example Output

```
▼ 15234 myproject
  ▼ 8921 src
    ▶ 4567 components
    ▶ 2341 utils
    ▼ 2013 services
      ▶ 1234 api
      ▶ 779 database
  ▶ 3456 tests
  ▶ 2857 docs
     0 config
```

## Appendix B: Build Instructions

```bash
# Clone repository
git clone <repository_url>
cd loctree

# Install dependencies
go mod init loctree
go get github.com/charmbracelet/bubbletea

# Build
go build -o loctree main.go

# Run
./loctree /path/to/directory
```