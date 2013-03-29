package main

import (
	"fmt"
	"github.com/hailiang/gosocks"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	conns := make([]net.Conn, 10)

	tor_con, err := net.Dial("tcp", "127.0.0.1:9051")
	if err != nil {
		fmt.Printf("Connecting to torcr failed: %s", err.Error())
		os.Exit(1)
	}
	//authenticate with torrc
	_, err = tor_con.Write([]byte("authenticate \"\"\n"))
	if err != nil {
		fmt.Printf("Writing auth to torcr failed: %s", err.Error())
		os.Exit(1)
	}
	//succesfully authenticated?
	tor_ok := make([]byte, 6)
	_, err = tor_con.Read(tor_ok)
	if err != nil {
		fmt.Printf("Reading torrc auth response failed: %s", err.Error())
		os.Exit(1)
	}
	if !strings.Contains(string(tor_ok), "250 OK") {
		fmt.Printf("Tor auth failed: %s", string(tor_ok))
		os.Exit(1)
	}

	//set up tor socks5 dialer
	Proxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")

	start := time.Now()

	//create 10 connections and do the semtex handshake
	for i := 0; i < len(conns); {
		conn, err := hs_viaproxy(Proxy)
		//was the handshake succesful?
		if err == nil {
			step := time.Now()
			fmt.Printf("%d Elapsed Time: %v\n", i+1, step.Sub(start))
			conns[i] = conn
			i += 1
		}
		//tell tor to route the next connection
		//through a new circuit
		new_circ(tor_con)
	}

	reply := make([]byte, 15)
	//loop over all the connections and try to read the password
	for {
		for i := 0; i < len(conns); i += 1 {
			//set short read timeout
			//only one connection will give us somthing to read
			//and waiting for the default timeout takes to long
			conns[i].SetReadDeadline(time.Now().Add(500 * time.Millisecond))
			_, err := conns[i].Read(reply)
			switch err {
			case nil:
				fmt.Printf("%d -> PASSWORD from server: \"%s\"\n", i, string(reply))
				for j := 0; j < len(conns); j += 1 {
					conns[i].Close()
				}
				os.Exit(0)
			case io.EOF:
				break
			default:
				fmt.Printf("%d. connection failed: %s\n", i, err.Error())
				os.Exit(1)
			}
		}
	}
}

func hs_viaproxy(proxy (func(string, string) (net.Conn, error))) (net.Conn, error) {
	semtext5 := "HELICOTRMA"
	servAddr := "semtex.labs.overthewire.org:24027"

	//connect to the semtext server
	conn, err := proxy("tcp", servAddr)

	if err != nil {
		fmt.Printf("Dial failed: %s\n", err.Error())
		os.Exit(1)
	}

	//read the 10 byte challenge
	reply := make([]byte, 10)

	_, err = conn.Read(reply)
	if err != nil {
		fmt.Printf("Read from server failed: %s (Circ. not ready OR same IP)\n", err.Error())
		return nil, err
	}

	fmt.Printf("From server: %s\n", reply)

	//compute the response...
	for i := 0; i < 10; i += 1 {
		reply[i] = reply[i] ^ semtext5[i]
	}

	//...and send it
	_, err = conn.Write([]byte(fmt.Sprintf("%sHARRYHACKT", reply)))
	if err != nil {
		fmt.Printf("Write to server failed: %s\n", err.Error())
		return nil, err
	}

	fmt.Print("Response sent.\n")

	return conn, nil
}

func new_circ(tor_con net.Conn) bool {
	//tell torcr to route the next connection through a new circ.
	_, err := tor_con.Write([]byte("signal newnym \"\"\n"))
	if err != nil {
		fmt.Printf("Writing newnym to torcr failed: %s", err.Error())
		os.Exit(1)
	}
	tor_ok := make([]byte, 6)
	_, err = tor_con.Read(tor_ok)
	if err != nil {
		fmt.Printf("Reading torcr newnym response failed: %s", err.Error())
		os.Exit(1)
	}
	//wait until the new circuit is established
	//8 seconds because TOR has a 7 sec NEWNYM delay
	time.Sleep(8 * time.Second)
	return true
}
