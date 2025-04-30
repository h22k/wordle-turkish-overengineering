package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/h22k/wordle-turkish-overengineering/server/config"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/bootstrap"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

func main() {
	cfg := config.LoadConfig()
	ctx := context.Background()
	app := bootstrap.InitApplication(ctx, cfg)

	gameService := app.AppService().GameService()

	word := make(chan domain.Word, 500)
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range word {
				if w.Len() >= 5 && w.Len() <= 7 {
					_ = gameService.AddWord(ctx, w)
				}
			}
		}()
	}

	readLine(word)

	wg.Wait()

	app.Close()

	fmt.Println("Done")
}

func readLine(word chan<- domain.Word) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := path.Join(wd, "storage", "turkish_words.txt")
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, " ") {
			continue
		}
		word <- domain.Word(strings.ToLower(text))
	}

	close(word)
}
