<h1> Interfaces in practice </h1>

<h2> Introduction </h2>

- Go Docs
- Interface Examples
- Stringer
- HTTP Module
- Reader Interface
- Writer Interface
- JSON

<h2> Stringer </h2>
Open the docs for the fmt package and find the Stringer interface
- One of the most common interfaces
- Anything that can describe itself as a string

```
func (d dog) String() string {
    return fmt.Sprintf("Name: %v,  Age: (%d)", d.name, d.age)
}
```

<h2> HTTP Module </h2>
Let’s look at the Net/HTTP package in Go - https://pkg.go.dev/net/http#pkg-overview

`resp, err := http.Get("http://example.com/")`

- That seems simple, let’s dig deeper
- Where can we find info about Get?
- What about Response?
- What’s a Reader?

<h2> Reader </h2>

- Let’s look at the Reader interface in Go - `https://pkg.go.dev/io#Reader`
- The entire reader interface

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

- Pass a byte slice to a reader and it will mutate that slice with its source of data – read the data onto the byte slice
- This allows a common interface for reading from many different sources with a single type
- The data sources could be wildly varied – HTTP, the disk, the keyboard, a image, sensors
- Reader describes something that can provide data to our program
- You might expect the you get returned some data when you call `read()`
- In Go, we are taking advantage of mutation and pointers
- We have to make and pass our own byte slice to the Reader
- This allows us more control
- We keep the declaration and management of the data structure on “our” side – with the consumer, rather than the producer
- Common pattern in Go

```
resp, err := http.Get("http://google.com")
bs := make([]byte, 99999)
resp.Body.Read(bs)
fmt.Println(string(bs))
```

<h2> Writer </h2>

Let’s go to the docs again - `https://pkg.go.dev/io#Writer`

```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

- The “opposite” of a reader
- Take a byte slice of data and sends it to an output source
- This allows a byte slice to go to any variety of output sources

<h3> io.Copy </h3>

- Let’s look at Copy in the io package - https://pkg.go.dev/io#Copy
- Let’s look at Stdout from the os package - https://pkg.go.dev/os#pkg-variables
- Let’s try to make our own writer to print our HTTP response

<h3> JSON </h3>

- If were using HTTP on the back end, chances are we care about JSON
- Let’s take a look at the encoding/JSON docs
- Marshal and Unmarshal for encoding and decoding
- Specify json field names with struct field tags
- Must be `tagName: "tagValue"` in backticks, one line

<h3> Marshal/Unmarshal vs Encode/decode </h3>

Related and sometimes used interchangeably
- Encoding - converting to binary, operates on streams of bytes
- Decoding - converting from binary to string, operates on streams of bytes
- Marshal - Map some other types (like JSON) to Go types
- Unmarshal - Map Go types to some other types (like JSON)
- Marshaling — change the in memory representation of data to a different format

<h2> Warmup </h2>

<h3> Warmup </h3>
Let’s take a look at the Pokemon API
pokedex/kanto is the API we’ll use

```
type Response struct {
    Name string `json:"name"`
        Pokemon []Pokemon `json:"pokemon_entries"`
}
```

<h2> Coming up </h2>

- Wordle lab exercise (continued)
- Intro to Concurrency