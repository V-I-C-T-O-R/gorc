package gorc

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func Test_NewFile(t *testing.T) {
	var url = "http://golangtc.com/static/go/1.8beta1/go1.8beta1.darwin-amd64.pkg"
	l, _ := sendHead(url)
	ROOT, _ = os.Getwd()
	p := path.Join(ROOT, "lib", "dsfds")

	fmt.Println(l)
	fmt.Println(ROOT)
	fmt.Println(p)
}
