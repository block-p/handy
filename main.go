package main

import (
	"flag"
	"io"
	"log"
	"os"
)

type appState int

const (
	serverMode appState = iota
	clientMode
)

var w io.Writer

func main() {
	w = os.Stdout
	serverflag := flag.Bool("s", false, "server mode usage: handy -s 127.0.0.1:8080")
	clientflag := flag.Bool("c", false, "client mode usage: handy -c method url")
	headerflag := flag.Bool("h", false, "set headers")
	flag.Parse()
	switch {
	case *serverflag:
		var server *serverState
		if len(flag.Args()) == 0 {
			server = newServer("")
		} else {
			server = newServer(flag.Arg(0))
		}
		server.run()
	case *clientflag:
		var client *clientState
		if len(flag.Args()) == 1 {
			client = newClient("get", flag.Arg(0))
		} else if len(flag.Args()) == 2 && flag.Arg(0) == "get" || flag.Arg(0) == "post" {
			client = newClient(flag.Arg(0), flag.Arg(1))
		} else {
			log.Fatal("invalid usage")
		}
		client.do()
	default:
		flag.PrintDefaults()
	}

}
