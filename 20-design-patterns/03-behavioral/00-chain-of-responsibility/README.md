# Chain of Responsibility

- It lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

```go
type Request struct {
    // things
}

type handler interface {
    handle(r *Request)
    setNext(h handler)
}
```
