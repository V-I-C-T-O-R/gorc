package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type session struct {
	Session []*Data `json:"session"`
}
type Data struct {
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
	Even      *event `json:"event"`
}
type event struct {
	Exits       int    `json:"exits"`
	Sessions    int    `json:"sessions"`
	Session_id  string `json:"session_id"`
	Page_url_id string `json:"page_url_id"`
	Bounces     int    `json:"bounces"`
	Pageviews   int    `json:"pageviews"`
}
type ids struct {
	Id string `json:"id"`
}

const PATH = "/home/wenhuanhuan/go/src/github.com/V-I-C-T-O-R/gorc/lib"

var sess session = session{}
var d ids = ids{}

func main() {
	LoadConf(PATH+"/ses.log", &sess)
	LoadConf(PATH+"/session.json", &d)
	fmt.Println("data == ", d.Id)
	str := strings.Split(d.Id, ",")
	fmt.Println(str)
	for _, value := range str {
		for _, event := range sess.Session {
			if strings.TrimSpace(value) == strings.TrimSpace(event.Even.Session_id) {
				content := event.Version + "," + event.Timestamp + "," + strconv.Itoa(event.Even.Exits) + "," + strconv.Itoa(event.Even.Sessions) + "," + event.Even.Session_id + "," + event.Even.Page_url_id + "," + strconv.Itoa(event.Even.Bounces) + "," + strconv.Itoa(event.Even.Pageviews) + "\r\n"
				appendToFile(PATH+"/session.csv", content)
			}
		}

	}
}
func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

//加载json（可配置扩展字段）配置文件
func LoadConf(filePath string, v interface{}) error {
	data, err := readFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func readFile(filePath string) (data []byte, err error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return
	}
	if fi.IsDir() {
		err = fmt.Errorf(filePath + " is not a file.")
		return
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	return b, err
}
