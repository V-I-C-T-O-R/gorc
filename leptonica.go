package gorc

/*
#cgo LDFLAGS: -llept
#include "leptonica/allheaders.h"
#include <stdlib.h>
l_uint8* uglycast(void* value) { return (l_uint8*)value; }
*/
import "C"
import "sync"

type ImgType int32

const (
	UNKNOWN ImgType = iota
	BMP
	PNG
	GIF
	TIFF
	JPG
)

type Pix struct {
	GPix *C.PIX
	Lock sync.Mutex
}

func (p *Pix) getGPix() *C.PIX {
	return p.GPix
}
