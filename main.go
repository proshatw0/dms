package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"dms/src/structs"
)

// file
func Read_Line_Fromfile(path string, line_num int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		if lineCount == line_num {
			return scanner.Text(), nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return "", fmt.Errorf("-->line not found")
}

func WriteLineFromFile(path string, lineNum int, data string) error {
	lines, err := ReadLinesFromFile(path)
	if err != nil {
		return err
	}

	if lineNum < 1 || lineNum > len(lines) {
		return errors.New("-->table not found")
	}

	lines[lineNum-1] = data

	return WriteLinesToFile(path, lines)
}

func ReadLinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func WriteLinesToFile(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func Search_Number_Table(line string, name_table string) int {
	startIndex := strings.Index(line, name_table) + len(name_table) + 2
	endIndex := strings.Index(line[startIndex:], "}")
	valueStr := line[startIndex : startIndex+endIndex]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0
	}
	return value
}

func Search_Table(filepath string, command string, name_table string) (string, int) {
	switch command {
	case "sadd", "srem", "sismember":
		line, err := Read_Line_Fromfile(filepath, 1)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "set", 0
		}
		return "set", number_line
	case "spush", "spop":
		line, err := Read_Line_Fromfile(filepath, 2)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "stack", 0
		}
		return "stack", number_line
	case "qpush", "qpop":
		line, err := Read_Line_Fromfile(filepath, 3)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "queue", 0
		}
		return "queue", number_line
	case "hset", "hdel", "hget":
		line, err := Read_Line_Fromfile(filepath, 4)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "hash_table", 0
		}
		return "hash_table", number_line
	default:
		return "", 0
	}
}

func Scan_Table_Set(filepath string, line_number int) structs.Set {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return structs.Set{}
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "}")
	sizeIndex := strings.Index(line, ",")
	size := line[startIndex:sizeIndex]
	val := line[sizeIndex+len(size)+2 : endIndex]
	arr := strings.Split(val, ",")
	size_int, _ := strconv.Atoi(size)
	set := structs.NewSet(size_int)
	for i := 0; i < len(arr); i++ {
		set.Sadd(strings.ReplaceAll(arr[i], " ", ""))
	}
	return set
}

func Scan_Table_Stack(filepath string, line_number int) structs.Stack {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return structs.Stack{}
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "]")
	val := line[startIndex:endIndex]
	arr := strings.Split(val, ",")
	stack := structs.Stack{}
	for i := len(arr) - 1; i >= 0; i-- {
		stack.Spush(strings.ReplaceAll(arr[i], " ", ""))
	}
	return stack
}
func Scan_Table_Queue(filepath string, line_number int) structs.Queue {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return structs.Queue{}
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "]")
	val := line[startIndex:endIndex]
	arr := strings.Split(val, ",")
	queue := structs.Queue{}
	for i := 0; i < len(arr); i++ {
		queue.Qpush(strings.ReplaceAll(arr[i], " ", ""))
	}
	return queue
}

func Scan_Table_Hash_Table(filepath string, line_number int) structs.Hash_Table {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return structs.Hash_Table{}
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "}")
	sizeIndex := strings.Index(line, ",")
	size := line[startIndex:sizeIndex]
	val := line[sizeIndex+len(size)+2 : endIndex]
	val = strings.ReplaceAll(val, ")", "")
	val = strings.ReplaceAll(val, "(", "")
	arr := strings.Split(val, ",")
	size_int, _ := strconv.Atoi(size)
	hash_table := structs.NewHashTable(size_int)
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			key := strings.ReplaceAll(arr[i], " ", "")
			i++
			value := strings.ReplaceAll(arr[i], " ", "")
			hash_table.Hset(key, value)
		}
	}
	return hash_table
}

func Print_Table_Set(filepath string, line_number int, name_table string, set structs.Set) error {
	out := name_table
	out += ": [" + strconv.Itoa(set.Size) + ", {"
	for index := 0; index < set.Size; index++ {
		currentNode := set.Table[index].Head
		for currentNode != nil {
			out += currentNode.Data + ", "
			currentNode = currentNode.Next
		}
	}
	if len(out) == len(name_table)+len(strconv.Itoa(set.Size))+6 {
		out += "}]"
	} else {
		out = out[0:len(out)-2] + "}]"
	}
	err := WriteLineFromFile(filepath, line_number, out)
	if err != nil {
		return err
	}
	return nil
}

