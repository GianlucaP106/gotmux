# gotmux

Gotmux is a library to interact with tmux. It communicates using the tmux CLI. It aims to be a completed library for tmux.

# Installation

```bash
go get github.com/GianlucaP106/gotmux
```

# Basic usage

```go
import (
    "fmt"
    "log"

    "github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
    tmux := gotmux.DefaultTmux()
    sessions, err := tmux.ListSessions()
    if err != nil {
        log.Fatal(err)
    }

    for _, s := range sessions {
        fmt.Println(s.Name)
    }
}

```
