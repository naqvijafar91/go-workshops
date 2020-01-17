# Running Go Programs

## Single File
To run a go file, simply use **go run** command followed by **filename**

```go run example.go```

## Package

To run a package containing multiple go files 

```go run *.go```

## Package containing tests

Test files in go have name with '*_test.go' pattern. When you use ```go run *.go``` command in a package containing tests, it gives an error ```go run: cannot run *_test.go files```. 
To avoid running test cases, you can use 

```go run $(ls -1 *.go | grep -v _test.go)```