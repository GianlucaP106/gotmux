# gotmux

Gotmux is a library to interact with tmux. It communicates using the tmux CLI. It aims to be a completed library for tmux.

## Installation

```bash
go get github.com/GianlucaP106/gotmux
```

## Dependencies

- tmux

## Features

This library aims to be feature complete with tmux. Currently not all features are supported but they are planned to be impelemented. Contributions are also welcome, see CONTRIBUTING.

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

## Data snapshots

The objects returned (Session, Window, Pane...), contain data attributes and methods. Some methods are mutating. These methods do not update themselves or other generated objects, the data in the objects are snapshots from when they were first fetched. Therefore their data will be outdated as soon as a mutating action occurs (through the library or in through other clients). This may be something to consider while building your applications.
