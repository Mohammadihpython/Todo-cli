package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"todocli/delivery/deliveryparam"
	"todocli/repository/memorystore"
	"todocli/service/task"
)

func main() {
	const (
		network = "tcp"
		address = ":2001"
	)

	// create new Listener
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("can not listen:", err)
	}
	defer listener.Close()

	fmt.Println("listening on", listener.Addr())

	taskMemoryRepo := memorystore.NewTaskStore()

	taskCategoryRepo := memorystore.TaskCategory{
		Task:     taskMemoryRepo,
		Category: nil,
	}
	taskService := task.NewService(taskCategoryRepo)

	for {
		connection, aERR := listener.Accept()
		if aERR != nil {
			log.Println("accept err:", aERR)

			// we put continue key  to continue to listening for other requests
			continue
		}
		var rawRequest = make([]byte, 1024)
		numberOfReadBytes, err := connection.Read(rawRequest)
		if err != nil {
			log.Println("read err:", err)
			continue
		}
		fmt.Printf("client address: %s, numOfReadBytes: %d, data: %s\n",
			connection.RemoteAddr(), numberOfReadBytes, string(rawRequest))
		req := &deliveryparam.Request{}
		if uErr := json.Unmarshal(rawRequest[:numberOfReadBytes], req); uErr != nil {
			log.Println("bad request", uErr)

			continue
		}
		switch req.Command {
		case "create_task":
			response, cErr := taskService.Create(task.CreateRequest{
				Title:               req.CreateTaskRequest.Title,
				DuDate:              req.CreateTaskRequest.DuDate,
				CategoryID:          req.CreateTaskRequest.CategoryId,
				AuthenticatedUserID: 0,
			})
			if cErr != nil {
				_, wErr := connection.Write([]byte(cErr.Error()))
				if wErr != nil {
					log.Println("can't write data to connection", wErr)

					continue
				}
			}
			data, mErr := json.Marshal(&response)
			if mErr != nil {
				_, wErr := connection.Write([]byte(mErr.Error()))
				if wErr != nil {
					log.Println("can't write data to connection", mErr)
					continue
				}
				continue

			}
			_, wErr := connection.Write(data)
			if wErr != nil {
				log.Println("can't write data to connection", wErr)

				continue
			}

		}
		connection.Close()
	}
}
