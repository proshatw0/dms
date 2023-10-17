package main

import (
	"flag"
	"fmt"
	"strings"

	"dms/src/workfile"
)

func main() {
	filepath := flag.String("file", "", "filepath")
	command := flag.String("query", "", "command")

	flag.Parse()
	if *command != "" && *filepath != "" && strings.ReplaceAll(*command, " ", "") != "" {
		arr := strings.Fields(*command)
		arr[0] = strings.ToLower(arr[0])
		commands := [4]string{"", "", "", ""}
		for i := 0; i < len(arr); i++ {
			commands[i] = arr[i]
		}
		workfile.Processing_Request("../data/"+*filepath, commands)
	} else {
		fmt.Println("-->invalid request")
		fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
	}
}
