FROM golang:1.23.4 AS api

WORKDIR /usr/api

COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

COPY . .

EXPOSE 8080
CMD [ "go", "run", "src/main.go" ]