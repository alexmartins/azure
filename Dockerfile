FROM golang:alpine

RUN apk add --no-cache bash gcc musl-dev

WORKDIR /usr/src/myapp

COPY ./src/desafio/main.go /usr/src/myapp
COPY ./src/desafio/main_test.go /usr/src/myapp
COPY ./start/go.sh /usr/src/myapp

ENTRYPOINT [ "sh", "go.sh" ]

EXPOSE 8000