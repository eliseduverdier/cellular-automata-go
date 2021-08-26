# cellular automata generator (in Go)

Generates [elementary cellular automata](https://en.wikipedia.org/wiki/Elementary_cellular_automaton) images, with 2 to 9 states, or 1st and 2nd order, with customizable size, colors, and points distribution, using Go

<img src="images/s2-o1-r73.png">

## Use locally

- `go build main.go` will compile the program (and make it easier to access shelll arguments)
- `./main.go -s=2 -o=1 -r=123 -w=10 -h=20` create an in `images/s2-o1-r123.png`,

## Run tests

```
go test tests/*
```

## Etc / Help

### Todo

- [] add web server
- [] add docker

### install latest go version on ubuntu

```
sudo apt-get purge golang\*
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```
