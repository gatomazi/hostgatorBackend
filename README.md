# Hostgator Backend

You are going to need docker-compose and Go 1.13

[Golang Download](https://golang.org/dl/)

##

[Docker Compose Download](https://docs.docker.com/compose/install/)

## Installation

With the repo cloned, enter on it and make the mysql container up with the command:

```bash
docker-compose up
```

Now, to get all golang dependencies, do a:

```bash
go get
```

To up the Go server, you just need to run the command

```bash
go run main.go
```

To run the tests

```bash
go test
```

<!--
But if you want to start the server with a database in a different ip run

```bash
go run main.go -dockerIp="xx.xx.xx.xx"
``` -->
