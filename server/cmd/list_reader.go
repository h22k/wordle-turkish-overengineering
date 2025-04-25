package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("../storage/turkish_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numWorkers := runtime.NumCPU()

	lines := make(chan string, numWorkers*5) // buffered channel
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for line := range lines {
				processLine(workerID, line)
			}
		}(i)
	}

	counter := 0
	// Read lines and send to workers
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if len(strings.Split(text, " ")) > 1 {
			continue // Skip lines with more than one word
		}
		// TODO:: there is a bug in the scanner, some 4 letter words are also processed
		if len(text) == 5 || len(text) == 6 || len(text) == 7 {
			lines <- text
			counter++
		}
	}
	close(lines) // Close channel to stop workers

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Wait()                                      // Wait for all workers to finish
	fmt.Println("Total lines processed:", counter) // 15030
}

func processLine(workerID int, line string) {
	fmt.Println(workerID, " -> ", strings.ToLower(line))
}
