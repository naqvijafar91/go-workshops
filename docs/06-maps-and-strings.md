# Maps


### What is a map?

A map is a builtin type in Go which associates a value to a key. The value can be retrieved using the corresponding key.

### How to create a map?

A map can be created by passing the type of key and value to the  `make`  function.  `make(map[type of key]type of value)`  is the syntax to create a map.

```go
personSalary := make(map[string]int)  
```
The above line of code creates a map named  `personSalary`  which has  `string`  keys and  `int`  values.

The zero value of a map is  `nil`. If you try to add items to the nil map, a run-time panic will occur. Hence the map has to be initialized using  `make`  function.

```go
package main

import (  
    "fmt"
)

func main() {  
    var personSalary map[string]int
    if personSalary == nil {
        fmt.Println("map is nil. Going to make one.")
        personSalary = make(map[string]int)
    }
}

```

[Run in playground](https://play.golang.org/p/IwJnXMGc1M)

The program will output  `map is nil. Going to make one.`

### Adding items to a map



```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := make(map[string]int)
    // We do not have put and get 
    personSalary["steve"] = 12000
    personSalary["jamie"] = 15000
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)
}
```

[Run in playground](https://play.golang.org/p/V1lnQ4Igw1)

The above program outputs,  `personSalary map contents: map[steve:12000 jamie:15000 mike:9000]`

It is also possible to initialize a map during the declaration itself.

```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int {
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)
}

```

[Run in playground](https://play.golang.org/p/nlH_ADhO9f)


```
personSalary map contents: map[steve:12000 jamie:15000 mike:9000]  
```

Only string types don't need to be keyed. All comparable types such as boolean, integer, float, complex, string, ... can also be keys. If you would like to know more about comparable types, please visit  [http://golang.org/ref/spec#Comparison_operators](http://golang.org/ref/spec#Comparison_operators)
  

### Finding if a key is present


What if we want to know whether a  `key`  is present in a map or not.

```go
value, ok := map[key]  
```

```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    newEmp := "joe"
    value, ok := personSalary[newEmp]
    if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
    } else {
        fmt.Println(newEmp,"not found")
    }

}

```

[Run in playground](https://play.golang.org/p/q8fL6MeVZs)


```
joe not found  
```

The  `range`  form of the  `for`  loop is used to iterate over all elements of a map.

```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("All items of a map")
    for key, value := range personSalary {
        fmt.Printf("personSalary[%s] = %d\n", key, value)
    }
}
```

[Run in playground](https://play.golang.org/p/gq9ZOKsI9b)

The above program outputs,

```
All items of a map  
personSalary[mike] = 9000  
personSalary[steve] = 12000  
personSalary[jamie] = 15000  
```

**One important fact is that the order of the retrieval of values from a map when using  `for range`  is not guaranteed to be the same for each execution of the program.**

### Deleting items

_[delete(map, key)](https://golang.org/pkg/builtin/#delete)_  is the syntax to delete  `key`  from a  `map`. The delete function does not return any value.

```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("map before deletion", personSalary)
    delete(personSalary, "steve")
    fmt.Println("map after deletion", personSalary)

}

```

[Run in playground](https://play.golang.org/p/nroJzeF-a7)


```
map before deletion map[steve:12000 jamie:15000 mike:9000]  
map after deletion map[mike:9000 jamie:15000]  
```

### Length of the map

The length of the map can be determined using the  [len](https://golang.org/pkg/builtin/#len)  function.

```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("length is", len(personSalary))

}

```

[Run in playground](https://play.golang.org/p/8O1WnKUuDP)

 `length is 3`

### Maps are reference types

Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same internal data structure. Hence changes made in one will reflect in the other.

```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("Original person salary", personSalary)
    newPersonSalary := personSalary
    newPersonSalary["mike"] = 18000
    fmt.Println("Person salary changed", personSalary)

}

```

[Run in playground](https://play.golang.org/p/OGFl3addq1)



```
Original person salary map[steve:12000 jamie:15000 mike:9000]  
Person salary changed map[steve:12000 jamie:15000 mike:18000]  

```

### Maps equality

Maps can't be compared using the  `==`  operator. The  `==`  can be only used to check if a map is  `nil`.

```go
package main

func main() {  
    map1 := map[string]int{
        "one": 1,
        "two": 2,
    }
    map2 := map1
    if map1 == map2 {
    }
}
```

[Run in playground](https://play.golang.org/p/MALqDyWkcT)

The above program will throw compilation error  **invalid operation: map1 == map2 (map can only be compared to nil)**.

One way to check whether two maps are equal is to compare each one's elements one by one.

# Strings

Strings deserve a special mention in Go as they are different in implementation when compared to other languages.

### What is a String?

**A string in Go is a slice of bytes**. 




The above program will output  `Hello World`.

Strings in Go are Unicode compliant and are UTF-8 Encoded.

### Accessing individual bytes of a string

Since a string is a slice of bytes, it's possible to access each byte of a string.

```go
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}


func printChars(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}

func main() {  
    name := "Hello World"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
}

```
[Run in playground](https://play.golang.org/p/Jss0HG1q80)

```
48 65 6c 6c 6f 20 57 6f 72 6c 64  
H e l l o   W o r l d  

```

Although the above program looks like a legitimate way to access the individual characters of a string, this has a serious bug. 

```go
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}

func main() {  
    name := "Hello World"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
    fmt.Printf("\n")
    name = "Señor"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
}

```

[Run in playground](https://play.golang.org/p/UQOVvRVaFH)

The output of the above program is

```
48 65 6c 6c 6f 20 57 6f 72 6c 64  
H e l l o   W o r l d  
53 65 c3 b1 6f 72  
S e Ã ± o r  
```

The reason is that the Unicode code point of  `ñ`  is  `U+00F1`  and it's [UTF-8 encoding](https://mothereff.in/utf-8#%C3%B1)  occupies 2 bytes. We are trying to print characters assuming that each code point will be one byte long which is wrong. 

### rune

A rune is a builtin type in Go and it's the alias of int32. the rune represents a Unicode code point in Go. It does not matter how many bytes the code point occupies, it can be represented by a rune. 

```go
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
	// Convert to a runes array
    runes := []rune(s)
    for i:= 0; i < len(runes); i++ {
        fmt.Printf("%c ",runes[i])
    }
}

func main() {  
    name := "Hello World"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
    fmt.Printf("\n\n")
    name = "Señor"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
}

```

[Run in playground](https://play.golang.org/p/t4z-f8I-ih)

```
48 65 6c 6c 6f 20 57 6f 72 6c 64  
H e l l o   W o r l d 

53 65 c3 b1 6f 72  
S e ñ o r  
```


### Strings are immutable

Strings are immutable in Go. Once a string is created it's not possible to change it.

```go
package main

import (  
    "fmt"
)

func mutate(s string)string {  
    s[0] = 'a'//any valid unicode character within single quote is a rune 
    return s
}
func main() {  
    h := "hello"
    fmt.Println(mutate(h))
}

```

[Run in playground](https://play.golang.org/p/bv4SlSd_hp)

Program throws error  **main.go:8: cannot assign to s[0]**

To workaround this string immutability, strings are converted to a slice of runes. 

```go
package main

import (  
    "fmt"
)

func mutate(s []rune) string {  
    s[0] = 'a' 
    return string(s)
}
func main() {  
    h := "hello"
    fmt.Println(mutate([]rune(h)))
}

```

[Run in playground](https://play.golang.org/p/GL1cm17IP1)

This program outputs  `aello`