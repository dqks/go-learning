package main

import (
	"parallel-file-processor/internal/file"
	"parallel-file-processor/internal/services"
)

func main() {
	files := file.NewFilledVirtualFiles()
	services.ProcessFiles(files)
}
