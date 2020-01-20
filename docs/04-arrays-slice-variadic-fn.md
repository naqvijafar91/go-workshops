
# Arrays and Slices


### Arrays

An array is a collection of elements that belong to the same type. For example, the collection of integers 5, 8, 9, 79, 76 forms an array. Just like Java, mixing values of different types, for example, an array that contains both strings and integers is not allowed in Go.

##### Declaration



There are different ways to declare arrays. Let's look at them one by one.

```go
package main

import (  
    "fmt"
)


func main() {  
    var a [3]int //int array with length 3
    fmt.Println(a)
}

```

[Run in playground](https://play.golang.org/p/Zvgh82u0ej)

**var a [3]int**  declares an integer array of length 3.  **All elements in an array are automatically assigned the zero value of the array type**. 



```go
package main

import (  
    "fmt"
)

func main() {  
    var a [3]int //int array with length 3
    a[0] = 12 // array index starts at 0
    a[1] = 78
    a[2] = 50
    fmt.Println(a)
}
```

[Run in playground](https://play.golang.org/p/WF0Uj8sv39)

a[0] assigns value to the first element of the array. The program will  **output**  `[12 78 50]`

Using  **short hand declaration**.

```go
package main 

import (  
    "fmt"
)

func main() {  
    a := [3]int{12, 78, 50} // short hand declaration to create array
    fmt.Println(a)
}

```

[Run in playground](https://play.golang.org/p/NKOV04zgI6)

The program above will print the same  **output**  `[12 78 50]`

It is not necessary that all elements in an array have to be assigned a value during a short hand declaration.

```go
package main

import (  
    "fmt"
)

func main() {  
    a := [3]int{12} 
    fmt.Println(a)
}

```

[Run in playground](https://play.golang.org/p/AdPH0kXRly)

The remaining 2 elements are assigned  `0`  automatically. The program will  **output**  `[12 0 0]`


**The size of the array is a part of the type.**  Hence  `[5]int`  and  `[25]int`  are distinct types. Because of this, arrays cannot be resized. Don't worry about this restriction since  `slices`  exist to overcome this.

```go
package main

func main() {  
    a := [3]int{5, 78, 8}
    var b [5]int
    b = a //not possible since [3]int and [5]int are distinct types
}

```

[Run in playground](https://play.golang.org/p/kBdot3pXSB)

the compiler will throw error  _main.go:6: cannot use a (type [3]int) as type [5]int in assignment_.

#### Arrays are value types not refernce types

This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. If changes are made to the new variable, it will not be reflected in the original array.

```go
package main

import "fmt"

func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}

```

[Run in playground](https://play.golang.org/p/-ncGk1mqPd)


```
a is [USA China India Germany France]  
b is [Singapore China India Germany France]  

```

Similarly, when arrays are passed to functions as parameters, they are passed by value and the original array is unchanged.

```go
package main

import "fmt"

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}
func main() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}

```

[Run in playground](https://play.golang.org/p/e3U75Q8eUZ)


```
before passing to function  [5 6 7 8 8]  
inside function  [55 6 7 8 8]  
after passing to function  [5 6 7 8 8]  
```

##### Length of an array

The length of the array is found by passing the array as a parameter to the  `len`  function.

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of a is",len(a))
}
```

[Run in playground](https://play.golang.org/p/UrIeNlS0RN)

The  **output**  of the above program is  `length of a is 4`

##### Iterating arrays using range

The  `for`  loop can be used to iterate over elements of an array.

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}
```

[Run in playground](https://play.golang.org/p/80ejSTACO6)

```
0 th element of a is 67.70  
1 th element of a is 89.80  
2 th element of a is 21.00  
3 th element of a is 78.00  

```

Go provides a better and concise way to iterate over an array by using the  **range**  form of the  `for`  loop.  `range`  returns both the index and the value at that index. Let's rewrite the above code using the range.

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for i, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)
}
```

[Run in playground](https://play.golang.org/p/Ji6FRon36m)

 The  **output**  of the program is,

```
0 the element of a is 67.70  
1 the element of a is 89.80  
2 the element of a is 21.00  
3 the element of a is 78.00

sum of all elements of a 256.5  
```


```go
for _, v := range a { //ignores index  
}
```

The above for loop ignores the index. Similarly, the value can also be ignored.

##### Multidimensional arrays



```go
package main

import (  
    "fmt"
)

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {  
    a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(b)
}

```

[Run in playground](https://play.golang.org/p/InchXI4yY8)



```
lion tiger  
cat dog  
pigeon peacock 

apple samsung  
microsoft google  
AT&T T-Mobile  
```

## Slices

**A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are the just references to existing arrays.**

##### Creating a slice

A slice with elements of type T is represented by  `[]T`

```go
package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b)
}

```

[Run in playground](https://play.golang.org/p/Za6w5eubBB)

The syntax  `a[start:end]`  creates a slice from array  `a`  starting from index  `start`  to index  `end - 1`. 

Another way to create a slice.

```go
package main

import (  
    "fmt"
)

func main() {  
    c := []int{6, 7, 8} //creates and array and returns a slice reference
    fmt.Println(c)
}
```

[Run in playground](https://play.golang.org/p/_Z97MgXavA)

In the above program in line no. 9,  `c:= []int{6, 7, 8}`  creates an array with 3 integers and returns a slice reference which is stored in c.

##### modifying a slice

> **A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications done to the
> slice will be reflected in the underlying array.**

```go
package main

import (  
    "fmt"
)

func main() {  
    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59} // This is an array
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after",darr) 
}

```

[Run in playground](https://play.golang.org/p/6FinudNf1k)


```
array before [57 89 90 82 100 78 67 69 59]  
array after [57 89 91 83 101 78 67 69 59]  

```

When many slices share the same underlying array, the changes that each one makes will be reflected in the array.

```go
package main

import (  
    "fmt"
)

func main() {  
    numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2", numa)
}

```

[Run in playground](https://play.golang.org/p/mdNi4cs854)


```
array before change 1 [78 79 80]  
array after modification to slice nums1 [100 79 80]  
array after modification to slice nums2 [100 101 80]  
```

From the output, it's clear that when slices share the same array, the modifications which each one makes are reflected in the array.

#### Length and capacity of a slice

The length of the slice is the number of elements in the slice. 

 **The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.**

```go
package main

import (  
    "fmt"
)

func main() {  
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
}

```

[Run in playground](https://play.golang.org/p/a1WOcdv827)



  **The capacity of  `fruitslice`  is the no of elements in  `fruitarray`  starting from index  `1`  i.e from  `orange`  and that value is  `6`.** Hence the capacity of fruitslice is 6. 

The  [program](https://play.golang.org/p/a1WOcdv827)  outputs  **length of slice 2 capacity 6**.

A slice can be re-sliced up to its capacity. Anything beyond that will cause the program to throw a run time error.

```go
package main

import (  
    "fmt"
)

func main() {  
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}

```

[Run in playground](https://play.golang.org/p/GcNzOOGicu)

In line no. 11 of the above program,`fruitslice`  is re-sliced to its capacity. The above program outputs,

```
length of slice 2 capacity 6  
After re-slicing length is 6 and capacity is 6  

```
#### creating a slice using make

_func make([]T, len, cap) []T_  can be used to create a slice bypassing the type, length, and capacity. The capacity parameter is optional and defaults to the length. 

```go
package main

import (  
    "fmt"
)

func main() {  
    i := make([]int, 5, 5)
    fmt.Println(i)
}

```

[Run in playground](https://play.golang.org/p/M4OqxzerxN)

The values are zeroed by default when a slice is created using make. The above program will output  `[0 0 0 0 0]`.

#### Appending to a slice

The definition of append function is  `func append(s []T, x ...T) []T`.

**x ...T**  in the function definition means that the function accepts a variable number of arguments for the parameter x. These type of functions are called  [variadic functions](https://golangbot.com/variadic-functions/).

> ### Q) If slices are backed by arrays and arrays themselves are of fixed length then how come a slice is of dynamic length?

A) When new elements are appended to the slice, a new array is created. The elements of the existing array are copied to this new array and a new slice reference for this new array is returned. The capacity of the new slice is now twice that of the old slice. 

```go
package main

import (  
    "fmt"
)

func main() {  
    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}

```

[Run in playground](https://play.golang.org/p/VUSXCOs1CF)


```
cars: [Ferrari Honda Ford] has old length 3 and capacity 3  
cars: [Ferrari Honda Ford Toyota] has new length 4 and capacity 6  

```


```go
package main

import (  
    "fmt"
)

func main() {  
	// Append works on nil slices as well
    var names []string //zero value of a slice is nil
    if names == nil {
        fmt.Println("slice is nil going to append")
        names = append(names, "John", "Sebastian", "Vinay")
        fmt.Println("names contents:",names)
    }
}

```

[Run in playground](https://play.golang.org/p/x_-4XAJHbM)

```
slice is nil going to append  
names contents: [John Sebastian Vinay]  

```

It is also possible to append one slice to another using the  `...`  operator. You can learn more about this operator in the  [variadic functions](https://golangbot.com/variadic-functions/)  tutorial.

```go
package main

import (  
    "fmt"
)

func main() {  
    veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)
    fmt.Println("food:",food)
}

```

[Run in playground](https://play.golang.org/p/UnHOH_u6HS)

The output of the program is  `food: [potatoes tomatoes brinjal oranges apples]`

#### Internal Representation of a slice

Slices can be thought of as being represented internally by a structure type. This is how it looks,

```
// A slice contains the length, capacity and a pointer to the zeroth
// element of the array. 
type slice struct {  
    Length        int
    Capacity      int
    ZerothElement *byte
}
```

When a slice is passed to a function, even though it's passed by value, the pointer variable will refer to the same underlying array. When a slice is passed to a function as a parameter, changes made inside the function are visible outside the function too. 

```go
package main

import (  
    "fmt"
)

func subtactOne(numbers []int) {  
    for i := range numbers {
        numbers[i] -= 2
    }
    numbers = append(numbers, 786)
	fmt.Println("slice inside function call", numbers)

}
func main() {  
    nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
    fmt.Println("slice after function call", nos) //modifications are visible outside
}

```

[Run in playground](https://play.golang.org/p/IzqDihNifq)

Output of the above  [program](https://play.golang.org/p/bWUb6R-1bS)  is,

```
slice before function call [8 7 6]
slice inside function call [6 5 4 786]
slice after function call [6 5 4] 

```

#### Q) Why that happened?

A) Refer to the internal representation.

#### Memory Optimisation

Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected. This might be of concern when it comes to memory management. Let's assume that we have a very large array and we are interested in processing only a small part of it. Henceforth we create a slice from that array and start processing the slice. The important thing to be noted here is that the array will still be in memory since the slice references it.

One way to solve this problem is to use the  [copy](https://golang.org/pkg/builtin/#copy)  function  `func copy(dst, src []T) int`  to make a copy of that slice. This way we can use the new slice and the original array can be garbage collected.

```go
package main

import (  
    "fmt"
)

func countries() []string {  
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}
func main() {  
    countriesNeeded := countries()
    fmt.Println(countriesNeeded)
}

```

[Run in playground](https://play.golang.org/p/35ayYBhcDE)

Now  `countries`  array can be garbage collected since  `neededCountries`  is no longer referenced.

# Variadic Functions

### What is a variadic function?

Functions, in general, accept only a fixed number of arguments.  _A variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by ellipsis  **...**, then the function can accept any number of arguments for that parameter._

**Only the last parameter of a function can be variadic. We will learn why this is the case in the next section of this tutorial.**

### Syntax

```go
func hello(a int, b ...int) {  
}

```

 This function can be called by using the syntax.

```go
hello(1, 2) //passing one argument "2" to b  
hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b  

```
It is also possible to pass zero arguments to a variadic function.

```lanaguage
hello(1)  
```
Let's try to make the first parameter of the  `hello`  function variadic.

The syntax will look like

```go
// Not allowed
func hello(b ...int, a int) {  
}
```

 `syntax error: cannot use ... with non-final parameter b`

### Examples and understanding how variadic functions work

Program to find whether an integer exists in an input list of integers.

```go
package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}

