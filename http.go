package gorc

import (
	"crypto/tls"
	"github.com/coreos/go-log/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	HCP_ORC    string = "https://kyfw.12306.cn/otn/passcodeNew/getPassCodeNew?module=other&rand=sjrand&0.21191171556711197"
	HCP_LOGIN  string = ""
	HCP_QUERY  string = ""
	HCP_SUBMIT string = ""
	UA         string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.89 Safari/537.36"
)

type hvp struct {
	Cookies []*http.Cookie
	Content []byte
	Err     error
}

func httpRequest(request *http.Request) *hvp {
	h := &hvp{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Debug("get body failed")
		h.Err = err
		return h
	}
	h.Content = body
	h.Cookies = resp.Cookies()
	return h
}

func SendGet(url string) (h *hvp, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Debug("create request failed")
		return
	}
	h = httpRequest(req)
	return
}

func SendPost(url string, values url.Values) (h *hvp, err error) {
	var req *http.Request
	req, err = http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	if err != nil {
		log.Debug("create request failed")
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", UA)
	h = httpRequest(req)
	return
}
