package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 0, "tcp port to connect to")
}

// generateRequest will create an HTTP request every 30 seconds
// and write it to the connection, it will then read the response.
func generateRequests(client net.Conn) {
	request, _ := http.NewRequest("GET", "/", nil)
	for {
		if err := request.Write(client); err != nil {
			fmt.Println("error: " + err.Error())
			return
		} else {
			reader := bufio.NewReader(client)
			resp, err := http.ReadResponse(reader, request)
			if err != nil {
				fmt.Println("error: " + err.Error())
				return
			} else {
				fmt.Println("debug: response received")
				resp.Body.Close()
			}
		}
		time.Sleep(time.Second * 30)
	}
}

func main() {
	flag.Parse()
	if port == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("error: " + err.Error())
		os.Exit(1)
	}

	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println("error: " + err.Error())
			os.Exit(1)
		}
		go generateRequests(client)
	}
}
