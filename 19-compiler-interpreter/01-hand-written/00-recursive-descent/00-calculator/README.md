# Calculator

## grammer

```
<expression> ::= <term> (("+" | "-") <term>)*

<term>       ::= <factor> (("*" | "/") <factor>)*

<factor>     ::= <integer> | "(" <expression> ")"

<integer>    ::= <digit>+

<digit>      ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```