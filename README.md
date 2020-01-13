# Golang Workshop

This repository contains files needed for managing Go language workshop - it's some kind of quite complete walk-through of Go language. Feel free to look on the code, there are many comments which could be useful for beginners and semi-intermediate Go developers.

## What will be covered?
1. Introduction
2. [Variables and Constants](docs/01-variables-types-constants.md)
3. [Functions](docs/02-functions-and-packages.md)
4. [Packages](docs/02-functions-and-packages.md)
5. [Conditional Statements](docs/03-conditional-and-loops.md)
6. [Loops](docs/03-conditional-and-loops.md)
7. [Arrays](docs/04-arrays-slice-variadic-fn.md)
8. [Slice](docs/04-arrays-slice-variadic-fn.md)
9. [Variadic Functions](docs/04-arrays-slice-variadic-fn.md)
10. [Maps](docs/07-maps-and-strings.md)
11. [Strings](docs/07-maps-and-strings.md)
12. [Pointers](docs/08-pointers.md)
13. [Structs](docs/09-structs-and-interface.md)
14. [Interface](docs/09-structs-and-interface.md)
15. [Goroutines, channels](docs/10-concurrency.md)
16. [sync.WaitGroup and sync.Mutex](docs/10-concurrency.md)
17. [Unit Testing](docs/12-testing.md)
18. [Object Oriented Programming in Go](docs/11-object-oriented-programming.md)


## About Go

Who designed Go language?

- Rob Pike (Unix, UTF-8)

- Ken Thompson (Unix author, UTF-8, B lang)

- Robert Griesemer (V8, Java Hotspot (Oracle Java), GFS)

but those above are only ignitors of whole contributions: https://golang.org/AUTHORS

Why `go` was developed by Google? They have a lot of problems in C/Python/Java codebases:

- speed up development

- speed up compiling

- multicore systems

sources:

- https://golang.org/doc/faq#What_is_the_purpose_of_the_project

- https://talks.golang.org/2012/splash.article

## Go language characteristics

- statically compiled (one fat binary with all dependencies)

- Garbage Collected

- Strong types

- Functions as first class citizens

- Object Oriented (but without inheritance and classes)

## Why it's worth of your time?

- easy deployment (no dependencies, single binary statically linked)

- no more code style wars - `gofmt`

- integrated package downloader `go get`

- integrated code validation `go vet` and `golint` (github.com/golang/lint)

- nice playground (https://play.golang.org/)

- `gocode` intellisense server - you don't need fat IDE to write go code, you can use now editor which you love (Sublime, Atom, Vim, Emacs, VSCode)

- very fast compilation - if you're going from JAVA world you'll be really surprised

- quite complete standard library - template/html, performant www servers, json, xml, streams, io, buffers, first class citizen concurrency

- easy to use cross-compilation (x64, ARM, 386, Mac, Windows)

- easy start, bunch of editors, all things simply work

- http2 in core

- testing included

- benchmarking of code included

- very low entry barrier

- hype, one of fastest growing language, many new projects are in Go recently

- concurrency

- great documentation generator

- and many many more ...

## Workshop prerequisites

You can install `golang` and `docker` using your preferred way i.e. your OS package manager (brew, pacman, apt, snap or other) or you can simply follow installation instruction on go and docker sites.

### Golang installation

For recent installation instructions please refer to [Golang installation guide](https://golang.org/doc/install)

You'll need `git` to be installed

### Docker install

> Docker is the company driving the container movement and the only container platform provider to address every application across the hybrid cloud. Today’s businesses are under pressure to digitally transform but are constrained by existing applications and infrastructure while rationalizing an increasingly diverse portfolio of clouds, datacenters and application architectures. Docker enables true independence between applications and infrastructure and developers and IT ops to unlock their potential and creates a model for better collaboration and innovation.

For recent docker installation please follow [Docker installation guide](https://docs.docker.com/install/) for your OS.

Docker is needed for some parts of workshops for running databases or other needed dependencies. Also will be needed to complete homework.

### Docker Compose Installation

> Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration.

To install docker compose please follow [Docker compose installation guide](https://docs.docker.com/compose/install/)

## Init workshops to play on your machine with code

go get github.com/naqvijafar91/go-workshops

cd \$GOPATH/src/github.com/naqvijafar91/go-workshops

## Go Tools

### Included

- go test - included testing framework

- go fmt - code formater - only one valid coding standard - [fmt library website](https://golang.org/pkg/fmt/)

- go vet - code validator and fixer - [vet library website](https://golang.org/cmd/vet/)

- go oracle - dependencies analyser (can be integrated in editor) [oracle web site](http://golang.org/s/oracle-user-manual)

- godoc - great documentation generator and viewer

- gocode - autocomplete service - [gocode library website](https://github.com/mdempsky/gocode)

### IDEs

- IntelliJ Go plugin https://plugins.jetbrains.com/plugin/9568-go

- GoLand - https://www.jetbrains.com/go/

- Emacs - go-mode

- Vim - vim-go

- Visual Studio Code (for the workshop)

- LiteIDE

- SublimeText

- Atom


## Github style - project structure

In go, idiomatic way is to organise code in "github style", so part of the path is looking like server address to library. Of course if you want you don't need to do this, but whole ecosystem works that way.

```
bin/
    hello                          # command executable
    outyet                         # command executable
src/
    [github.com/golang/example/](https://github.com/golang/example/)
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    [golang.org/x/image/](https://golang.org/x/image/)
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...
```

Environment variable `$GOPATH` is responsible for path to the root dir of `src`, `bin` and `pkg` directories.

## Packages and Importing

`package` in go is in form of files with directive `package package_name` on top of each file. Package by default is imported as full path to package.
```go
import "github.com/naqvijafar91/go-workshops/010-basics-importing/sub"
```
## Getting and installing external packages

To get external package we need to run go install which will get sources and binaries and put them to `src`/`bin`/`pkg` directories

```sh

go install external.package.com/uri/to/package

```

## Package managers

Currently most advanced in go ecosystem is `dep` https://github.com/golang/dep

You can init your project:

\$ dep init

\$ ls

Gopkg.toml Gopkg.lock vendor/

after that dep will add vendor dir where all dependencies will be loaded (In go after 1.5 `vendor` dir have preference over "github style `$GOPATH` based directories - when compiler will not find library in vendor dir it'll try to take it from `$GOPATH`).

For more details please refer to `dep` documentation at https://golang.github.io/dep/docs/daily-dep.html

([code for: Package Management](https://github.com/naqvijafar91/go-workshops/tree/master/011-package-management))
