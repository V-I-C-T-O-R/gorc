package main

import "github.com/V-I-C-T-O-R/gorc"

func main() {
	//var url = "https://github-windows.s3.amazonaws.com/GitHubSetup.exe"
	var url = "http://down.360safe.com/se/360se8.1.1.246.exe"
	//var url = "https://cdn.mysql.com//Downloads/MySQL-5.7/mysql-5.7.17-win32.zip"
	gorc.Download(url)
}
