package gorc

import (
	"crypto/tls"
	"github.com/coreos/go-log/log"
	"io/ioutil"
	"net/http"
)

const (
	HCP_ORC    string = "https://kyfw.12306.cn/otn/passcodeNew/getPassCodeNew?module=other&rand=sjrand&0.21191171556711197"
	HCP_LOGIN  string = ""
	HCP_QUERY  string = ""
	HCP_SUBMIT string = ""
	UA         string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.89 Safari/537.36"
)

func sendGet(url string, start int64, end int64) (content string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func sendHead(url string) (l string) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		log.Debug("create HEAD failed")
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Debug("HEAD response failed")
		return
	}
	defer resp.Body.Close()
	l = resp.Header.Get("Content-Length")
	return
}
