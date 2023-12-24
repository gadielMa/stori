# Stori Challenge

<img align="right" width="159px" src="https://www.storicard.com/static/images/thumbnail_storicard_small.png" alt="">

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)

This is a project developed by [Gadiel Malagrino](https://github.com/gadielMa). The system processes a file `txns.csv` which contain a list of debit and credit transactions on an account. The system process the file and send summary information to a user in the form of an email.

### Description

No `unit tests` requested.

In this project I decided on a 3-layer scaffold.

The first layer `cmd`, contains the initial configuration of the project and a `write` folder where the POST, PUT, etc. are.

The GET or DELETE, etc., would be in the "read" folder. This is to separate the flows into different golang instances and help the processor.

The second layer is `service`, where the business logic is, there are also the `interfaces` to the repositories. This way we always put the interfaces on a layer before using them.

The last layer is `repositories`, where the uses of external services and interactions with the databases are.

I didn't need to create my own Docker images, since my `Dockerfile` was very simple.

Although it is bad practice, I decided to write the `.env` file for reference. And only by modifying the `MAIL_TO` will you receive the email.

I also wrote down my `EMAIL_FROM PASSWORD` from my old university email, in case you don't know how to get one...

The email service chosen was `GMAIL`, so please choose a Gmail account.

The name of the database used is "postgres" since it is a database that is already created within the `postgres:latest` image.

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
- **[Git](https://github.com/gadielMa/stori)**
- **[Docker](https://www.docker.com/get-started/)**: You must have docker installed on your system.

<img src="https://2.bp.blogspot.com/-uL7xdajcC8I/XFL3cNJ6PTI/AAAAAAAAKLM/DLpmLx1W_CUXZST_7CIHWC8uNqt2enVNwCLcBGAs/s1600/docker.png" alt="">

### Steps

```sh
$ git clone https://github.com/gadielMa/stori.git
```

```sh
$ cd stori
```

```sh
$  docker compose up
```

Wait a minute and...

Use a REST Client like [Postman](https://www.postman.com) to run the following curl:

```sh
curl --location --request POST 'http://localhost:3000/stori/summary'
```

If you receive a code 200, you will receive an email in your `MAIL_TO` box.