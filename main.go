package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <target_url> <num_requests> <concurrency>\n", os.Args[0])
		return
	}

	targetURL := os.Args[1]
	numRequests := atoi(os.Args[2])
	concurrency := atoi(os.Args[3])

	fmt.Printf("Starting DoS simulation on %s with %d requests and %d concurrency\n", targetURL, numRequests, concurrency)

	var g errgroup.Group
	var mu sync.Mutex
	totalTime := 0.0

	sem := make(chan struct{}, concurrency)

	for i := 0; i < numRequests; i++ {
		sem <- struct{}{}

		g.Go(func() error {
			start := time.Now()
			resp, err := http.Get(targetURL)
			if err != nil {
				log.Printf("Request failed: %v", err)
				<-sem
				return err
			}
			resp.Body.Close()
			elapsed := time.Since(start).Seconds()

			mu.Lock()
			totalTime += elapsed
			mu.Unlock()

			<-sem
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Printf("Errors occurred during the requests: %v", err)
	}

	averageTime := totalTime / float64(numRequests)
	fmt.Printf("DoS simulation completed. Average response time: %.2f seconds\n", averageTime)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid number: %s", s)
	}
	return i
}
