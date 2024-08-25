// Copyright (c) 2024 Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"strconv"

	"github.com/GianlucaP106/gotmux/gotmux/vars"
)

type Server struct {
	Pid       int32
	Socket    *Socket
	StartTime string
	// StartTime  time.Time
	Uid     string
	User    string
	Version string

	tmux *Tmux
}

func (q *query) serverVars() *query {
	return q.vars(
		vars.Pid,
		vars.SocketPath,
		vars.StartTime,
		vars.Uid,
		vars.User,
		vars.Version,
	)
}

func (q queryResult) toServer(t *Tmux) *Server {
	pid, _ := strconv.Atoi(q.get(vars.Pid))
	socketPath := q.get(vars.SocketPath)
	socket, _ := newSocket(socketPath)
	startTime := q.get(vars.StartTime)
	uid := q.get(vars.Uid)
	user := q.get(vars.User)
	version := q.get(vars.Version)

	s := &Server{
		Pid:       int32(pid),
		Socket:    socket,
		StartTime: startTime,
		Uid:       uid,
		User:      user,
		Version:   version,

		tmux: t,
	}

	return s
}