```

[Run in playground](https://play.golang.org/p/7occymiS6s)


**The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter. The find function expects a variadic  `int`  argument. Hence these three arguments will be converted by the COMPILER to a slice of type int  `[]int{89, 90, 95}`  and then it will be passed to the  `find`  function.**


```
type of nums is []int  
89 found at index 0 in [89 90 95]

type of nums is []int  
45 found at index 2 in [56 67 45 90 109]

type of nums is []int  
78 not found in  [38 56 98]

type of nums is []int  
87 not found in  []  
```

### Append is a variadic function

Have you ever wondered how the  [append](https://golang.org/pkg/builtin/#append)  function in the standard library used to append values to a  [slice](https://golangbot.com/arrays-and-slices/)  accepts any number of arguments. It's because it's a variadic function.

```
func append(slice []Type, elems ...Type) []Type  
```

### Passing a slice to a variadic function

Let's pass a slice to a variadic function and find out what happens from the below example.

```go
package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    nums := []int{89, 90, 95}
    find(89, nums)
}

```

[Run in playground](https://play.golang.org/p/A-DNilpH2L)

The above program will fail with compilation error  `./prog.go:23:10: cannot use nums (type []int) as type int in argument to find`

 The signature of the  `find`  function is provided below,

```go
func find(num int, nums ...int)  

