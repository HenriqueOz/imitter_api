BIN_DIR = bin
API_EXE = $(BIN_DIR)/api-release

all: compile run

compile:
	@mkdir -p bin
	go build -o  ./$(API_EXE) ./src
	cp .env ./$(BIN_DIR)/.env

run:
	./$(API_EXE)