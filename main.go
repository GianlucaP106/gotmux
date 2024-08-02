package main

import (
	"log"

	"github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
	tmux, err := gotmux.NewTmux("/private/tmp/tmux-501/default")
	if err != nil {
		log.Panicln(err)
	}

	s, err := tmux.GetSessionByName("helloworld")
	if err != nil {
		log.Panicln(err)
	}

	windows, err := s.ListWindows()
	if err != nil {
		log.Panicln(err)
	}

	w := windows[0]
	w.Rename("somethingnew")
}
