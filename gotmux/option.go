package gotmux

import (
	"strings"
)

// Represents a tmux option which is a key value pair.
//
// Reference: https://man.openbsd.org/OpenBSD-current/man1/tmux.1#OPTIONS
type Option struct {
	Key   string
	Value string
}

func newOption(key, value string) *Option {
	return &Option{Key: key, Value: value}
}

func (q *queryOutput) toOptions() []*Option {
	lines := strings.Split(q.raw(), "\n")
	out := make([]*Option, 0)
	for _, line := range lines {
		s := strings.Split(line, " ")
		if len(s) != 2 {
			continue
		}
		key, val := s[0], s[1]
		out = append(out, newOption(key, val))
	}
	return out
}
