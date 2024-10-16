# Functional Programming

## First class functions

- passing functions as arguments to other functions
- returning functions as the values from other functions
- assigning functions to variables or storing them in data structure
- anonymous function / lambdas

## Pure functions

- same output for same function arguments/inputs.
- no side effects

## Higher order functions

- function that takes one or more functions as arguments or
- returns a function as its result

## Closure & state capturing

- technique for implementing lexically scoped name binding in a language with first class functions
- combination of a function bundled together with references to its surrounding state.

```go
func get_me_closure(v int) func(int) bool {
    return func(x int) bool {
        return v == x
    }
}
```

## Function currying

- technique of translating a function that takes multiple arguments into a sequence of families of function, each taking a single argument
- eg. option pattern
