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

	// create new window
	window, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	// move the window to index 10
	err = window.Move(session.Name, 10)
	if err != nil {
		log.Fatal(err)
	}

	// list windows
	windows, err := session.ListWindows()
	if err != nil {
		log.Fatal(err)
	}

	// toggle selection
	for _, w := range windows {
		if !w.Active {
			err = w.Select()
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	// set layout
	err = window.SelectLayout(gotmux.WindowLayoutEvenVertical)
	if err != nil {
		log.Fatal(err)
	}
}
