FILES = main.go
TARGET = bin/dms_server

.PHONY: clean

TARGET: 
	go build -o $(TARGET) $(FILES)

clean:
	rm -r $(TARGET)
