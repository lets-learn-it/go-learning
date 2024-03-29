# Variables

## Lifetime of Variable

- Lifetime of package level variable is the entire execution of the program.
- Local variables have dynamic lifetimes. new instance is created each time declaration statements is executed, and variable lives on until it becomes unreachable, at which point its storage may be recycled.
- Function parametes and results are local variables too; they are created each time their enclosing function is called.

## Declare 

```go
// syntax
// var <variable_name> <type>
var name string
name = "parikshit"

// shorhand syntax. compiler will infer
village := "Kavathe Ekand"
```
## stack & heap

- A compiler may choose to allocate local variables on the heap or on the stack but, perhaps surprisingly, this choice is not determined by whether `var` or `new` was used to declare the variable.
- It is determined by whether the variable is **reachable or not reachable**.
- Variables which escape their local/function scope are allocated on heap which has **extra memory allocation cost**. Variables which are not useful outside local/function scope, make sure they are not escaping their scope.


## Basic Types

- `bool` ( default value is `false` )
- `string` ( default value is "" )
- `int`, `int8`, `int16`, `int32`, `int64` 
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- ( default value for numeric types is `0` )
- `byte` ( alias for `uint8` )
- `rune` ( alias for `int32`. represents a Unicode code point )
- `float32`, `float64`
- `complex64`, `complex128`
