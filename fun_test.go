package gorc

import (
	"fmt"
	"testing"
)

func Test_NewFile(t *testing.T) {
	var url = "http://golangtc.com/static/go/1.8beta1/go1.8beta1.darwin-amd64.pkg"
	l := sendHead(url)
	fmt.Println(l)
}
