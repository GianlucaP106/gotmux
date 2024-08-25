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

	session, err := tmux.New()
	if err != nil {
		log.Fatal(err)
	}

	window, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	pane, err := window.GetPaneByIndex(0)
	if err != nil {
		log.Fatal(err)
	}

	err = pane.Split()
	if err != nil {
		log.Fatal(err)
	}
}
