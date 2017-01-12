package gorc

import (
	"log"
	"sync"
)

var group sync.WaitGroup

func Download(url string) {
	assign(url)
	for key, meta := range Context.fileNames {
		group.Add(1)
		go goBT(Context.file.url, key, meta)
	}
	group.Wait()
	err := createFileOnly(Context.file.filePath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	for _, value := range Context.tempList {
		appendToFile(Context.file.filePath, readFile(value))
	}
}
func goBT(url string, address string, b *block) {
	l, err := sendGet(url, address, b.start, b.end)
	if err != nil || l != (b.end-b.start+1) {
		group.Add(1)
		goBT(url, address, b)
	}
	defer group.Done()
}
