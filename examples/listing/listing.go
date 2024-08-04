package main

import (
	"fmt"
	"log"

	"github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
	tmux := gotmux.DefaultTmux()
	sessions, err := tmux.ListSessions()
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range sessions {
		fmt.Println("Session: ", s.Name)
		windows, err := s.ListWindows()
		if err != nil {
			log.Fatal(err)
		}

		for _, w := range windows {
			fmt.Println("Window: ", w.Index)
			panes, err := w.ListPanes()
			if err != nil {
				log.Fatal(err)
			}

			for _, p := range panes {
				fmt.Println("Pane: ", p.Path)
			}
		}
	}
}
