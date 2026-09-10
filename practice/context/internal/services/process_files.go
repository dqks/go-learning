package services

import (
	"context"
	"errors"
	"fmt"
	"parallel-file-processor/internal/file"
	"sync"
	"time"
)

type processedFile struct {
	virtualFile *file.VirtualFile
	err         error
}

func ProcessFiles(files []*file.VirtualFile) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(files) * 2)
	resChan := make(chan processedFile, len(files))

	for i := range files {
		go fileHandler(ctx, &wg, files[i], resChan)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	resultProcessor(resChan)
}

func fileHandler(ctx context.Context, wg *sync.WaitGroup, file *file.VirtualFile, resChan chan processedFile) {
	defer wg.Done()
	select {
	case f := <-fileProcessor(ctx, wg, *file):
		resChan <- f
	case <-ctx.Done():
		resChan <- processedFile{virtualFile: file, err: ctx.Err()}
	}
}

func fileProcessor(ctx context.Context, wg *sync.WaitGroup, f file.VirtualFile) <-chan processedFile {
	result := make(chan processedFile, 1)

	// 2. Затем мы в канал спустя time.Millisecond * timeToProcess
	// передадим значение
	go func() {
		defer wg.Done()
		// Для того, чтобы отменить горутину по истечению
		// таймаута мы используем также и здесь select
		// отслеживая ctx.Done()
		// т.к. таймаут может истечь быстрее времени на выполнение
		select {
		case <-ctx.Done():
		case <-time.After(time.Millisecond * f.TimeToProcess()):
			f.SetProcessed(true)
			result <- processedFile{virtualFile: &f, err: nil}
		}
	}()

	// 1. Мы сначала вернем значение
	return result
}

func resultProcessor(res chan processedFile) {
	completedSum := 0
	canceledSum := 0
	otherSum := 0
	for f := range res {
		if f.err == nil {
			completedSum++
		}
		if errors.Is(f.err, context.DeadlineExceeded) || errors.Is(f.err, context.Canceled) {
			canceledSum++
		}
		if !errors.Is(f.err, context.DeadlineExceeded) && f.err != nil {
			otherSum++
		}
		fmt.Printf("Файл %s со временм обработки %d\n", f.virtualFile.Name(), f.virtualFile.TimeToProcess())
	}

	fmt.Printf("Завершено: %d\nОтменено: %d\nC другими ошибками: %d\n", completedSum, canceledSum, otherSum)
}
