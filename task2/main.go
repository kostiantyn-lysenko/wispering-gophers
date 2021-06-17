package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Message struct {
	Body string
	To   string
}

func main() {
	message := new(Message)
	fmt.Println("Enter remote host address: ")
	fmt.Fscan(os.Stdin, &message.To)

	fmt.Println("Enter a message: ")
	fmt.Fscan(os.Stdin, &message.Body)

	conn, err := net.Dial("tcp", message.To)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	out, err := json.Marshal(message.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	outString := string(out)

	fmt.Fprintf(conn, "%s", outString)
}
