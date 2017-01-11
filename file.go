package gorc

import (
	"os"
	"sync"
)

const (
	LEVEL = 1024
)

type File struct {
	name   string
	length int64
	index  int64
	file   *os.File
}
type block struct {
	previous *block
	start    int64
	end      int64
	subList  *block
}
type context struct {
	fileNames map[string]block
	lock      *sync.Mutex
	files     []File
}

func openfile(c *context) *os.File {
	return
}
func partFile(length int64, start int64) *block {
	if length/(LEVEL*LEVEL*10) > 0 && length/(LEVEL*LEVEL*LEVEL) == 0 {
		return &block{start: length / 2, end: length, previous: partFile(length/2-1, start), subList: partFile(length-length/2+1, length/2)}
	}
	if length/(LEVEL*LEVEL*LEVEL) > 0 && length/(LEVEL*LEVEL*LEVEL*LEVEL) == 0 {
		return &block{start: length / 2, end: length, previous: partFile(length/2-1, start), subList: partFile(length-length/2+1, length/2)}
	}
	if length/(LEVEL*LEVEL*LEVEL*LEVEL) > 0 {
		return &block{start: length / 2, end: length, previous: partFile(length/2-1, start), subList: partFile(length-length/2+1, length/2)}
	}
	return &block{start: start, end: length}
}
