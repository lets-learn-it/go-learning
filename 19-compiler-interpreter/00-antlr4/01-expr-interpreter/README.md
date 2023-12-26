# Expr Interpreter

- Only basic arithmetic operators (add, subtract, multiply, and divide), parenthesized expressions, integer numbers, and variables.

**Example Sample**

```
193     # 193
a = 5   #
b = 6   #
a+b*2   # 17
(1+2)*3 # 9
```

## Commands

```sh
### Generate Go code
antlr4 -Dlanguage=Go -no-listener -visitor -o parser Expr.g4

### create module
go mod init github.com/pskp-95/antlr-expr

### get required module
go mod tidy
```

## Reference

[00] Definitive Antlr4 Reference - Parr, Terence