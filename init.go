package gorc

import (
	"flag"
	"os"
)

var thread int64
var manual bool
var root string
var cacheSize int

func init() {
	flag.Int64Var(&thread, "thread", 5, "concurrent thread number")
	flag.BoolVar(&manual, "manual", false, "specific thread number or not")
	root, _ = os.Getwd()
	flag.IntVar(&cacheSize, "cacheSize", 1024, "cache area size")
}
