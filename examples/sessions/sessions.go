package main

import (
	"log"

	"github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
	tmux, err := gotmux.DefaultTmux()
	if err != nil {
		log.Fatal(err)
	}

	s, err := tmux.NewSession(&gotmux.SessionOptions{
		StartDirectory: "/home",
		Name:           "somename",
	})
	if err != nil {
		log.Fatal(err)
	}

	// takes over the client (terminal) if possible
	err = s.Attach()
	if err != nil {
		log.Fatal(err)
	}

	// when detached this will continue...

	// next and prev cycle
	err = s.NextWindow()
	if err != nil {
		log.Fatal(err)
	}

	// kill the session
	s.Kill()
}
