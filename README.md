# golaunch — go project init

## Usage
`golaunch <name>` to create a project
- Creates a new directory
- Fails if a file or directory already exists with the name specified

## Project Structure
```
example <- name used for root dir
├─ app
│  ├─ main.go
│  └─ main_test.go
├─ go.mod
├─ justfile <- build/run/test scripts
├─ LICENSE <- auto generated from username/current year
└─ README.md
```
- License is MIT by default, may add future options to specify a license
- CLI tool [just](https://github.com/casey/just) is highly recommended

## Installation
- I'll add to some package repos eventually
- Until then, there is binaries available in [releases](https://github.com/Sophed/golaunch/releases)

## Building
- Requires `just` and `go` installed
```
git clone https://github.com/Sophed/golaunch
cd golaunch
just build
```