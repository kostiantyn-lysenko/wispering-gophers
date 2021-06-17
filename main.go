package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Body string
}

func main() {
	message := new(Message)
	fmt.Fscan(os.Stdin, &message.Body)

	out, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(out))
}
