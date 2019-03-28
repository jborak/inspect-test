package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

var (
	port int
	host string
)

func init() {
	flag.StringVar(&host, "host", "", "hostname or ip address to connect to")
	flag.IntVar(&port, "port", 0, "tcp port to connect to")
}

func main() {
	flag.Parse()
	if len(host) == 0 || port == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Connect to the host on the specified port.
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("error: " + err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(conn)
	numErrors := 0
	message := bytes.NewBufferString("Can you see this?")

	for numErrors < 5 {
		req, err := http.ReadRequest(reader)
		if err != nil {
			fmt.Println("error: " + err.Error())
			numErrors++
			continue
		}

		// Send a response.
		resp := http.Response{
			StatusCode: http.StatusOK,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Request:    req,
			Body:       ioutil.NopCloser(message),
		}
		if err := resp.Write(conn); err != nil {
			fmt.Println("error: " + err.Error())
			continue
		}
	}
}
