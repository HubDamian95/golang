<h1>Golang 21/11/2024</h1>

<h2> Introduction </h2>

<h2> Agenda </h2>

- Concurrency
- Parallelism

<h2>Concurrency And Parallelism</h2>

- <b>This is just a taste</b> – we don’t have time to dig in more
- Go’s most famous feature is its concurrency (and thus scaling) features
- Simple (and wrong) definition – computer doing more stuff at once
- Intentionally built for servers – not all problems respond well to a concurrent approach but serving requests responds very well
- Adding concurrency will <b>not</b> automatically make things faster
- It may even make them <b>slower</b>
- It depends on the type of problem and the quality of your concurrent design
- Go’s model is based on the paper Communicating Sequential Processes (https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf) also known as CSP by C.A.R. Hoare
- Same person who invented quicksort!
- Extremely influential 12 page paper–Highly recommended read! Has impact far outside of just Go

<h2> Concurrency is not parallelism </h2>

<h2> Concurrency </h2>

- Tasks which start, run, and complete in <b>overlapping</b> time periods
- Same time–but not simultaneous
- If you use JavaScript a lot and know the event loop, you know it is single-threaded (one thing at a time) but concurrent (many things at once)
- Another way to think about this is your tasks having the property of <b>interruptibility</b>
- Composition (design, structure) of independently executing processes
- <b>Dealing</b> with a lot of things at once

<h2> Parallelism </h2>

- Tasks which run at the same time – simultaneous
- Another way to think about this is your tasks having the property of <b>independence</b>
- <b>Doing</b> a lot of things at once

<h2> Concurrency and Parallelism </h2>

- Very related ideas, but not the same
- Concurrency is about <b>structure</b> and parallelism is about <b>execution</b>
- Doing concurrency well will allow us to <b>maybe</b> also add parallelism to improve performance
- Parallelism is not the explicit goal of concurrency, the goal is a good structure for a problem
- Not worth doing on things that are already fast – this all has overhead

<h2> Huh? - Gopghers at work </h2>

- Let’s go through some examples to see these two different ideas at work
- Examples are from Rob Pike’s talk (https://www.youtube.com/watch?v=oV9rvDllKEg)
- Gopher burning C++ language manuals to forward it’s own interests
- With only one Gopher this will take too long

<h2> But I don't want to burn the C++ Manuals </h2>

- That's OK. 
- The example is actually fairly accurate of web server architecture
- Gopher - CPU
- Book pile - Web content
- Cart - Networking, marshalling (however you are moving data)
- Incinerator - Consumer of the data like a web browser, proxy

<h2> Coming Up </h2>

- Goroutines
- Channels
- Further Study