package main

import (
	"bufio"
	"io"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	i := 0

	for {
		b, err := in.ReadByte()

		// handle error
		switch err {
		case nil:
			break
		case io.EOF:
			// write out data left in buffer
			out.Flush()
			os.Exit(0)
		default:
			panic(err)
		}

		// skip odd bytes
		if i%2 == 0 {
			out.WriteByte(b)
		}
		i += 1
	}

}
