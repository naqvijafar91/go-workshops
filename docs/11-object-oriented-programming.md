
# Object Oriented programming 

Go doesn’t have inheritance – instead composition, embed­ding and inter­faces support code reuse and poly­morphism.

![](https://yourbasic.org/golang/improvement.jpg)


## Object-oriented programming with inheritance

Inheritance in traditional object-oriented languages offers three features in one. When a `Dog`  inherits from an `Animal`

1.  the  `Dog`  class reuses code from the  `Animal` class,
2.  a variable  `x`  of type  `Animal`  can refer to either a  `Dog`  or an  `Animal`,
3.  `x.Eat()`  will choose an  `Eat`  method based on what type of object  `x`  refers to.

In object-oriented lingo, these features are known as  **code reuse**,  **poly­mor­phism**  and  **dynamic dispatch**.

All of these are available in Go, using separate constructs:

-   **composition**  and  **embedding**  provide code reuse,
-   interfaces  take care of polymorphism and dynamic dispatch.

## Code reuse by composition

> Don't worry about type hierarchies when starting a new Go project –  
> it's easy to introduce polymorphism and dynamic dispatch later on.

If a  `Dog`  needs some or all of the functionality of an  `Animal`, simply use  **composition**.

```
type Animal struct {
	// …
}

type Dog struct {
	beast Animal
	// …
}
```

This gives you full freedom to use the  `Animal`  part of your  `Dog`  as needed. Yes, it’s that simple.

## Code reuse by embedding

If the  `Dog`  class inherits  **the exact behavior**  of an  `Animal`, this approach can result in some tedious coding.

```
type Animal struct {
	// …
}

func (a *Animal) Eat()   { … }
func (a *Animal) Sleep() { … }
func (a *Animal) Breed() { … }

type Dog struct {
	beast Animal
	// …
}

func (a *Dog) Eat()   { a.beast.Eat() }
func (a *Dog) Sleep() { a.beast.Sleep() }
func (a *Dog) Breed() { a.beast.Breed() }
```

This code pattern is known as  **delegation**.

Go uses  **embedding**  for situations like this. The declaration of the  `Dog`  struct and it’s three methods can be reduced to:

```
type Dog struct {
	Animal
	// …
}
```

## Polymorphism and dynamic dispatch with interfaces

> Keep your interfaces short, and introduce them only when needed.

Further down the road your project might have grown to include more animals. At this point you can introduce polymorphism and dynamic dispatch using  **[interfaces](https://yourbasic.org/golang/interfaces-explained/)**.

If you need to put all your pets to sleep, you can define a  `Sleeper`  interface.

```
type Sleeper interface {
	Sleep()
}

func main() {
	pets := []Sleeper{new(Cat), new(Dog)}
	for _, x := range pets {
		x.Sleep()
	}
}
```

No explicit declaration is required by the  `Cat`  and  `Dog`  types. Any type that provides the methods named in an inter­face may be treated as an imple­mentation.

# Constructors deconstructed

Go doesn't have explicit constructors. The idiomatic way to set up new data structures is to use proper  **zero values**  coupled with  **factory**  functions.

## Zero value

Try to make the default  [zero value](https://yourbasic.org/golang/default-zero-value/)  useful and document its behavior. Sometimes this is all that’s needed.

```
// A StopWatch is a simple clock utility.
// Its zero value is an idle clock with 0 total time.
type StopWatch struct {
    start   time.Time
    total   time.Duration
    running bool
}

var clock StopWatch // Ready to use, no initialization needed.
```

-   `StopWatch`  takes advantage of the useful zero values of  `time.Time`,  `time.Duration`  and  `bool`.
-   In turn, users of  `StopWatch`  can benefit from  _its_  useful zero value.

## Factory

If the zero value doesn’t suffice, use factory functions named  `NewFoo`  or just `New`.

```
scanner := bufio.NewScanner(os.Stdin)
err := errors.New("Houston, we have a problem")
```
