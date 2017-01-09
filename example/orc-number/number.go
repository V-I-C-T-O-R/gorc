package main

import (
	"fmt"
	"github.com/V-I-C-T-O-R/gorc"
	"github.com/coreos/go-log/log"
	"github.com/otiai10/gosseract"
)

func main() {
	hvp, err := gorc.SendGet(gorc.HCP_ORC)
	if err != nil && hvp.Err != nil {
		log.Info("get picture failed")
		return
	}
	err = gorc.ResolvePng(hvp.Content)
	if err != nil {
		log.Info("resolve picture failed")
	}
	out := gosseract.Must(gosseract.Params{
		Src:       gorc.File,
		Languages: "eng+heb",
	})
	fmt.Println(out)
	client, _ := gosseract.NewClient()
	out, _ = client.Src(gorc.File).Out()
	fmt.Println(out)
}
