package gorc

import (
	"crypto/tls"
	"github.com/coreos/go-log/log"
	"github.com/hu17889/go_spider/core/common/request"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func sendGet(url string, start int64, end int64, file *os.File) (len int64, err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("Range", "bytes="+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var resp *http.Response
	resp, err = client.Do(req)
	defer resp.Body.Close()
	len, err = io.Copy(file, resp.Body)
	return len, err
}

func sendHead(url string) (l string, err error) {
	var req *http.Request
	req, err = http.NewRequest("HEAD", url, nil)
	if err != nil {
		log.Debug("create HEAD failed")
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		log.Debug("HEAD response failed")
		return
	}
	defer resp.Body.Close()
	l = resp.Header.Get("Content-Length")
	return
}
