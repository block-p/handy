package main

import (
	"flag"
	"io"
	"log"
	"os"
)

var W io.Writer

func main() {
	W = os.Stdout
	var headerlist headers
	var datavar data
	serverflag := flag.Bool("s", false, "server mode usage: handy -s 127.0.0.1:8080")
	clientflag := flag.Bool("c", false, "client mode usage: handy -c -h 'header' -d 'body data' method url")
	flag.Var(&headerlist, "h", "headers")
	flag.Var(&datavar, "d", "data var")
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
			client = newClient("get", flag.Arg(0), headerlist, datavar)
		} else if len(flag.Args()) == 2 && flag.Arg(0) == "get" || flag.Arg(0) == "post" {
			client = newClient(flag.Arg(0), flag.Arg(1), headerlist, datavar)
		} else {
			log.Fatal("invalid usage")
		}
		client.do()
	default:
		flag.PrintDefaults()
	}

}
