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
)

// Set
type Set struct {
	Table []*Doubly_Connected_Set
	Size  int
}

type Node_Set struct {
	data     string
	next     *Node_Set
	previous *Node_Set
}

type Doubly_Connected_Set struct {
	lenght int
	head   *Node_Set
	tail   *Node_Set
}

// Hash_Table
type Hash_Table struct {
	Table []*Doubly_Connected_Table
	Size  int
}

type Pair struct {
	key   string
	value string
}

type Node_Table struct {
	data     Pair
	next     *Node_Table
	previous *Node_Table
}

type Doubly_Connected_Table struct {
	lenght int
	head   *Node_Table
	tail   *Node_Table
}

// Stack/Queue
type Node struct {
	data string
	next *Node
}

type Stack struct {
	head *Node
}

type Queue struct {
	head *Node
	tail *Node
}

// Set
func NewSet(size int) Set {
	set := make([]*Doubly_Connected_Set, size)
	for i := range set {
		set[i] = &Doubly_Connected_Set{}
	}
	return Set{
		Table: set,
		Size:  size,
	}
}

func (ht *Set) Hash_Set(key string) int {
	key_int := 0
	for _, symbol := range key {
		key_int += int(symbol)
	}
	return key_int % ht.Size
}

func (ht *Set) sadd(key string) error {
	hash := ht.Hash_Set(key)
	if ht.Table[hash].lenght < 20 {
		return ht.Table[hash].dspush(key)
	} else {
		oldSize := ht.Size
		newHT := NewSet(oldSize * 2)
		for i := 0; i < oldSize; i++ {
			currentNode := ht.Table[i].head
			for currentNode != nil {
				new_hash := newHT.Hash_Set(currentNode.data)
				newHT.Table[new_hash].dspush(currentNode.data)
				currentNode = currentNode.next
			}
		}
		*ht = newHT
		new_hash := ht.Hash_Set(key)
		return ht.Table[new_hash].dspush(key)
	}
}

func (ht *Set) srem(key string) error {
	hash := ht.Hash_Set(key)
	return ht.Table[hash].dsdel(key)
}

func (ht *Set) sismember(key string) error {
	hash := ht.Hash_Set(key)
	currentNode := ht.Table[hash].head
	for currentNode != nil {
		if currentNode.data == key {
			return nil
		}
		currentNode = currentNode.next
	}
	return errors.New("--> False")
}

func (doubly_connected *Doubly_Connected_Set) dspush(val string) error {
	node_list := &Node_Set{data: val}
	if doubly_connected.head == nil {
		doubly_connected.head = node_list
		doubly_connected.tail = node_list
	} else {
		currentNode := doubly_connected.head
		for currentNode != nil {
			if currentNode.data == val {
				return errors.New("--> key already exists")
			}
			currentNode = currentNode.next
		}
		doubly_connected.tail.next = node_list
		node_list.previous = doubly_connected.tail
		doubly_connected.tail = node_list
	}
	doubly_connected.lenght++
	return nil
}

func (doubly_connected *Doubly_Connected_Set) dsdel(val string) error {
	currentNode := doubly_connected.head
	for currentNode != nil {
		if currentNode.data == val {
			if currentNode == doubly_connected.head {
				doubly_connected.head = currentNode.next
				if doubly_connected.head != nil {
					doubly_connected.head.previous = nil
				}
			} else if currentNode == doubly_connected.tail {
				doubly_connected.tail = currentNode.previous
				if doubly_connected.tail != nil {
					doubly_connected.tail.next = nil
				}
			} else {
				currentNode.previous.next = currentNode.next
				currentNode.next.previous = currentNode.previous
			}
			doubly_connected.lenght--
			return nil
		}

		currentNode = currentNode.next
	}
	return errors.New("--> key not founde")
}

// Hash_Table
func NewHashTable(size int) Hash_Table {
	table := make([]*Doubly_Connected_Table, size)
	for i := range table {
		table[i] = &Doubly_Connected_Table{}
	}
	return Hash_Table{
		Table: table,
		Size:  size,
	}
}

func (ht *Hash_Table) hash(key string) int {
	key_int := 0
	for _, symbol := range key {
		key_int += int(symbol)
	}
	return key_int % ht.Size
}

func (ht *Hash_Table) hset(key string, value string) error {
	val := &Pair{key: key, value: value}
	hash := ht.hash(val.key)
	if ht.Table[hash].lenght < 20 {
		return ht.Table[hash].dpush(*val)
	} else {
		oldSize := ht.Size
		newHT := NewHashTable(oldSize * 2)
		for i := 0; i < oldSize; i++ {
			currentNode := ht.Table[i].head
			for currentNode != nil {
				new_hash := newHT.hash(currentNode.data.key)
				newHT.Table[new_hash].dpush(currentNode.data)
				currentNode = currentNode.next
			}
		}
		*ht = newHT
		new_hash := ht.hash(val.key)
		return ht.Table[new_hash].dpush(*val)
	}
}

func (ht *Hash_Table) hdel(key string) (string, error) {
	hash := ht.hash(key)
	pair, err := ht.Table[hash].ddel(key)
	return pair.value, err
}

func (ht *Hash_Table) hget(key string) (string, error) {
	hash := ht.hash(key)
	currentNode := ht.Table[hash].head
	for currentNode != nil {
		if currentNode.data.key == key {
			return currentNode.data.value, nil
		}
		currentNode = currentNode.next
	}
	return "", errors.New("-->element not found")
}

func (pair *Pair) ppush(key string, value string) {
	pair.key = key
	pair.value = value
}

