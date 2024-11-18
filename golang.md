<h1>Golang 18/11/2024</h1>


<h2>Starter</h2>
go run - compile and run
go build - compile
go install - compile and move result to $gopath/bin

<h2>Packages</h2>
Go programs are made up of packages
First Statement must be the package
Module is a collection of packages versioned together
Always at least one package, maim, and the name is reserved for this purpose
The main package must have a main function

<h2>Importing/Exporting</h2>
Only names starting with capital letter are exported
Unexported names cannot be accessed outside of the package
Import using import "package name"
Alias with import [alias] "package name"
Packages share scope across files, regardless of naming


Example project repo:
https://github.com/golang-standards/project-layout

Variables
https://go.dev/ref/spec#Variables

<h2>Basic Types:</h2>
bool
string
int - integer (positive)
uint - positive or negative
uintptr - unassigned int pointer
All integers are 64 bits if possible. 

int8, int16, int32, int64
uint8, uint16, uint32, uint64

byte - alias for uint8
rune - alias for int32, represents a unicode point
float32,float64 - IEEE-754 
complex64,complex128

<h2>Declaring Vars</h2>
Valid names contain a-z, A-Z, 0-9 and _
_ is special and for ignoring values
Declare with keyword var
Requires a type declaration var cat string
Can leave off type if declaring with initial value var cat = "mittens"
Short assignment syntax := means we can leave off var - cat := "mittens"
All values must have a type
Variables are MixedCase or camelCase //public/private case.

<h2>Strings</h2>
Double quotes "hello \n" - normal interpreted string literal, escape characters and interpreted
Backticks `hello \n` - raw string literal, escape characters not interpreted
Single quotes 'hello' - rune literal, this will error
- Single quotes can only hold a single character - one rune's worth
You probably want to use double quotes.

<h2>Printing </h2>
fmt package has a lot of lovely ways of pringting that support formatting similar to C's printf
```
fmt.Prinf("%T", someValue) //%T will print the value
fmt.Println("hello") //simple print to stdout ends with new line.
```

<h2>Zero Values</h2>
There is no undefined value in Go. Instead we have a "zero value"
Boolean: False
Numeric: 0
String: ""
Everything Else: Nil

<h2>Functions:</h2>
Declared with func keyword
Must declare type of parameters and return values:
```
func add(x int, y int) int {
    return x + y
}
```

Can provide type only once if shared by params
```
func add(x,y int) (int, int, int) {
    return x + y, y + y, x + x
}
```

You can return multiple values.

Conditionals
if [condition] {...}
if [condition] {...} else {...}

Can declare variables in the condition, scoped to the if block
```
if x := 10 * y; x < 8 {
    //code snippet
} else {
    // code snippet
}
```

<h2>Math ops:</h2>

add + 
subract - 
multiply *
divide /
modulo %
increment ++ 
decrement -- 

Bitwise:
XOR ^
OR |
AND &
shift left << 
shift right >>

<h2>Logical Ops:</h2>
not !
or ||
and && 
comparisons !, ==, <, >, >=, <=

<h2>Swich statements</h2>
Essentially a chained if else
stops upojn first successful case
```
switch [val] {
    case "go":
    //code snippet
    case f():
    //code snippet
    default:
    //code snippet
}
```

Switch without a condition is the same as switch true:
Good way to write log if-else chains

```
switch {
    case x < y:
    //
    case x > t: 
    // 
    default:
    //
}
```

<h2>Loops:</h2>
Only one loop syntax, the for loop

```
for [init]; [condition]; [post iteration] {
    //
} 

for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

The initialization statement and the post iteration statement are optional, Thus Go has a while loop it's just spelled "for"
```
i := 0
for i < 10 {
    fmt.Println(i)
    i++;
}
```


Infinite loop for {...}
continue and break keywords
```
i = 0
for {
    if i == 10 {break}
    fmt.Println(i)
    i++
}
```

<h2>Slice and Arrays</h2>
```
var [size][type] name
var arr = [6] int{2, 3, 5, 7, 11, 13}
```

The size is part of the type; it cannot be resized later
Because of this, they are not very common
Can loop and index

<h2>Slices</h2>
```
var [] [type] name
var slc = []int{2, 3, 5, 7, 11, 13}
```

- Dynamically sized reference to an array
- Under the hood, a slice creates two data structure
- The first slice is called "SLICE" it contains a pointer to the head of an array, a capacity and length
- It also makes an array where it stores the values.
- Check capacity and length with cap() and len() built-in functions
- Very Common

```
var [][type] name
```
- Nil slice - no underlying array

```
var slc = []int {2, 3, 5, 7, 11, 13}
```
- Slice literal

<h2>make Keyword</h2>
There is another way, using the make keyword
```
mySlice := make([]int, 5)
```
- Creates a new slice with length 5, capacity 5
- Values are filled in with the zero value
- Accepts optional third argument to set capacity
- Never set capacity to less than length

<h2>Slice Rules</h2>
Can declare on an existing array
```
 primes := [6]int{2,3,5,7,11,13}
 var s []int = primes[1:4]
```

 - Range indexes are incluse : exclusive
 - Omitting the first index defaults to 0
 - Omitting the last index defaults to length -1
 - Slicing out of bounds will error
 - Slices on the same array all point to the same underlying array

<h2>Slice Looping</h2>

- Looping over a slice can use a counter and our basic loop but more commonly the range keyword.
- Name the index _ if you aren't going to use it or Go will error
```
for index, value := range someSlice {
    fmt.Println(index, value)
}
```

<h2>Slice Appending</h2>
- Add to a slice with the built in function append
```
 append(someSlice, value)
```

- Resizes the underlying array if it's too small to hold a new value
- If there is enough space, Go will mutate the array with the value
- Returns updated slice
- If append causes a new array to be created, the returned slice will no longer point at the same array
- Prevent ending up with the same underlying array
```
 copy(slice1, slice2)
```

<h2>Warm Up</h2>
- Write a program that loops over a list of ints and prints each one out
- Loop over a slice of strings and print out each word
- Write a program that prints out each letter of a string

<h2>Encodings Quiz</h2>
s := "Senor"
fmt.Println(len(s))
fmt.Println(s[2])
fmt.Printf("%x ", s[2]) //hex
fmt.Printf("%c ", s[2]) //character
fmt.Printf(string(s[2]))

<h2>Encodings:</h2>
- Strings are just byte slices
- UTF-8 (byte) can handle any UTF char, not fixed size
- UTF-32 (rune) can handle any UTF char, fixed size
- Go assumes strings are UTF-8
- Go is a low level language and knowing how encoding works is not optional
- Ken Thompson, one of the two creators of Go, also made UTF-8
- Generally don't inder or slice strings - use the strings and utf8 packages

<h2>Encodings: Further Study</h2>
- Now that' we've gone over UTF encoding, look at the float IEEE 754 encoding specification
- BONUS: Use your understanding of string and float encodings to hide string messages inside a valid float value - specifically, a NaN.
