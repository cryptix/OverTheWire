package main

import (
	"bufio"
	"fmt"
	"github.com/hailiang/gosocks"
	"io"
	"net"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func checkProxy(doneChan chan bool, address string) {
	start := time.Now()

	address = address[:len(address)-1]
	// fmt.Fprintf(os.Stderr, "Trying: %s\n", address)

	proxy := socks.DialSocksProxy(socks.SOCKS5, address)
	conn, err := proxy("tcp", "www.whatsmyip.net:80")

	switch err {
	case nil:
		fmt.Fprintf(os.Stderr, "Trying HTTP Get on: %s\n", address)

		fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: www.whatsmyip.net\r\nAccept: text/plain\r\n\r\n")
		content := bufio.NewReader(conn)

	httpGetLoop:
		for {
			line, err := content.ReadString('\n')
			// fmt.Fprintf(os.Stderr, "Body:%s", line)
			switch err {

			case nil:
				if strings.Contains(line, "Address is:") {
					done := time.Now()
					found := strings.Split(line, "\"")
					// fmt.Fprintf(os.Stderr, "Split:%#v\n", found)
					// fmt.Printf("IP: %s\n", found[7])
					fmt.Fprintf(os.Stderr, "ip: %s request took %v\n", found[7], done.Sub(start))
					fmt.Println(address)
					conn.Close()
					doneChan <- true
					break httpGetLoop
				}

			case io.EOF:
				fmt.Fprintln(os.Stderr, "Done Reading Get.")
				doneChan <- false
				return

			default:
				fmt.Fprintln(os.Stderr, "Error During HTTP Get:%v\n", err)
				doneChan <- false
				return
			}

		}

	case io.EOF:
		fmt.Fprintf(os.Stderr, "Proxy %s EOF!\n", address)
		doneChan <- false
		return
	default:
		if e, ok := err.(*net.OpError); ok {
			// thanks to http://bravenewmethod.wordpress.com/2011/03/17/interpreting-go-socket-errors/
			fmt.Fprintf(os.Stderr, "ErrorType: %T : %v\n", e.Error(), e.Error())
			if e.Timeout() {
				fmt.Fprintln(os.Stderr, "TIMEOUT")
				doneChan <- false
				return
			}
			if e.Temporary() {
				fmt.Fprintln(os.Stderr, "TEMPORARY")
				doneChan <- false
				return
			}

			// specific granular error codes in case we're interested
			switch e.Err {
			case syscall.EAGAIN:
				fmt.Fprintln(os.Stderr, "EAGAIN")
			case syscall.EPIPE:
				fmt.Fprintln(os.Stderr, "EPIPE") // broken pipe (e.g. on connection reset)
			case syscall.ECONNREFUSED:
				fmt.Fprintln(os.Stderr, "ECONNREFUSED")
			case syscall.ECONNRESET:
				fmt.Fprintln(os.Stderr, "ECONNRESET")
			default:
				fmt.Fprintf(os.Stderr, "%#v\n", e)

			}
			doneChan <- false
			return

		}

	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <proxyList>\n", os.Args[0])
		os.Exit(1)
	}

	proxyList, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer proxyList.Close()

	proxies := bufio.NewReader(proxyList)

	doneChan := make(chan bool, 1)
	proxyCount := 0
ReadLoop:
	for {
		proxyAddress, err := proxies.ReadString('\n')
		switch err {
		case nil:
			proxyCount += 1
			go checkProxy(doneChan, proxyAddress)
		case io.EOF:
			fmt.Fprint(os.Stderr, "Done Testing.")
			break ReadLoop
		default:
			fmt.Fprint(os.Stderr, "Reading proxy list error")
			panic(err)
		}
	}

	working := 0
	failed := 0
	for {
		select {
		case worked := <-doneChan:
			if worked {
				working += 1
			} else {
				failed += 1
			}
			proxyCount -= 1
			if proxyCount == 0 {
				fmt.Fprintf(os.Stderr, "Done! working:%d failed:%d\n", working, failed)
				os.Exit(0)
			}
		case <-time.After(1 * time.Second):
			fmt.Fprintf(os.Stderr, "Waiting.. %d left\n", proxyCount)
		}
	}
}
