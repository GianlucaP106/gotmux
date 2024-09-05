package main

import (
	"log"

	"github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
	tmux, _ := gotmux.DefaultTmux()
	_, err := tmux.GetCurrentClient()
	if err != nil {
		log.Fatal(err)
	}
}
