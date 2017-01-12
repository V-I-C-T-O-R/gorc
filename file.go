package gorc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	LEVEL = 1024
)

type File struct {
	url      string
	name     string
	length   int64
	filePath string
}
type block struct {
	previous *block
	id       string
	start    int64
	end      int64
}
type context struct {
	fileNames map[string]*block
	lock      *sync.Mutex
	file      *File
	tempList  []string
}

var Context *context = new(context)

func openfile(c *context) *os.File {
	return nil
}
func assign(url string) {
	tName, fName := searchName(url)
	length, err := sendHead(url)
	if err != nil {
		log.Println("get file length failed")
		return
	}
	f := &File{url: url, name: fName, length: length, filePath: path.Join(root, "lib", fName)}
	Context.file = f
	l, _ := strconv.ParseInt(length, 10, 64)
	var element *block
	if manual {
		element = partFileManual(l, thread, tName)
		assignBlock(element)
		return
	}
	if l/(LEVEL*LEVEL*32) == 0 {
		element = partFileManual(l, thread, tName)
		assignBlock(element)
		return
	}
	element = partFile(l, 0, l-1)
	assignBlock(element)
}
func searchName(url string) (tmpName, fullName string) {
	u := []byte(url)
	s := strings.LastIndex(url, "/")
	if s == -1 {
		s = 0
		fullName = u[s:]
	} else {
		fullName = u[s+1:]
	}
	t := []byte(fullName)
	d := strings.LastIndex(fullName, ".")
	if d == -1 {
		d = len(t)
		tmpName = string(t[:])
	} else {
		tmpName = string(t[:d])
	}
	return
}

func assignBlock(b *block) {
	if b == nil {
		return
	}
	m := make(map[string]*block)
	listId := []string{}
	p := path.Join(root, "lib", b.id)
	m[p] = b
	listId = append(listId, b.id)
	if b.previous != nil {
		b = b.previous
		p = path.Join(root, "lib", b.id)
		m[p] = b
		listId = append(listId, p)
	}
	Context.fileNames = m
	listNames := []string{}
	for i := len(listId) - 1; i >= 0; i++ {
		addr, _ := Context.fileNames[listId[i]]
		listNames = append(listNames, addr)
	}
	Context.tempList = listNames
}
func partFile(length int64, start int64, end int64) *block {
	if length/(LEVEL*LEVEL*32) > 0 && length/(LEVEL*LEVEL*LEVEL) == 0 {
		length = length - LEVEL*LEVEL*32 + 1
		return &block{id: GetRandomSalt(), start: length - 1, end: end, previous: partFile(length, start, length-1)}
	}
	if length/(LEVEL*LEVEL*LEVEL) > 0 && length/(LEVEL*LEVEL*LEVEL*LEVEL) == 0 {
		length = length - LEVEL*LEVEL*32 + 1
		return &block{id: GetRandomSalt(), start: length - 1, end: end, previous: partFile(length, start, length-1)}
	}
	if length/(LEVEL*LEVEL*LEVEL*LEVEL) > 0 {
		length = length - LEVEL*LEVEL*32 + 1
		return &block{id: GetRandomSalt(), start: length - 1, end: end, previous: partFile(length, start, length-1)}
	}
	return &block{id: GetRandomSalt(), start: start, end: end}
}

func partFileManual(length int64, thread int64, name string) (b *block) {
	blockSize := length / thread
	b = new(block)
	var start int64
	var i int64
	for i = 1; i <= thread; i++ {
		var seg = new(block)
		r := MD5(name + strconv.FormatInt(i, 10))
		seg.id = r
		seg.previous = b
		seg.start = start
		if blockSize*i < length {
			seg.end = blockSize * i
		} else {
			seg.end = blockSize*i - 1
		}
		start = blockSize * i
		b = seg
	}
	return b
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func createFile(file string) (f *os.File, err error) {
	if checkFileStat(file) {
		log.Println(file, "文件存在")
		f, err = os.OpenFile(file, os.O_RDWR, 0666)
		return f, err
	}
	f, err = os.Create(file)
	if err != nil {
		log.Println(file, "文件创建失败")
	}
	return file, err
}
func createFileOnly(file string) error {
	if checkFileStat(file) {
		log.Println(file, "文件存在")
		return nil
	}
	f, err := os.Create(file)
	if err != nil {
		log.Println(file, "文件创建失败")
	}
	defer f.Close()
	return err
}
func deleteFile(file string) error {
	if !checkFileStat(file) {
		log.Println(file, "文件不存在")
		return nil
	}
	err := os.Remove(file)
	if err != nil {
		log.Println(file, "文件删除失败")
	}
	return err
}
func checkFileStat(file string) bool {
	var exist = true
	if _, err := os.Stat(file); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		log.Println("file append failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
func readFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}
