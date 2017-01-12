package gorc

import "sync"

var group sync.WaitGroup

func doGet(url string) {

}
func download() {
	for key, meta := range Context.fileNames {
		group.Add(1)
		go goBT(Context.file.url, key, meta)
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
