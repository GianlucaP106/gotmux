// Copyright (c) Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import "errors"

// Tmux Socket object.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#S
type Socket struct {
	Path string
}

// Creates a new socket object. Verifies its validtity.
func newSocket(path string) (*Socket, error) {
	s := &Socket{}
	if !s.validateSocket(path) {
		return nil, errors.New("invalid socket")
	}
	s.Path = path
	return s, nil
}

// Valides a sockets validity
func (s *Socket) validateSocket(path string) bool {
	_, err := newQuery().
		cmd("-S", path, "list-clients").
		run()
	return err == nil
}
