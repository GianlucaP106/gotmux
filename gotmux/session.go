// Copyright (c) Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"errors"
	"io"
	"strconv"
)

// Tmux session object.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#Variable
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

// List clients attached to this session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#list-clients
func (s *Session) ListClients() ([]*Client, error) {
	clients, err := s.tmux.ListClients()
	if err != nil {
		return nil, err
	}

	out := make([]*Client, 0)
	for _, c := range clients {
		if c.Session == s.Name {
			out = append(out, c)
		}
	}

	return out, nil
}

// Attach session options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#attach-session
type AttachSessionOptions struct {
	WorkingDir    string
	DetachClients bool

	// Stdout/Stderr will be redirected to these writers if set.
	Output, Error io.Writer
}

// Attaches the current client to the session.
// Since this requires a client, it will attach to the terminal by redirecting.
// Provides the option to pipe output and error to a custom location.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#attach-session
func (s *Session) AttachSession(op *AttachSessionOptions) error {
	q := s.tmux.query().
		cmd("attach-session").
		fargs("-t", s.Name)

	if op != nil {
		if op.DetachClients {
			q.fargs("-d")
		}

		if op.WorkingDir != "" {
			q.fargs("-c", op.WorkingDir)
		}

		q.pipeOut(op.Output)
		q.pipeErr(op.Error)
	}

	err := q.runTty()
	if err != nil {
		return errors.New("failed to attach session")
	}

	return nil
}

// Attaches the current client to the session.
// Since this requires a client, it will attach to the terminal by redirecting.
// Shorthand for 'AttachSession', but with default configuration.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#attach-session
func (s *Session) Attach() error {
	return s.AttachSession(nil)
}

// Detaches all clients attached to this session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#detach-client
func (s *Session) Detach() error {
	_, err := s.tmux.query().
		cmd("detach-client").
		fargs("-s", s.Name).run()
	if err != nil {
		return errors.New("failed to detach session")
	}

	return nil
}

// Kills the session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#kill-session
func (s *Session) Kill() error {
	_, err := s.tmux.query().
		cmd("kill-session").
		fargs("-t", s.Name).
		run()
	if err != nil {
		return errors.New("failed to kill session")
	}

	return nil
}

// Renames the session to a new name.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#rename-session
func (s *Session) Rename(name string) error {
	_, err := s.tmux.query().
		cmd("rename-session").
		fargs("-t", s.Name).
		pargs(name).
		run()
	if err != nil {
		return errors.New("failed to rename session")
	}

	return nil
}

// Lists all windows for this session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#list-windows
func (s *Session) ListWindows() ([]*Window, error) {
	o, err := s.tmux.query().
		cmd("list-windows").
		fargs("-t", s.Name).
		windowVars().
		run()
	if err != nil {
		return nil, errors.New("failed to list windows")
	}

	qr := o.collect()
	out := make([]*Window, 0)
	for _, item := range qr {
		w := item.toWindow(s.tmux)
		out = append(out, w)
	}

	return out, nil
}

// List panes for this session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#list-panes
func (s *Session) ListPanes() ([]*Pane, error) {
	o, err := s.tmux.query().
		cmd("list-panes").
		fargs("-s", "-t", s.Name).
		paneVars().
		run()
	if err != nil {
		return nil, errors.New("failed to list panes")
	}

	out := make([]*Pane, 0)
	for _, item := range o.collect() {
		pane := item.toPane(s.tmux)
		out = append(out, pane)
	}

	return out, nil
}

// Gets a window by index in this session.
func (s *Session) GetWindowByIndex(idx int) (*Window, error) {
	windows, err := s.ListWindows()
	if err != nil {
		return nil, errors.New("failed to get window by idx")
	}

	for _, w := range windows {
		if w.Index == idx {
			return w, nil
		}
	}

	return nil, nil
}

// New window options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#new-window
type NewWindowOptions struct {
	StartDirectory string
	WindowName     string
	DoNotAttach    bool
}

// Creates a new window in this session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#new-window
func (s *Session) NewWindow(op *NewWindowOptions) (*Window, error) {
	q := s.tmux.query().
		cmd("new-window").
		fargs("-P", "-t", s.Name).
		windowVars()

	if op != nil {
		if op.StartDirectory != "" {
			q.fargs("-c", op.StartDirectory)
		}

		if op.WindowName != "" {
			q.fargs("-n", op.WindowName)
		}

		if op.DoNotAttach {
			q.fargs("-d")
		}
	}

	o, err := q.run()
	if err != nil {
		return nil, errors.New("failed to create window")
	}

	w := o.one().toWindow(s.tmux)
	return w, nil
}

