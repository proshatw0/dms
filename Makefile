SRC = src
FILES = main.go
TARGET = bin/dms

.PHONY: clean

TARGET: 
	go build -o $(TARGET) $(FILES)

clean:
	rm -r $(TARGET)
