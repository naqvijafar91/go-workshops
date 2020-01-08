# Golang Training Resources

This repository contains files needed for managing Go language workshop - it's some kind of quite complete walk-through of Go language. Feel free to look on the code, there are many comments which could be useful for beginners and semi-intermediate Go developers.

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

- Visual Studio Code (really great UX)

- LiteIDE

- SublimeText

- Atom

### Auto reload

- BRA (Brilliant Ridiculous Assistant) https://github.com/Unknwon/bra - it's good to setup it when you're working on some webservers to auto reload your app when changes in code are made.

## Github style - project structure

In go, idiomatic way is to organise code in "github style", so part of the path is looking like server address to library. Of course if you want you don't need to do this, but whole ecosystem works that way.

```sh

src/

github.com

naqvijafar91

mysuperproject

ioki.com.pl

mnmel

nmelinium

bin/

superappbinary

pkg/

compiled packages

```

Environment variable `$GOPATH` is responsible for path to the root dir of `src`, `bin` and `pkg` directories.

## Packages and Importing

`package` in go is in form of files with directive `package package_name` on top of each file. Package by default is imported as full path to package.

import "github.com/naqvijafar91/go-workshops/010-basics-importing/sub"

## Getting and installing external packages

To get external package we need to run go install which will get sources and binaries and put them to `src`/`bin`/`pkg` directories

```sh

go install external.package.com/uri/to/package

```

## `sub` example package

As we can see our `sub` package is in directory `sub` (obvious) and have two files; `sub1.go` and `sub2.go` each of them also have `package sub` directive which tells compiler that they are in one package.

