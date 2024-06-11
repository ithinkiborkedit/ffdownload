package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

func main() {

	const (
		maxConcurrentDownloads = 5
		chunkSize              = 1024 * 1024
	)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run download_accelerator.go <url>")
		os.Exit(1)
	}
	url := os.Args[1]
	filename := path.Base(url)
	// fileExt := path.Ext(filename)

	start := time.Now()

	resp, err := http.Head(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	contentLength := resp.ContentLength
	numChunks := contentLength / chunkSize

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	var wg sync.WaitGroup
	chunkCh := make(chan int)

	for i := 0; i < maxConcurrentDownloads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for chunk := range chunkCh {
				start := int64(chunk) * chunkSize
				end := start + chunkSize

				if end > contentLength {
					end = contentLength
				}

				req, err := http.NewRequest("GET", url, nil)
				if err != nil {

					fmt.Println("Error:", err)
					return
				}
				req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end-1))
				fmt.Println(req)
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				defer resp.Body.Close()

				var chunkData []byte
				chunkData, err = io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				_, err = file.WriteAt(chunkData, start)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			}
		}()
	}

	for i := 0; i <= int(numChunks); i++ {
		chunkCh <- i
	}
	close(chunkCh)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Download completed in %s\n", elapsed)
}
