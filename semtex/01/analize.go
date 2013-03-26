package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// lastLine := ""
	lineCnt := 0
	for {
		line, err := in.ReadString('\n')

		switch err {
		case nil:
			break
		case io.EOF:
			os.Exit(0)
		default:
			panic(err)
		}

		if lineCnt == 0 && strings.HasPrefix(line, "encrypting") {
			firstLine := strings.Split(line, "\"")
			fmt.Fprintf(os.Stderr, "input:%s\n", firstLine[1])
		}
		if strings.HasPrefix(line, "encryption finished:") {
			finishLine := strings.Split(line, ":")
			fmt.Fprintf(os.Stderr, "output:%s", finishLine[1])
		}

		// lastLine = line
		//fmt.Printf("%s", lastLine)

		lineCnt += 1
	}
}
