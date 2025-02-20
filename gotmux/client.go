// Copyright (c) Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"strconv"
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
		varClientActivity,
		varClientCellHeight,
		varClientCellWidth,
		varClientControlMode,
		varClientCreated,
		varClientDiscarded,
		varClientFlags,
		varClientHeight,
		varClientKeyTable,
		varClientLastSession,
		varClientName,
		varClientPid,
		varClientPrefix,
		varClientReadonly,
		varClientSession,
		varClientTermname,
		varClientTermfeatures,
		varClientTermtype,
		varClientTty,
		varClientUid,
		varClientUser,
		varClientUtf8,
		varClientWidth,
		varClientWritten,
	)
}

// Converts a QueryResult to a Client.
func (q queryResult) toClient(t *Tmux) *Client {
	activity := q.get(varClientActivity)
	cellHeight, _ := strconv.Atoi(q.get(varClientCellHeight))
	cellWidth, _ := strconv.Atoi(q.get(varClientCellWidth))
	controlMode := isOne(q.get(varClientControlMode))
	created := q.get(varClientCreated)
	discarded := q.get(varClientDiscarded)
	flags := q.get(varClientFlags)
	height, _ := strconv.Atoi(q.get(varClientHeight))
	keyTable := q.get(varClientKeyTable)
	lastSession := q.get(varClientLastSession)
	name := q.get(varClientName)
	pid, _ := strconv.Atoi(q.get(varClientPid))
	prefix := isOne(q.get(varClientPrefix))
	readonly := isOne(q.get(varClientReadonly))
	session := q.get(varClientSession)
	termname := q.get(varClientTermname)
	termfeatures := q.get(varClientTermfeatures)
	termtype := q.get(varClientTermtype)
	tty := q.get(varClientTty)
	uid, _ := strconv.Atoi(q.get(varClientUid))
	user := q.get(varClientUser)
	utf8 := isOne(q.get(varClientUtf8))
	width, _ := strconv.Atoi(q.get(varClientWidth))
	written := q.get(varClientWritten)

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
