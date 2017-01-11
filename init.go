package gorc

import "flag"

var thread int

func init() {
	flag.IntVar(&thread, "thread", 5, "concurrent thread number")
}
