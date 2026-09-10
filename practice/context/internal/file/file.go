package file

import (
	"math/rand/v2"
	"time"
)

type VirtualFile struct {
	name          string
	timeToProcess time.Duration
	processed     bool
}

func NewFilledVirtualFiles() []*VirtualFile {
	files := []*VirtualFile{
		{
			name:          "file1.txt",
			timeToProcess: time.Duration(300 + rand.IntN(2701)),
			processed:     false,
		},
		{
			name:          "file2.txt",
			timeToProcess: time.Duration(300 + rand.IntN(2701)),
			processed:     false,
		},
		{
			name:          "file3.txt",
			timeToProcess: time.Duration(300 + rand.IntN(2701)),
			processed:     false,
		},
	}
	return files
}

func (v VirtualFile) TimeToProcess() time.Duration {
	return v.timeToProcess
}

func (v VirtualFile) Name() string {
	return v.name
}

func (v *VirtualFile) SetProcessed(p bool) {
	v.processed = p
}

// func NewVirtualFile(name string, timeToProcess time.Duration) (*VirtualFile, error) {
// 	file := VirtualFile{
// 		Name:          name,
// 		TimeToProcess: timeToProcess,
// 	}

// 	return &file, nil
// }
