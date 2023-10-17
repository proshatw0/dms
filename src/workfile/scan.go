package workfile

import (
	"strconv"
	"strings"

	"dms/src/structs"
)

func Scan_Table_Array(filepath string, line_number int) structs.Array {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return structs.Array{}
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "}")
	sizeIndex := strings.Index(line, ",")
	size := line[startIndex:sizeIndex]
	val := line[sizeIndex+len(size)+1 : endIndex]
	arr := strings.Split(val, ",")
	size_int, _ := strconv.Atoi(size)
	Array := structs.NewArray(size_int)
	for i := 0; i < len(arr); i++ {
		Array.Aset(i, strings.ReplaceAll(arr[i], " ", ""))
	}
	return *Array
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
	return *set
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
	return *hash_table
}
