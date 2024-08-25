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

	server, err := tmux.GetServerInformation()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(server.Version)
}
