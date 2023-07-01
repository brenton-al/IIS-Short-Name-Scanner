package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	requestTimeout = 5 * time.Second
	userAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
)

func main() {
	targetURL := "http://example.com/"

	inputFile := "paths.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pathsChan := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range pathsChan {
				fullURL := fmt.Sprintf("%s%s", targetURL, path)

				var client *http.Client

				if strings.HasPrefix(targetURL, "https") {
					tr := &http.Transport{
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					}
					client = &http.Client{
						Transport: tr,
						Timeout:   requestTimeout,
					}
				} else {
					client = &http.Client{
						Timeout: requestTimeout,
					}
				}

				req, err := http.NewRequest(http.MethodGet, fullURL, nil)
				if err != nil {
					log.Printf("Error creating request for path %s: %s", path, err)
					continue
				}

				req.Header.Set("User-Agent", userAgent)

				resp, err := client.Do(req)
				if err != nil {
					log.Printf("Error scanning path %s: %s", path, err)
					continue
				}

				if resp.StatusCode == http.StatusOK {
					fmt.Printf("Found existing path: %s\n", path)
				}

				resp.Body.Close()
			}
		}()
	}

	totalPaths := 0
	donePaths := 0
	progressBarWidth := 50

	for scanner.Scan() {
		totalPaths++
	}

	file.Seek(0, 0)

	fmt.Println("Scanning paths...")

	for scanner.Scan() {
		path := scanner.Text()
		pathsChan <- path
		donePaths++

		progress := (float64(donePaths) / float64(totalPaths)) * 100
		barLength := int((progress / 100) * float64(progressBarWidth))
		fmt.Printf("\r[%s%s] %.1f%%", strings.Repeat("=", barLength), strings.Repeat(" ", progressBarWidth-barLength), progress)
	}

	close(pathsChan)

	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %s", err)
	}
}
