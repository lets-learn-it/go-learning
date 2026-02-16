# TUI

Standard escape codes are prefixed with **Escape**:

- Ctrl-Key: `^[`
- Octal: `\033`
- Unicode: `\u001b`
- Hexadecimal: `\x1B`
- Decimal: `27`

```sh
# bold & underline to text \x1b[<num>m
#
printf '\x1b[1;4m%s\n' 'text'
```

## References

[[01] Everything you never wanted to know about ANSI escape codes](https://notes.burke.libbey.me/ansi-escape-codes/)

[[01] ANSI Escape Codes](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797)

[[02] ANSI/VT100 Terminal Control Escape Sequences](https://www.cse.psu.edu/~kxc104/class/cse472/09f/hw/hw7/vt100ansi.htm)