// Creates a new window in this session.
// Shorthand for 'NewWindow', but with default options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#new-window
func (s *Session) New() (*Window, error) {
	return s.NewWindow(nil)
}

// Selects the next window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#next-window
func (s *Session) NextWindow() error {
	q := s.tmux.query().
		cmd("next-window").
		fargs("-t", s.Name)

	_, err := q.run()
	if err != nil {
		return errors.New("failed to select next window")
	}

	return nil
}

// Selects the previous window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#previous-window
func (s *Session) PreviousWindow() error {
	q := s.tmux.query().
		cmd("previous-window").
		fargs("-t", s.Name)

	_, err := q.run()
	if err != nil {
		return errors.New("failed to select the previous window")
	}

	return nil
}

// Sets an option with a given key.
// Note that custom options must begin with '@'.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#set-option
func (s *Session) SetOption(key, option string) error {
	return s.tmux.SetOption(s.Name, key, option, "")
}

// Retrieves an option from this session.
//
// https://man.openbsd.org/OpenBSD-current/man1/tmux.1#show-options
func (s *Session) Option(key string) (*Option, error) {
	return s.tmux.Option(s.Name, key, "")
}

// Retrieves all options in this session.
//
// https://man.openbsd.org/OpenBSD-current/man1/tmux.1#show-options
func (s *Session) Options() ([]*Option, error) {
	return s.tmux.Options(s.Name, "")
}

// Deletes an option from this session.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#set-option
func (s *Session) DeleteOption(key string) error {
	return s.tmux.DeleteOption(s.Name, key, "")
}

// Sets the session variables in the query.
func (q *query) sessionVars() *query {
	return q.vars(
		varSessionActivity,
		varSessionAlerts,
		varSessionAttached,
		varSessionAttachedList,
		varSessionCreated,
		varSessionFormat,
		varSessionGroup,
		varSessionGroupAttached,
		varSessionGroupAttachedList,
		varSessionGroupList,
		varSessionGroupManyAttached,
		varSessionGroupSize,
		varSessionGrouped,
		varSessionId,
		varSessionLastAttached,
		varSessionManyAttached,
		varSessionMarked,
		varSessionName,
		varSessionPath,
		varSessionStack,
		varSessionWindows,
	)
}

// Converts a QueryResult to a Session.
func (q queryResult) toSession(t *Tmux) *Session {
	activity := q.get(varSessionActivity)
	alerts := q.get(varSessionAlerts)
	attached, _ := strconv.Atoi(q.get(varSessionAttached))
	attachedList := parseList(q.get(varSessionAttachedList))
	created := q.get(varSessionCreated)
	format := isOne(q.get(varSessionFormat))
	group := q.get(varSessionGroup)
	groupAttached, _ := strconv.Atoi(q.get(varSessionGroupAttached))
	groupAttachedList := parseList(q.get(varSessionGroupAttachedList))
	groupList := parseList(q.get(varSessionGroupList))
	groupManyAttached := isOne(q.get(varSessionGroupManyAttached))
	groupSize, _ := strconv.Atoi(q.get(varSessionGroupSize))
	grouped := isOne(q.get(varSessionGrouped))
	id := q.get(varSessionId)
	lastAttached := q.get(varSessionLastAttached)
	manyAttached := isOne(q.get(varSessionManyAttached))
	marked := isOne(q.get(varSessionMarked))
	name := q.get(varSessionName)
	path := q.get(varSessionPath)
	stack := q.get(varSessionStack)
	windows, _ := strconv.Atoi(q.get(varSessionWindows))

	s := &Session{
		Activity:          activity,
		Alerts:            alerts,
		Attached:          attached,
		AttachedList:      attachedList,
		Created:           created,
		Format:            format,
		Group:             group,
		GroupAttached:     groupAttached,
		GroupAttachedList: groupAttachedList,
		GroupList:         groupList,
		GroupManyAttached: groupManyAttached,
		GroupSize:         groupSize,
		Grouped:           grouped,
		Id:                id,
		LastAttached:      lastAttached,
		ManyAttached:      manyAttached,
		Marked:            marked,
		Name:              name,
		Path:              path,
		Stack:             stack,
		Windows:           windows,

		tmux: t,
	}

	return s
}
