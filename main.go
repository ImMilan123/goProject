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
		"http://localhost:8080",
		nil,
	)
	if err != nil {
	}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
	}
	fmt.Println(string(bodyBytes))
}
