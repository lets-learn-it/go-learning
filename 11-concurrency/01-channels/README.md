# Channels

- A means of allowing communication to and from a GoRoutine
- Channels can be buffered, or unbuffered
- Once you're done with a channel, you must close it
- Channels typically only accept a given type or interface
- When we caopy channel or pass one as an argument to function, we are copying a reference,s o caller & callee refer to same data structure.

```go
ch = make(chan int)    // unbuffered channel of int
ch = make(chan int, 3) // buffered channel with capacity 3
```

## Channel Operations

- Channel has 3 operations. send, recieve & close.
```go
ch <- x   // send x to channel ch
x = <- ch // receive value in x from channel ch
<-ch      // receiving from channel ch but value is discarded.
close(ch) // close channel ch
```

- attempt to send values to closed channel will result in panic.
- attempt to read from closed channel retuns zero value of channel's element type.

## Unbuffered channels

- send/recieve on unbuffered channel blocks sending/recieving goroutine until another goroutine executes a corresponding receive/send on same channel.
- Communication over an unbuffered channel causes the sending and recieving goroutines to synchronize. also called synchronous channels.