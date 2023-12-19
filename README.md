# Stori Challenge

<img align="right" width="159px" src="https://www.storicard.com/static/images/thumbnail_storicard_small.png" alt="">

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)

This is a project developed by [Gadiel Malagrino](https://github.com/gadielMa). The system processes a file `txns.csv` which contain a list of debit and credit transactions on an account. The system process the file and send summary information to a user in the form of an email.

**Language:** [go1.21.x](https://tip.golang.org/doc/go1.21)

<img src="https://i.imgur.com/3elNhQu.png" alt="">

**Chosen Architecture:** Clean Architecture
<img src="https://miro.medium.com/v2/resize:fit:800/1*0R0r00uF1RyRFxkxo3HVDg.png">

**Framework:** [go-chi](https://github.com/go-chi/chi)

<img src="https://camo.githubusercontent.com/f72d07b7d898f8935d557867df17416a1b430a2572f8ea1bae57d1700f5c754b/68747470733a2f2f63646e2e7261776769742e636f6d2f676f2d6368692f6368692f6d61737465722f5f6578616d706c65732f6368692e737667" alt="">

**Relational Database:** [gorm.io/driver/postgres](https://github.com/go-gorm/postgres)

<img src="https://es-wiki.ikoula.com/images/a/a3/Postgre.png" alt="">

**ORM:** [gorm.io/gorm](https://gorm.io)

**Mail Service:** [go-mail](https://gopkg.in/mail.v2)

## Getting started

### Prerequisites

- **[Docker](https://www.docker.com/get-started/)**: any one of the latest versions

<img src="https://2.bp.blogspot.com/-uL7xdajcC8I/XFL3cNJ6PTI/AAAAAAAAKLM/DLpmLx1W_CUXZST_7CIHWC8uNqt2enVNwCLcBGAs/s1600/docker.png" alt="">

### Run

We use the db automatically created called equal as the user db. // TODO

```sh
$ docker run --name mypostgres -e POSTGRES_USER=malagrino -e POSTGRES_PASSWORD=malagrino -p 5432:5432 -d postgres:16-alpine
docker exec -it mypostgres bash
```
go install github.com/cosmtrek/air@latest
curl --location --request POST 'http://localhost:3000/stori/summary'