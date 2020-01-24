## Functions

### What is a function?

A function is a block of code that performs a specific task. A function takes an input, performs some calculations on the input and generates an output.

### Function declaration

The general syntax for declaring a function in go is

```go
func functionname(parametername type) returntype {  
 //function body
}
```

The function declaration starts with a  `func`  keyword followed by the  `functionname`. The parameters are specified between  `(`  and  `)`  followed by the  `returntype`  of the function. The syntax for specifying a parameter is parameter name followed by the type. Any number of parameters can be specified like  `(parameter1 type, parameter2 type)`. Then there is a block of code between  `{`  and  `}`  which is the body of the function.

The parameters and return types are optional in a function. Hence the following syntax is also a valid function declaration.

```go
func functionname() {  
}

```

### Sample Function

Let's write a function which takes the price of a single product and number of products as input parameters and calculates the total price by multiplying these two values and returns the output.

```go
func calculateBill(price int, no int) int {  
    var totalPrice = price * no
    return totalPrice
}

```

The above function has two input parameters  `price`  and  `no`  of type int and it returns the  `totalPrice`  which is the product of price and no. The return value is also of type int.

**If consecutive parameters are of the same type, we can avoid writing the type each time and it is enough to be written once at the end._ie_  `price int, no int`  can be written as  `price, no int`**. The above function can hence be rewritten as,

```go
func calculateBill(price, no int) int {  
    var totalPrice = price * no
    return totalPrice
}

```

Now that we have a function ready, let's call it from somewhere in the code. The syntax for calling a function is  `functionname(parameters)`. The above function can be called using the code.

```go
calculateBill(10, 5)  

```

Here is the complete program which uses the above function and prints the total price.

```go
package main

import (  
    "fmt"
)

func calculateBill(price, no int) int {  
    var totalPrice = price * no
    return totalPrice
}
func main() {  
    price, no := 90, 6
    totalPrice := calculateBill(price, no)
    fmt.Println("Total price is", totalPrice)
}

```

