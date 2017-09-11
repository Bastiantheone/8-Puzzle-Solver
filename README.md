# 8-Puzzle-Solver
A program that solves the 8-puzzle problem using the A* algorithm written in [Go](https://golang.org/).
The program solves the puzzle in far under a second.

## Installation
```
$ go get github.com/Bastiantheone/8-Puzzle-Solver/main
```

## How to run it
In the `%GOPATH%/bin` folder:
```
$ main {path to start file} {optional: path to goal file}
```
or
```
$ go run {path to}main.go {path to start file} {optional: path to goal file}
```
Default goal is:   <br>
&nbsp;0 1 2<br>
&nbsp;6 7 8<br>
&nbsp;3 4 5
