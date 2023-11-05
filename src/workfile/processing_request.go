package workfile

import (
	"errors"
	"strconv"

	"dms/src/checkURL"
	"dms/src/structs"
)

func Processing_Request(filepath string, commands [4]string) (error, string, int) {
	err, mod, number_line := Search_Table(filepath, commands[0], commands[1])
	if err != nil || number_line == -1 {
		return err, "", -1
	}
	if mod == "" {
		return errors.New("--> unknown command"), "", -1
	}
	if number_line != -1 {
		switch mod {
		case "dl_list":
			dl_list := Scan_Table_Dl_list(filepath, number_line)
			if dl_list.Head != nil {
				if dl_list.Head.Data == "error1456&789" {
					return errors.New("-->table not found"), "", -1
				}
			}
			switch commands[0] {
			case "dlpush_end":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1
				}
				dl_list.Dlpush_end(commands[2])
				err := Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "dlpush_begin":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				dl_list.Dlpush_begin(commands[2])
				err := Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "dldel_end":
				err, value := dl_list.Dldel_end()
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					return err, "", -1
				}
				return nil, value, -1
			case "dldel_begin":
				err, value := dl_list.Dldel_begin()
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					return err, "", -1
				}
				return nil, value, -1
			case "dldel":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				err := dl_list.Dldel(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Dl_list(filepath, number_line, commands[1], dl_list)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "dlcout":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				err := dl_list.Dlcout(commands[2])
				if err != nil {
					return nil, "False", -1
				}
				return nil, "True", -1
			case "dllen":
				len := dl_list.Dllen()
				return nil, "", len
			}
		case "tree":
			tree := Scan_Table_Tree(filepath, number_line)
			if tree.Root == nil {
				return errors.New("-->table not found"), "", -1
			}
			switch commands[0] {
			case "tins":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				tree.Tins(commands[2])
				err := Print_Table_Tree(filepath, number_line, commands[1], tree)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "tdel":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				err := tree.Tdel(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Tree(filepath, number_line, commands[1], tree)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "tcon":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				err := tree.Tcon(commands[2])
				if err != nil {
					return nil, "False", -1
				}
				return nil, "True", -1
			case "tmax":
				return nil, structs.Tmax(tree.Root), -1
			case "tmin":
				return nil, structs.Tmin(tree.Root), -1
			}
		case "array":
			array := Scan_Table_Array(filepath, number_line)
			if array.Lenght <= 0 {
				return errors.New("-->table not found"), "", -1
			}
			switch commands[0] {
			case "aset":
				index, err := strconv.Atoi(commands[2])
				if err != nil || commands[3] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				err = array.Aset(index, commands[3])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[3], index
			case "aget":
				index, err := strconv.Atoi(commands[2])
				if err != nil {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				data, err := array.Aget(index)
				if err != nil {
					return err, "", -1
				}
				return nil, data, -1
			case "aindex":
				index, err := array.Aindex(commands[2])
				if err != nil {
					return err, "", -1
				}
				return nil, "", index
			case "adel":
				index, err := strconv.Atoi(commands[2])
				if err != nil {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				value, err := array.Adel(index)
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					return err, "", -1
				}
				return nil, value, index
			case "adel_value":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				value, err := array.Adel_value(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					return err, "", -1
				}
				return nil, value, -1
			case "apush":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				err := array.Apush(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "apop":
				value, err := array.Apop()
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Array(filepath, number_line, commands[1], array)
				if err != nil {
					return err, "", -1
				}
				return nil, value, -1
			}
		case "set":
			set := Scan_Table_Set(filepath, number_line)
			if set.Size <= 0 {
				return errors.New("-->table not found"), "", -1
			}
			switch commands[0] {
			case "sadd":
				if commands[2] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				err := set.Sadd(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Set(filepath, number_line, commands[1], set)
				if err != nil {
					return err, "", -1
				}
				out := "--> " + commands[2]
				return nil, out, -1
			case "srem":
				err := set.Srem(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Set(filepath, number_line, commands[1], set)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "sismember":
				err := set.Sismember(commands[2])
				if err != nil {
					return err, "", -1
				}
				return nil, "True", -1
			}
		case "stack":
			stack := Scan_Table_Stack(filepath, number_line)
			switch commands[0] {
			case "spush":
				err := stack.Spush(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Stack(filepath, number_line, commands[1], stack)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "spop":
				value, err := stack.Spop()
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Stack(filepath, number_line, commands[1], stack)
				if err != nil {
					return err, "", -1
				}
				return nil, value, -1
			}
		case "queue":
			queue := Scan_Table_Queue(filepath, number_line)
			switch commands[0] {
			case "qpush":
				err := queue.Qpush(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Queue(filepath, number_line, commands[1], queue)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2], -1
			case "qpop":
				value, err := queue.Qpop()
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Queue(filepath, number_line, commands[1], queue)
				if err != nil {
					return err, "", -1
				}
				return nil, value, -1
			}
		case "hash_table":
			if commands[2] == "" {
				return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

			}
			hash_table := Scan_Table_Hash_Table(filepath, number_line)
			if hash_table.Size <= 0 {
				return errors.New("-->table not found"), "", -1
			}
			switch commands[0] {
			case "hset":
				if commands[3] == "" {
					return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

				}
				err := hash_table.Hset(commands[2], commands[3])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Hash_Table(filepath, number_line, commands[1], hash_table)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2] + " ~ " + commands[3], -1
			case "hdel":
				value, err := hash_table.Hdel(commands[2])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Hash_Table(filepath, number_line, commands[1], hash_table)
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2] + " ~ " + value, -1
			case "hget":
				value, err := hash_table.Hget(commands[2])
				if err != nil {
					return err, "", -1
				}
				return nil, commands[2] + " ~ " + value, -1
			}
		case "link":
			if commands[1] == "" {
				return errors.New("-->Example request: ./<name of your program> --file <path to the data file> --query <operation table_name element>"), "", -1

			}
			links, len := Scan_Table_Links(filepath, number_line)
			if links.Size <= 0 && len == -1 {
				return errors.New("-->table not found"), "", -1
			}
			switch commands[0] {
			case "post":
				if !checkURL.CheckURL(commands[1]) {
					return errors.New("-->original link is not available"), "", -1
				}
				base := "localhost/"
				base += strconv.Itoa(len + 1)
				err := links.Hset(base, commands[1])
				if err != nil {
					return err, "", -1
				}
				err = Print_Table_Hash_Table(filepath, number_line, "links", links)
				if err != nil {
					return err, "", -1
				}
				return nil, base, -1
			case "get":
				value, err := links.Hget(commands[1])
				if err != nil {
					return err, "", -1
				}
				return nil, value, -1
			}
		default:
			return errors.New("-->table not found"), "", -1
		}
	} else {
		return errors.New("-->table not found"), "", -1
	}
	return errors.New("-->table not found"), "", -1
}

func CheckURL(s string) {
	panic("unimplemented")
}