([code for: Basics Importing](https://github.com/naqvijafar91/go-workshops/tree/master/010-basics-importing))

## Package managers

Currently most advanced in go ecosystem is `dep` https://github.com/golang/dep

You can init your project:

\$ dep init

\$ ls

Gopkg.toml Gopkg.lock vendor/

after that dep will add vendor dir where all dependencies will be loaded (In go after 1.5 `vendor` dir have preference over "github style `$GOPATH` based directories - when compiler will not find library in vendor dir it'll try to take it from `$GOPATH`).

For more details please refer to `dep` documentation at https://golang.github.io/dep/docs/daily-dep.html

([code for: Package Management](https://github.com/naqvijafar91/go-workshops/tree/master/011-package-management))

### Tools

As in middle 2019 editors have still issues with new modules, you need

to install proper `gocode` or can use `gopls` (language server) for

autocompletition

([code for: Modules](https://github.com/naqvijafar91/go-workshops/tree/master/012-modules))

## Data structures: Arrays

In Go, an `array` is a **numbered sequence of elements** of a specific

length. Arrays are "low level" data structures with slices over them

which simplifies creating and managing.

([code for: Basics Arrays](https://github.com/naqvijafar91/go-workshops/tree/master/050-basics-arrays))

## Slices

Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).

sources:

- https://blog.golang.org/slices

- https://github.com/golang/go/wiki/SliceTricks

- https://blog.golang.org/go-slices-usage-and-internals

- http://research.swtch.com/godata

- http://blog.golang.org/go-slices-usage-and-internals

- http://www.dotnetperls.com/slice-go

([code for: Basics Slices](https://github.com/naqvijafar91/go-workshops/tree/master/051-basics-slices))

## Loops

In go there is only one loop keyword: `for`. It's often used with `range` keyword to iterate over array like elements.

([code for: Basics Loops](https://github.com/naqvijafar91/go-workshops/tree/master/052-basics-loops))

## Data structures: Maps

One of the most useful data structures in computer science is the hash table. Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes. Go provides a built-in map type that implements a hash table.

([code for: Basics Maps](https://github.com/naqvijafar91/go-workshops/tree/master/055-basics-maps))

## Basics Overriding Internal Types

There is no README.md for Basics Overriding Internal Types use ([code for: Basics Overriding Internal Types](https://github.com/naqvijafar91/go-workshops/tree/master/057-basics-overriding-internal-types)) link to follow code examples

## `make` and `new` keywords

`new(T)` allocates zeroed storage for a new item of type T and returns its address. In Go terminology, it returns a pointer to a newly allocated zero value of type T.

`make(T)` For slices, maps, and channels, make initializes the internal data structure and prepares the value for use.

### `make`

Things you can do with make that you can't do any other way:

- Create a channel

- Create a map with space preallocated

- Create a slice with space preallocated or with len != cap

### `new`

It's a little harder to justify new. The main thing it makes easier is creating pointers to non-composite types. The two functions below are equivalent. One's just a little more concise:

func newInt1() \*int { return new(int) }

func newInt2() \*int {

var i int

return &i

}

Go has multiple ways of memory allocation and value initialization:

&T{...}, &someLocalVar, new, make

Allocation can also happen when creating composite literals.

new can be used to allocate values such as integers, &int is illegal:

new(Point)

&Point{} // OK

&Point{2, 3} // Combines allocation and initialization

new(int)

&int // Illegal

// Works, but it is less convenient to write than new(int)

var i int

&i

In terms of channels there you can use make and new

p := new(chan int) // p has type: \*chan int

c := make(chan int) // c has type: chan int

([code for: Basics New And Make](https://github.com/naqvijafar91/go-workshops/tree/master/059-basics-new-and-make))

## Structs

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.

([code for: Basics Structs Defining](https://github.com/naqvijafar91/go-workshops/tree/master/060-basics-structs-defining))

([code for: Basics Struct Composition](https://github.com/naqvijafar91/go-workshops/tree/master/062-basics-struct-composition))

## Struct tags (annotations like)

A tag for a field allows you to attach meta-information to the field which can be acquired using reflection. Usually it is used to provide transformation info on how a struct field is encoded to or decoded from another format (or stored/retrieved from a database), but you can use it to store whatever meta-info you want to, either intended for another package or for your own use.

([code for: Basics Struct Tags](https://github.com/naqvijafar91/go-workshops/tree/master/064-basics-struct-tags))

## Basics Anonymous Structs

There is no README.md for Basics Anonymous Structs use ([code for: Basics Anonymous Structs](https://github.com/naqvijafar91/go-workshops/tree/master/065-basics-anonymous-structs)) link to follow code examples

## Interface

Go have "implicit interfaces". To implement an interface in Go, we

just need to implement all the methods in the interface.

So what is an interface? An interface is two things: it is a set of

methods, but it is also a type.

### The `interface{}` type

The `interface{}` type, the empty interface, is the source of much

confusion. The interface{} type is the interface that has no

methods. Since there is no implements keyword, all types implement at

least zero methods, and satisfying an interface is done automatically,

all types satisfy the empty interface. That means that if you write a

function that takes an `interface{}` value as a parameter, you can

supply that function with any value. So this function:

func DoSomething(v interface{}) {

// ...

}

will accept any parameter whatsoever.

### Type assertions

- https://medium.com/golangspec/type-assertions-in-go-e609759c42e1

([code for: Basics Interfaces](https://github.com/naqvijafar91/go-workshops/tree/master/065-basics-interfaces))

## Error handling

There is no exceptions in Go, errors are returned by value, or aggregated in intermediate objects. In go error is simply value which should be handled programatically as quick as possible.

Sources:

- https://blog.golang.org/errors-are-values

- http://blog.golang.org/error-handling-and-go

- http://davidnix.io/post/error-handling-in-go/

([code for: Basics Errors](https://github.com/naqvijafar91/go-workshops/tree/master/067-basics-errors))

## Panics

- Used when we want to stop the program.

- We can check if there was a `panic` occurence in function defer chain

([code for: Basics Panics](https://github.com/naqvijafar91/go-workshops/tree/master/068-basics-panics))

## Strings

In Go, a string is in effect a **read-only slice of bytes**.

It's important to state right up front that a string holds arbitrary bytes. It is not required to hold Unicode text, UTF-8 text, or any other predefined format. As far as the content of a string is concerned, it is exactly equivalent to a slice of bytes.

More on https://blog.golang.org/strings

([code for: Basics Strings](https://github.com/naqvijafar91/go-workshops/tree/master/070-basics-strings))

## Go routines

A goroutine is a lightweight thread of execution.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives (channels).

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

([code for: Basics Goroutines](https://github.com/naqvijafar91/go-workshops/tree/master/080-basics-goroutines))

## Using 3rd parties

In go we can get packages to our \$GOPATH with use of `go get` command.

## DEP

```sh

go dep init

```

`dep` is fully-fledged dependency manager. It downloads all dependencies source code to `vendor` directory and then compilator includes those code during compilation process.

([code for: Basics 3rd Party Packages](https://github.com/naqvijafar91/go-workshops/tree/master/090-basics-3rd-party-packages))

## Basics Pointers

There is no README.md for Basics Pointers use ([code for: Basics Pointers](https://github.com/naqvijafar91/go-workshops/tree/master/090-basics-pointers)) link to follow code examples

## Other concurrency patterns

You've lerned about channels primitive on previous chapter of this workshops, now it's time to learn about some other ways of creating concurrent programs in go.

We'll be doing some code with:

- data races errors

- atomic counters

- mutexes

- wait groups

([code for: Concurrency Other](https://github.com/naqvijafar91/go-workshops/tree/master/100-concurrency-other))

## Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

ch <- v // Send v to channel ch.

v := <-ch // Receive from ch, and

// assign value to v.

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

eventsChannel := make(chan int)

By default, **sends and receives block until the other side is ready**. This allows goroutines to synchronize without explicit locks or condition variables.

### Buffered Channels

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)

Sends to a buffered channel **block only when the buffer is full**. Receives block when the buffer is empty.

### In this workshops you'll learn about:

- channels mechanics

- receiving data from channel

- channels length

- buffers

- channel closing

- generate data with channels

- timers, tickers with channels

- worker pool

- rate limiting with channels

- selecting proper channel stream

- pipeline and fan-in/fan-out patterns

Feel free to browse [channels chapter source code](100-concurrency-channels/)

([code for: Concurrency Channels](https://github.com/naqvijafar91/go-workshops/tree/master/101-concurrency-channels))

## Date time

Go has very powerful standard library, first and one of awesome library are is Date time.

([code for: Stdlib Date Time](https://github.com/naqvijafar91/go-workshops/tree/master/110-stdlib-date-time))

## Running commands in shell

([code for: Stdlib Os Processes](https://github.com/naqvijafar91/go-workshops/tree/master/110-stdlib-os-processes))

## Args and flags

One of important thing when runnein your program is to get some arguments from user.

in Go you can easily get those data using `os.Args` slice or more powerful package `flags`.

Additionally you can get environment variables from `os.Getenv` function in `os` package

([code for: Stdlib Args](https://github.com/naqvijafar91/go-workshops/tree/master/114-stdlib-args))

# Streams

Streams in go are very powerful feature, very large part of standard library is written as some kind of stream reader or stream writer.

Go have two basic interfaces shaping all world of data streams `io.Reader` and `io.Writer`.

In this section of workhops we'll try to cover some of use cases for each of them.

([code for: Stdlib Streams](https://github.com/naqvijafar91/go-workshops/tree/master/115-stdlib-streams))

## Basic IO operations

- `bufio` examples

- directory traversal

- merging files with use of `buffers`

([code for: Stdlib Io](https://github.com/naqvijafar91/go-workshops/tree/master/116-stdlib-io))

## Stdlib Logging package

Go has basic logging package to log what's happening in your program.

([code for: Stdlib Logging](https://github.com/naqvijafar91/go-workshops/tree/master/120-stdlib-logging))

## HTTP library

Package http provides HTTP client and server implementations.

HTTP is one of fundamental lib in programming. Go has implemented very powerful HTTP primitives which allows creation of full-fledged HTTP powered servers.

Info: there is no routing in stdlib so you need to implement your own or use third party libraries (Gorilla mux, Gin, Echo are ones who can help you)

([code for: Stdlib Http](https://github.com/naqvijafar91/go-workshops/tree/master/140-stdlib-http))

## HTTP Middlewares

Go is using the term `middleware`, but each language/framework calls the concept differently. `NodeJS` and `Rails` calls it `middleware`. In the `Java EE` (i.e. Java Servlet), it’s called `filters`. `C#` calls it `delegate handlers`.

Essentially, the middleware **performs some specific function** on the HTTP request or response **at a specific stage in the HTTP pipeline** before or after the user defined controller. Middleware is a design pattern to eloquently add cross cutting concerns like logging, handling authentication, or gzip compression without having many code contact points.

([code for: Stdlib Http Middlewares](https://github.com/naqvijafar91/go-workshops/tree/master/141-stdlib-http-middlewares))

## Stdlib Encoding Json

There is no README.md for Stdlib Encoding Json use ([code for: Stdlib Encoding Json](https://github.com/naqvijafar91/go-workshops/tree/master/150-stdlib-encoding-json)) link to follow code examples

## Stdlib Encoding Xml

There is no README.md for Stdlib Encoding Xml use ([code for: Stdlib Encoding Xml](https://github.com/naqvijafar91/go-workshops/tree/master/151-stdlib-encoding-xml)) link to follow code examples

## Templates

In programming we need often some meta templates who help us with interoperability between our code and many output formats. One example could be template engine for generating HTML files for web sites.

In Go there are template engines (yes plural!) implemented in `stdlib`!

We'll go in this chapter by some `html` and `text` template engines.

([code for: Stdlib Templates](https://github.com/naqvijafar91/go-workshops/tree/master/170-stdlib-templates))

## Stdlib Math

There is no README.md for Stdlib Math use ([code for: Stdlib Math](https://github.com/naqvijafar91/go-workshops/tree/master/180-stdlib-math)) link to follow code examples

## Stdlib Regexp

There is no README.md for Stdlib Regexp use ([code for: Stdlib Regexp](https://github.com/naqvijafar91/go-workshops/tree/master/180-stdlib-regexp)) link to follow code examples

## Context package

Context is very powerful package, in this section i've implemented HTTP client and server which handles cancelling both sides when client e.g. press Ctrl+C during request.

([code for: Stdlib Context](https://github.com/naqvijafar91/go-workshops/tree/master/181-stdlib-context))

## Stdlib Sort

There is no README.md for Stdlib Sort use ([code for: Stdlib Sort](https://github.com/naqvijafar91/go-workshops/tree/master/181-stdlib-sort)) link to follow code examples

## Stdlib Signal

There is no README.md for Stdlib Signal use ([code for: Stdlib Signal](https://github.com/naqvijafar91/go-workshops/tree/master/182-stdlib-signal)) link to follow code examples

## MySQL

Install instructions to run this section:

```sh

cd docker/mysql

docker-compose up

docker exec -it mysql_app_1 mysql -uroot -proot -P7701 -hlocalhost

```

([code for: Databases Mysql](https://github.com/naqvijafar91/go-workshops/tree/master/210-databases-mysql))
