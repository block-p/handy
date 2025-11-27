package main

import "strings"

type headers []string

func (h *headers) String() string { return strings.Join(*h, ", ") }
func (h *headers) Set(v string) error {
	*h = append(*h, v)
	return nil
}

type data []byte

func (d *data) String() string { return string(*d) }
func (d *data) Set(v string) error {
	*d = []byte(v)
	return nil
}
