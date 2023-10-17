package workfile
import "fmt"

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
