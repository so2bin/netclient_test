package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {
	service := "127.0.0.1:6543"
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	//checkError(err)
	fmt.Println("connect ", service)
	conn, err := net.DialTimeout("tcp", service, 10*time.Second)
	checkError(err)
	_, err = conn.Write([]byte("hello world"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	fmt.Println("close socket connection...")
	conn.Close()
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err)
		os.Exit(1)
	}
}
