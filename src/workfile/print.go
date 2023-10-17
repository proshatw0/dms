package workfile

import (
	"strconv"

	"dms/src/structs"
)

func Print_Table_Array(filepath string, line_number int, name_table string, array structs.Array) error {
	out := name_table
	out += ": [" + strconv.Itoa(array.Lenght) + ", {"
	if array.Data[0] == "" {
		out += " "
	}
	for index := 0; index < array.Lenght; index++ {
		out += array.Data[index] + ", "
	}
	if len(out) == len(name_table)+len(strconv.Itoa(array.Lenght))+6 {
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
