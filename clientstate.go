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
	method  string
	url     string
	headers []string
	body    []byte
	client  http.Client
}

func newClient(method, url string, headers []string, body []byte) *clientState {
	return &clientState{
		method:  method,
		url:     url,
		headers: headers,
		body:    body,
	}

}

func (c *clientState) do() {
	fmt.Println(c.url)
	var body io.Reader
	if c.body != nil {
		body = bytes.NewBuffer(c.body)
	}

	req, err := http.NewRequest(c.method, c.url, body)
	for _, h := range c.headers {
		kv := strings.Split(h, " ")

		req.Header.Set(kv[0], kv[1])
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal("error: ", err)
	}
	var buf bytes.Buffer
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	responseBody := resp.Body
	data, err := io.ReadAll(responseBody)
	if err != nil {
		log.Fatal("error: ", err)
	}
	if s := strings.Split(resp.Header.Get("Content-Type"), ";"); s[0] == "application/json" {

		err := json.Indent(&buf, data, "", "  ")
		if err != nil {
			log.Fatal("error: ", err)
		}

		w.Write(buf.Bytes())

	} else {
		w.Write(data)
	}
}
