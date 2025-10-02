BINARY = humangl
CMD_DIR = cmd/humangl

all: build

build:
	@xhost +local:docker > /dev/null 2>&1
	@docker-compose build
	@xhost -local:docker > /dev/null 2>&1

run: build
	@xhost +local:docker > /dev/null 2>&1
	@docker-compose up
	@xhost -local:docker > /dev/null 2>&1

fclean:
	@docker-compose down --rmi all > /dev/null 2>&1
	@docker system prune -f > /dev/null 2>&1
	@rm -f $(BINARY)

re: fclean build

.PHONY: all build run fclean re
