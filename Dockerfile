FROM golang:1.21.5

LABEL maintainer="Gadiel Malagrino <gadiel.malagrino@gmail.com>"

WORKDIR /usr/src/app

COPY . .

EXPOSE 3000