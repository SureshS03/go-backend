package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"testing"
)

const (
	baseURL     = "http://localhost:8080"
	totalReq    = 1000
	concurrency = 10
	token       = "41ed25e92acf8686bbcddcd9ff5084042ae5285c"
)

func BenchmarkGetUser(b *testing.B) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	start := time.Now()

	for i := 0; i < totalReq; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()

			url := fmt.Sprintf("%s/user/", baseURL)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("request error:", err)
				return
			}

			req.Header.Set("Authorization", "Bearer "+token)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("Request %d failed: %v\n", i, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Request %d: %s\n", i, resp.Status)
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Finished %d requests in %v\n", totalReq, duration)
}
