<h1>Golang 19/11/2024</h1>

<h2>Pointers</h2>
- A Pointer holds the memory address of a value
- Pointers show up two main ways, in types and for pointer operations
  - as a type var myPointer *int
  - for dereferencing fmt.Println(*myPointer)
  - & will generate a pointer from a value
  - x := 7 
  - &x will generate where the memory address where 7 lives.


```
func main() {
    var myPointer *int
    x := 7
    myPointer = &x
    fmt.Println(myPointer)
    fmt.Println(&myPointer)
    fmt.Println(*myPointer)
}
```

- everything in Go is pass by value
- when a function is called a copy is made of all arguments being passed in
- However some data structures, like slices are just references to underlying structures
- The reference is copied but the underlying sctructure remains the same

- Value types: int, float, bool, struct, string
- Reference Types: slices, maps, channels, pointers, functions
- If you want to mutate a value type then you need a pass pointer to it
- Changes to reference types will mutate the original
https://go.dev/playp/p/6cRvSCVxgh-

Pointers

```
func main() {
    name := "Auden"
    namePointer := &name
    fmt.Println(&namePointer)
    printPointer(namePointer)
}

func printPointer(namePointer *string){
    fmt.Println(&namePointer)
    fmt.Println(*namePointer)
    fmt.Println(*&namePointer)
    fmt.Println(namePointer)
}
```

<h2>Loop Variables</h2>

<b>Now we know a little about pointers, a quick not about Go pre- and post-V1.22</b>
- Go 1.22 introduced a breaking change to loop scoping
- Changed loop vars to have a per-iteration instead of per-loop scope

<h2>Receivers</h2>

Go isn't OOP but it does have methods
- in GO they are called receivers or receiver functions
- can be declaredon any type
- Functions with an extra field, the receiver argument
- Must be declared on types in the same package

```
func (c Cat) pet(intensity int) {
    petService(c, intensity)
}

felix.pet(8)
```

<h2> Pointer Receivers </h2>

- it's common to have pointer receivers-receivers who take a pointer to a value type so they can modify it
- as a convenience, Go will allow you to pass either a value or pointer to a receiver and covert for you. Works for value receivers, too
- All methods for a given type should uniformly take either value or pointer receivers

```
func (p *Cat) feed(amount int){
    foodService(p, amount)
}

catPointer := &felix
catPointer.feed(3)

felix.feed(1)
```

<h2>Maps</h2>
Collection of  key value pairs
- All keys must be of the same type and unique
- All values must be of the same type (can be different than key)
- Missing values return the zero type of the value
- Indexable, iterable
- Reference type
- Go has no Set type, you can use a map with bool values for many of the same things

Check membership with value, present := m["key"]
<i>whether or not something was actually there</i>

```
m := make(map[string]int)
m["count"] = 5
v, ok := m["count"]
fmt.Println(v, ok)
delete(m, "count")
fmt.Println(m["count"])
```

<h2>Structs</h2>

Most similar to an object (class) in other languages
- Values can be of any type, but must be declared at compile time
- Errors if you try to access something that is not present
- More common than maps
- Value type in Go - a copy is made when passed

```
type contactInfo struct {
    email string
    zipcode int
}
type person struct {
    name string
    contact contactInfo
}
```

- Since they are value type, pointers to structs are very common. 
- As a convenience, we can access struct fields without dereferencing.

```
andrew := person{"Andrew", contactInfo{"a@a.com", 11111}}
fmt.Println(andrew.name)
p := &andrew
fmt.Println(p.name)
```

<h2>Interfaces</h2>

- Must specify types of everything in Go
- Makes code reuse hard, as you would have to copy/paste and change every type
- Interfaces associate different kinds of data together based on shared receiver functions (methods)
- Anything that matches ainterface can be passed in where that interface is specified

```
type interface animal() {
    speak()
}
```

- A set of method signatures
- Method Names
- Parameter types
- Return value types

```
type interface animal() {
    speak()
    feed(int, string) int
}
```

- This makes no guarantess about actual behaviour or correctness

```
type interface animal () {
    feed(int, string) int
}
```

```
func (a animal) feed (amt int, food string) int {
    for const i := 0 i < amt; i++ {
        lightOnFire(food)
        thrownAtChildren(food)
    }
    return -99
}
```

- Interfaces are implicit-connection  between different types is never directly state anywhere
- Satisfying an interface is just based on what implements the interface, in the case of animal anything that has a speak method
- Go dev tools can show you what satisfies an interface

```
type interface animal () {
    speak()
}
```

- We can now use the interface as a type, giving us polymorphism 

```
func printGreeting(a animal) {
    fmt.Prinln(a.speak())
}
```

<h2> Interface Design </h2>

- Writing your own interfaces well is deceptively tricky
- A common error made by programmers coming from high level languages is interface pollution - making too many interfaces
- "Interfaces should be discovered, not created" - Rob Pike
- General advise - your functions should accept interfaces and return structs.

The bigger the interface, the weaker the abstraction

Consider a rectangle here: 
- Width: Size of interface - price to inderstand it
- Heigh: Value gained from the abstraction - what you get to do once you understant it

---------------
|             |
|             |
| Abstraction | 
|             |
|             |
---------------

Best Case: 
- Small and easy to understand, provides little value
_____
|   |
|   |
|   |
|   |
|   |
|   |
|___|

- Worst case - Large and hard to understand, provides little value
_______________________________
|                             |
|_____________________________|

- you might come from a language with lots of getters and setters
- go usually does not use getters and setters - you access values directly 
- Given the rectangle model, this is almost a square!
- You take the same amount of time to understand and use a setter/getter as you would to do it yourrself, without added value
- We want to strive for long tall rectangles. 


<h2> The smallest interface </h2>

- You can make an empty interface `type veryAgreeable interface {}`
- You can also just use `interface{}` as a type directly
- Everything satisfiest this interface
- There is an alias for this interface - `any`
- Any says nothing; ti  carries no information
- Be very careful with this as it leads to bugs and is usually not needed. 

