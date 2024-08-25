// Copyright (c) 2024 Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/GianlucaP106/gotmux/gotmux/vars"
)

// Tmux window object.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#Variable
type Window struct {
	Active             bool
	ActiveClients      int
	ActiveClientsList  []string
	ActiveSessions     int
	ActiveSessionsList []string
	Activity           string
	ActivityFlag       bool
	BellFlag           bool
	Bigger             bool
	CellHeight         int
	CellWidth          int
	EndFlag            bool
	Flags              string
	Format             bool
	Height             int
	Id                 string
	Index              int
	LastFlag           bool
	Layout             string
	Linked             bool
	LinkedSessions     int
	LinkedSessionsList []string
	MarkedFlag         bool
	Name               string
	OffsetX            int
	OffsetY            int
	Panes              int
	RawFlags           string
	SilenceFlag        int
	StackIndex         int
	StartFlag          bool
	VisibleLayout      string
	Width              int
	ZoomedFlag         bool

	tmux *Tmux
}

// Window layout.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#WINDOWS_AND_PANES
type WindowLayout string

// Enumeration of window layouts.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#WINDOWS_AND_PANES
const (
	WindowLayoutEvenHorizontal WindowLayout = "even-horizontal"
	WindowLayoutEvenVertical   WindowLayout = "even-vertical"
	WindowLayoutMainVertical   WindowLayout = "main-horizontal"
	WindowLayoutTiled          WindowLayout = "tiled"
)

// List panes for this window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#list-panes
func (w *Window) ListPanes() ([]*Pane, error) {
	o, err := w.tmux.query().
		cmd("list-panes").
		fargs("-t", w.Id).
		paneVars().
		run()
	if err != nil {
		return nil, errors.New("failed to list panes")
	}

	out := make([]*Pane, 0)
	for _, item := range o.collect() {
		pane := item.toPane(w.tmux)
		out = append(out, pane)
	}

	return out, nil
}

// Kills the window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#kill-window
func (w *Window) Kill() error {
	_, err := w.tmux.query().
		cmd("kill-window").
		fargs("-t", w.Id).
		run()
	if err != nil {
		return errors.New("failed to kill window")
	}

	return nil
}

// Renames a window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#rename-window
func (w *Window) Rename(newName string) error {
	_, err := w.tmux.query().
		cmd("rename-window").
		fargs("-t", w.Id).
		pargs(newName).
		run()
	if err != nil {
		return errors.New("failed to rename window")
	}

	return nil
}

// Selects this window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#select-window
func (w *Window) Select() error {
	_, err := w.tmux.query().
		cmd("select-window").
		fargs("-t", w.Id).
		run()
	if err != nil {
		return errors.New("failed to select window")
	}

	return nil
}

// Selects the layout for this window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#select-layout
func (w *Window) SelectLayout(layout WindowLayout) error {
	_, err := w.tmux.query().
		cmd("select-layout").
		fargs("-t", w.Id).
		pargs(string(layout)).
		run()
	if err != nil {
		return errors.New("failed to select layout")
	}

	return nil
}

// Move this window to a different location.
// This will return an error if the window already exists.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#move-window
func (w *Window) Move(targetSession string, targetIdx int) error {
	_, err := w.tmux.query().
		cmd("move-window").
		fargs("-s", w.Id).
		fargs("-t", fmt.Sprintf("%s:%d", targetSession, targetIdx)).
		run()
	if err != nil {
		return errors.New("failed to move window")
	}

	return nil
}

// Gets a pane by index in this window.
func (w *Window) GetPaneByIndex(idx int) (*Pane, error) {
	panes, err := w.ListPanes()
	if err != nil {
		return nil, errors.New("failed to get pane by index")
	}

	for _, p := range panes {
		if p.Index == idx {
			return p, nil
		}
	}

	return nil, nil
}

// Lists the sessions linked to this window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#window_linked_sessions_list
func (w *Window) ListLinkedSessions() ([]*Session, error) {
	out := make([]*Session, 0)
	for _, s := range w.LinkedSessionsList {
		session, err := w.tmux.GetSessionByName(s)
		if err != nil {
			return nil, err
		}

		out = append(out, session)
	}

	return out, nil
}

// List the sessions on which this window is active.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#window_active_sessions_list
func (w *Window) ListActiveSessions() ([]*Session, error) {
	out := make([]*Session, 0)
	for _, s := range w.ActiveSessionsList {
		session, err := w.tmux.GetSessionByName(s)
		if err != nil {
			return nil, err
		}

		out = append(out, session)
	}

	return out, nil
}

// List the clients viewing this window.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#window_active_clients_list
func (w *Window) ListActiveClients() ([]*Client, error) {
	out := make([]*Client, 0)
	for _, c := range w.ActiveClientsList {
		client, err := w.tmux.GetClientByTty(c)
		if err != nil {
			return nil, err
		}

		out = append(out, client)
	}

	return out, nil
}

