BINARY = humangl
CMD_DIR = cmd/humangl
	
all: build

build:
	go build -o $(BINARY) ./$(CMD_DIR)

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)

.PHONY: all build run clean
