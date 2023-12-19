FROM golang:1.21.5

LABEL maintainer="Gadiel Malagrino <gadiel.malagrino@gmail.com>"

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy

EXPOSE 3000