package workfile

import (
	"fmt"
	"strconv"

	"dms/src/structs"
)

func Processing_Request(filepath string, commands [4]string) {
	mod, number_line := Search_Table(filepath, commands[0], commands[1])
	if mod == "" {
		fmt.Println("--> unknown command")
		return
	}
	if number_line != 0 {
		switch mod {
		case "dl_list":
			dl_list := Scan_Table_Dl_list(filepath, number_line)
			if dl_list.Head != nil {
				if dl_list.Head.Data == "error1456&789" {
					fmt.Println("-->table not found", dl_list.Head)
					return
				}
			}
			switch commands[0] {
			case "dlpush_end":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				dl_list.Dlpush_end(commands[2])
				err := Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "dlpush_begin":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				dl_list.Dlpush_begin(commands[2])
				err := Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "dldel_end":
				err, value := dl_list.Dldel_end()
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", value)
				return
			case "dldel_begin":
				err, value := dl_list.Dldel_begin()
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", value)
				return
			case "dldel":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				err := dl_list.Dldel(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "dlcout":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				err := dl_list.Dlcout(commands[2])
				if err != nil {
					fmt.Println("--> False")
					return
				}
				fmt.Println("--> True")
				return
			case "dllen":
				len := dl_list.Dllen()
				fmt.Println("-->", len)
				return
			}
		case "tree":
			tree := Scan_Table_Tree(filepath, number_line)
			if tree.Root == nil {
				fmt.Println("-->table not found")
				return
			}
			switch commands[0] {
			case "tins":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				tree.Tins(commands[2])
				err := Print_Table_Tree(filepath, number_line, commands[1], tree)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "tdel":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				err := tree.Tdel(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Tree(filepath, number_line, commands[1], tree)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "tcon":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				err := tree.Tcon(commands[2])
				if err != nil {
					fmt.Println("--> False")
					return
				}
				fmt.Println("--> True")
			case "tmax":
				fmt.Println("-->", structs.Tmax(tree.Root))
				return
			case "tmin":
				fmt.Println("-->", structs.Tmin(tree.Root))
				return
			}
		case "array":
			array := Scan_Table_Array(filepath, number_line)
			if array.Lenght <= 0 {
				fmt.Println("-->table not found")
				return
			}
			switch commands[0] {
			case "aset":
				index, err := strconv.Atoi(commands[2])
				if err != nil || commands[3] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				err = array.Aset(index, commands[3])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2], "~", commands[3])
				return
			case "aget":
				index, err := strconv.Atoi(commands[2])
				if err != nil {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				data, err := array.Aget(index)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", data)
				return
			case "aindex":
				index, err := array.Aindex(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", index)
				return
			case "adel":
				index, err := strconv.Atoi(commands[2])
				if err != nil {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				value, err := array.Adel(index)
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2], "~", value)
				return
			case "adel_value":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				value, err := array.Adel_value(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", value)
				return
			case "apush":
				if commands[2] == "" {
					fmt.Println("-->invalid request")
					fmt.Println("Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>")
					return
				}
				err := array.Apush(commands[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", commands[2])
				return
			case "apop":
				value, err := array.Apop()
				if err != nil {
					fmt.Println(err)
					return
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("-->", value)
				return
			}
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