```

In this case,  `nums`  is already a  `[]int`  slice and the compiler tries to create a new  `[]int`  i.e the compiler tries to do

```
find(89, []int{nums})  
```

which will fail since  `nums`  is a  `[]int`  and not a  `int`.

So is there a way to pass a slice to a variadic function? The answer is yes.

**There is a syntactic sugar which can be used to pass a slice to a variadic function. You have to suffix the slice with ellipsis  `...`  If that is done, the slice is directly passed to the function without a new slice being created.**



```go
package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    nums := []int{89, 90, 95}
    find(89, nums...)
}

```
```
type of nums is []int  
89 found at index 0 in [89 90 95]  
```
[Run in playground](https://play.golang.org/p/IvzwhzhFsT)

### WAT

Just be sure you know what you are doing when you are modifying a slice inside a variadic function.

Let's look at a simple example.

```go
package main

import (  
    "fmt"
)

func change(s ...string) {  
    s[0] = "Go"
}

func main() {  
    welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)
}

```

[Run in playground](https://play.golang.org/p/R0GsuW7rdd)


```
[Go world]
```

Here is one more program to understand variadic functions.

```go
package main

import (  
    "fmt"
)

func change(s ...string) {  
    s[0] = "Go"
    s = append(s, "playground")
    fmt.Println(s)
}

func main() {  
    welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)
}

```
### What will be the output?

[Run in playground](https://play.golang.org/p/WdbFIkdLoe)