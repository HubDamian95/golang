<h1>Golang 22/11/2024</h1>

<h2> Introduction </h2>

<h2> Agenda </h2>

- Goroutines
- Channels
- Further Study

<h2> Review </h2>

- Concurrency and Parallelism
- Gophers at work...

<h2> Goroutines </h2>

- When we run our Go program, it automatically creates one goroutine, called the main routine
- When you use the <b> go </b> keyword before a function invocation it will launch its own goroutine.
- Lightweight thread manage by the Go rountime, it's not a system level thread.
- Common and supported to spawn many (100's - 10000's) of goroutines.

`go func () {}`

- Your program will keep running and not wait for the goroutine to finish - probably. 
- Concurrent vs parallel - Maybe you only have one gopher (or all your gophers are busy) so they have to finish the goroutine before they are available to continue running the rest of your code. 
- Let's see an example https://go.dev/play/p/f0sLl05L4f6

<h2> Channels </h2>

- We need a way for our gophers to communicate
- Channels are built in type in Go.
`ch := make(chan int)`
- Send or receive from a channel with <b><-</b>
`ch <- value`
- End value to the channel ch
`data_store := <- ch`
- Receive from channel ch and assign value to <b>data_store</b>
- Values written to a channel can be read only <b>once</b>
- If multiple people are listening to the same channel, only one will read the value.
- <b> Sends and receives are blocking </b>
- Close a channel with `close(ch)`
- The job of closing a channel falls to the routine that writes to it
- Writing to a closed channel will error, reading will not. 
- Use <b>comma ok</b> pattern to know if a channel has been closed. 
</br>

- Only need to close if someone is waiting for it to close, such as reading from a channel in a for loop. 
- The loop will sit (blocking) and receive values from the channel until it's closed

```
for v := range ch {
    fmt.Println(v)
}
```

- Classic example - https://go.dev/play/p/lSM5G7RRYbv

<h2> Deadlocks </h2>

- `<- ch` blocks until we receive a value
- `ch <- v` also blocks until a value being sent is received. 
- if two (or more) goroutines are both waiting on each other, that is a deadlock 
    - The Catch-22 of programming!
- If every goroutine is deadlocked, Go will kill your program
- Let's look at another example: https://go.dev/play/p/YM87xOFSJCj
    - Go <b>main</b> is waiting on <b>ch2</b> being read
    - The <b>func goroutine</b> is waiting on <b>ch1</b> being read

<h2> Channels </h2>

- Channels can also be buffered - provide a buffer size as the 2nd argument to make a buffered channel. 
- Sending to a buffered channels only block once the buffer is full
- Receiving form a buffered channel only blocks when the buffer is empty
- Buferred channels add a lot of complication
- Using them well is hard 
- Mostly use unbuffered channels

```
ch := make(chan int, 10)
```

<h2>Select</h2>

- The <b>select</b> statement is like a switch for concurrent operations
- Lets us read or write on multiple channels
- Will do the first one that's available
- If more than one are available, chooses randomly

```
select {
    case v := <- ch:
            fmt.Println(v)
    case v := <- ch2:
            fmt.Println(v)
}
```

<h2>For-Select</h2>

- Will move on after doing a single case
- Can wrap in a for loop to keep doing cases as they are satisfied
- Called for-select

```
for  {
    select {
    case v := <- ch:
            fmt.Println(v)
    case v := <- ch2:
            fmt.Println(v)
    }
}
```

<h2>Select</h2>

- `select` can help prevent deadlocks since the order does not matter
- Let's revisit the deadlock example (https://go.dev/play/p/NJUwZTpxAPy) from earlier but with a select statement.

<h2>Demo</h2>
Let's code out an example that will check if websites are up, concurrently

```
package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    links := []string{
        "http://google.com",
        "http://golang.org",
        "http://reddit.com",
        "http://pokeapi.co/",
        "http://amazon.com",
    }

    c := make(chan string)

    for _, link := range links {
        go checkLink(link, c)
    }

    for l := range c {
        go func() {
            time.Sleep(10 * time.Second)
            checkLink(l, c)
        }()
    }

}

func checkLink(link string, c chan string) {
    _, err := http.Get(link)
    if err != nil {
        fmt.Println(link, " might be down")
        c <- link
        return
    }
    fmt.Println(link, " is ok")
    c <- link
}
```

<h2> Further Study </h2>

<h2> Channels Further Study </h2>

- Try this Leetcode in Go (https://leetcode.com/problems/same-tree/)
- Now design your approach to be concurrent!
- <b>BONUS:</b> Make the concurrent Web Crawler from Tour of Go. (https://go.dev/tour/concurrency/10)

<h2> General Further Study </h2>

- Write your own container from scratch (and learn how they really work in <b>76 lines in Go</b>) [https://www.youtube.com/watch?v=8fi7uSYlOdc]

- The best Go Book: Learning Go, 2nd Edition by Jon Bodner (https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/)
- Go by example for (very) terse reference (https://gobyexample.com/)
- <b>Build anything! Do it now while it’s fresh.</b> Learning a modern, low level language like Go is useful but will also make you a better developer overall by strengthening your higher level abstractions in your languages of choice.
