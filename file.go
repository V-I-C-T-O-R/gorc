package gorc

import (
	"bytes"
	"github.com/coreos/go-log/log"
	"image/png"
	"math/rand"
	"os"
)

const (
	StdLen  = 16
	UUIDLen = 20
)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var File string

//实践证明，这丫的就只生成一个不变的数
func ResolvePng(content []byte) error {
	img, err := png.Decode(bytes.NewReader(content))
	filename := "./pic/" + New() + ".jpg"
	log.Info("create file name == ", filename)
	file, _ := os.Create(filename)
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		log.Info("create file failed")
		return err
	}
	File = filename
	return err
}

func New() string {
	return NewLenChars(StdLen, StdChars)
}

func NewLen(length int) string {
	return NewLenChars(length, StdChars)
}

func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
