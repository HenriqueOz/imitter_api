FROM golang:1.23.4 AS api-build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src src
COPY .env .
RUN mkdir ./bin

RUN go build -o ./bin/api-release ./src

FROM debian:latest AS api-final-stage

WORKDIR /app

COPY --from=api-build-stage /app/bin/api-release /app/api-release
COPY --from=api-build-stage /app/.env /app/.env

RUN addgroup appgroup && adduser appuser
RUN usermod -g appgroup appuser
RUN chown -R appuser:appgroup  /app

USER appuser

EXPOSE 8080
CMD [ "/app/api-release" ]