[Run in playground](https://play.golang.org/p/YJlW3g-VZH)

The above program will print

```
Total price is 540  

```

### Multiple return values

It is possible to return multiple values from a function. Let's write a function  `rectProps`  which takes the length and width of a rectangle and returns both the area and perimeter of the rectangle. The area of the rectangle is the product of length and width and the perimeter is twice the sum of the length and width.

```go
package main

import (  
    "fmt"
)

func rectProps(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

func main() {  
     area, perimeter := rectProps(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f", area, perimeter) 
}

```

[Run in playground](https://play.golang.org/p/qAftE_yke_)

If a function returns multiple return values then they must be specified between  `(`  and  `)`.  `func rectProps(length, width float64)(float64, float64)`  has two float64 parameters  `length and width`  and also returns two float64 values. The above program outputs

```
Area 60.480000 Perimeter 32.800000  

```

### Named return values

It is possible to return named values from a function. If a return value is named, it can be considered as being declared as a variable in the first line of the function.

The above rectProps can be rewritten using named return values as

```go
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}

```

**area**  and  **perimeter**  are the named return values in the above function. Note that the return statement in the function does not explicitly return any value. Since  `area`  and  `perimeter`  are specified in the function declaration as return values, they are automatically returned from the function when a return statement is encountered.

### Blank Identifier

**_**  is known as the blank identifier in Go. It can be used in place of any value of any type. Let's see what's the use of this blank identifier.

The  `rectProps`  function returns the area and perimeter of the rectangle. What if we only need the  `area`  and want to discard the  `perimeter`. This is where  `_`  is of use.

The program below uses only the  `area`  returned from the  `rectProps`  function.

```go
package main

import (  
    "fmt"
)

func rectProps(length, width float64) (float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
func main() {  
    area, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ", area)
}

```

[Run in playground](https://play.golang.org/p/IkugSH1jIt)

In line no. 13 we use only the area and the  `_`  identifier is used to discard the parameter.

  
  

([code for: Basics Functions](https://github.com/naqvijafar91/go-workshops/tree/master/040-basics-functions))

## Packages 

### What are packages and why are they used?

**Packages are used to organize go source code for better reusability and readability.**  Packages offer compartmentalization of code and hence it becomes easy to maintain go applications.

We will step by step create an application that calculates the area and diagonal of a rectangle.

### main function and main package

Every executable go application must contain the main function. This function is the entry point for execution. The main function should reside in the main package.

**The line of code to specify that a particular source file belongs to a package is  `package package-name`. This should be the first line of every go source file.**


```go
//geometry.go
package main 

import "fmt"

func main() {  
    fmt.Println("Geometrical shape properties")
}

```

The line of code  `package main`  specifies that this file belongs to the main package. The  `import "package-name"`  statement is used to import an existing package. In this case, we import the  `fmt`  package which contains the Println method. Then there is a main function which prints  `Geometrical shape properties`


### Creating a custom package

We will structure the code in such a way that all functionalities related to a rectangle are in the `rectangle`  package.

**Source files belonging to a package should be placed in separate folders of their own. It is a convention in Go to name this folder with the same name as the package.**

```go
//rectprops.go
package rectangle

import "math"

func Area(len, wid float64) float64 {  
    area := len * wid
    return area
}

func Diagonal(len, wid float64) float64 {  
    diagonal := math.Sqrt((len * len) + (wid * wid))
    return diagonal
}

```

In the above code, we have created two functions that calculate  `Area`  and  `Diagonal`. The area of the rectangle is the product of the length and width. The diagonal of the rectangle is the square root of the sum of squares of the length and width. The  `Sqrt`  function in the  `math`  package is used to calculate the square root.

_Note that the function names  **Area**  and  **Diagonal** start with caps. This is essential and we will explain shortly why this is needed._

### Importing a custom package

To use a custom package we must first import it.  `import path`  is the syntax to import a custom package. We must specify the path to the custom package concerning the  `src`  folder inside the workspace. 

The line  `import "geometry/rectangle"`  will import the rectangle package.

Importing custom package in geometry.go

```go
//geometry.go
package main 

import (  
    "fmt"
    "geometry/rectangle" //importing custom package
)

func main() {  
    var rectLen, rectWidth float64 = 6, 7
    fmt.Println("Geometrical shape properties")
        /*Area function of rectangle package used
        */
    fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
        /*Diagonal function of rectangle package used
        */
    fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
}

```

The above code imports the  `rectangle`  package and uses the  `Area`  and  `Diagonal`  function of it to find the area and diagonal of the rectangle. The  `%.2f`  format specifier in Printf is to truncate the floating-point to two decimal places. The output of the application is

```
Geometrical shape properties  
area of rectangle 42.00  
diagonal of the rectangle 9.22  

```

### Exported Names

We capitalised the functions  `Area`  and  `Diagonal`  in the rectangle package. This has a special meaning in Go.

 **Any variable or function which starts with a capital letter are exported names in go. Only exported functions and variables can be accessed from other packages.** 

In this case, we need to access  `Area`  and  `Diagonal`  functions from the main package. Hence they are capitalized.

If the function names are changed from  `Area(len, wid float64)`  to  `area(len, wid float64)`  in  `rectprops.go`  and from  `rectangle.Area(rectLen, rectWidth)`  to  `rectangle.area(rectLen, rectWidth)`  in  `geometry.go`  and if the program is run, the compiler will throw error  `geometry.go:11: cannot refer to unexported name rectangle.area`. Hence if you want to access a function outside of a package, it should be capitalised.

### init function

Every package can contain an `init`  function. The init function should not have any return type and should not have any parameters. The init function cannot be called explicitly in our source code. The init function looks like below

```
func init() {  
}

```

The init function can be used to perform initialisation tasks and can also be used to verify the correctness of the program before the execution starts.

The order of initialisation of a package is as follows

1.  Package level variables are initialised first
2.  init function is called next. A package can have multiple init functions (either in a single file or distributed across multiple files) and they are called in the order in which they are presented to the compiler.

If a package imports other packages, the imported packages are initialised first.

**A package will be initialised only once even if it is imported from multiple packages.**

  `rectprops.go`  with init function:

```go
//rectprops.go
package rectangle

import "math"  
import "fmt"

/*
 * init function added
 */
func init() {  
    fmt.Println("rectangle package initialized")
}
func Area(len, wid float64) float64 {  
    area := len * wid
    return area
}

func Diagonal(len, wid float64) float64 {  
    diagonal := math.Sqrt((len * len) + (wid * wid))
    return diagonal
}

```

We have added a simple init function which just prints  `rectangle package initialised`

### Use of the blank identifier

It is illegal in Go to import a package and not to use it anywhere in the code. The compiler will complain if you do so. The reason for this is to avoid bloating of unused packages which will significantly increase the compilation time. Replace the code in  `geometry.go`  with the following,

```go
//geometry.go
package main 

import (   

     "geometry/rectangle" //importing custom package

)
func main() {

}

```

The above program will throw error  `geometry.go:6: imported and not used: "geometry/rectangle"`

The  `_`  blank identifier saves us in those situations.

The error in the above program can be silenced by the following code,

```go
package main

import (  
    "geometry/rectangle" 
)

var _ = rectangle.Area //error silencer

func main() {

}

```

The line  `var _ = rectangle.Area`  silences the error. 

**We should keep track of these kinds of error silencers and remove them including the imported package at the end of application development if the package is not used. Hence it is recommended to write error silencers in the package level just after the import statement.**

Sometimes we need to import a package just to make sure the initialization takes place even though we do not need to use any function or variable from the package. For example, we might need to ensure that the init function of the rectangle package is called even though we do not use that package anywhere in our code. The _ blank identifier can be used in this case too as shown below.

```go
package main 

import (   

     _ "geometry/rectangle" 

)
func main() {

}

```

  

([code for: Basics Init](https://github.com/naqvijafar91/go-workshops/tree/master/041-basics-init))

  

## Closures aka anonymous functions

  

Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

  

Functions in Go are first-class citizens so it can be:

- passed as parameters

- created as types

- assigned as values to variables

- called anonymously

   

([code for: Basics Closures](https://github.com/naqvijafar91/go-workshops/tree/master/042-basics-closures))
