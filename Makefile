BINARY = humangl
CMD_DIR = cmd/humangl

# Check if Go is installed
GO_INSTALLED := $(shell command -v go 2> /dev/null)

# Source files
SOURCES = $(wildcard cmd/humangl/*.go) \
          $(wildcard internal/*/*.go) \
          go.mod go.sum

all: $(BINARY)

$(BINARY): $(SOURCES)
ifdef GO_INSTALLED
	@echo "Go found on host, building locally..."
	@go build -o $(BINARY) ./$(CMD_DIR)
else
	@echo "Go not found, using Docker for compilation..."
	@DOCKER_BUILDKIT=1 docker build -t humangl-builder .
	@docker create --name temp-container humangl-builder
	@docker cp temp-container:/app/$(BINARY) ./$(BINARY)
	@docker rm temp-container
	@echo "Binary extracted to host: $(BINARY)"
endif

run: $(BINARY)
	./$(BINARY)

fclean:
ifdef GO_INSTALLED
	@echo "Cleaning local build..."
	@rm -f $(BINARY)
else
	@echo "Cleaning Docker images..."
	@docker-compose down --rmi all > /dev/null 2>&1
	@docker system prune -f > /dev/null 2>&1
	@docker rmi humangl-builder 2>/dev/null || true
	@rm -f $(BINARY)
endif

re: fclean $(BINARY)

.PHONY: all run fclean re
