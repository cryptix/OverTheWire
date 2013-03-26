package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"syscall"
)

func main() {

	conn, err := net.Dial("tcp", "semtex.labs.overthewire.org:24073")

	switch err {
	case nil:

		content := bufio.NewReader(conn)

		for {
			line, err := content.ReadString('\n')
			fmt.Fprintf(os.Stderr, "%s", line)
			switch err {

			case nil:
				// do nothing
			case io.EOF:
				fmt.Fprintln(os.Stderr, "Done Reading Get.")
				os.Exit(0)

			default:
				fmt.Fprintln(os.Stderr, "Error During HTTP Get:%v\n", err)
				os.Exit(1)
			}

		}

	case io.EOF:
		fmt.Fprintln(os.Stderr, "Proxy EOF!")

	default:
		if e, ok := err.(*net.OpError); ok {
			// thanks to http://bravenewmethod.wordpress.com/2011/03/17/interpreting-go-socket-errors/
			fmt.Fprintf(os.Stderr, "ErrorType: %T : %v\n", e.Error(), e.Error())
			if e.Timeout() {
				fmt.Fprintln(os.Stderr, "TIMEOUT")
			}
			if e.Temporary() {
				fmt.Fprintln(os.Stderr, "TEMPORARY")
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
		}
	}
}