func Print_Table_Stack(filepath string, line_number int, name_table string, stack structs.Stack) error {
	out := name_table
	out += ": ["
	currentNode := stack.Head
	for currentNode != nil {
		out += currentNode.Data + ", "
		currentNode = currentNode.Next
	}
	if len(out) == len(name_table)+3 {
		out = out + "]"
	} else {
		out = out[0:len(out)-2] + "]"
	}
	err := WriteLineFromFile(filepath, line_number, out)
	if err != nil {
		return err
	}
	return nil
}

func Print_Table_Queue(filepath string, line_number int, name_table string, queue structs.Queue) error {
	out := name_table
	out += ": ["
	currentNode := queue.Head
	for currentNode != nil {
		if currentNode.Data != "" {
			out += currentNode.Data + ", "
		}
		currentNode = currentNode.Next
	}
	if len(out) == len(name_table)+3 {
		out = out + "]"
	} else {
		out = out[0:len(out)-2] + "]"
	}
	err := WriteLineFromFile(filepath, line_number, out)
	if err != nil {
		return err
	}
	return nil
}

func Print_Table_Hash_Table(filepath string, line_number int, name_table string, ht structs.Hash_Table) error {
	out := name_table
	out += ": [" + strconv.Itoa(ht.Size) + ", {"
	for index := 0; index < ht.Size; index++ {
		currentNode := ht.Table[index].Head
		for currentNode != nil {
			out += "(" + currentNode.Data.Key + ", " + currentNode.Data.Value + "), "
			currentNode = currentNode.Next
		}
	}
	if len(out) == len(name_table)+len(strconv.Itoa(ht.Size))+6 {
		out += "}]"
	} else {
		out = out[0:len(out)-2] + "}]"
	}
	err := WriteLineFromFile(filepath, line_number, out)
	if err != nil {
		return err
	}
	return nil
}

func Processing_Request(filepath string, commands [4]string) {
	mod, number_line := Search_Table(filepath, commands[0], commands[1])
	if mod == "" {
		fmt.Println("--> unknown command")
		return
	}
	if number_line != 0 {
		switch mod {
		case "set":
			set := Scan_Table_Set(filepath, number_line)
			if set.Size <= 0 {
				fmt.Println("-->table not found")
				return
			}
			switch commands[0] {
			case "sadd":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				err := set.Sadd(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Set(filepath, number_line, commands[1], set)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "srem":
				err := set.Srem(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Set(filepath, number_line, commands[1], set)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "sismember":
				err := set.Sismember(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("--> True")
			}
		case "stack":
			stack := Scan_Table_Stack(filepath, number_line)
			switch commands[0] {
			case "spush":
				err := stack.Spush(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Stack(filepath, number_line, commands[1], stack)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "spop":
				value, err := stack.Spop()
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Stack(filepath, number_line, commands[1], stack)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", value)
				return
			}
		case "queue":
			queue := Scan_Table_Queue(filepath, number_line)
			switch commands[0] {
			case "qpush":
				err := queue.Qpush(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Queue(filepath, number_line, commands[1], queue)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "qpop":
				value, err := queue.Qpop()
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Queue(filepath, number_line, commands[1], queue)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", value)
				return
			}
		case "hash_table":
			if commands[2] == "" {
				fmt.Println("-->invalid request")
				fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
				return
			}
			hash_table := Scan_Table_Hash_Table(filepath, number_line)
			if hash_table.Size <= 0 {
				fmt.Println("-->table not found")
				return
			}
			switch commands[0] {
			case "hset":
				if commands[3] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				err := hash_table.Hset(commands[2], commands[3])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Hash_Table(filepath, number_line, commands[1], hash_table)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2], "~", commands[3])
				return
			case "hdel":
				value, err := hash_table.Hdel(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Hash_Table(filepath, number_line, commands[1], hash_table)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2], "~", value)
				return
			case "hget":
				value, err := hash_table.Hget(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2], "~", value)
			}
		default:
			fmt.Println("-->table not found")
			return
		}
	} else {
		fmt.Println("-->table not found")
		return
	}
}

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
		Processing_Request("../data/"+*filepath, commands)
	} else {
		fmt.Println("-->invalid request")
		fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
	}
}
