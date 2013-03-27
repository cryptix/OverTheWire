// 654f 2fe8
// fa89 07f4
// 1d79 7f91
// 0971 5609
package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

type pwInts struct {
	Int1 uint32
	Int2 uint32
	Int3 uint32
	Int4 uint32
}

func (p *pwInts) sum() uint32 {
	return p.Int1 + p.Int2 + p.Int3 + p.Int4
}

func main() {
	conn, err := net.Dial("tcp", "vortex.labs.overthewire.org:5842")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during connect: %s\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	bufRead := bufio.NewReader(conn)

	var ints pwInts

IntLoop:
	for {
		err := binary.Read(bufRead, binary.LittleEndian, &ints)

		switch err {
		case nil:

			fmt.Printf("ints:%#d\n", ints)
			fmt.Println("Sum:", ints.sum())
			binary.Write(conn, binary.LittleEndian, ints.sum())
			break IntLoop

		case io.EOF:
			fmt.Fprintln(os.Stderr, "conn eof!")

		default:
			fmt.Fprintf(os.Stderr, "Error?! %T - %s\n", err, err.Error())
			os.Exit(1)
		}
	}

	pwBuf := make([]byte, 4096)
	for {
		_, err := bufRead.Read(pwBuf)

		switch err {
		case nil:
			fmt.Printf("%s\n", pwBuf)
			os.Exit(0)

		case io.EOF:
			fmt.Fprintln(os.Stderr, "PW EOF! Done?")
			os.Exit(0)
		default:
			fmt.Fprintf(os.Stderr, "Error?! %s\n", err.Error())
		}

	}

}
