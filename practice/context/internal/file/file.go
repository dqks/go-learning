package file

import (
	"time"
)

type VirtualFile struct {
	Name          string
	TimeToProcess time.Duration
}

func NewFilledVirtualFiles() []*VirtualFile {
	files := []*VirtualFile{
		{
			Name:          "file1.txt",
			TimeToProcess: 1999,
		},
		{
			Name:          "file2.txt",
			TimeToProcess: 1000,
		},
		{
			Name:          "file3.txt",
			TimeToProcess: 3000,
		},
	}
	return files
}

// func NewVirtualFile(name string, timeToProcess time.Duration) (*VirtualFile, error) {
// 	file := VirtualFile{
// 		Name:          name,
// 		TimeToProcess: timeToProcess,
// 	}

// 	return &file, nil
// }
