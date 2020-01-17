## Github style - project structure

In Go, the idiomatic way is to organize code in "GitHub style", so part of the path is looking like a server address to the library. Of course, if you want you don't need to do this, but the whole ecosystem works that way.

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

Environment variable `$GOPATH` is responsible for the path to the root dir of `src`, `bin` and `pkg` directories.


## Packages and Importing

`package` in go is in the form of files with directive `package package_name` on top of each file. Package by default is imported as the full path to the package.

    import "github.com/naqvijafar91/go-workshops/010-basics-importing/sub"

## Getting and installing external packages

To get the external package we need to run go install which will get sources and binaries and put them to `src`/`bin`/`pkg` directories

```sh
go install external.package.com/uri/to/package
```

## `sub` example package

As we can see our `sub` package is in directory `sub` (obvious) and have two files; `sub1.go` and `sub2.go` each of them also have `package sub` directive which tells the compiler that they are in one package.
