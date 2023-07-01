package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"log"
)

const (
	requestTimeout = 5 * time.Second
	userAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
)

func main() {
	targetURL := "http://example.com/"

	paths := generatePaths()

	outputFile := "paths.txt"
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}
	defer file.Close()

	pathsChan := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range pathsChan {
				fullURL := fmt.Sprintf("%s%s", targetURL, path)

				client := &http.Client{
					Timeout: requestTimeout,
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
					file.WriteString(fullURL + "\n")
				}

				resp.Body.Close()
			}
		}()
	}

	totalPaths := len(paths)
	donePaths := 0
	progressBarWidth := 50

	fmt.Println("Scanning paths...")

	for _, path := range paths {
		pathsChan <- path
		donePaths++

		progress := (float64(donePaths) / float64(totalPaths)) * 100
		barLength := int((progress / 100) * float64(progressBarWidth))
		fmt.Printf("\r[%s%s] %.1f%%", strings.Repeat("=", barLength), strings.Repeat(" ", progressBarWidth-barLength), progress)
	}

	close(pathsChan)

	wg.Wait()

	fmt.Println("\nPaths exported to", outputFile)
}

func generatePaths() []string {
	var paths []string

	shortnames := []string{"CON", "AUX", "PRN", "NUL", "COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9", "LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9"}

	for _, shortname := range shortnames {
		paths = append(paths, "~1"+shortname)
	}

	for _, shortname := range shortnames {
		paths = append(paths, shortname)
	}

	return paths
}
