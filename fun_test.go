package gorc

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func Test_NewFile(t *testing.T) {
	/*var url = "http://golangtc.com/static/go/1.8beta1/go1.8beta1.darwin-amd64.pkg"
	l, _ := sendHead(url)*/
	root, _ = os.Getwd()
	p := path.Join(root, "lib", "dsfds")
	x := filepath.Join(root, "lib", "dsfds")
	//fmt.Println(l)
	fmt.Println(root)
	fmt.Println(p)
	fmt.Println(runtime.GOOS)
	fmt.Println(x)

	result := fmt.Sprintf("%.f", float64(111111111)/float64(2222222222)*100)
	fmt.Println(result)
}
