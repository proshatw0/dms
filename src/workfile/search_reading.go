package workfile

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	case "dlpush_end", "dlpush_begin", "dldel_end", "dldel_begin", "dldel", "dlcout", "dllen":
		line, err := Read_Line_Fromfile(filepath, 6)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "tree", 0
		}
		return "dl_list", number_line
	case "tins", "tdel", "tcon", "tmax", "tmin":
		line, err := Read_Line_Fromfile(filepath, 7)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "tree", 0
		}
		return "tree", number_line
	case "aset", "aget", "aindex", "adel", "adel_value", "apush", "apop":
		line, err := Read_Line_Fromfile(filepath, 1)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "array", 0
		}
		return "array", number_line
	case "sadd", "srem", "sismember":
		line, err := Read_Line_Fromfile(filepath, 2)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "set", 0
		}
		return "set", number_line
	case "spush", "spop":
		line, err := Read_Line_Fromfile(filepath, 3)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "stack", 0
		}
		return "stack", number_line
	case "qpush", "qpop":
		line, err := Read_Line_Fromfile(filepath, 4)
		if err != nil {
			log.Fatal(err)
		}
		number_line := Search_Number_Table(line, name_table)
		if number_line == 0 {
			return "queue", 0
		}
		return "queue", number_line
	case "hset", "hdel", "hget":
		line, err := Read_Line_Fromfile(filepath, 5)
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
