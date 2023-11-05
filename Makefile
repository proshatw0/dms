FILES = main.go
TARGET = bin/dms_server.exe

.PHONY: clean

TARGET: 
	go build -o $(TARGET) $(FILES)

clean:
	rm -r $(TARGET)
