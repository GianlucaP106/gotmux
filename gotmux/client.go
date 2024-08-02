// Copyright (c) 2024 Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"strconv"

	"github.com/GianlucaP106/gotmux/gotmux/vars"
)

// Tmux client object.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#Variable
type Client struct {
	Activity     string
	CellHeight   int
	CellWidth    int
	ControlMode  bool
	Created      string
	Discarded    string
	Flags        string
	Height       int
	KeyTable     string
	LastSession  string
	Name         string
	Pid          int32
	Prefix       bool
	Readonly     bool
	Session      string
	Termname     string
	Termfeatures string
	Termtype     string
	Tty          string
	Uid          int32
	User         string
	Utf8         bool
	Width        int
	Written      string

	tmux *Tmux
}

// Gets the session that this client is attached to.
func (c *Client) GetSession() (*Session, error) {
	return c.tmux.GetSessionByName(c.Session)
}

// Sets the client variables in the query.
func (q *query) clientVars() *query {
	return q.vars(
		vars.ClientActivity,
		vars.ClientCellHeight,
		vars.ClientCellWidth,
		vars.ClientControlMode,
		vars.ClientCreated,
		vars.ClientDiscarded,
		vars.ClientFlags,
		vars.ClientHeight,
		vars.ClientKeyTable,
		vars.ClientLastSession,
		vars.ClientName,
		vars.ClientPid,
		vars.ClientPrefix,
		vars.ClientReadonly,
		vars.ClientSession,
		vars.ClientTermname,
		vars.ClientTermfeatures,
		vars.ClientTermtype,
		vars.ClientTty,
		vars.ClientUid,
		vars.ClientUser,
		vars.ClientUtf8,
		vars.ClientWidth,
		vars.ClientWritten,
	)
}

// Converts a QueryResult to a Client.
func (q queryResult) toClient(t *Tmux) *Client {
	activity := q.get(vars.ClientActivity)
	cellHeight, _ := strconv.Atoi(q.get(vars.ClientCellHeight))
	cellWidth, _ := strconv.Atoi(q.get(vars.ClientCellWidth))
	controlMode := isOne(q.get(vars.ClientControlMode))
	created := q.get(vars.ClientCreated)
	discarded := q.get(vars.ClientDiscarded)
	flags := q.get(vars.ClientFlags)
	height, _ := strconv.Atoi(q.get(vars.ClientHeight))
	keyTable := q.get(vars.ClientKeyTable)
	lastSession := q.get(vars.ClientLastSession)
	name := q.get(vars.ClientName)
	pid, _ := strconv.Atoi(q.get(vars.ClientPid))
	prefix := isOne(q.get(vars.ClientPrefix))
	readonly := isOne(q.get(vars.ClientReadonly))
	session := q.get(vars.ClientSession)
	termname := q.get(vars.ClientTermname)
	termfeatures := q.get(vars.ClientTermfeatures)
	termtype := q.get(vars.ClientTermtype)
	tty := q.get(vars.ClientTty)
	uid, _ := strconv.Atoi(q.get(vars.ClientUid))
	user := q.get(vars.ClientUser)
	utf8 := isOne(q.get(vars.ClientUtf8))
	width, _ := strconv.Atoi(q.get(vars.ClientWidth))
	written := q.get(vars.ClientWritten)

	c := &Client{
		Activity:     activity,
		CellHeight:   cellHeight,
		CellWidth:    cellWidth,
		ControlMode:  controlMode,
		Created:      created,
		Discarded:    discarded,
		Flags:        flags,
		Height:       height,
		KeyTable:     keyTable,
		LastSession:  lastSession,
		Name:         name,
		Pid:          int32(pid),
		Prefix:       prefix,
		Readonly:     readonly,
		Session:      session,
		Termname:     termname,
		Termfeatures: termfeatures,
		Termtype:     termtype,
		Tty:          tty,
		Uid:          int32(uid),
		User:         user,
		Utf8:         utf8,
		Width:        width,
		Written:      written,

		tmux: t,
	}

	return c
}
