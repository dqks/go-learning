package services

import (
	"context"
	"errors"
	"fmt"
	"parallel-file-processor/internal/file"
	"sync"
	"time"
)

func ProcessFiles(files []*file.VirtualFile) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(files))
	resChan := make(chan error, len(files))

	for i := range files {
		go fileHandler(ctx, &wg, files[i], resChan)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	resultProcessor(resChan)
}

func fileHandler(ctx context.Context, wg *sync.WaitGroup, file *file.VirtualFile, resChan chan error) {
	defer wg.Done()
	select {
	case <-fileProcessor(*file):
		resChan <- nil
	case <-ctx.Done():
		resChan <- ctx.Err()
	}
}

func fileProcessor(f file.VirtualFile) <-chan file.VirtualFile {
	result := make(chan file.VirtualFile, 1)

	// 2. Затем мы в канал спустя time.Millisecond * timeToProcess
	// передадим в канал значение
	go func() {
		time.Sleep(time.Millisecond * time.Duration(f.TimeToProcess))
		result <- f
	}()

	// 1. Мы сначала вернем значение
	return result
}

func resultProcessor(resChan chan error) {
	completedSum := 0
	canceledSum := 0
	otherSum := 0
	for res := range resChan {
		if res == nil {
			completedSum++
		}
		if errors.Is(res, context.DeadlineExceeded) {
			canceledSum++
		}
		if !errors.Is(res, context.DeadlineExceeded) && res != nil {
			otherSum++
		}
	}

	fmt.Printf("Завершено: %d\nОтменено: %d\nC другими ошибками: %d\n", completedSum, canceledSum, otherSum)
}