func (doubly_connected *Doubly_Connected_Table) dpush(val Pair) error {
	node_hesh := &Node_Table{data: val}
	if doubly_connected.head == nil {
		doubly_connected.head = node_hesh
		doubly_connected.tail = node_hesh
	} else {
		currentNode := doubly_connected.head
		for currentNode != nil {
			if currentNode.data.key == val.key {
				return errors.New("--> key already exists")
			}
			currentNode = currentNode.next
		}
		doubly_connected.tail.next = node_hesh
		node_hesh.previous = doubly_connected.tail
		doubly_connected.tail = node_hesh
	}
	doubly_connected.lenght++
	return nil
}

func (doubly_connected *Doubly_Connected_Table) ddel(val string) (Pair, error) {
	currentNode := doubly_connected.head
	if currentNode == nil {
		return Pair{}, errors.New("-->list is clear")
	}
	for currentNode != nil {
		if currentNode.data.key == val {
			if currentNode == doubly_connected.head {
				doubly_connected.head = currentNode.next
				if doubly_connected.head != nil {
					doubly_connected.head.previous = nil
				}
			} else if currentNode == doubly_connected.tail {
				doubly_connected.tail = currentNode.previous
				if doubly_connected.tail != nil {
					doubly_connected.tail.next = nil
				}
			} else {
				currentNode.previous.next = currentNode.next
				currentNode.next.previous = currentNode.previous
			}
			doubly_connected.lenght--
			return currentNode.data, nil
		}

		currentNode = currentNode.next
	}
	return Pair{}, errors.New("--> key not founde")
}

// Stack
func (stack *Stack) spush(val string) error {
	if val == "" {
		return errors.New("-->unknown command")
	}
	node := &Node{data: val}
	if stack.head == nil {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
	return nil
}

func (stack *Stack) spop() (string, error) {
	if stack.head == nil {
		return "", errors.New("--> stack is empty")
	} else {
		val := stack.head.data
		stack.head = stack.head.next
		return val, nil
	}
}

// Queue
func (queue *Queue) qpush(val string) error {
	if val == "" {
		return errors.New("-->unknown command")
	}
	node := &Node{data: val}
	if queue.head == nil {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
	return nil
}

func (queue *Queue) qpop() (string, error) {
	if queue.head == nil {
		return "", errors.New("--> queue is empty")
	} else {
		val := queue.head.data
		queue.head = queue.head.next
		return val, nil
	}
}

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

func Scan_Table_Set(filepath string, line_number int) Set {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return Set{}
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "}")
	sizeIndex := strings.Index(line, ",")
	size := line[startIndex:sizeIndex]
	val := line[sizeIndex+len(size)+2 : endIndex]
	arr := strings.Split(val, ",")
	size_int, _ := strconv.Atoi(size)
	set := NewSet(size_int)
	for i := 0; i < len(arr); i++ {
		set.sadd(strings.ReplaceAll(arr[i], " ", ""))
	}
	return set
}

func Scan_Table_Stack(filepath string, line_number int) Stack {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return Stack{}
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "]")
	val := line[startIndex:endIndex]
	arr := strings.Split(val, ",")
	stack := Stack{}
	for i := len(arr) - 1; i >= 0; i-- {
		stack.spush(strings.ReplaceAll(arr[i], " ", ""))
	}
	return stack
}
func Scan_Table_Queue(filepath string, line_number int) Queue {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return Queue{}
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "]")
	val := line[startIndex:endIndex]
	arr := strings.Split(val, ",")
	queue := Queue{}
	for i := 0; i < len(arr); i++ {
		queue.qpush(strings.ReplaceAll(arr[i], " ", ""))
	}
	return queue
}

func Scan_Table_Hash_Table(filepath string, line_number int) Hash_Table {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return Hash_Table{}
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
	hash_table := NewHashTable(size_int)
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			key := strings.ReplaceAll(arr[i], " ", "")
			i++
			value := strings.ReplaceAll(arr[i], " ", "")
			hash_table.hset(key, value)
		}
	}
	return hash_table
}

func Print_Table_Set(filepath string, line_number int, name_table string, set Set) error {
	out := name_table
	out += ": [" + strconv.Itoa(set.Size) + ", {"
	for index := 0; index < set.Size; index++ {
		currentNode := set.Table[index].head
		for currentNode != nil {
			out += currentNode.data + ", "
			currentNode = currentNode.next
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

func Print_Table_Stack(filepath string, line_number int, name_table string, stack Stack) error {
	out := name_table
	out += ": ["
	currentNode := stack.head
	for currentNode != nil {
		out += currentNode.data + ", "
		currentNode = currentNode.next
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

func Print_Table_Queue(filepath string, line_number int, name_table string, queue Queue) error {
	out := name_table
	out += ": ["
	currentNode := queue.head
	for currentNode != nil {
		if currentNode.data != "" {
			out += currentNode.data + ", "
		}
		currentNode = currentNode.next
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

func Print_Table_Hash_Table(filepath string, line_number int, name_table string, ht Hash_Table) error {
	out := name_table
	out += ": [" + strconv.Itoa(ht.Size) + ", {"
	for index := 0; index < ht.Size; index++ {
		currentNode := ht.Table[index].head
		for currentNode != nil {
			out += "(" + currentNode.data.key + ", " + currentNode.data.value + "), "
			currentNode = currentNode.next
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
				err := set.sadd(commands[2])
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
				err := set.srem(commands[2])
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
				err := set.sismember(commands[2])
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
				err := stack.spush(commands[2])
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
				value, err := stack.spop()
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
				err := queue.qpush(commands[2])
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
				value, err := queue.qpop()
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
				err := hash_table.hset(commands[2], commands[3])
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
				value, err := hash_table.hdel(commands[2])
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
				value, err := hash_table.hget(commands[2])
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
