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
	root, _ = os.Getwd()
	p := path.Join(root, "lib", "dsfds")

	fmt.Println(l)
	fmt.Println(root)
	fmt.Println(p)
}
