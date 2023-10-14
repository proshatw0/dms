SRC = src
FILES = main.go
TARGET = bin/dms


TARGET: 
	go build -o $(TARGET) $(FILES)
