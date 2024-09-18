# gotmux

A comprehensive library designed to interact with tmux, communicating through the tmux CLI and aiming to offer complete functionality for tmux integration.

## Installation

```bash
go get github.com/GianlucaP106/gotmux
```

## Dependencies

- tmux

## Features

This library aims to be feature complete with tmux. Currently not all features are supported but they are planned to be impelemented. Contributions are also welcome.

## Documentation

- [pkg.go.dev](https://pkg.go.dev/github.com/GianlucaP106/gotmux/gotmux)

## Usage

### Basic example

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

#### See the list of [examples](https://github.com/GianlucaP106/gotmux/tree/main/examples)

# Features

This library aims to support all tmux feature, but currently not everything is impelemented.

### Impelemented features

#### Full session, window and pane management (with helper methods)

Create, update, delete and view sessions, windows and panes in an object oriented fashion. For example:

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

#### View all data

The returned objects contain ALL attributes provided by tmux. For example, this is the session type:

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

You can refer to the tmux documentation to fully understand what these attributes represent (<https://man.openbsd.org/OpenBSD-current/man1/tmux.1#Variable>).

#### Get server and client information

View information about the tmux server and active clients (terminals). For example:

- Clients

```go
func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    clients, err := tmux.ListClients()
    if err != nil {
        log.Fatal(err)
    }

    for _, c := range clients {
        fmt.Println(c.Tty, c.Session)
    }
}

```

- Server

```go
func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    server, err := tmux.GetServerInformation()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(server.Version)
}
```
