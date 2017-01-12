package main

import "github.com/V-I-C-T-O-R/gorc"

func main() {
	var url = "http://golangtc.com/static/go/1.8beta1/go1.8beta1.darwin-amd64.pkg"
	gorc.Download(url)
}
