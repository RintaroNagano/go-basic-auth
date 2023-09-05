FROM golang:1.18.3 as builder

RUN mkdir /app
WORKDIR /app
ADD ./ /app/

RUN go env -w GO111MODULE=on

RUN go mod tidy
RUN go build -o server

FROM debian:11-slim

WORKDIR /app
COPY --from=builder /app/server /app/server
COPY --from=builder /app/.env /app/.env

ENV DOCKERIZE_VERSION v0.6.1
RUN apt-get update && apt-get install -y wget \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
RUN chmod +x server