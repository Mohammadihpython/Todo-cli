package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"todocli/delivery/deliveryparam"
)

func main() {
	fmt.Println("command", os.Args[0])
	if len(os.Args) < 2 {
		log.Fatalln("you should set ip address of server")
	}
	serverAddress := os.Args[1]

	message := "default message"
	if len(os.Args) > 2 {
		message = os.Args[2]
	}
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalln("cant dial the given address")

	}
	defer conn.Close()
	fmt.Println("local address", conn.LocalAddr())
	req := deliveryparam.Request{Command: message}

	if req.Command == "create-task" {
		req.CreateTaskRequest = deliveryparam.CreateTaskRequest{
			Title:      "test",
			DuDate:     "test",
			CategoryId: 1,
		}
		serializedData, mErr := json.Marshal(req.CreateTaskRequest)
		if mErr != nil {
			log.Fatalln("cant marshal create task request", mErr)

		}
		numberOfWriteBytes, cErr := conn.Write(serializedData)
		if cErr != nil {
			log.Fatalln("cant write create task request", cErr)

		}
		fmt.Println("number of write bytes", numberOfWriteBytes)

		var data = make([]byte, 1024)
		_, rErr := conn.Read(data)
		if rErr != nil {
			log.Fatalln("cant read create task request", rErr)
		}
		fmt.Println("server Response ", string(data))

	}
}
