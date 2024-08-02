// Copyright (c) 2024 Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"errors"
	"strconv"

	"github.com/GianlucaP106/gotmux/gotmux/vars"
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
}

// Attaches the current client to the session.
// Since this requires a client, it will attach to the terminal by redirecting.
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
	}

	err := q.runTty()
	if err != nil {
		return errors.New("failed to attach session")
	}

	return nil
}

// Attaches the current client to the session.
// Since this requires a client, it will attach to the terminal by redirecting.
// Sorthand for 'AttachSession', but runs with default configuration.
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

	s.Name = name
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

// Selects the next window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#next-window
func (s *Session) NextWindow() error {
	_, err := s.tmux.query().
		cmd("next-window").
		fargs("-t", s.Name).
		run()
	if err != nil {
		return errors.New("failed to select next window")
	}

	return nil
}

// Selects the previous window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#previous-window
func (s *Session) PreviousWindow() error {
	_, err := s.tmux.query().
		cmd("previous-window").
		fargs("-t", s.Name).
		run()
	if err != nil {
		return errors.New("failed to select the previous window")
	}

	return nil
}

// Sets the session variables in the query.
func (q *query) sessionVars() *query {
	return q.vars(
		vars.SessionActivity,
		vars.SessionAlerts,
		vars.SessionAttached,
		vars.SessionAttachedList,
		vars.SessionCreated,
		vars.SessionFormat,
		vars.SessionGroup,
		vars.SessionGroupAttached,
		vars.SessionGroupAttachedList,
		vars.SessionGroupList,
		vars.SessionGroupManyAttached,
		vars.SessionGroupSize,
		vars.SessionGrouped,
		vars.SessionId,
		vars.SessionLastAttached,
		vars.SessionManyAttached,
		vars.SessionMarked,
		vars.SessionName,
		vars.SessionPath,
		vars.SessionStack,
		vars.SessionWindows,
	)
}

// Converts a QueryResult to a Session.
func (q queryResult) toSession(t *Tmux) *Session {
	activity := q.get(vars.SessionActivity)
	alerts := q.get(vars.SessionAlerts)
	attached, _ := strconv.Atoi(q.get(vars.SessionAttached))
	attachedList := parseList(q.get(vars.SessionAttachedList))
	created := q.get(vars.SessionCreated)
	format := isOne(q.get(vars.SessionFormat))
	group := q.get(vars.SessionGroup)
	groupAttached, _ := strconv.Atoi(q.get(vars.SessionGroupAttached))
	groupAttachedList := parseList(q.get(vars.SessionGroupAttachedList))
	groupList := parseList(q.get(vars.SessionGroupList))
	groupManyAttached := isOne(q.get(vars.SessionGroupManyAttached))
	groupSize, _ := strconv.Atoi(q.get(vars.SessionGroupSize))
	grouped := isOne(q.get(vars.SessionGrouped))
	id := q.get(vars.SessionId)
	lastAttached := q.get(vars.SessionLastAttached)
	manyAttached := isOne(q.get(vars.SessionManyAttached))
	marked := isOne(q.get(vars.SessionMarked))
	name := q.get(vars.SessionName)
	path := q.get(vars.SessionPath)
	stack := q.get(vars.SessionStack)
	windows, _ := strconv.Atoi(q.get(vars.SessionWindows))

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
