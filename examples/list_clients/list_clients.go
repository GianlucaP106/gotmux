package main

import (
	"fmt"
	"log"

	"github.com/GianlucaP106/gotmux/gotmux"
)

func main() {
	tmux, err := gotmux.DefaultTmux()
	if err != nil {
		log.Fatal(err)
	}

	clients, err := tmux.ListClients()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range clients {
		fmt.Println(c.Tty, c.Session)
	}
}
