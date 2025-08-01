// Copyright (c) Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"errors"
	"fmt"
	"strconv"
)

// Tmux pane object.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#Variable
type Pane struct {
	Active         bool
	AtBottom       bool
	AtLeft         bool
	AtRight        bool
	AtTop          bool
	Bg             string
	Bottom         string
	CurrentCommand string
	CurrentPath    string
	Dead           bool
	DeadSignal     int
	DeadStatus     int
	DeadTime       string
	Fg             string
	Format         bool
	Height         int
	Id             string
	InMode         bool
	Index          int
	InputOff       bool
	Last           bool
	Left           string
	Marked         bool
	MarkedSet      bool
	Mode           string
	Path           string
	Pid            int32
	Pipe           bool
	Right          string
	SearchString   string
	SessionName    string
	StartCommand   string
	StartPath      string
	Synchronized   bool
	Tabs           string
	Title          string
	Top            string
	Tty            string
	UnseenChanges  bool
	Width          int
	WindowIndex    int

	tmux *Tmux
}

// Pane relative position.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#select-pane
type PanePosition string

// Enumeration of pane positions.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#select-pane
const (
	PanePositionUp    PanePosition = "-U"
	PanePositionRight PanePosition = "-R"
	PanePositionDown  PanePosition = "-D"
	PanePositionLeft  PanePosition = "-L"
)

// Pane split direction.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#split-window
type PaneSplitDirection string

// Enumeration of pane split directions.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#split-window
const (
	PaneSplitDirectionHorizontal PaneSplitDirection = "-h"
	PaneSplitDirectionVertical   PaneSplitDirection = "-v"
)

// Pane send-keys.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#send-keys
func (p *Pane) SendKeys(line string) error {
	_, err := p.tmux.query().
		cmd("send-keys").
		fargs("-t", p.Id).
		pargs(line).
		run()
	if err != nil {
		return errors.New("failed to send keys")
	}

	return nil
}

// Kills the pane.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#kill-pane
func (p *Pane) Kill() error {
	_, err := p.tmux.query().
		cmd("kill-pane").
		fargs("-t", p.Id).
		run()
	if err != nil {
		return errors.New("failed to kill pane")
	}

	return nil
}

// Options for select pane.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#select-pane
type SelectPaneOptions struct {
	TargetPosition PanePosition
}

// Selects the pane with the provided options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#select-pane
func (p *Pane) SelectPane(op *SelectPaneOptions) error {
	q := p.tmux.query().
		cmd("select-pane").
		fargs("-t", p.Id)

	if op != nil {
		if op.TargetPosition != "" {
			q.fargs(string(op.TargetPosition))
		}
	}

	_, err := q.run()
	if err != nil {
		return errors.New("failed to select pane")
	}

	return nil
}

// Selects the pane with the provided options.
// Shorthand 'SelectPane' but with default options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#select-pane
func (p *Pane) Select() error {
	return p.SelectPane(nil)
}

// Split window options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#split-window
type SplitWindowOptions struct {
	SplitDirection PaneSplitDirection
	StartDirectory string
	ShellCommand   string
}

// Split the window (pane).
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#split-window
func (p *Pane) SplitWindow(op *SplitWindowOptions) error {
	q := p.tmux.query().
		cmd("split-window").
		fargs("-t", p.Id)

	if op != nil {
		if op.SplitDirection != "" {
			q.fargs(string(op.SplitDirection))
		}

		if op.StartDirectory != "" {
			q.fargs("-c", op.StartDirectory)
		}

		if op.ShellCommand != "" {
			q.pargs(fmt.Sprintf("'%s'", op.ShellCommand))
		}
	}

	_, err := q.run()
	if err != nil {
		return errors.New("failed to split pane")
	}

	return nil
}

// Split the window (pane).
// Shorthand for 'SplitWindow' but with default options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#split-window
func (p *Pane) Split() error {
	return p.SplitWindow(nil)
}

// Choose tree options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#choose-tree
type ChooseTreeOptions struct {
	SessionsCollapsed bool
	WindowsCollapsed  bool
}

// Puts the pane in choose tree mode.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#choose-tree
func (p *Pane) ChooseTree(op *ChooseTreeOptions) error {
	q := p.tmux.query().
		cmd("choose-tree").
		fargs("-t", p.Id)

	if op != nil {
		if op.SessionsCollapsed {
			q.fargs("-s")
		}

		if op.WindowsCollapsed {
			q.fargs("-w")
		}
	}

	_, err := q.run()
	if err != nil {
		return errors.New("failed to put the pane in choose tree mode")
	}

	return nil
}

// Capture pane command options.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#capture-pane
type CaptureOptions struct {
	EscTxtNBgAttr    bool
	EscNonPrintables bool
	IgnoreTrailing   bool
	PreserveTrailing bool
	PreserveAndJoin  bool
}

// Captures the content of the pane
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#capture-pane
func (p *Pane) CapturePane(op *CaptureOptions) (string, error) {
	q := p.tmux.query().
		cmd("capture-pane").
		fargs("-t", p.Id).
		fargs("-p")

	if op != nil {
		if op.EscTxtNBgAttr {
			q.fargs("-e")
		}

		if op.EscNonPrintables {
			q.fargs("-C")
		}

		if op.IgnoreTrailing {
			q.fargs("-T")
		}

		if op.PreserveTrailing {
			q.fargs("-N")
		}

		if op.PreserveTrailing {
			q.fargs("-J")
		}
	}

	o, err := q.run()
	if err != nil {
		return "", errors.New("failed to capture pane")
	}

	return o.result, nil
}

// Captures the pane with background and text atrributes escaped.
// Shorthand for `CapturePane`.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#capture-pane
func (p *Pane) Capture() (string, error) {
	return p.CapturePane(&CaptureOptions{
		EscTxtNBgAttr: true,
	})
}

