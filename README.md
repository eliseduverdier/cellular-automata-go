# cellular automata generator (in Go)

> **More documentation about Automata and their implementation on [eliseduverdier.github.io/cellular-automata-go/](https://eliseduverdier.github.io/cellular-automata-go/).**

![GitHub top language](https://img.shields.io/github/languages/top/eliseduverdier/cellular-automata-go?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/eliseduverdier/cellular-automata-go?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/eliseduverdier/cellular-automata-go?style=flat-square)

![CI status](https://github.com/eliseduverdier/cellular-automata-go/actions/workflows/quality.yml/badge.svg)
[![codecov](https://codecov.io/gh/eliseduverdier/cellular-automata-go/branch/main/graph/badge.svg?token=OIW9T63RAI)](https://codecov.io/gh/eliseduverdier/cellular-automata-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/eliseduverdier/cellular-automata-go)](https://goreportcard.com/report/github.com/eliseduverdier/cellular-automata-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/c24524e60f6020b406a3/maintainability)](https://codeclimate.com/github/eliseduverdier/cellular-automata-go/maintainability)

Generates [elementary cellular automata](https://en.wikipedia.org/wiki/Elementary_cellular_automaton) images, with 2 to 9 states, or 1st and 2nd order, with customizable size, ~colors~, and points distribution, using Go

- As PNG images (rule 73)

  <img src="_doc/images/sample.png">

- Or as text (rule 110)

<pre style="line-height:8px;">
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
</pre>

## Use locally

```shell
PORT=8888 go run main.go
```

Then go to http://localhost:8888/text or http://localhost:8888/image and use GET parameters (listed below)

## Exemple of nice automatas

![Order 1  /  2 states  /  Rule # 150](_doc/images/o1-s2-r150.png)
![Order 1  /  2 states  /  Rule # 73](_doc/images/o1-s2-r73.png)
![Order 1  /  3 states  /  Rule # 78271914](_doc/images/o1-s3-r78271914.png)
![Order 1  /  3 states  /  Rule # 76148092](_doc/images/o1-s3-r76148092.png)

![Order 2  /  2 states  /  Rule # 24987](_doc/images/o2-s2-r24987.png)
![Order 2  /  2 states  /  Rule # 56922](_doc/images/o2-s2-r56922.png)
![Order 2  /  2 states  /  Rule # 5198](_doc/images/o2-s2-r5198.png)
![Order 2  /  2 states  /  Rule # 58937](_doc/images/o2-s2-r58937.png)
![Order 2  /  2 states  /  Rule # 37486](_doc/images/o2-s2-r37486.png)

### parameters

| option | role                                                                   |
| ------ | ---------------------------------------------------------------------- |
| s      | number of states (>2)                                                  |
| o      | order (1 or 2)                                                         |
| r      | rule number                                                            |
| w      | width in pixels                                                        |
| h      | height in pixels                                                       |
| p      | size of one cell in pixels (1 \< p \< 20)                              |
| start  | first line strategy ("random", "centered", "left", "right", see below) |
| line   | a sequence of numbers representing the first line like "0000010..."    |

## Run tests

```shell
# run tests
make test
# generate code coverage
make coverage
# lint (needs https://golangci-lint.run/)
make lint
```

# Etc

## Todo

- [ ] add docker
- [ ] add option to get dots bigger than 1px
- [ ] add option to choose custom colors
- [ ] Cache when rule isn't random

### install latest go version on ubuntu

```
sudo apt-get purge golang\*
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```
