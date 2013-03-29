// http://www.thesprawl.org/research/tor-control-protocol/#creating-custom-circuits
// http://www.thesprawl.org/projects/tor-autocircuit/
// http://schoolofprivacy.eu/post/43978479406/howto-use-tor-for-all-network-traffic-by-default-on

package main

import (
	"bufio"
	"fmt"
	"github.com/hailiang/gosocks"
	"io"
	"time"
	// "net/http"
	"os"
	"strings"
)

func main() {
	torProxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")

	// http method
	// tr := &http.Transport{Dial: torProxy}
	// httpClient := &http.Client{Transport: tr}
	// resp, err := httpClient.Get("http://www.whatsmyip.net")
	// if err != nil {
	// 	panic(err)
	// }
	// defere resp.Close
	// if resp.StatusCode != 200 {
	// 	fmt.Printf("http code: %d\n", resp.StatusCode)
	// 	os.Exit(1)
	// }
	// //io.Copy(os.Stderr, resp.Body)
	// content := bufio.NewReader(resp.Body)
	start := time.Now()
	// tcp method
	conn, err := torProxy("tcp", "www.whatsmyip.net:80")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: www.whatsmyip.net\r\nAccept: text/plain\r\n\r\n")
	content := bufio.NewReader(conn)

	for {
		line, err := content.ReadString('\n')
		// fmt.Fprintf(os.Stderr, "Body:%s", line)
		switch err {
		case nil:
			if strings.Contains(line, "Address is:") {
				done := time.Now()
				found := strings.Split(line, "\"")
				// fmt.Fprintf(os.Stderr, "Split:%#v\n", found)
				fmt.Printf("IP: %s\n", found[7])
				fmt.Fprintf(os.Stderr, "ip: %s request took %v\n", found[7], done.Sub(start))
				conn.Close()
				os.Exit(0)
			}
		case io.EOF:
			fmt.Fprint(os.Stderr, "Done.")
			os.Exit(0)
		default:
			panic(err)
		}

	}
}
