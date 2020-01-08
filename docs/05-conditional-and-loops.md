# If else statement

**if**  is a conditional statement. The syntax of the if statement is

```go
if condition {  
}

```

If the  `condition`  is true, the lines of code between  `{`  and  `}`  is executed.

Unlike in other languages like C, the  `{ }`  are mandatory even if there is only one statement between the  `{ }`.

The if statement also has optional  `else if`  and  `else`  components.

```go
if condition {  
} else if condition {
} else {
}

```

There can be any number of  `else if`s. The condition is evaluated for truth from the top to bottom. Which ever  `if`  or  `else if`'s condition evaluates to true, the corresponding block of code is executed. If none of the conditions are true then  `else`  block is executed.


```go
package main

import (  
    "fmt"
)

func main() {  
    num := 10
    if num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  else {
        fmt.Println("the number is odd")
    }
}

```

[Run in playground](https://play.golang.org/p/vWfN8UqZUr)

The  `if num % 2 == 0`  statement checks whether the reminder of dividing a number by 2 is zero. It it is, then "the number is even" is printed else "the number is odd" is printed. In the above program  `the number is even`  is printed.

There is one more variant of  `if`  which includes a optional  `statement`  component which is executed before the condition is evaluated. Its syntax is

```go
if statement; condition {  
}

```

Lets rewrite the program which finds whether the number is even or odd using the above syntax.

```go
package main

import (  
    "fmt"
)

func main() {  
    if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println(num,"is even") 
    }  else {
        fmt.Println(num,"is odd")
    }
}

```

[Run in playground](https://play.golang.org/p/_X9q4MWr4s)

In the above program  `num`  is initialised in the  `if`  statement. One thing to be noted is that  `num`  is available only for access from inside the  `if`  and  `else`. i.e. the scope of  `num`  is limited to the  `if`  `else`  blocks. If we try to access  `num`  from outside the  `if`  or  `else`, the compiler will complain.

Lets write one more program which uses  `else if`.

```go
package main

import (  
    "fmt"
)

func main() {  
    num := 99
    if num <= 50 {
        fmt.Println("number is less than or equal to 50")
    } else if num >= 51 && num <= 100 {
        fmt.Println("number is between 51 and 100")
    } else {
        fmt.Println("number is greater than 100")
    }
}
```

[Run in playground](https://play.golang.org/p/Eji7vmb17Q)

In the above program  `else if num >= 51 && num <= 100`  is true and hence the program will output  `number is between 51 and 100`

  

Get the free Golang tools cheat sheet

  

### WAT: Go and semicolons

The  `else`  statement should start in the same line after the closing curly brace  `}`  of the if statement. If not the compiler will complain.

Let's understand this by means of a program.

```go
package main

import (  
    "fmt"
)

func main() {  
    num := 10
    if num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  
    else {
        fmt.Println("the number is odd")
    }
}

```

[Run in playground](https://play.golang.org/p/RYNqZZO2F9)

In the program above, the  `else`  statement does not start in the same line after the closing  `}`  of the if statement. Instead it starts in the next line. This is not allowed in Go. If you run this program, the compiler will output the error,

```
main.go:12:5: syntax error: unexpected else, expecting }  

```

**The reason is because of the way Go inserts semicolons automatically. You can read about the semicolon insertion rule here  [https://golang.org/ref/spec#Semicolons](https://golang.org/ref/spec#Semicolons).**

In the rules, it's specified that a semicolon will be inserted after  `}`, if that is the final token of the line. So a semicolon is automatically inserted after the if statement's  `}`.

So our program actually becomes

```go
if num%2 == 0 {  
      fmt.Println("the number is even") 
};  //semicolon inserted by Go
else {  
      fmt.Println("the number is odd")
}

```

after semicolon insertion. You can see the semicolon insertion in line no. 3 of the above snippet.

Since  `if{...} else {...}`  is one single statement, a semicolon should not be present in the middle of it. Hence there is a requirement to place the  `else`  in the same line after the closing  `}`.

I have rewritten the program by moving the else after the closing  `}`  of the if statement to prevent the automatic semicolon insertion.

```go
package main

import (  
    "fmt"
)

func main() {  
    num := 10
    if num%2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    } else {
        fmt.Println("the number is odd")
    }
}

```

[Run in playground](https://play.golang.org/p/hv_27vbIBC)

Now the compiler will be happy and so are we 😃.

A loop statement is used to execute a block of code repeatedly.

`for`  is the only loop available in Go. Go doesn't have while or do while loops which are present in other languages like C.

# Loops

### for loop syntax

```go
for initialisation; condition; post {  
}
```


All the three components namely initialisation, condition and post are optional in Go. Let's look at an example to understand for loop better.

### Example

 For loop to print all numbers from 1 to 10.

```go
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        fmt.Printf(" %d",i)
    }
}

```

[Run in playground](https://play.golang.org/p/mV6Zgcx2DY)

In the above program,  `i`  is initialised to 1. The conditional statement will checks if  `i <= 10`. If the condition is true, the value of i is printed, else the loop is terminated. The post statement increments i by 1 at the end of each iteration. Once  `i`  becomes greater than 10, the loop terminates.

The above program will print  `1 2 3 4 5 6 7 8 9 10`

The variables declared in a for loop are only available within the scope of the loop. Hence  `i`  cannot be accessed outside the body for loop, same as Java.

### break

The  `break`  statement is used to terminate the for loop abruptly before it finishes its normal execution and move the control to the line of code just after the for loop.

Let's write a program which prints numbers from 1 to 5 using break.

```go
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i > 5 {
            break //loop is terminated if i > 5
        }
        fmt.Printf("%d ", i)
    }
    fmt.Printf("\nline after for loop")
}

```

[Run in playground](https://play.golang.org/p/sujKy92f--)

In the above program, the value of i is checked during each iteration. If i is greater than 5 then  `break`  executes and the loop is terminated. The print statement just after the for loop is then executed. The above program will output,

```
1 2 3 4 5  
line after for loop  

```
 

### continue

The  `continue`  statement is used to skip the current iteration of the for loop. All code present in a for loop after the continue statement will not be executed for the current iteration. The loop will move on to the next iteration.

Oodd numbers from 1 to 10 using continue.

```go
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Printf("%d ", i)
    }
}

```

[Run in playground](https://play.golang.org/p/DRLN2ZHwVS)

In the above program the line  `if i%2 == 0`  checks if the reminder of dividing i by 2 is 0. If it is zero, then the number is even and  `continue`  statement is executed and the control moves to the next iteration of the loop. Hence the print statement after the continue will not be called and the loop proceeds to the next iteration. The output of the above program is  `1 3 5 7 9`

### Nested for loops

A  `for`  loop which has another  `for`  loop inside it is called a nested for loop.

Let's understand nested for loops by writing a program which prints the sequence below.

```
*
**
***
****
*****

```

The program below uses nested for loops to print the sequence. The variable  `n`  in line no. 8 stores the number of lines in the sequence. In our case it's  `5`. The outer for loop iterates  `i`  from  `0`  to  `4`  and the inner for loop iterates  `j`  from  `0`  to the current value of  `i`. The inner loop prints  `*`  for each iteration and the outer loop prints a new line at the end of each iteration. Run this program and you see the sequence printed as output.

```go
package main

import (  
    "fmt"
)

func main() {  
    n := 5
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

```

[Run in playground](https://play.golang.org/p/0rq8fWjVDLb)

### Labels

Labels can be used to break the outer loop from inside the inner for loop. Let's understand what I mean by using a simple example.

```go
package main

import (  
    "fmt"
)

func main() {  
    for i := 0; i < 3; i++ {
        for j := 1; j < 4; j++ {
            fmt.Printf("i = %d , j = %d\n", i, j)
        }

    }
}

```

[Run in playground](https://play.golang.org/p/BnCKho2x5hM)

The above program is self explanatory and it will print

```
i = 0 , j = 1  
i = 0 , j = 2  
i = 0 , j = 3  
i = 1 , j = 1  
i = 1 , j = 2  
i = 1 , j = 3  
i = 2 , j = 1  
i = 2 , j = 2  
i = 2 , j = 3  

```

Nothing special in this :)

What if we want to stop printing when  `i`  and  `j`  are equal. To do this we need to  `break`  from the outer  `for`  loop. Adding a  `break`  in the inner  `for`  loop when  `i`  and  `j`  are equal will only break from the inner  `for`  loop.

```go
package main

import (  
    "fmt"
)

func main() {  
    for i := 0; i < 3; i++ {
        for j := 1; j < 4; j++ {
            fmt.Printf("i = %d , j = %d\n", i, j)
            if i == j {
                break
            }
        }

    }
}

```

[Run in playground](https://play.golang.org/p/uMjbF8Ii41d)

In the program above, we have added a  `break`  inside the inner  `for`  loop when  `i`  and  `j`  are equal in line no. 10. This will  `break`  only from the inner for loop and the outer loop will continue. This program will print.

```
i = 0 , j = 1  
i = 0 , j = 2  
i = 0 , j = 3  
i = 1 , j = 1  
i = 2 , j = 1  
i = 2 , j = 2 
```

This is not the intended output. We need to stop printing when both  `i`  and  `j`  are equal  _i.e_  when they are equal to  `1`.

This is where labels come to our rescue. A label can be used to break from an outer loop. Let's rewrite the program above using labels.

```go
package main

import (  
    "fmt"
)

func main() {  
outer:  
    for i := 0; i < 3; i++ {
        for j := 1; j < 4; j++ {
            fmt.Printf("i = %d , j = %d\n", i, j)
            if i == j {
                break outer
            }
        }

    }
}

```

[Run in playground](https://play.golang.org/p/BI10Rmp_Z3y)

In the program above, we have added a label  `outer`  in line no. 8 on the outer for loop and in line no. 13 we break the outer for loop by specifying the label. This program will stop printing when both  `i`  and  `j`  are equal. This program will output

```
i = 0 , j = 1  
i = 0 , j = 2  
i = 0 , j = 3  
i = 1 , j = 1  
```

 
### More examples

Let's write some more code to cover all variations of  `for`  loop.

The program below prints all even numbers from 0 to 10.

```go
package main

import (  
    "fmt"
)

func main() {  
    i := 0
    for ;i <= 10; { // initialisation and post are omitted
        fmt.Printf("%d ", i)
        i += 2
    }
}

```

[Run in playground](https://play.golang.org/p/PNXliGINku)



```go
package main

import (  
    "fmt"
)

func main() {  
    i := 0
    // While loop in Go
    for i <= 10 { //semicolons are ommitted and only condition is present
        fmt.Printf("%d ", i)
        i += 2
    }
}

```

[Run in playground](https://play.golang.org/p/UYiz-Wtnoj)

### infinite loop

The syntax for creating an infinite loop is,

```go
for {  
}
```

  
  
The following program will keep printing  `Hello World`  continuously without terminating.

```go
package main

import "fmt"

func main() {  
    for {
        fmt.Println("Hello World")
    }
}

```

If you try to run the above program in the  [go playground](https://play.golang.org/p/kYQZw1AWT4)  you will get error "process took too long". Please try running it in your local system to print "Hello World" infinitely.

There is one more construct  **range**  which can be used in  `for`  loops for array manipulation.

# Switch Statement



```go
package main

import (  
    "fmt"
)

func main() {  
    finger := 4
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")

    }
}

```

[Run in playground](https://play.golang.org/p/q4kjm2kpVe)

Duplicate cases with the same constant value are not allowed. If you try to run the program below, the compiler will complain  `main.go:18:2: duplicate case 4 in switch previous case at tmp/sandbox887814166/main.go:16:7`

```go
package main

import (  
    "fmt"
)

func main() {  
    finger := 4
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 4://duplicate case not allowed
        fmt.Println("Another Ring")
    case 5:
        fmt.Println("Pinky")

    }
}

```

[Run in playground](https://play.golang.org/p/SfXdChWdoN)

### Default case

```go
package main

import (  
    "fmt"
)

func main() {  
	// Switch can also contain an optional statement,just like if
    switch finger := 8; finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    default: //default case
        fmt.Println("incorrect finger number")
    }
}

```

[Run in playground](https://play.golang.org/p/Fq7U7SkHe1)


**The scope of  `finger`  in this case is limited to the switch block.**  

### Multiple expressions in case

It is possible to include multiple expressions in a case by separating them with comma.

```go
package main

import (  
    "fmt"
)

func main() {  
    letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u": //multiple expressions in case
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
    }
}

```

[Run in playground](https://play.golang.org/p/Zs9Ek5SInh)

The above program checks whether  `letter`  is a vowel of not. The line  `case "a", "e", "i", "o", "u":`  matches all the vowels. This program outputs  `vowel`.

### Expressionless switch

The expression in a switch is optional and it can be omitted. If the expression is omitted, the switch is considered to be  `switch true`  and each of the  `case`  expression is evaluated for truth and the corresponding block of code is executed.

```go
package main

import (  
    "fmt"
)

func main() {  
    num := 75
    switch { // expression is omitted and it is equal to switch true
    // This type of switch can be considered an alternative to multiple if else 
    // clauses
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100")
    case num >= 101:
        fmt.Println("num is greater than 100")
    }

}

```
[Run in playground](https://play.golang.org/p/mMJ8EryKbN)

### Fallthrough Statement

In Go the control comes out of the switch statement immediately after a case is executed. A  `fallthrough`  statement is used to transfer control to the first statement of the case that is present immediately after the case which has been executed.

Lets write an program to understand fallthrough. Our program will check whether the input number is lesser than 50, 100 or 200. For instance if we input 75, the progam will print that 75 is less than both 100 and 200. We will achieve this output using fallthrough.

```go
package main

import (  
    "fmt"
)

func number() int {  
        num := 15 * 5 
        return num
}

func main() {

    switch num := number(); { //num is not a constant
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough // It should be the last statement for each case
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
    }

}

```

[Run in playground](https://play.golang.org/p/svGJAiswQj)


```
75 is lesser than 100  
75 is lesser than 200  

```

**`fallthrough`  should be the last statement in a  `case`. If it present somewhere in the middle, the compiler will throw error  `fallthrough statement out of place`**
