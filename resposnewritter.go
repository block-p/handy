package main

import "net/http"

type reswrwrapper struct {
	rw         http.ResponseWriter
	statuscode int
	data       []byte
}

func Newreswrwrapper(r http.ResponseWriter) *reswrwrapper {

	return &reswrwrapper{
		rw:         r,
		statuscode: http.StatusOK,
	}

}
func (r *reswrwrapper) Header() http.Header {
	return r.rw.Header()
}

func (r *reswrwrapper) Write(data []byte) (int, error) {
	r.data = data
	return len(data), nil

}
func (r *reswrwrapper) WriteHeader(code int) {
	r.statuscode = code
}
