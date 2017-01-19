package main

import "github.com/V-I-C-T-O-R/gorc"

func main() {
	//var url = "http://down.xp696.com/17.1/win10_64/DEEP_Win10x64_cjb201701.rar"
	//var url = "https://github.com/yangyangwithgnu/goagent_out_of_box_yang/archive/master.zip"
	//var url = "http://down.ylmf123.com/17.1/win7_32/YLMF123_Win7x86_201701.rar"
	var url = "http://down.360safe.com/se/360se8.2.1.332.exe"
	gorc.Download(url)
}
