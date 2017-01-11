package gorc

import "os"

type File struct {
	content []byte
	length  int64
	index   int64
	file    *os.File
}
