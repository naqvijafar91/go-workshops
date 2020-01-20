
# Writing unit tests

## 1. Testing in Go

Go has a built-in testing command called  `go test`  and a package  `testing`  which combine to give a minimal but complete testing experience.

The standard tool-chain also includes benchmarking and statement-based code coverage similar to NCover (.NET) or Istanbul (Node.js).

### 1.2 Writing tests

Here is an example of a method we want to test in the  `main`  package. We have defined an exported function called  `Sum`  which takes in two integers and adds them together.

```go
package main

func Sum(x int, y int) int {
    return x + y
}

func main() {
    Sum(5, 5)
}

```

We then write our test in a separate file. The text file can be in a different package (and folder) or the same one (`main`). Here's a unit test to check addition:

```go
package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

```

Characteristics of a Golang test function:

-   The first and only parameter needs to be  `t *testing.T`
-   It begins with the word Test followed by a word or phrase starting with a capital letter.
-   (usually the method under test i.e.  `TestValidateClient`)
-   Calls  `t.Error`  or  `t.Fail`  to indicate a failure (I called t.Errorf to provide more details)
-   `t.Log`  can be used to provide non-failing debug information
-   Must be saved in a file named  `something_test.go`  such as:  `addition_test.go`

> If you have code and tests in the same folder then you cannot execute your program with  `go run *.go`. I tend to use  `go build`  to create a binary and then I run that.


**Test tables**

The concept of "test tables" is a set (slice array) of test input and output values. Here is an example of the  `Sum`  function:

```golang
package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

```

If you want to trigger the errors to break the test then alter the  `Sum`  function to return  `x * y`.

```
$ go test -v
=== RUN   TestSum
--- FAIL: TestSum (0.00s)
	table_test.go:19: Sum of (1+1) was incorrect, got: 1, want: 2.
	table_test.go:19: Sum of (1+2) was incorrect, got: 2, want: 3.
	table_test.go:19: Sum of (5+2) was incorrect, got: 10, want: 7.
FAIL
exit status 1
FAIL	github.com/alexellis/t6	0.013s

```

**Launching tests:**

There are two ways to launch tests for a package. These methods work for unit tests and integration tests alike.

1.  Within the same directory as the test:

```
go test

```

_This picks up any files matching packagename_test.go_

or

2.  By fully-qualified package name

```
go test github.com/alexellis/golangbasics1

```

You have now run a unit test in Go, for a more verbose output type in  `go test -v`  and you will see the PASS/FAIL result of each test including any extra logging produced by  `t.Log`.

> The difference between unit and integration tests is that unit tests usually isolate dependencies that communicate with the network, disk, etc. Unit tests normally test only one thing such as a function.

## 1.3 More on  `go test`

**Statement coverage**

The  `go test`  tool has built-in code-coverage for statements. To try it without example above type in:

```
$ go test -cover
PASS
coverage: 50.0% of statements
ok  	github.com/alexellis/golangbasics1	0.009s

```

 If you delete the "if" a statement from our previous test it will retain 50% test coverage but lose its usefulness in verifying the behavior of the "Sum" method.

**Generating an HTML coverage report**

If you use the following two commands you can visualize which parts of your program have been covered by the tests and which statements are lacking:

```
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 

```

Then open coverage.html in a web-browser.

**Go doesn't ship your tests**

Also, it may feel unnatural to leave files named  `addition_test.go`  in the middle of your package. Rest assured that the Go compiler and linker will not ship your test files in any binaries it produces.

For more on the basics read the  [Golang testing docs](https://golang.org/pkg/testing/).

### 1.4 Isolating dependencies

The key factor that defines a unit test in isolation from runtime-dependencies or collaborators.

This is achieved in Golang through interfaces, but if you're coming from a C# or Java background, they look a little different in Go. Interfaces are implied rather than enforced which means that concrete classes don't need to know about the interface ahead of time.

That means we can have very small interfaces such as  [io.ReadCloser](https://golang.org/src/io/io.go?s=4977:5022#L116)  which has only two methods made up of the Reader and Closer interfaces:

```
        Read(p []byte) (n int, err error)

```

_Reader interface_

```
        Close() error

```

_Closer interface_

If you are designing a package to be consumed by a third-party then it makes sense to design interfaces so that others can write unit tests to isolate your package when needed.

An interface can be substituted in a function call. So if we wanted to test this method, we'd just have to supply a fake / test-double class that implemented the Reader interface.

```
package main

import (
	"fmt"
	"io"
)

type FakeReader struct {
}

func (FakeReader) Read(p []byte) (n int, err error) {
	// return an integer and error or nil
}

func ReadAllTheBytes(reader io.Reader) []byte {
	// read from the reader..
}

func main() {
	fakeReader := FakeReader{}
	// You could create a method called SetFakeBytes which initialises canned data.
	fakeReader.SetFakeBytes([]byte("when called, return this data"))
	bytes := ReadAllTheBytes(fakeReader)
	fmt.Printf("%d bytes read.\n", len(bytes))
}

```

Before implementing your abstractions (as above) it is a good idea to search the Golang docs to see if there is already something you can use. In the case above we could also use the standard library in the  [bytes](https://golang.org/pkg/bytes/)  package:

```
    func NewReader(b []byte) *Reader

```

The Golang  [testing/iotest](https://golang.org/pkg/testing/iotest/)  package provides some Reader implementations which are slow or which cause errors to be thrown halfway through reading. These are ideal for resilience testing.

-   Golang docs:  [testing/iotest](https://golang.org/pkg/testing/iotest/)

### 1.5 Worked example



We'll start with the test file:

```
package main

import "testing"

type testWebRequest struct {
}

func (testWebRequest) FetchBytes(url string) []byte {
	return []byte(`{"number": 2}`)
}

func TestGetAstronauts(t *testing.T) {
	amount := GetAstronauts(testWebRequest{})
	if amount != 1 {
		t.Errorf("People in space, got: %d, want: %d.", amount, 1)
	}
}

```

I have an exported method called GetAstronauts which calls into an HTTP endpoint, reads the bytes from the result and then parses this into a struct and returns the integer in the "number" property.

My fake / test-double in the test only returns the bare minimum of JSON needed to satisfy the test and to begin with I had it return a different number so that I knew the test worked. It's hard to be sure whether a test that passes the first time has worked.

Here's the application code where we run our  `main`  function. The  `GetAstronauts`  function takes an interface as its first argument allowing us to isolate and abstract away any HTTP logic from this file and its import list.

```
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func GetAstronauts(getWebRequest GetWebRequest) int {
	url := "http://api.open-notify.org/astros.json"
	bodyBytes := getWebRequest.FetchBytes(url)
	peopleResult := people{}
	jsonErr := json.Unmarshal(bodyBytes, &peopleResult)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return peopleResult.Number
}

func main() {
	liveClient := LiveGetWebRequest{}
	number := GetAstronauts(liveClient)

	fmt.Println(number)
}

```

The GetWebRequest interface specifies the following function:

```
type GetWebRequest interface {
	FetchBytes(url string) []byte
}

```

> Interfaces are inferred on rather than explicitly decorated onto a struct. This is different from languages like C# or Java.

The complete file named types.go looks like this:

```
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:"number"`
}

type GetWebRequest interface {
	FetchBytes(url string) []byte
}

type LiveGetWebRequest struct {
}

func (LiveGetWebRequest) FetchBytes(url string) []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

```

**Choosing what to abstract**

The above unit test is effectively only testing the  `json.Unmarshal`  function and our assumptions about what a valid HTTP response body would look like. This abstracting may be OK for our example, but our code coverage score will be lower.