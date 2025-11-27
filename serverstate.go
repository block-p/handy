package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type serverState struct {
	server  *http.Server
	addrptr string
}

func newServer(addrptr string) *serverState {
	if addrptr == "" {
		addrptr = "127.0.0.1:8080"
	}
	return &serverState{
		addrptr: addrptr,
		server:  &http.Server{Addr: addrptr},
	}
}

func (s *serverState) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := Newreswrwrapper(w)
		start := time.Now()
		next.ServeHTTP(rec, r)

		rec.rw.WriteHeader(rec.statuscode)
		rec.rw.Write(rec.data)
		w.Write([]byte(fmt.Sprintf("%d %s  %s |> %s  %dms\n", rec.statuscode, http.StatusText(rec.statuscode), r.URL.Path, r.Method, time.Since(start).Milliseconds())))

	})
}

func (s *serverState) run() {
	mux := http.NewServeMux()
	mux.Handle("/", s.logger(http.FileServer(http.Dir("."))))
	s.server.Handler = mux
	log.Println("start server on " + s.addrptr)
	log.Println(s.server.ListenAndServe())

}
