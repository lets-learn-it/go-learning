


```antlr4
grammar Expr;

prog: stat+;
stat: expr NEWLINE
    | ID '=' expr NEWLINE
    | NEWLINE
    ;

expr: expr op=('*'|'/') expr
    | expr op=('+'|'-') expr
    | INT
    | ID
    | '(' expr ')'
    ;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
LPAREN: '(';
RPAREN: ')';
ASSIGN: '=';
ID: [a-zA-Z]+;
INT: [0-9]+;
NEWLINE: '\r'?'\n';
WS: [ \t]+ -> skip;
```

## References

[[00] Precedence and Associativity in Grammar Rules](https://arman-dev-blog.hashnode.dev/precedence-and-associativity-in-grammar-rules)
