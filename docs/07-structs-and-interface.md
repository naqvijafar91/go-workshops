
# Structs and Interfaces

Although it would be possible for us to write programs only using Go's built-in data types, at some point it would become quite tedious. Consider a program that interacts with shapes:
```go
package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 – x1
  b := y2 – y1
  return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
  l := distance(x1, y1, x1, y2)
  w := distance(x1, y1, x2, y1)
  return l * w
}
func circleArea(x, y, r float64) float64 {
  return math.Pi * r*r
}
func main() {
  var rx1, ry1 float64 = 0, 0
  var rx2, ry2 float64 = 10, 10
  var cx, cy, cr float64 = 0, 0, 5

  fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
  fmt.Println(circleArea(cx, cy, cr))
}
```

Keeping track of all the coordinates makes it difficult to see what the program is doing and will likely lead to mistakes.

## Structs

An easy way to make this program better is to use a struct. A struct is a type which contains named fields. For example we could represent a Circle like this:
```go
type Circle struct {
  x float64
  y float64
  r float64
}
```

The  `type`  keyword introduces a new type. It's followed by the name of the type (`Circle`), the keyword  `struct`  to indicate that we are defining a  `struct`  type and a list of fields inside of curly braces. Each field has a name and a type. Like with functions we can collapse fields that have the same type:

```go
type Circle struct {
  x, y, r float64
}
```
### Initialization

We can create an instance of our new Circle type in a variety of ways:
```go
var c Circle
**```**
Like with other data types, this will create a local Circle variable that is by default set to zero. For a  `struct`  zero means each of the fields is set to their corresponding zero value (`0`  for  `int`s,  `0.0`  for  `float`s,  `""`  for  `string`s,  `nil`  for pointers, …) We can also use the new function:
```go
c := new(Circle)
```
This allocates memory for all the fields, sets each of them to their zero value and returns a pointer. (`*Circle`) More often we want to give each of the fields a value. We can do this in two ways. Like this:
```go
c := Circle{x: 0, y: 0, r: 5}
```
Or we can leave off the field names if we know the order they were defined:
```go
c := Circle{0, 0, 5}
```
### Fields

We can access fields using the  `.`  operator:
```go
fmt.Println(c.x, c.y, c.r)
c.x = 10
c.y = 5
```
Let's modify the  `circleArea`  function so that it uses a  `Circle`:
```go
func circleArea(c Circle) float64 {
  return math.Pi * c.r*c.r
}
```
In main we have:
```go
c := Circle{0, 0, 5}
fmt.Println(circleArea(c))
```
One thing to remember is that arguments are always copied in Go. If we attempted to modify one of the fields inside of the  `circleArea`  function, it would not modify the original variable. Because of this we would typically write the function like this:
```go
func circleArea(c *Circle) float64 {
  return math.Pi * c.r*c.r
}
```
And change main:
```go
c := Circle{0, 0, 5}
fmt.Println(circleArea(&c))
```
## Methods

Although this is better than the first version of this code, we can improve it significantly by using a special type of function known as a method:
```go
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}
```

In between the keyword  `func`  and the name of the function we've added a “receiver”. The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the  `.`  operator:
```go
fmt.Println(c.area())
```
This is much easier to read, we no longer need the  `&`  operator (Go automatically knows to pass a pointer to the circle for this method) and because this function can only be used with  `Circle`s we can rename the function to just  `area`.

Let's do the same thing for the rectangle:
```go
type Rectangle struct {
  x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}
```

`main`  has:
```go
r := Rectangle{0, 0, 10, 10}
fmt.Println(r.area())
```
### Embedded Types

A struct's fields usually represent the has-a relationship. For example a  `Circle`  has a  `radius`. Suppose we had a person struct:
```go
type Person struct {
  Name string
}
func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}
```
And we wanted to create a new  `Android`  struct. We could do this:
```go
type Android struct {
  Person Person
  Model string
}
```
This would work, but we would rather say an Android is a Person, rather than an Android has a Person. Go supports relationships like this by using an embedded type. Also known as anonymous fields, embedded types look like this:
```go
type Android struct {
  Person
  Model string
}
```
We use the type (`Person`) and don't give it a name. When defined this way the  `Person`  struct can be accessed using the type name:
```go
a := new(Android)
a.Person.Talk()
```
But we can also call any  `Person`  methods directly on the  `Android`:
```go
a := new(Android)
a.Talk()
```
The is-a relationship works this way intuitively: People can talk, an android is a person, therefore an android can talk.

## Interfaces

You may have noticed that we were able to name the  `Rectangle`'s  `area`  method the same thing as the  `Circle`'s  `area`  method. This was no accident. In both real life and in programming, relationships like these are commonplace. Go has a way of making these accidental similarities explicit through a type known as an Interface. Here is an example of a  `Shape`  interface:
```go
type Shape interface {
  area() float64
}
```
Like a struct an interface is created using the  `type`  keyword, followed by a name and the keyword  `interface`. But instead of defining fields, we define a “method set”. A method set is a list of methods that a type must have in order to “implement” the interface.

In our case both  `Rectangle`  and  `Circle`  have area methods which return  `float64`s so both types implement the  `Shape`  interface. By itself this wouldn't be particularly useful, but we can use interface types as arguments to functions:
```go
func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}
```
We would call this function like this:
```go
fmt.Println(totalArea(&c, &r))
```
Interfaces can also be used as fields:
```go
type MultiShape struct {
  shapes []Shape
}
```
We can even turn  `MultiShape`  itself into a  `Shape`  by giving it an area method:
```go
func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}
```
Now a  `MultiShape`  can contain  `Circle`s,  `Rectangle`s or even other  `MultiShape`s.