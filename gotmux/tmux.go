// Copyright (c) 2024 Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Entrypoint object to the library.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#DESCRIPTION
type Tmux struct {
	Socket *Socket
}

// Initializes the tmux client with a socket path.
// Entry point to the library.
func NewTmux(socketPath string) (*Tmux, error) {
	if !isTmuxInstalled() {
		return nil, errors.New("tmux is not installed on the system")
	}
	t := &Tmux{}
	s, err := newSocket(socketPath)
	if err != nil {
		return nil, err
	}
	t.Socket = s

	return t, nil
}

// Initializes the tmux client with default socket.
// Entry point to the library.
func DefaultTmux() (*Tmux, error) {
	if !isTmuxInstalled() {
		return nil, errors.New("tmux is not installed on the system")
	}
	return &Tmux{
		Socket: nil,
	}, nil
}

// Get server information.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#Variable
func (t *Tmux) GetServerInformation() (*Server, error) {
	o, err := t.query().
		cmd("display-message").
		serverVars().
		run()
	if err != nil {
		return nil, err
	}

	server := o.one().toServer(t)
	return server, nil
}

// List all clients.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#list-clients
func (t *Tmux) ListClients() ([]*Client, error) {
	output, err := t.query().
		cmd("list-clients").
		clientVars().
		run()
	if err != nil {
		return nil, errors.New("failed to list clients")
	}

	result := output.collect()
	out := make([]*Client, 0)
	for _, item := range result {
		c := item.toClient(t)
		out = append(out, c)
	}

	return out, nil
}

// List sessions.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#list-sessions
func (t *Tmux) ListSessions() ([]*Session, error) {
	output, err := t.query().
		cmd("list-sessions").
		sessionVars().
		run()
	if err != nil {
		return nil, errors.New("failed to list sessions")
	}

	result := output.collect()
	out := make([]*Session, 0)
	for _, item := range result {
		s := item.toSession(t)
		out = append(out, s)
	}

	return out, nil
}

// Returns true if the session exists, false otherwise.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#has-session
func (t *Tmux) HasSession(session string) bool {
	_, err := t.query().
		cmd("has-session").
		fargs("-t", session).
		run()

	return err == nil
}

// Gets a session by name.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#session_name
func (t *Tmux) GetSessionByName(name string) (*Session, error) {
	sessions, err := t.ListSessions()
	if err != nil {
		return nil, errors.New("failed to get session by name")
	}

	for _, s := range sessions {
		if s.Name == name {
			return s, nil
		}
	}

	return nil, nil
}

// Gets a client by tty.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#client_tty
func (t *Tmux) GetClientByTty(tty string) (*Client, error) {
	clients, err := t.ListClients()
	if err != nil {
		return nil, err
	}

	for _, c := range clients {
		if c.Tty == tty {
			return c, nil
		}
	}

	return nil, nil
}

// Options object for creating a session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#new-session
type SessionOptions struct {
	Name           string
	ShellCommand   string
	StartDirectory string
	Width          int
	Height         int
}

// Creates a new session without attaching to it.
// Pass nil to create a session with default options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#new-session
func (t *Tmux) NewSession(op *SessionOptions) (*Session, error) {
	q := t.query().
		cmd("new-session").
		fargs("-d", "-P").
		sessionVars()

	if op != nil {
		if op.Name != "" {
			if !checkSessionName(op.Name) {
				return nil, errors.New("invalid tmux session name")
			}

			q.fargs("-s", op.Name)
		}

		if op.StartDirectory != "" {
			q.fargs("-c", op.StartDirectory)
		}

		if op.Width != 0 {
			w := strconv.Itoa(op.Width)
			q.fargs("-x", w)
		}

		if op.Height != 0 {
			h := strconv.Itoa(op.Height)
			q.fargs("-y", h)
		}

		if op.ShellCommand != "" {
			s := fmt.Sprintf("'%s'", op.ShellCommand)
			q.pargs(s)
		}
	}

	o, err := q.run()
	if err != nil {
		return nil, errors.New("failed to create session")
	}

	s := o.one().toSession(t)
	return s, nil
}

// Creates a new session without attaching.
// Shorthand for 'NewSession', but with default options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#new-session
func (w *Tmux) New() (*Session, error) {
	return w.NewSession(nil)
}

// Options object for detaching a session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#detach-client
type DetachClientOptions struct {
	TargetClient  string
	TargetSession string
}

// Detaches current client, a target client or all the clients of a target session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#detach-client
func (t *Tmux) DetachClient(op *DetachClientOptions) error {
	q := t.query().
		cmd("detach-client")

	if op != nil {
		if op.TargetClient != "" {
			q.fargs("-t", op.TargetClient)
		} else if op.TargetSession != "" {
			q.fargs("-s", op.TargetSession)
		}
	}

	_, err := q.run()
	if err != nil {
		return errors.New("failed to detach client")
	}

	return nil
}

// Kills the server. Kills all clients and servers.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#kill-server
func (t *Tmux) KillServer() error {
	_, err := t.query().cmd("kill-server").run()
	if err != nil {
		return errors.New("failed to kill server")
	}

	return nil
}

// Lists all windows.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#list-windows
func (t *Tmux) ListAllWindows() ([]*Window, error) {
	o, err := t.query().
		cmd("list-windows").
		fargs("-a").
		windowVars().
		run()
	if err != nil {
		return nil, errors.New("failed to list all windows")
	}

	out := make([]*Window, 0)
	for _, res := range o.collect() {
		w := res.toWindow(t)
		out = append(out, w)
	}

	return out, nil
}

// Lists all panes in the server.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#list-panes
func (t *Tmux) ListAllPanes() ([]*Pane, error) {
	o, err := t.query().
		cmd("list-panes").
		fargs("-a").
		paneVars().
		run()
	if err != nil {
		return nil, errors.New("failed to list all panes")
	}

	out := make([]*Pane, 0)
	for _, r := range o.collect() {
		p := r.toPane(t)
		out = append(out, p)
	}

	return out, nil
}

// Returns the window with the given Id.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#window_id
func (t *Tmux) GetWindowById(id string) (*Window, error) {
	windows, err := t.ListAllWindows()
	if err != nil {
		return nil, err
	}

	for _, w := range windows {
		if w.Id == id {
			return w, nil
		}
	}

	return nil, nil
}

// Returns the pane with the given Id.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#pane_id
func (t *Tmux) GetPaneById(id string) (*Pane, error) {
	panes, err := t.ListAllPanes()
	if err != nil {
		return nil, err
	}

	for _, p := range panes {
		if p.Id == id {
			return p, nil
		}
	}

	return nil, nil
}

// Adds socket argument.
func (t *Tmux) query() *query {
	q := newQuery()
	if t.Socket != nil {
		q.cmd("-S", t.Socket.Path)
	}
	return q
}

// Checks if a string is 1.
func isOne(s string) bool {
	return s == "1"
}

// Splits a string by comma.
func parseList(l string) []string {
	return strings.Split(l, ",")
}

// Checks the validity of the tmux session name.
func checkSessionName(name string) bool {
	if len(name) == 0 {
		return false
	}

	if strings.Contains(name, ":") {
		return false
	}

	if strings.Contains(name, ".") {
		return false
	}

	return true
}

// Returns true if tmux is installed on the system otherwise false.
func isTmuxInstalled() bool {
	_, err := exec.LookPath("tmux")
	return err == nil
}
