# Concurrency 

### Topics:
1. Concurrency vs Parallellism
2. Goroutines
3. sync.WaitGroup
4. channels
5. select
6. sync.Mutex
7. Mutex vs Channels


# Concurrency Vs Parallelism

**Concurrency:** You want to execute 50 database queries at once and the database server is hosted on a separate machine, this is an I/O bound operation and your Go program will not need any CPU after your application has initiated the requests and the queries are being executed by the database’s machine. So, you can launch 50 Goroutines simultaneously, one for each query without draining any resources. Now assume that your application is hosted on a 4 core machine, so in this case, you will not be utilizing a single core completely.

**Parallelism:** You have to perform some heavy image processing or CSV parsing work and you have 50 images/CSV files, this is a CPU bound operation as each file is going to be processed by your Go application. You can still launch 50 Goroutines simultaneously on your same quad-core processor, but now only 4 tasks would be processed at any given time and all the remaining goroutines will be scheduled on any available CPU one after the other.

Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.

# Goroutines

1. Goroutines are functions or methods that run concurrently with other functions or methods. 
2. They can be thought of as lightweight threads.
3. Cost of creating and managing a goroutine is very small as compared to threads and that is why it is normal for an application to have thousands of goroutines running concurrently.
   
```go
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	// If you comment the below line, nothing will be printed
	// as the main goroutine will exit before the other has executed
	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine")
}

```
A new goroutine can be started just by adding the keyword  **“go”**  before a function call.

  **What if we could make this more efficient by only waiting till the other goroutine is executing rather than hardcoding the wait time?**  WaitGroups come to the rescue.

# sync.WaitGroup

We can modify the above program by simply replacing the sleep call with a WaitGroup.
```go
package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	fmt.Println("Hello world goroutine")
	wg.Done()
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go hello(wg)
	wg.Wait()
	fmt.Println("main goroutine")
}

```
Not only is the above program faster, as now we are not wasting any time and the program exits as soon as it is done, but it is more robust as well.

# Advantages of Goroutines over Threads

-   **Startup:**  Goroutines have a faster startup time than threads.
-   **Cheap:** They are only a few kb in stack size and the stack can grow and shrink according to needs of the application whereas in the case of threads the stack size has to be specified and is fixed.
-   **A single thread can execute multiple goroutines:**  Suppose there are 50 Goroutines which are I/O bound. If any Goroutine in that thread blocks for I/O say waiting for user input, then another Goroutine can be executed on that same thread.  **This avoids the need to create more threads at the OS level and thus saving a lot of time which is wasted in Context Switching**. All these are taken care of by the Go Runtime and an I/O bound task is converted to a CPU bound task magically.
-   **Channels:**  Channels are a way by which goroutines communicate with each other and it prevents race conditions from happening when accessing shared memory. Channels can be thought of as a pipe using which Goroutines communicate.

# Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine. It is the backbone behind Go’s approach to concurrency:

**_Do not communicate by sharing the memory; instead, share memory by communicating._**

So far we know how to launch goroutines and how to wait for that to finish using sync.WaitGroup, but  **what if we also want to return something from that goroutine?**  Let’s say that you have launched a batch job in a separate background goroutine and now you want to know whether the job failed or succeeded. For this tutorial, we will modify the above program to return the string rather than printing it, you can return anything from a goroutine using the same technique. Please go through the below program and read the comments.

```go
//code003.go  
package main

import (
	"fmt"
)

func hello(ch chan string) {
	ch <- "Hello world goroutine"
}
func main() {
	ch := make(chan string)
	go hello(ch)
	responseFromGoroutine := <-ch
	fmt.Println(responseFromGoroutine)
	fmt.Println("main goroutine")
}

```

# Sending and receiving from a channel

We can send and receive by placing <- in the appropriate position as explained:

```go
ch <- "Hello world goroutine" // Write to channel  
responseFromGoroutine := <-ch // Read from channel
```

Another  **very important thing to note here is that sends and receives are blocking**. This means, when data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly, when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.

# Deadlock

If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will  **panic**  at runtime with Deadlock.
```go
//code004.go  
package main  
func main() {  
  c:=make(chan string)  
  c<-"Hello"  
}
```
This will be the output:
```go
fatal error: all goroutines are asleep - deadlock!  
goroutine 1 [chan send]:
```
Now, let us try to fix this:
```go
//code005.go  
package main

import "fmt"

func write(c chan string) {
	c <- "Hello"
}

func main() {
	c := make(chan string)
	go write(c)
	fmt.Println(<-c)
}

```


So far, all the examples are using single capacity channels, Go also supports  **Buffered Channels**  which have capacity more than 1. You may not need it right away so just remember that there is something called Buffered Channels and when the need arises, you can read about it  [here](https://tour.golang.org/concurrency/3).

# Select

This statement is used to choose from multiple send/receive channel operations. It blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. The syntax is similar to switch except that each of the case statements will be a channel operation.
```go
//code006.go  
package main

import (
	"fmt"
	"time"
)

func task1(ch1 chan string) {
	time.Sleep(5 * time.Second)
	ch1 <- "Task 1 Complete"
}

func task2(ch2 chan string) {
	time.Sleep(2 * time.Second)
	ch2 <- "Task 2 Complete"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go task1(ch1)
	go task2(ch2)
	select {
	case str1 := <-ch1:
		fmt.Println(str1)
	case str2 := <-ch2:
		fmt.Println(str2)
	}

}

```
The select statement executes the case which occurs earlier. Task1 completes in 5 seconds and Task2 takes just 2 seconds, so the code will print “Task 2 complete” and exit.

# sync.Mutex

Although channels are good enough for all the use cases, it is an overkill for certain specific use cases, like incrementing a global variable from multiple Goroutines.
```go
//code007.go  
package main

import (
	"fmt"
	"sync"
)

var num = 0

func increment(wg *sync.WaitGroup) {
	num = num + 1
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	for i:=0;i<500;i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(num)
}

```

Try running the above program multiple times, you will get different output every time instead of the desired output, ie. 500.

You can still solve this problem using channels, but Mutex seems to be a better fit over here.
```go
//code008.go
package main

import (
	"fmt"
	"sync"
)

var num = 0

func increment(wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	num = num + 1
	mut.Unlock()
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go increment(wg, mut)
	}
	wg.Wait()
	fmt.Println(num)
}

```
Now, 500 is the guaranteed output every single time.

# Mutex vs Channels

As mentioned before, you could use channels wherever you can use a mutex. Go’s official documentation states that  _“A common Go newbie mistake is to over-use channels and goroutines just because it’s possible, and/or because it’s fun.”_

To summarise, use channels to distribute tasks to Goroutines and use Mutex when multiple Goroutines need to access or modify a common resource/critical section which in the above case is a global variable.