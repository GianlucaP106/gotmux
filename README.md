# gotmux

[![Go Report Card](https://goreportcard.com/badge/github.com/GianlucaP106/gotmux)](https://goreportcard.com/report/github.com/GianlucaP106/gotmux)
[![Go Reference](https://pkg.go.dev/badge/github.com/GianlucaP106/gotmux/gotmux.svg)](https://pkg.go.dev/github.com/GianlucaP106/gotmux/gotmux)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive Go library for programmatically interacting with tmux sessions, windows, and panes. gotmux provides a clean, type-safe interface to manage tmux through Go, making it easy to automate terminal multiplexing tasks and build tmux-based applications.

## Features

- **Complete tmux Integration**: Full access to tmux sessions, windows, and panes
- **Type-Safe API**: Strongly typed structures for all tmux entities
- **Comprehensive Data Access**: Retrieve all tmux fields and properties
- **Session Management**: Create, rename, delete, and manage sessions
- **Window Operations**: Handle windows, layouts, and navigation
- **Pane Control**: Split, resize, and manage panes programmatically
- **Server Information**: Query tmux server status and client connections
- **Error Handling**: Robust error handling throughout the API

## Requirements

- tmux installed on your system

## Installation

```bash
go get github.com/GianlucaP106/gotmux
```

## Quick Start

```go
package main

import (
    "log"
    "github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
    // Initialize tmux client
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    // Create a new session
    session, err := tmux.New()
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

## API Overview

### Session Management

Create and manage tmux sessions with full control over their properties:

```go
// Create a new session with custom options
session, err := tmux.NewSession(&gotmux.SessionOptions{
    StartDirectory: "/home/user",
    Name:          "my-session",
})

// List all sessions
sessions, err := tmux.ListSessions()

// Attach to a session
err = session.Attach()

// Rename a session
err = session.Rename("new-name")
```

### Window Operations

Manage windows within sessions:

```go
// Create a new window
window, err := session.New()

// Get window by index
window, err := session.GetWindowByIndex(0)

// Select a specific layout
err = window.SelectLayout(gotmux.WindowLayoutEvenVertical)

// Move window to another session
err = window.MoveWindow(targetSession)
```

### Pane Management

Control individual panes within windows:

```go
// Get pane by index
pane, err := window.GetPaneByIndex(0)

// Split pane horizontally
err = pane.Split()

// Split pane vertically
err = pane.SplitVertical()

// Resize pane
err = pane.Resize(10, 20)

// Execute command in pane
err = pane.SendKeys("ls -la")
```

### Server and Client Information

Query tmux server status and connected clients:

```go
// Get server information
server, err := tmux.GetServerInformation()
fmt.Printf("tmux version: %s\n", server.Version)

// List connected clients
clients, err := tmux.ListClients()
for _, client := range clients {
    fmt.Printf("Client: %s, Session: %s\n", client.Tty, client.Session)
}
```

## Data Structures

gotmux provides comprehensive data structures that capture all tmux fields. For example, the Session struct includes:

```go
type Session struct {
    Activity          string   // Last activity time
    Alerts            string   // Alert flags
    Attached          int      // Number of attached clients
    AttachedList      []string // List of attached clients
    Created           string   // Creation time
    Format            bool     // Format flag
    Group             string   // Session group
    GroupAttached     int      // Number of attached clients in group
    GroupAttachedList []string // List of attached clients in group
    GroupList         []string // List of sessions in group
    GroupManyAttached bool     // Multiple clients attached to group
    GroupSize         int      // Number of sessions in group
    Grouped           bool     // Whether session is grouped
    Id                string   // Session ID
    LastAttached      string   // Last attachment time
    ManyAttached      bool     // Multiple clients attached
    Marked            bool     // Marked flag
    Name              string   // Session name
    Path              string   // Working directory
    Stack             string   // Stack information
    Windows           int      // Number of windows
}
```

## Examples

The project includes comprehensive examples in the `examples/` directory:

- **create/**: Basic session, window, and pane creation
- **sessions/**: Session management operations
- **windows_panes/**: Window and pane manipulation
- **listing/**: Querying and listing tmux entities
- **get_server_info/**: Server information retrieval
- **list_clients/**: Client connection management

### Example: Advanced Session Management

```go
func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    // Create session with specific options
    session, err := tmux.NewSession(&gotmux.SessionOptions{
        StartDirectory: "/home/user/projects",
        Name:          "development",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Create a window and split it
    window, err := session.New()
    if err != nil {
        log.Fatal(err)
    }

    pane, err := window.GetPaneByIndex(0)
    if err != nil {
        log.Fatal(err)
    }

    // Split the pane and execute commands
    err = pane.Split()
    if err != nil {
        log.Fatal(err)
    }

    err = pane.SendKeys("cd /home/user/projects && ls -la")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Documentation

- **API Reference**: [pkg.go.dev](https://pkg.go.dev/github.com/GianlucaP106/gotmux/gotmux)
- **Examples**: [examples directory](https://github.com/GianlucaP106/gotmux/tree/main/examples)
- **Source Code**: [GitHub repository](https://github.com/GianlucaP106/gotmux)

## Contributing

Contributions are welcome! The project aims to provide complete tmux functionality through Go. Please feel free to:

1. Report bugs or request features
2. Submit pull requests
3. Improve documentation
4. Add new examples

## Development Status

gotmux is actively developed and aims to provide complete tmux functionality. While not all tmux features are implemented yet, the core functionality is stable and ready for production use.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
