# cellular automata generator (in Go)

> **More documentation about Automata and their implementation on [eliseduverdier.github.io/cellular-automata-go/](https://eliseduverdier.github.io/cellular-automata-go/).**

![GitHub top language](https://img.shields.io/github/languages/top/eliseduverdier/cellular-automata-go?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/eliseduverdier/cellular-automata-go?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/eliseduverdier/cellular-automata-go?style=flat-square)

![Tests status](https://github.com/eliseduverdier/cellular-automata-go/actions/workflows/quality.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/eliseduverdier/cellular-automata-go)](https://goreportcard.com/report/github.com/eliseduverdier/cellular-automata-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/c24524e60f6020b406a3/maintainability)](https://codeclimate.com/github/eliseduverdier/cellular-automata-go/maintainability)

Generates [elementary cellular automata](https://en.wikipedia.org/wiki/Elementary_cellular_automaton) images, with 2 to 9 states, or 1st and 2nd order, with customizable size, colors, and points distribution, using Go

- As PNG images

  <img src="images/s2-o1-r73.png">

- Or as text

```
██ █ █    ██  ███ █ ███   ██ █ █   ██   █  ██
       █  █  ██         █ █      █ █  █ █ ██
      █ ██ ██  █       █   █    █   ██      █
     █       ██ █     █ █ █ █  █ █ █  █    █ █
    █ █     █    █   █       ██     ██ █  █   █
   █   █   █ █  █ █ █ █     █  █   █    ██ █ █
  █ █ █ █ █   ██       █   █ ██ █ █ █  █
 █         █ █  █     █ █ █          ██ █
█ █       █   ██ █   █     █        █    █    █
   █     █ █ █    █ █ █   █ █      █ █  █ █  █
█ █ █   █     █  █     █ █   █    █   ██   ██
     █ █ █   █ ██ █   █   █ █ █  █ █ █  █ █  █
    █     █ █      █ █ █ █     ██     ██   ██
   █ █   █   █    █       █   █  █   █  █ █  █
█ █   █ █ █ █ █  █ █     █ █ █ ██ █ █ ██   ██ █
```

## Use locally

```shell
go run main.go
```

Then go to http://localhost:8888/text or http://localhost:8888/image and use GET parameters (listed below)

Or use in the shell with

```shell
go run main.go --shell
```

This will generate an image in `images/` (or a text directly with the `-render=text` flag)

### command line options & get parameters

Example: `go run main.go -s=5 -o=2 -r=12345 -w=30 -h=10 -render=text` will render the automata #12345 of 2nd order with 5 sates, in text directly to the console, showing 30 columns of characters, and 10 rows.
| option | role |
|---|---|
| s | number of states ()>2) |
| o | order (1 or 2) |
| r | rule number |
| w | width in pixels |
| h | height in pixels |
| random_start | first line |
| render | 'image' or 'text' |

## Run test

```shell
go test tests/*_test.go
```

# Etc

## Todo

- [ ] add docker
- [ ] code coverage
- [ ] add option to get dots bigger than 1px
- [ ] add stragety pattern for the first line (random|single point|word|custom)

### install latest go version on ubuntu

```
sudo apt-get purge golang\*
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```
