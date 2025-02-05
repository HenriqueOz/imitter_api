compile:
	go build -o  ./bin/api-release ./src
	cp .env ./bin/.env
