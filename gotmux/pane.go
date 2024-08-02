// Copyright (c) 2024 Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"errors"
	"strconv"

	"github.com/GianlucaP106/gotmux/gotmux/vars"
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
	StartCommand   string
	StartPath      string
	Synchronized   bool
	Tabs           string
	Title          string
	Top            string
	Tty            string
	UnseenChanges  bool
	Width          int

	tmux *Tmux
}

// Type representing the relative position of panes
//
// TODO: Reference:
type PanePosition string

// Enumeratioe pane positions
//
// TODO: Reference:
const (
	PaneUp    PanePosition = "-U"
	PaneRight PanePosition = "-R"
	PaneDown  PanePosition = "-D"
	PaneLeft  PanePosition = "-L"
)

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
// Shorthand of the method above but with default options
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#select-pane
func (p *Pane) Select() error {
	return p.SelectPane(nil)
}

// Sets the pane variables in the query.
func (q *query) paneVars() *query {
	return q.vars(
		vars.PaneActive,
		vars.PaneAtBottom,
		vars.PaneAtLeft,
		vars.PaneAtRight,
		vars.PaneAtTop,
		vars.PaneBg,
		vars.PaneBottom,
		vars.PaneCurrentCommand,
		vars.PaneCurrentPath,
		vars.PaneDead,
		vars.PaneDeadSignal,
		vars.PaneDeadStatus,
		vars.PaneDeadTime,
		vars.PaneFg,
		vars.PaneFormat,
		vars.PaneHeight,
		vars.PaneId,
		vars.PaneInMode,
		vars.PaneIndex,
		vars.PaneInputOff,
		vars.PaneLast,
		vars.PaneLeft,
		vars.PaneMarked,
		vars.PaneMarkedSet,
		vars.PaneMode,
		vars.PanePath,
		vars.PanePid,
		vars.PanePipe,
		vars.PaneRight,
		vars.PaneSearchString,
		vars.PaneStartCommand,
		vars.PaneStartPath,
		vars.PaneSynchronized,
		vars.PaneTabs,
		vars.PaneTitle,
		vars.PaneTop,
		vars.PaneTty,
		vars.PaneUnseenChanges,
		vars.PaneWidth,
	)
}

// Converts a QueryResult to a pane.
func (q queryResult) toPane(t *Tmux) *Pane {
	active := isOne(q.get(vars.PaneActive))
	atBottom := isOne(q.get(vars.PaneAtBottom))
	atLeft := isOne(q.get(vars.PaneAtLeft))
	atRight := isOne(q.get(vars.PaneAtRight))
	atTop := isOne(q.get(vars.PaneAtTop))
	bg := q.get(vars.PaneBg)
	bottom := q.get(vars.PaneBottom)
	currentCommand := q.get(vars.PaneCurrentCommand)
	currentPath := q.get(vars.PaneCurrentPath)
	dead := isOne(q.get(vars.PaneDead))
	deadSignal, _ := strconv.Atoi(q.get(vars.PaneDeadSignal))
	deadStatus, _ := strconv.Atoi(q.get(vars.PaneDeadStatus))
	deadTime := q.get(vars.PaneDeadTime)
	fg := q.get(vars.PaneFg)
	format := isOne(q.get(vars.PaneFormat))
	height, _ := strconv.Atoi(q.get(vars.PaneHeight))
	id := q.get(vars.PaneId)
	inMode := isOne(q.get(vars.PaneInMode))
	index, _ := strconv.Atoi(q.get(vars.PaneIndex))
	inputOff := isOne(q.get(vars.PaneInputOff))
	last := isOne(q.get(vars.PaneLast))
	left := q.get(vars.PaneLeft)
	marked := isOne(q.get(vars.PaneMarked))
	markedSet := isOne(q.get(vars.PaneMarkedSet))
	mode := q.get(vars.PaneMode)
	path := q.get(vars.PanePath)
	pid, _ := strconv.Atoi(q.get(vars.PanePid))
	pipe := isOne(q.get(vars.PanePipe))
	right := q.get(vars.PaneRight)
	searchString := q.get(vars.PaneSearchString)
	startCommand := q.get(vars.PaneStartCommand)
	startPath := q.get(vars.PaneStartPath)
	synchronized := isOne(q.get(vars.PaneSynchronized))
	tabs := q.get(vars.PaneTabs)
	title := q.get(vars.PaneTitle)
	top := q.get(vars.PaneTop)
	tty := q.get(vars.PaneTty)
	unseenChanges := isOne(q.get(vars.PaneUnseenChanges))
	width, _ := strconv.Atoi(q.get(vars.PaneWidth))

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
		StartCommand:   startCommand,
		StartPath:      startPath,
		Synchronized:   synchronized,
		Tabs:           tabs,
		Title:          title,
		Top:            top,
		Tty:            tty,
		UnseenChanges:  unseenChanges,
		Width:          width,

		tmux: t,
	}

	return p
}
