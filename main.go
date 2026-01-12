package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8081",
		nil,
	)
	req.Header.Set("User-Agent", "GoClient/1.0")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))
}
