## Variables

Variable is the name given to a memory location to store the value of a specific type. There are various syntaxes to declare variables in go.

### Declaration and Type inference

If a variable has an initial value, Go will automatically be able to infer the type of that variable using that initial value. Hence if a variable has an initial value, the _type_ in the variable declaration can be omitted.

```go

package main

import  "fmt"

func  main() {

var age int  =  29  // variable declaration with initial value
var  age2 = 29  // type will be inferred
fmt.Println("my age is", age)
fmt.Println("my age is", age2)
var  width, height  int = 100, 50  //declaring multiple variables
fmt.Println("width is", width, "height is", height)

var  width2, height2  int // Default value is 0
fmt.Println("width is", width2, "height is", height2)

// Block declaration
var (
name3 = "naveen"
age3 = 29
height3  int
)
fmt.Println("my name is", name3, ", age is", age3, "and height is", height3)
}

```

### Shorthand declaration

Go also provides another concise way of declaring variables. This is known as shorthand declaration and it uses **:=** operator.

**name := initial-value** is the shorthand syntax to declare a variable.

```go
package main

import  "fmt"

func  main() {
name, age := "naveen", 29  //short hand declaration
fmt.Println("my name is", name, "age is", age)
}

```

[Run in playground](https://play.golang.org/p/ctqgw4w6kx)

If you run the above program, you can see `my name is Naveen age is 29` as the output.

Shorthand syntax can only be used when at least one of the variables on the left side of **:=** is newly declared.

```go
package main

import  "fmt"

func  main() {
a, b := 20, 30  // declare variables a and b
fmt.Println("a is", a, "b is", b)
b, c := 40, 50  // b is already declared but c is new
fmt.Println("b is", b, "c is", c)
b, c = 80, 90  // assign new values to already declared variables b and c
fmt.Println("changed b is", b, "c is", c)
}

```

[Run in playground](https://play.golang.org/p/MSUYR8vazB)

In the above program, in line no. 8, **b** has already been declared but **c** is newly declared and hence it works and outputs

```

a is 20 b is 30

b is 40 c is 50

changed b is 80 c is 90
```

Whereas if we run the program below,

```go
package main

import "fmt"

func main() {

	a, b := 20, 30 //a and b declared

	fmt.Println("a is", a, "b is", b)

	a, b := 40, 50 //error, no new variables

}
```

[Run in playground](https://play.golang.org/p/EYTtRnlDu3)

it will throw error `no new variables on the left side of `:=` This is because both the variables **a** and **b** has already been declared and there are no new variables in the left side of **:=**

Since Go is strongly typed, variables declared as belonging to one type cannot be assigned a value of another type.

```go
package main

func main() {

	age := 29 // age is int

	age = "naveen" // error since we are trying to assign a string to a variable of type int

}

```

[Run in playground](https://play.golang.org/p/K5rz4gxjPj)

## Types

The following are the basic types available in go

- bool

* Numeric Types

- int8, int16, int32, int64, int

* uint8, uint16, uint32, uint64, uint -Unsigned integers only contain positive numbers (or zero).

- float32, float64

* complex64, complex128

- byte - same as uint8. Bytes are an extremely common unit of measurement used on computers (1 byte = 8 bits, 1024 bytes = 		1 kilobyte, 1024 kilobytes = 1 megabyte, …) and therefore Go's byte data type is often used in the definition of other types.
- 
* rune - the same as int32

- string

### bool

A bool type represents a boolean and is either _true_ or _false_.

```go
package main

import "fmt"

func main() {

	a := true

	b := false

	fmt.Println("a:", a, "b:", b)

	c := a && b

	fmt.Println("c:", c)

	d := a || b

	fmt.Println("d:", d)

}
```

[Run in playground](https://play.golang.org/p/v_W3HQ0MdY)

```
a: true b: false
c: false
d: true
```

### Signed integers

**int8:** represents 8 bit signed integers

**size:** 8 bits

**range:** -128 to 127

**int16:** represents 16 bit signed integers

**size:** 16 bits

**range:** -32768 to 32767

**int32:** represents 32 bit signed integers

**size:** 32 bits

**range:** -2147483648 to 2147483647

**int64:** represents 64 bit signed integers

**size:** 64 bits

**range:** -9223372036854775808 to 9223372036854775807

**int:** represents 32 or 64-bit integers depending on the underlying platform. You should generally be using _int_ to represent integers unless there is a need to use a specifically sized integer.

**size:** 32 bits in 32 bit systems and 64 bit in 64 bit systems.

**range:** -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

```go
package main

import "fmt"

func main() {

	var a int = 89

	b := 95

	fmt.Println("value of a is", a, "and b is", b)

}
```

[Run in playground](https://play.golang.org/p/NyDPsjkma3)

The above program will output `value of a is 89 and b is 95`

The type of a variable can be printed using a **%T** format specifier in the `Printf` function. Go has a package [unsafe](https://golang.org/pkg/unsafe/) which has a [Sizeof](https://golang.org/pkg/unsafe/#Sizeof) function which returns in bytes the size of the variable passed to it. _unsafe_ package should be used with care as the code using it might have portability issues.

```go
package main

import (
	"fmt"

	"unsafe"
)

func main() {

	var a int = 89

	b := 95

	fmt.Println("value of a is", a, "and b is", b)

	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) //type and size of a

	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b

}
```

[Run in playground](https://play.golang.org/p/mFsmjVk5oc)

The above program will produce the output

```
value of a is 89 and b is 95
type of a is int, size of a is 4
type of b is int, size of b is 4
```

We can infer from the above output that a and b are of type _int_ and they are _32 bit sized(4 bytes)_. The output will vary if you run the above program on a 64-bit system. In a 64 bit system, a and b occupy 64 bits (8 bytes).

### Unsigned integers

**uint8:** represents 8 bit unsigned integers

**size:** 8 bits

**range:** 0 to 255

**uint16:** represents 16 bit unsigned integers

**size:** 16 bits

**range:** 0 to 65535

**uint32:** represents 32 bit unsigned integers

**size:** 32 bits

**range:** 0 to 4294967295

**uint64:** represents 64 bit unsigned integers

**size:** 64 bits

**range:** 0 to 18446744073709551615

**uint:** represents 32 or 64-bit unsigned integers depending on the underlying platform.

**size:** 32 bits in 32 bit systems and 64 bits in 64 bit systems.

**range:** 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

### Floating-point types

**float32:** 32-bit floating-point numbers

**float64:** 64-bit floating-point numbers

### Complex types

**complex64:** complex numbers which have float32 real and imaginary parts

**complex128:** complex numbers with float64 real and imaginary parts

The builtin function **[complex](https://golang.org/pkg/builtin/#complex)** is used to construct a complex number with real and imaginary parts. The complex function has the following definition

```

func complex(r, i FloatType) ComplexType

```

It takes a real and imaginary part as a parameter and returns a complex type. _Both the real and imaginary parts must be of the same type. ie either float32 or float64. If both the real and imaginary parts are float32, this function returns a complex value of type complex64. If both the real and imaginary parts are of type float64, this function returns a complex value of type complex128_

Complex numbers can also be created using the shorthand syntax

```go

c := 6 + 7i

```

Let's write a small program to understand complex numbers.

```go
package main

import (
	"fmt"
)

func main() {

	c1 := complex(5, 7)

	c2 := 8 + 27i

	cadd := c1 + c2

	fmt.Println("sum:", cadd)

	cmul := c1 * c2

	fmt.Println("product:", cmul)

}

```

[Run in playground](https://play.golang.org/p/kEz1uKCdKs)

This program will output

```
sum: (13+34i)
product: (-149+191i)
```

### Other numeric types

**byte** is an alias of uint8

**rune** is an alias of int32

We will discuss bytes and runes in more detail when we learn about strings.

### string type

Strings are a collection of bytes in golang. It's alright if this definition doesn't make any sense.

### Type Conversion

Go is very strict about explicit typing. There is no automatic type promotion or conversion.

```go
package main

import (

"fmt"

)

func  main() {
i := 55  //int
j := 67.8  //float64
// sum := i + j //WRONG: This will give error int + float64 not allowed
sum := i + int(j) //j is converted to int
fmt.Println(sum)

sample :=  10
var jk float64  =  float64(sample)  //this statement will not work without explicit conversion
fmt.Println("j", jk)
}
```

## Constants

### Definition

Consider the following code,

```go

var  a  int = 50

var  b  string = "I love Go"

```

Constants, as the name indicate, cannot be reassigned again to any other value and hence the below program will not work and it will fail with compilation error **cannot assign to a**.

```go
package main


func  main() {

const  a = 55  //allowed

a = 89  //reassignment not allowed

}

```

[Run in playground](https://play.golang.org/p/b2J8_UQobb)

The value of a constant should be known at compile time. Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.

```go
package main

import (
	"fmt"

	"math"
)

func main() {
	fmt.Println("Hello, playground")
	var a = math.Sqrt(4) //allowed
	const b = math.Sqrt(4) //not allowed
}
```

[Run in playground](https://play.golang.org/p/dCON1LzCTw)

In the above program, `a` is a variable and hence it can be assigned to the result of the function `math.Sqrt(4)`

_b_ is a constant and the value of b needs to know at compile time. The function `math.Sqrt(4)` will be evaluated only during run time and hence `const b = math.Sqrt(4)` throws **error main.go:11: const initializer math.Sqrt(4) is not constant.**

### String Constants

Any value enclosed between double quotes is a string constant in Go. For example strings like "Hello World" or "Sam" are all constants in Go.

What type does a string constant belong to? The answer is they are **untyped**.

**A string constant like "Hello World" does not have any type**.

```go
const  hello = "Hello World"
```

In the above case, we have assigned "Hello World" to a named constant **hello**. Now does the constant _hello_ have a type? The answer is No. The constant still doesn't have a type.

```go

package main


import (
"fmt"
)

func  main() {
var  name = "Sam"
fmt.Printf("type %T value %v", name, name)

}
```

The following code creates a typed constant.

```go

const typedhello string = "Hello World"

```

**_typedhello_ in the above code is a constant of type string.**

Go is a strongly typed language. Mixing types during the assignment is not allowed.

```go
package main

func main() {

	var defaultName = "Sam" //allowed

	// Define a new type myString
	type myString string

	var customName myString = "Sam" //allowed

	customName = defaultName //not allowed
}
```

[Run in playground](https://play.golang.org/p/1Q-vudNn_9)

In the above code, we first create a variable _defaultName_ and assign it to the constant _Sam_. **The default type of the constant Sam is a string, so after the assignment default name is of type String.**

**Hence the assignment custom name = default name is not allowed and the compiler throws an error _main.go:7:20: cannot use default name (type string) as type myString in assignment_**
