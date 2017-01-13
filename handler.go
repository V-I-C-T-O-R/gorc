package gorc

import (
	"log"
	"path"
	"sync"
)

var group sync.WaitGroup
var pi chan string = make(chan string, 2)
var exit chan bool = make(chan bool, 1)

func Download(url string) (err error) {
	assign(url)
	group.Add(1)
	go removeCache()
	log.Println("start download")
	for key, meta := range Context.fileNames {
		log.Println("file", key, "start")
		group.Add(1)
		go goBT(Context.file.url, key, meta)
	}
	exit <- true
	group.Wait()
	err = createFileOnly(Context.file.filePath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	log.Println("start unzip")
	for i := len(Context.tempList) - 1; i >= 0; i-- {
		err = appendToFile(Context.file.filePath, readFile(Context.tempList[i]))
		if err != nil {
			log.Println(err.Error())
			panic(err)
		}
	}
	for _, file := range Context.tempList {
		deleteFile(file)
	}
	log.Println("download completed")
	return
}
func goBT(url string, address string, b *block) {
	l, err := sendGet(url, address, b.start, b.end)
	if err != nil || l != (b.end-b.start+1) {
		b.lock.Lock()
		if b.count > 3 {
			pi <- b.id
		}
		if b.count <= 3 {
			group.Add(1)
			b.count++
			goBT(url, address, b)
		}
		b.lock.Unlock()
	}
	defer group.Done()
}
func removeCache() {
	for {
		select {
		case str := <-pi:
			p := path.Join(root, "lib", str)
			log.Println("file ", p, "download failed")
			deleteFile(p)
		case <-exit:
			break
		}
	}
	group.Done()
}
