# gotmux 🚀

A **powerful and comprehensive** Go library for interacting with `tmux`, offering a clean and complete interface to manage sessions, windows, and panes efficiently via the `tmux` CLI.

## 📌 Requirements

- ✅ **tmux** installed on the system

## 📦 Installation

Install `gotmux` using:

```bash
go get github.com/GianlucaP106/gotmux
```

## 🎯 Features

`gotmux` provides a **rich** Go interface to `tmux`, covering:

### 🔹 Session Management

- Create, rename, and delete sessions
- Attach/detach from sessions
- List and query session details
- Group sessions together
- Monitor session activity

### 🔹 Window Operations

- Create and manage windows
- Move windows between sessions
- Set and customize layouts
- Navigate between windows
- Track window activity

### 🔹 Pane Control

- Split panes **horizontally** or **vertically**
- Resize and rearrange panes
- Capture pane content
- Sync panes
- Execute commands within panes

### 🔹 Server & Client Info

- Query server status and version
- List connected clients
- Retrieve terminal information
- Monitor client activity
- Manage sockets

### 📊 All Data is Returned

gotmux returns all the fields for sessions, windows, panes and more. For example, this is the session type:

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
}
```

## 🚀 Implementation Status

`gotmux` aims to be **feature-complete** with `tmux`. Not all features are implemented yet, but contributions are welcome! 🤝

## 📖 Documentation

- 📚 [pkg.go.dev](https://pkg.go.dev/github.com/GianlucaP106/gotmux/gotmux)
- 💡 [Examples Directory](https://github.com/GianlucaP106/gotmux/tree/main/examples)

---

## 🛠️ Usage

### ⚡ Quick Start

```go
import (
    "fmt"
    "log"
    "github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    session, err := tmux.New()
    if err != nil {
        log.Fatal(err)
    }

    err = session.Attach()
    if err != nil {
        log.Fatal(err)
    }
}
```

### 🖥️ Advanced Examples

#### 🎛️ Session Management

```go
func main() {
    tmux, err := gotmux.DefaultTmux()
    if err != nil {
        log.Fatal(err)
    }

    session, err := tmux.NewSession(&gotmux.SessionOptions{
        StartDirectory: "/home",
        Name:          "somename",
    })
    if err != nil {
        log.Fatal(err)
    }

    err = session.Attach()
    if err != nil {
        log.Fatal(err)
    }
}
```

#### 🪟 Window & Pane Management

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

#### 📋 Listing Clients & Server Info

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
        fmt.Println("Client:", c.Tty, "Session:", c.Session)
    }

    server, err := tmux.GetServerInformation()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("tmux version:", server.Version)
}
```

#### 🔄 Window Layout Management

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

    err = window.SelectLayout(gotmux.WindowLayoutEvenVertical)
    if err != nil {
        log.Fatal(err)
    }
}
```

👉 See the complete list of **[examples](https://github.com/GianlucaP106/gotmux/tree/main/examples)**.

---

## 📜 License

This project is licensed under the **MIT License** – see the [LICENSE](LICENSE) file for details.

---

🚀 **Start managing `tmux` with Go today!**
