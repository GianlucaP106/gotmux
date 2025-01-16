# gotmux

A comprehensive library designed to interact with tmux, communicating through the tmux CLI and aiming to offer complete functionality for tmux integration.

## Requirements

- Go 1.22.3 or higher
- tmux installed on the system

## Installation

```bash
go get github.com/GianlucaP106/gotmux
```

## Features

This library provides a comprehensive Go interface to tmux, offering:

### Core Features

- **Complete Session Management**
  - Create, rename, and delete sessions
  - Attach/detach from sessions
  - List and query session information
  - Group session management
  - Session activity monitoring
- **Window Operations**
  - Create and manage windows
  - Move windows between sessions
  - Set window layouts
  - Navigate between windows
  - Window layout customization
  - Window activity flags
- **Pane Control**
  - Split panes horizontally or vertically
  - Resize and rearrange panes
  - Capture pane content
  - Pane synchronization
  - Command execution in panes
- **Server & Client Information**
  - Query server status and version
  - List connected clients
  - Get terminal information
  - Monitor client activity
  - Socket management

### Implementation Status

This library aims to be feature complete with tmux. Currently not all features are supported but they are planned to be implemented. Contributions are welcome.

## Documentation

- [pkg.go.dev](https://pkg.go.dev/github.com/GianlucaP106/gotmux/gotmux)
- [Examples Directory](https://github.com/GianlucaP106/gotmux/tree/main/examples)

## Usage

### Basic Example

```go
import (
    "fmt"
    "log"

    // import gotmux
    "github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
    // construct tmux client with socket path
    tmux, err := gotmux.NewTmux("/private/tmp/tmux-501/default")
    if err != nil {
        log.Fatal(err)
    }

    // you can also construct a default tmux client
    tmux, err = gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    // create a new session
    session, err := tmux.New()
    if err != nil {
        log.Fatal(err)
    }

    // attach with current terminal (if possible)
    err = session.Attach()
    if err != nil {
        log.Fatal(err)
    }

    // kill the session
    err = session.Kill()
    if err != nil {
        log.Fatal(err)
    }
}
```

### Advanced Examples

#### Session Management

```go
func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    // Create a named session with specific directory
    session, err := tmux.NewSession(&gotmux.SessionOptions{
        StartDirectory: "/home",
        Name:          "somename",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Attach to the session
    err = session.Attach()
    if err != nil {
        log.Fatal(err)
    }
}
```

#### Window and Pane Management

```go
func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    session, err := tmux.New()
    if err != nil {
        log.Fatal(err)
    }

    // Create and manage windows
    window, err := session.New()
    if err != nil {
        log.Fatal(err)
    }

    // Get and split panes
    pane, err := window.GetPaneByIndex(0)
    if err != nil {
        log.Fatal(err)
    }

    err = pane.Split()
    if err != nil {
        log.Fatal(err)
    }
}
```

#### Listing and Information

```go
func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    // List clients
    clients, err := tmux.ListClients()
    if err != nil {
        log.Fatal(err)
    }

    for _, c := range clients {
        fmt.Println(c.Tty, c.Session)
    }

    // Get server information
    server, err := tmux.GetServerInformation()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(server.Version)
}
```

#### Window Layout Management

```go
func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    session, err := tmux.New()
    if err != nil {
        log.Fatal(err)
    }

    window, err := session.New()
    if err != nil {
        log.Fatal(err)
    }

    // Available layouts:
    // - WindowLayoutEvenHorizontal
    // - WindowLayoutEvenVertical
    // - WindowLayoutMainVertical
    // - WindowLayoutTiled
    err = window.SelectLayout(gotmux.WindowLayoutEvenVertical)
    if err != nil {
        log.Fatal(err)
    }
}
```

See the complete list of [examples](https://github.com/GianlucaP106/gotmux/tree/main/examples)

## API Overview

### Core Types

All types provide comprehensive access to tmux attributes:

#### Session Type

```go
type Session struct {
    Activity          string
    Alerts            string
    Attached          int
    AttachedList      []string
    Created           string
    Format            bool
    Group             string
    GroupAttached     int
    GroupAttachedList []string
    GroupList         []string
    GroupManyAttached bool
    GroupSize         int
    Grouped           bool
    Id                string
    LastAttached      string
    ManyAttached      bool
    Marked            bool
    Name              string
    Path              string
    Stack             string
    Windows           int

    tmux *Tmux
}
```

### Common Operations

#### Session Operations

- `New()` - Create a new session
- `Attach()` - Attach to a session
- `Kill()` - Kill a session
- `Rename()` - Rename a session
- `ListWindows()` - List windows in a session

#### Window Operations

- `New()` - Create a new window
- `Kill()` - Kill a window
- `Move()` - Move a window
- `SelectLayout()` - Select window layout
- `ListPanes()` - List panes in a window

#### Pane Operations

- `Split()` - Split a pane
- `Kill()` - Kill a pane
- `Capture()` - Capture pane content
- `Select()` - Select a pane

## Error Handling

The library uses Go's standard error handling patterns. All operations that can fail return an error as their last return value:

```go
session, err := tmux.New()
if err != nil {
    // Handle error
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
