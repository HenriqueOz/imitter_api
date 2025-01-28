FROM golang:1.23.4 AS api

WORKDIR /usr/api

COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

COPY . .

RUN go build -o ../bin/api-release ./src && cp .env /bin

EXPOSE 8080
CMD [ "/bin/api-release" ]

FROM mysql:latest AS mysql

ADD ./database/imitter_schema.sql ./docker-entrypoint-initdb.d