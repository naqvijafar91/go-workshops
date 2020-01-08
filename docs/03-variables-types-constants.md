## Variables

  

### What is a variable

  

Variable is the name given to a memory location to store a value of a specific type. There are various syntaxes to declare variables in go.

  

### Declaring a single variable

  

**var name type** is the syntax to declare a single variable.

  

```go

package main

  

import  "fmt"

  

func  main() {

var  age  int  // variable declaration

fmt.Println("my age is", age)

}

  

```

  

[Run in playground](https://play.golang.org/p/XrveIxw_YI)

  

The statement `var age int` declares a variable named _age_ of type _int_. We have not assigned any value for the variable. If a variable is not assigned any value, go automatically initialises it with the **zero value** of the variable's type. In this case, age is assigned the value `0`. 

  

```
my age is 0
```

  

A variable can be assigned to any value of its type. In the above program age can be assigned any integer value.

  

```go

package main

  

import  "fmt"

  

func  main() {

var  age  int  // variable declaration

fmt.Println("my age is ", age)

age = 29  //assignment

fmt.Println("my age is", age)

age = 54  //assignment

fmt.Println("my new age is", age)

}

  

```

  

[Run in playground](https://play.golang.org/p/z4nKMjBxLx)

  

The above program will produce the following output.

  

```

my age is 0

my age is 29

my new age is 54
```

  

### Declaring a variable with initial value

  

A variable can also be given a initial value when it is declared.

  

**var name type = initialvalue** is the syntax to declare a variable with initial value.

  

```go

package main

import  "fmt"

func  main() {
var  age  int = 29  // variable declaration with initial value
fmt.Println("my age is", age)
}

  

```

  

[Run in playground](https://play.golang.org/p/TFfpzsrchh)

  

In the above program, age is variable of type int and has initial value 29. If you run the above program, you can see the following output. It shows that age has been initialised with the value 29.

  

```
my age is 29

```

  

### Type inference

  

If a variable has an initial value, Go will automatically be able to infer the type of that variable using that initial value. Hence if a variable has an initial value, the _type_ in the variable declaration can be omitted.

  

If the variable is declared using the syntax **var name = initialvalue**, Go will automatically infer the type of that variable from the initial value.

  

In the following example, you can see that the type _int_ of the variable _age_ has been removed in line no. 6. Since the variable has a initial value of 29, go can infer that it is of type int.

  

```go

package main


import  "fmt"

  

func  main() {

var  age = 29  // type will be inferred

fmt.Println("my age is", age)

}

  

```

  

[Run in playground](https://play.golang.org/p/FgNbfL3WIt)

  

### Multiple variable declaration

  

Multiple variables can be declared in a single statement.

  

**var name1, name2 type = initialvalue1, initialvalue2** is the syntax for multiple variable declaration.

  

```go

package main

  

import  "fmt"

  

func  main() {

var  width, height  int = 100, 50  //declaring multiple variables

  

fmt.Println("width is", width, "height is", height)

}

  

```

  

[Run in playground](https://play.golang.org/p/4aOQyt55ah)

  

The _type_ can be omitted if the variables have initial value. The program below declares multiple variables using type inference.

  

```go

package main

  

import  "fmt"

  

func  main() {

var  width, height = 100, 50  //"int" is dropped

  

fmt.Println("width is", width, "height is", height)

}

  

```

  

[Run in playground](https://play.golang.org/p/sErofTJ6g-)

  

The above program will print `width is 100 height is 50` as the output.

  

If the initial value is not specified for width and height, they will have `0` assigned as their initial value.

  

```go

package main

  

import  "fmt"

  

func  main() {

var  width, height  int

fmt.Println("width is", width, "height is", height)

width = 100

height = 50

fmt.Println("new width is", width, "new height is ", height)

}

  

```

  

[Run in playground](https://play.golang.org/p/DM00pcBbsu)

  

The above program will print

  

```undefined

width is 0 height is 0

new width is 100 new height is 50

  

```

  

There might be cases where we would want to declare variables belonging to different types in a single statement. The syntax for doing that is

  

```go

var (

name1 = initialvalue1

name2 = initialvalue2

)

```

  

The following program uses the above syntax to declare variables of different types.

  

```go

package main

  

import  "fmt"

  

func  main() {

var (

name = "naveen"

age = 29

height  int

)

fmt.Println("my name is", name, ", age is", age, "and height is", height)

}

  

```

  

[Run in playground](https://play.golang.org/p/7pkp74h_9L)

  

Here we declare a variable **name of type string, age and height of type int**. Running the above program will produce the output `my name is naveen , age is 29 and height is 0`

  

### Short hand declaration

  

Go also provides another concise way for declaring variables. This is known as short hand declaration and it uses **:=** operator.

  

**name := initialvalue** is the short hand syntax to declare a variable.

  

```go

package main

  

import  "fmt"

  

func  main() {

name, age := "naveen", 29  //short hand declaration
fmt.Println("my name is", name, "age is", age)

}

  

```

  

[Run in playground](https://play.golang.org/p/ctqgw4w6kx)

  

If you run the above program, you can see `my name is naveen age is 29` as the output.

  

Short hand declaration requires initial values for all variables in the left hand side of the assignment. The following program will thrown an error `cannot assign 1 values to 2 variables`. This is because **age has not been assigned a value.**

  

```go

package main

  

import  "fmt"

  

func  main() {

name, age := "naveen"  //error

fmt.Println("my name is", name, "age is", age)

}

  

```

  

[Run in playground](https://play.golang.org/p/wZd2HmDvqw)

  

Short hand syntax can only be used when at least one of the variables in the left side of **:=** is newly declared. Consider the following program,

  

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

  

import  "fmt"

  

func  main() {

a, b := 20, 30  //a and b declared

fmt.Println("a is", a, "b is", b)

a, b := 40, 50  //error, no new variables

}

  

```

  

[Run in playground](https://play.golang.org/p/EYTtRnlDu3)

  

it will throw error `no new variables on left side of :=` This is because both the variables **a** and **b** have already been declared and there are no new variables in the left side of **:=**

Since Go is strongly typed, variables declared as belonging to one type cannot be assigned a value of another type. The following program will throw an error `cannot use "naveen" (type string) as type int in assignment` because age is declared as type int and we are trying to assign a string value to it.

  

```go

package main

  

func  main() {

age := 29  // age is int

age = "naveen"  // error since we are trying to assign a string to a variable of type int

}

  

```

  

[Run in playground](https://play.golang.org/p/K5rz4gxjPj)


  

## Types

  
  

The following are the basic types available in go

  

- bool

- Numeric Types

- int8, int16, int32, int64, int

- uint8, uint16, uint32, uint64, uint

- float32, float64

- complex64, complex128

- byte

- rune

- string

  

### bool

  

A bool type represents a boolean and is either _true_ or _false_.

  

```go

package main

  

import  "fmt"

  

func  main() {

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

  

In the program above, a is assigned to true and b is assigned a false value.

  

c is assigned the value of a && b. _The && operator returns true only when both a and b are true_. So in this case c is false.

  

_The || operator returns true when either a or b is true_. In this case d is assigned to true since a is true. We will get the following output for this program.

  

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

  

**int:** represents 32 or 64 bit integers depending on the underlying platform. You should generally be using _int_ to represent integers unless there is a need to use a specific sized integer.

**size:** 32 bits in 32 bit systems and 64 bit in 64 bit systems.

**range:** -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

  

```go

package main

  

import  "fmt"

  

func  main() {

var  a  int = 89

b := 95

fmt.Println("value of a is", a, "and b is", b)

}

  

```

  

[Run in playground](https://play.golang.org/p/NyDPsjkma3)

  

The above program will output `value of a is 89 and b is 95`

  

In the above program _a is of type int_ and _the type of b is inferred from the value assigned to it (95)._ As we have stated above, the size of _int is 32 bit in 32 bit systems and 64 bit in 64 bit systems_. Lets go ahead and verify this claim.

  

The type of a variable can be printed using **%T** format specifier in `Printf` function. Go has a package [unsafe](https://golang.org/pkg/unsafe/) which has a [Sizeof](https://golang.org/pkg/unsafe/#Sizeof) function which returns in bytes the size of the variable passed to it. _unsafe_ package should be used with care as the code using it might have portability issues.

  

The following program outputs the type and size of both variables a and b. `%T` is the format specifier to print the type and `%d` is used to print the size.

  

```go

package main

  

import (

"fmt"

"unsafe"

)

  

func  main() {

var  a  int = 89

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

  

We can infer from the above output that a and b are of type _int_ and they are _32 bit sized(4 bytes)_. The output will vary if you run the above program on a 64 bit system. In a 64 bit system, a and b occupy 64 bits (8 bytes).

  

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

  

**uint :** represents 32 or 64 bit unsigned integers depending on the underlying platform.

**size :** 32 bits in 32 bit systems and 64 bits in 64 bit systems.

**range :** 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

  

### Floating point types

  

**float32:** 32 bit floating point numbers

**float64:** 64 bit floating point numbers

  

The following is a simple program to illustrate integer and floating point types

  

```go

package main

  

import (

"fmt"

)

  

func  main() {

a, b := 5.67, 8.97

fmt.Printf("type of a %T b %T\n", a, b)

sum := a + b

diff := a - b

fmt.Println("sum", sum, "diff", diff)

  

no1, no2 := 56, 89

fmt.Println("sum", no1+no2, "diff", no1-no2)

}

  

```

  

[Run in playground](https://play.golang.org/p/upwUCprT-j)

  

The type of a and b is inferred from the value assigned to them. In this case a and b are of type float64.(float64 is the default type for floating point values). We add a and b and assign it to a variable sum. We subtract b from a and assign it to diff. Then the sum and diff is printed. Similar computation is done with no1 and no2. The above program will print

  

```

type of a float64 b float64

sum 14.64 diff -3.3000000000000007

sum 145 diff -33

  

```

  

### Complex types

  

**complex64:** complex numbers which have float32 real and imaginary parts

**complex128:** complex numbers with float64 real and imaginary parts

  

The builtin function **[complex](https://golang.org/pkg/builtin/#complex)** is used to construct a complex number with real and imaginary parts. The complex function has the following definition

  

```
func complex(r, i FloatType) ComplexType
```

  

It takes a real and imaginary part as parameter and returns a complex type. _Both the real and imaginary parts must be of the same type. ie either float32 or float64. If both the real and imaginary parts are float32, this function returns a complex value of type complex64. If both the real and imaginary parts are of type float64, this function returns a complex value of type complex128_

  

Complex numbers can also be created using the shorthand syntax

  

```go
c := 6 + 7i
```

  

Lets write a small program to understand complex numbers.

  

```go

package main

  

import (

"fmt"

)

  

func  main() {

c1 := complex(5, 7)

c2 := 8 + 27i

cadd := c1 + c2

fmt.Println("sum:", cadd)

cmul := c1 * c2

fmt.Println("product:", cmul)

}

  

```

  

[Run in playground](https://play.golang.org/p/kEz1uKCdKs)

  

In the above program, c1 and c2 are two complex numbers. c1 has 5 as real part and 7 as the imaginary part. c2 has real part 8 and imaginary part 27. `cadd` is assigned the sum of c1 and c2 and `cmul` is assigned the product of c1 and c2. This program will output

  

```

sum: (13+34i)

product: (-149+191i)

  

```

  

### Other numeric types

  

**byte** is an alias of uint8

**rune** is an alias of int32

  

We will discuss bytes and runes in more detail when we learn about strings.

  

### string type

  

Strings are a collection of bytes in golang. It's alright if this definition doesn't make any sense. For now we can assume a string to be a collection of characters. 

  

Lets write a program using strings.

  

```go

package main

  

import (

"fmt"

)

  

func  main() {

first := "Naveen"

last := "Ramanathan"

name := first +" "+ last

fmt.Println("My name is",name)

}

  

```

  

[Run in playground](https://play.golang.org/p/CI6phwSVel)

  

In the above program, _first_ is assigned the string _"Naveen"_, _last_ is assigned the string "Ramanathan". Strings can be concatenated using the + operator. _name_ is assigned the value of _first_ concatenated to a _space_ followed by _last_. The above program will print `My name is Naveen Ramanathan` as the output.


  

### Type Conversion

  

Go is very strict about explicit typing. There is no automatic type promotion or conversion. Lets look at what this means with an example.

  

```go

package main

  

import (

"fmt"

)

  

func  main() {

i := 55  //int

j := 67.8  //float64

sum := i + j //int + float64 not allowed

fmt.Println(sum)

}

  

```

  

[Run in playground](https://play.golang.org/p/m-TMRpmmnm)

  

The above code is perfectly legal in C language. But in the case of go, this wont work. i is of type int and j is of type float64. We are trying to add 2 numbers of different types which is not allowed. When you run the program, you will get `main.go:10: invalid operation: i + j (mismatched types int and float64)`

  

To fix the error, both _i_ and _j_ should be of the same type. Let's convert _j_ to int. _T(v) is the syntax to convert a value v to type T_

  

```go

package main

  

import (

"fmt"

)

  

func  main() {

i := 55  //int

j := 67.8  //float64

sum := i + int(j) //j is converted to int

fmt.Println(sum)

}

```

  

[Run in playground](https://play.golang.org/p/mweu3n3jMy)

  

Now when you run the above program, you can see `122` as the output.

  

The same is the case with assignment. Explicit type conversion is required to assign a variable of one type to another. This is explained in the following program.

  

```go

package main

  

import (

"fmt"

)

 
func  main() {

i := 10

var  j  float64 = float64(i) //this statement will not work without explicit conversion

fmt.Println("j", j)

}

  

```

  

[Run in playground](https://play.golang.org/p/Y2uSYYr46c)

  

In line no. 9, i is converted to float64 and then assigned to j. When you try to assign i to j without any type conversion, the compiler will throw an error.

## Constants

  

### Definition

The term constant is used in Go to denote fixed values such 5, -89, "I love Go", 67.89 and so on.

Consider the following code,

```go
var a int = 50  
var b string = "I love Go"  
```

**In the above code  _a_  and  _b_  are assigned to constants  _50_  and  _I love Go_  respectively**. The keyword  **const**  is used to denote constants such as  _50_  and  _I love Go_. Even though we do not explicitly use the keyword const anywhere in the above code, internally they are constants in Go.

Constants as the name indicate cannot be reassigned again to any other value and hence the below program will not work and it will fail with compilation error  **cannot assign to a**.

```go
package main

func main() {  
    const a = 55 //allowed
    a = 89 //reassignment not allowed
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
    var a = math.Sqrt(4)//allowed
    const b = math.Sqrt(4)//not allowed
}

```

[Run in playground](https://play.golang.org/p/dCON1LzCTw)

In the above program,  `a`  is a variable and hence it can be assigned to the result of the function  `math.Sqrt(4)`  

_b_  is a constant and the value of b needs to be know at compile time. The function  `math.Sqrt(4)`  will be evaluated only during run time and hence  `const b = math.Sqrt(4)`  throws  **error main.go:11: const initializer math.Sqrt(4) is not a constant.**

### String Constants

Any value enclosed between double quotes is a string constant in Go. For example strings like "Hello World" or "Sam" are all constants in Go.

What type does a string constant belong to? The answer is they are  **untyped**.

**A string constant like "Hello World" does not have any type**.

```go
const hello = "Hello World"  

```

In the above case we have assigned "Hello World" to a named constant  **hello**. Now does the constant  _hello_  have a type? The answer is No. The constant still doesn't have a type.

Go is a strongly typed language. All variables require a explicit type. Then how does the following program which assigns a variable  `name`  to an untyped constant  `Sam`  work?

```go
package main

import (  
    "fmt"
)

func main() {  
    var name = "Sam"
    fmt.Printf("type %T value %v", name, name)

}

```

[Run in playground](https://play.golang.org/p/xhYV4we_Jz)

**The answer is untyped constants have a default type associated with them and they supply it if and only if a line of code demands it. In the statement  `var name = "Sam"`,  `name`  needs a type and it gets it from the default type of the string constant "Sam" which is a  _string._**

Is there a way to create a typed constant? The answer is yes. The following code creates a typed constant.

```go
const typedhello string = "Hello World"  
```

**_typedhello_  in the above code is a constant of type string.**

Go is a strongly typed language. Mixing types during assignment is not allowed. Let's see what this means by the help of a program.

```go
package main

func main() {  
        var defaultName = "Sam" //allowed
        type myString string
        var customName myString = "Sam" //allowed
        customName = defaultName //not allowed

}

```

[Run in playground](https://play.golang.org/p/1Q-vudNn_9)

In the above code, we first create a variable  _defaultName_  and assign it to the constant  _Sam_.  **The default type of the constant Sam is string, so after the assignment defaultName is of type String.**

In the next line we create a new type myString which is an alias of string.

Then we create a variable customName of type myString and assign it to the constant  _Sam_. Since the constant  _Sam_  is untyped it can be assigned to any string variable. Hence this assignment is allowed and customName gets the type myString.

Now we have a variable defaultName of type string and another variable customName of type myString. Even though we know that myString is an alias of string, Go's strong typing policy disallows variables of one type to be assigned to another.  **Hence the assignment customName = defaultName is not allowed and the compiler throws an error  _main.go:7:20: cannot use defaultName (type string) as type myString in assignment_**

  
  
  

([code for: Basics Constants](https://github.com/naqvijafar91/go-workshops/tree/master/030-basics-constants))

  