// Sets the window variables in the query.
func (q *query) windowVars() *query {
	return q.vars(
		vars.WindowActive,
		vars.WindowActiveClients,
		vars.WindowActiveClientsList,
		vars.WindowActiveSessions,
		vars.WindowActiveSessionsList,
		vars.WindowActivity,
		vars.WindowActivityFlag,
		vars.WindowBellFlag,
		vars.WindowBigger,
		vars.WindowCellHeight,
		vars.WindowCellWidth,
		vars.WindowEndFlag,
		vars.WindowFlags,
		vars.WindowFormat,
		vars.WindowHeight,
		vars.WindowId,
		vars.WindowIndex,
		vars.WindowLastFlag,
		vars.WindowLayout,
		vars.WindowLinked,
		vars.WindowLinkedSessions,
		vars.WindowLinkedSessionsList,
		vars.WindowMarkedFlag,
		vars.WindowName,
		vars.WindowOffsetX,
		vars.WindowOffsetY,
		vars.WindowPanes,
		vars.WindowRawFlags,
		vars.WindowSilenceFlag,
		vars.WindowStackIndex,
		vars.WindowStartFlag,
		vars.WindowVisibleLayout,
		vars.WindowWidth,
		vars.WindowZoomedFlag,
	)
}

// Converts a QueryResult to a Window.
func (q queryResult) toWindow(t *Tmux) *Window {
	active := isOne(q.get(vars.WindowActive))
	activeClients, _ := strconv.Atoi(q.get(vars.WindowActiveClients))
	activeClientsList := parseList(q.get(vars.WindowActiveClientsList))
	activeSessions, _ := strconv.Atoi(q.get(vars.WindowActiveSessions))
	activeSessionsList := parseList(q.get(vars.WindowActiveSessionsList))
	activity := q.get(vars.WindowActivity)
	activityFlag := isOne(q.get(vars.WindowActivityFlag))
	bellFlag := isOne(q.get(vars.WindowBellFlag))
	bigger := isOne(q.get(vars.WindowBigger))
	cellHeight, _ := strconv.Atoi(q.get(vars.WindowCellHeight))
	cellWidth, _ := strconv.Atoi(q.get(vars.WindowCellWidth))
	endFlag := isOne(q.get(vars.WindowEndFlag))
	flags := q.get(vars.WindowFlags)
	format := isOne(q.get(vars.WindowFormat))
	height, _ := strconv.Atoi(q.get(vars.WindowHeight))
	id := q.get(vars.WindowId)
	index, _ := strconv.Atoi(q.get(vars.WindowIndex))
	lastFlag := isOne(q.get(vars.WindowLastFlag))
	layout := q.get(vars.WindowLayout)
	linked := isOne(q.get(vars.WindowLinked))
	linkedSessions, _ := strconv.Atoi(q.get(vars.WindowLinkedSessions))
	linkedSessionsList := parseList(q.get(vars.WindowLinkedSessionsList))
	markedFlag := isOne(q.get(vars.WindowMarkedFlag))
	name := q.get(vars.WindowName)
	offsetX, _ := strconv.Atoi(q.get(vars.WindowOffsetX))
	offsetY, _ := strconv.Atoi(q.get(vars.WindowOffsetY))
	panes, _ := strconv.Atoi(q.get(vars.WindowPanes))
	rawFlags := q.get(vars.WindowRawFlags)
	silenceFlag, _ := strconv.Atoi(q.get(vars.WindowSilenceFlag))
	stackIndex, _ := strconv.Atoi(q.get(vars.WindowStackIndex))
	startFlag := isOne(q.get(vars.WindowStartFlag))
	visibleLayout := q.get(vars.WindowVisibleLayout)
	width, _ := strconv.Atoi(q.get(vars.WindowWidth))
	zoomedFlag := isOne(q.get(vars.WindowZoomedFlag))

	w := &Window{
		Active:             active,
		ActiveClients:      activeClients,
		ActiveClientsList:  activeClientsList,
		ActiveSessions:     activeSessions,
		ActiveSessionsList: activeSessionsList,
		Activity:           activity,
		ActivityFlag:       activityFlag,
		BellFlag:           bellFlag,
		Bigger:             bigger,
		CellHeight:         cellHeight,
		CellWidth:          cellWidth,
		EndFlag:            endFlag,
		Flags:              flags,
		Format:             format,
		Height:             height,
		Id:                 id,
		Index:              index,
		LastFlag:           lastFlag,
		Layout:             layout,
		Linked:             linked,
		LinkedSessions:     linkedSessions,
		LinkedSessionsList: linkedSessionsList,
		MarkedFlag:         markedFlag,
		Name:               name,
		OffsetX:            offsetX,
		OffsetY:            offsetY,
		Panes:              panes,
		RawFlags:           rawFlags,
		SilenceFlag:        silenceFlag,
		StackIndex:         stackIndex,
		StartFlag:          startFlag,
		VisibleLayout:      visibleLayout,
		Width:              width,
		ZoomedFlag:         zoomedFlag,

		tmux: t,
	}

	return w
}
