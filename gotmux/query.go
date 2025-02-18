// Copyright (c) 2024 Gianluca Piccirillo
// This software is licensed under the MIT License.
// See the LICENSE file in the root directory for more information.

package gotmux

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Seperator used to query the data.
const sep = "-:-"

// Represents a query to tmux.
// Includes flag arguments, positional arguments, tmux command and tmux vars.
type query struct {
	fArgs     []string
	pArgs     []string
	command   []string
	variables []string
}

// Returns a new newQuery.
func newQuery() *query {
	return &query{
		command: make([]string, 0),
		fArgs:   make([]string, 0),
		pArgs:   make([]string, 0),
	}
}

// Append a command to the query.
func (q *query) cmd(c ...string) *query {
	q.command = append(q.command, c...)
	return q
}

// Append flag arguments to the query.
func (q *query) fargs(a ...string) *query {
	q.fArgs = append(q.fArgs, a...)
	return q
}

// Append positional arguments to the query.
func (q *query) pargs(a ...string) *query {
	q.pArgs = append(q.pArgs, a...)
	return q
}

// Adds tmux variables to the query.
func (q *query) vars(v ...string) *query {
	q.variables = v
	return q
}

// Prepares the query to be ran.
func (q *query) prepare() *exec.Cmd {
	query := []string{}

	query = append(query, "tmux")
	query = append(query, q.command...)
	query = append(query, q.fArgs...)

	vars := func() string {
		if len(q.variables) == 0 {
			return ""
		}

		out := []string{}
		for _, vr := range q.variables {
			out = append(out, fmt.Sprintf("#{%s}", vr))
		}

		joined := strings.Join(out, sep)
		joined = fmt.Sprintf("'%s'", joined)

		return joined
	}()

	if vars != "" {
		if q.command[0] == "display-message" {
			query = append(query, "-p", vars)
		} else {
			query = append(query, "-F", vars)
		}
	}

	query = append(query, q.pArgs...)
	return exec.Command(query[0], query[1:]...)
}

// Runs the query with output.
func (q *query) run() (*queryOutput, error) {
	b, err := q.prepare().Output()
	if err != nil {
		return nil, err
	}

	o := &queryOutput{
		result:    string(b),
		variables: q.variables,
	}

	return o, nil
}

// Runs the query and attaches to the terminal by redirecting.
func (q *query) runTty() error {
	cmd := q.prepare()
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Query Output object.
type queryOutput struct {
	result    string
	variables []string
}

// Collects an output into a resut.
func (q *queryOutput) collect() []queryResult {
	lines := strings.Split(q.result, "\n")
	out := make([]queryResult, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		stripped := strings.Trim(line, "'")
		vars := strings.Split(stripped, sep)

		if len(vars) != len(q.variables) {
			log.Panicln("invalid query output")
		}

		result := make(queryResult)
		for idx, v := range q.variables {
			result[v] = vars[idx]
		}

		out = append(out, result)
	}

	return out
}

// Returns one element from the result.
func (q *queryOutput) one() queryResult {
	return q.collect()[0]
}

// Returns the raw result.
func (q *queryOutput) raw() string {
	return q.result
}

// Represents a query result.
type queryResult map[string]string

// Get a value by key in the result.
func (q queryResult) get(key string) string {
	return q[key]
}
