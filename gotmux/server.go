// Copyright (c) Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"strconv"
)

type Server struct {
	Pid       int32
	Socket    *Socket
	StartTime string
	Uid       string
	User      string
	Version   string

	tmux *Tmux
}

func (q *query) serverVars() *query {
	return q.vars(
		varPid,
		varSocketPath,
		varStartTime,
		varUid,
		varUser,
		varVersion,
	)
}

func (q queryResult) toServer(t *Tmux) *Server {
	pid, _ := strconv.Atoi(q.get(varPid))
	socketPath := q.get(varSocketPath)
	socket, _ := newSocket(socketPath)
	startTime := q.get(varStartTime)
	uid := q.get(varUid)
	user := q.get(varUser)
	version := q.get(varVersion)

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
