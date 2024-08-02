package gotmux

import "strings"

// Checks if a string is 1.
func isOne(s string) bool {
	return s == "1"
}

// Splits a string by comma.
func parseList(l string) []string {
	return strings.Split(l, ",")
}

// Checks the validity of the tmux session name.
func checkSessionName(name string) bool {
	if len(name) == 0 {
		return false
	}

	if strings.Contains(name, ":") {
		return false
	}

	if strings.Contains(name, ".") {
		return false
	}

	return true
}
