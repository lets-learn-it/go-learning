# Reflection

- It allows to examine and manipulate variables and types while your program is running.
- reflection may affect performance
- **reflect** package has 2 main types, `reflect.Type` and `reflect.Value`
- `reflect.Type` represents type of value while `reflect.Value` represents an instance of value.
- `reflect.TypeOf()` and `reflect.ValueOf()` function can be used to get them

## Common Usages

- Custom struct tag
- dynamic type checking and type assertions
- iterating over struct fields
- implementing dependency injection: using `ValueOf` and `Set` methods to set fields of structs at runtime.

## Reference

[[00] https://levelup.gitconnected.com/reflection-in-go-everything-you-need-to-know-to-use-it-effectively-52c78da1f4ff](https://levelup.gitconnected.com/reflection-in-go-everything-you-need-to-know-to-use-it-effectively-52c78da1f4ff)
