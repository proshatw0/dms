package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"

	"dms/src/workfile"
)

func main() {
	address := "10.241.125.222:6379"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error when starting the server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("The server is listening on:", address)

	var mutex sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Error accepting connection:", err)
					continue
				}
				go handleConnection(conn, &mutex)
			}
		}()
	}
	wg.Wait()
}

func handleConnection(conn net.Conn, mutex *sync.Mutex) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	filepath, err := reader.ReadString('\n')
	filepath = strings.ReplaceAll(filepath, "\n", "")
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	commands := [4]string{"", "", "", ""}
	for i := 0; i < 4; i++ {
		commands[i], err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}
		commands[i] = strings.ReplaceAll(commands[i], "\n", "")
	}
	mutex.Lock()

	err, value, integer := workfile.Processing_Request("../data/"+filepath, commands)

	mutex.Unlock()

	if err != nil {
		response := fmt.Sprintln(err)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
	}

	if integer != -1 && value != "" {
		out := "--> " + strconv.Itoa(integer) + " ~ " + value
		response := fmt.Sprintln(out)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
	}

	if integer != -1 && value == "" {
		out := "--> " + strconv.Itoa(integer)
		response := fmt.Sprintln(out)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
	}
	if value != "" && integer == -1 {
		out := "--> " + value
		response := fmt.Sprintln(out)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
	}
}
