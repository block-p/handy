package main

import (
	"os"
)

type appState int

const (
	serverMode appState = iota
	clientMode
)

func main() {

	args := os.Args[1:]
	switch {
	case len(args) < 2:

		server := newServer("")
		server.run()
	case args[0] == "s":
		server := newServer(args[1])
		server.run()
	case args[0] == "get" || args[0] == "post" && len(args) > 2:

		client := newClient(args[0], args[1])
		client.do()
	case len(args) == 2 && args[0] != "get" || args[0] != "post":
		client := newClient("GET", args[1])
		client.do()
	}

}