// Sets an option with a given key.
// Note that custom options must begin with '@'.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#set-option
func (p *Pane) SetOption(key, option string) error {
	return p.tmux.SetOption(p.Id, key, option, "-p")
}

// Retrieves an option from this pane.
//
// https://man.openbsd.org/OpenBSD-current/man1/tmux.1#show-options
func (p *Pane) Option(key string) (*Option, error) {
	return p.tmux.Option(p.Id, key, "-p")
}

// Retrieves all options in this pane.
//
// https://man.openbsd.org/OpenBSD-current/man1/tmux.1#show-options
func (p *Pane) Options() ([]*Option, error) {
	return p.tmux.Options(p.Id, "-p")
}

// Deletes an option from this pane.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#set-option
func (p *Pane) DeleteOption(key string) error {
	return p.tmux.DeleteOption(p.Id, key, "-p")
}

// Sets the pane variables in the query.
func (q *query) paneVars() *query {
	return q.vars(
		varPaneActive,
		varPaneAtBottom,
		varPaneAtLeft,
		varPaneAtRight,
		varPaneAtTop,
		varPaneBg,
		varPaneBottom,
		varPaneCurrentCommand,
		varPaneCurrentPath,
		varPaneDead,
		varPaneDeadSignal,
		varPaneDeadStatus,
		varPaneDeadTime,
		varPaneFg,
		varPaneFormat,
		varPaneHeight,
		varPaneId,
		varPaneInMode,
		varPaneIndex,
		varPaneInputOff,
		varPaneLast,
		varPaneLeft,
		varPaneMarked,
		varPaneMarkedSet,
		varPaneMode,
		varPanePath,
		varPanePid,
		varPanePipe,
		varPaneRight,
		varPaneSearchString,
		varPaneSessionName,
		varPaneStartCommand,
		varPaneStartPath,
		varPaneSynchronized,
		varPaneTabs,
		varPaneTitle,
		varPaneTop,
		varPaneTty,
		varPaneUnseenChanges,
		varPaneWidth,
		varPaneWindowIndex,
	)
}

// Converts a QueryResult to a pane.
func (q queryResult) toPane(t *Tmux) *Pane {
	active := isOne(q.get(varPaneActive))
	atBottom := isOne(q.get(varPaneAtBottom))
	atLeft := isOne(q.get(varPaneAtLeft))
	atRight := isOne(q.get(varPaneAtRight))
	atTop := isOne(q.get(varPaneAtTop))
	bg := q.get(varPaneBg)
	bottom := q.get(varPaneBottom)
	currentCommand := q.get(varPaneCurrentCommand)
	currentPath := q.get(varPaneCurrentPath)
	dead := isOne(q.get(varPaneDead))
	deadSignal, _ := strconv.Atoi(q.get(varPaneDeadSignal))
	deadStatus, _ := strconv.Atoi(q.get(varPaneDeadStatus))
	deadTime := q.get(varPaneDeadTime)
	fg := q.get(varPaneFg)
	format := isOne(q.get(varPaneFormat))
	height, _ := strconv.Atoi(q.get(varPaneHeight))
	id := q.get(varPaneId)
	inMode := isOne(q.get(varPaneInMode))
	index, _ := strconv.Atoi(q.get(varPaneIndex))
	inputOff := isOne(q.get(varPaneInputOff))
	last := isOne(q.get(varPaneLast))
	left := q.get(varPaneLeft)
	marked := isOne(q.get(varPaneMarked))
	markedSet := isOne(q.get(varPaneMarkedSet))
	mode := q.get(varPaneMode)
	path := q.get(varPanePath)
	pid, _ := strconv.Atoi(q.get(varPanePid))
	pipe := isOne(q.get(varPanePipe))
	right := q.get(varPaneRight)
	searchString := q.get(varPaneSearchString)
	startCommand := q.get(varPaneStartCommand)
	startPath := q.get(varPaneStartPath)
	synchronized := isOne(q.get(varPaneSynchronized))
	tabs := q.get(varPaneTabs)
	title := q.get(varPaneTitle)
	top := q.get(varPaneTop)
	tty := q.get(varPaneTty)
	unseenChanges := isOne(q.get(varPaneUnseenChanges))
	width, _ := strconv.Atoi(q.get(varPaneWidth))
	windowIndex, _ := strconv.Atoi(q.get(varPaneWindowIndex))
	sessionName := q.get(varPaneSessionName)

	p := &Pane{

		Active:         active,
		AtBottom:       atBottom,
		AtLeft:         atLeft,
		AtRight:        atRight,
		AtTop:          atTop,
		Bg:             bg,
		Bottom:         bottom,
		CurrentCommand: currentCommand,
		CurrentPath:    currentPath,
		Dead:           dead,
		DeadSignal:     deadSignal,
		DeadStatus:     deadStatus,
		DeadTime:       deadTime,
		Fg:             fg,
		Format:         format,
		Height:         height,
		Id:             id,
		InMode:         inMode,
		Index:          index,
		InputOff:       inputOff,
		Last:           last,
		Left:           left,
		Marked:         marked,
		MarkedSet:      markedSet,
		Mode:           mode,
		Path:           path,
		Pid:            int32(pid),
		Pipe:           pipe,
		Right:          right,
		SearchString:   searchString,
		SessionName:    sessionName,
		StartCommand:   startCommand,
		StartPath:      startPath,
		Synchronized:   synchronized,
		Tabs:           tabs,
		Title:          title,
		Top:            top,
		Tty:            tty,
		UnseenChanges:  unseenChanges,
		Width:          width,
		WindowIndex:    windowIndex,
		tmux:           t,
	}

	return p
}
