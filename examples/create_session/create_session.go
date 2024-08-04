package main

import (
	"log"

	"github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
	tmux := gotmux.DefaultTmux()

	s, err := tmux.NewSession(&gotmux.SessionOptions{
		StartDirectory: "/home",
		// ...
	})
	if err != nil {
		log.Fatal(err)
	}

	// takes over the client (terminal) if possible
	s.Attach()

	// kill when detached
	s.Kill()
}
