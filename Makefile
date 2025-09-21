BINARY = humangl

all: build

build:
    go build -o $(BINARY)

run: build
    ./$(BINARY)

clean:
    rm -f $(BINARY)
