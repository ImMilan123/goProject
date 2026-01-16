package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	//Burp proxy URL
	proxyURL, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "GoClient/1.0")

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