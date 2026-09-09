package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Millisecond int

type VirtualFile struct {
	name          string
	timeToProcess Millisecond
}

func ProcessFile(ctx context.Context, wg *sync.WaitGroup, file VirtualFile) {
	defer wg.Done()
	select {
	case file := <-fileHandler(file):
		fmt.Println("Обработка файла ", file)
		// Отправка данных в результирующий канал
	case <-time.After(time.Second * 2):
		fmt.Println("Таймаут, прошло 2 секунды")
	case <-ctx.Done():
		fmt.Println("Операция отменилась, таймаут")
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
	ctx, cancel := context.WithCancel(context.Background())
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
	// fileChan := make(chan VirtualFile, len(files))

	for i := range files {
		go ProcessFile(ctx, &wg, files[i])
	}

	wg.Wait()
}
