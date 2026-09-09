package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Millisecond int

type VirtualFile struct {
	name          string
	timeToProcess Millisecond
}

func ProcessFile(ctx context.Context, wg *sync.WaitGroup, file VirtualFile, resChan chan error) {
	defer wg.Done()
	select {
	case <-fileHandler(file):
		resChan <- nil
	case <-ctx.Done():
		resChan <- ctx.Err()
	}
}

func fileHandler(file VirtualFile) <-chan VirtualFile {
	result := make(chan VirtualFile, 1)

	// 2. Затем мы в канал спустя time.Millisecond * timeToProcess
	// передадим в канал значение
	go func() {
		time.Sleep(time.Millisecond * time.Duration(file.timeToProcess))
		result <- file
	}()

	// 1. Мы сначала вернем значение
	return result
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	files := []VirtualFile{
		{
			name:          "file1.txt",
			timeToProcess: 1999,
		},
		{
			name:          "file2.txt",
			timeToProcess: 1000,
		},
		{
			name:          "file3.txt",
			timeToProcess: 3000,
		},
	}

	var wg sync.WaitGroup
	wg.Add(len(files))
	resChan := make(chan error, len(files))

	for i := range files {
		go ProcessFile(ctx, &wg, files[i], resChan)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	completedSum := 0
	canceledSum := 0
	otherSum := 0
	for res := range resChan {
		if res == nil {
			completedSum++
		}
		if errors.Is(res, context.Canceled) {
			canceledSum++
		}
		if errors.Is(res, context.DeadlineExceeded) {
			otherSum++
		}
	}

	fmt.Printf("Завершено: %d\nОтменено: %d\nC другими ошибками: %d\n", completedSum, canceledSum, otherSum)
}
