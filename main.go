package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	var conn net.Conn
	var err error
	var lastNotice time.Time

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("[IP]:[Port] required")
		return
	}

	for {
		conn, err = net.Dial("tcp", args[0])

		if err == nil {
			break
		}

		if lastNotice.IsZero() || time.Now().Sub(lastNotice).Seconds() >= 10 {
			fmt.Println("unable to obtain connection: " + err.Error())
			lastNotice = time.Now()
		}
	}
	defer conn.Close()
	fmt.Println("connected to server")
	reader := bufio.NewReader(conn)

	for {
		input, err := reader.ReadString('\n')

		if err != nil {

			if err == io.EOF {
				fmt.Println("connection closed by server")
			} else {
				fmt.Println(err)
			}
			break
		}
		message := string(input)
		fmt.Printf("%s", message)
	}

}
