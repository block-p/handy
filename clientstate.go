package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type clientState struct {
	method string
	url    string
	client http.Client
}

func newClient(method string, url string) *clientState {
	return &clientState{
		method: method,
		url:    url,
	}

}

func (c *clientState) do() {
	fmt.Println(c.url)
	req, err := http.NewRequest(c.method, c.url, nil)
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal("error: ", err)
	}
	var buf bytes.Buffer
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	body := resp.Body
	data, err := io.ReadAll(body)
	if err != nil {
		log.Fatal("error: ", err)
	}
	if s:= strings.Split(resp.Header.Get("Content-Type"), ";"); s[0] == "application/json"{

		err := json.Indent(&buf, data, "", "  ")
		if err != nil{
			log.Fatal("error: ", err)
		}

	fmt.Println(buf.String())

	}else{
		fmt.Println(string(data))
	}
}
