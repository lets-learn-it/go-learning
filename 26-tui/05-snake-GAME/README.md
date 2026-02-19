# Snake GAME

**ANSI escape codes**

- hide cursor   `\033[?25l`
- show cursor   `\033[?25h`
- clear         `\033[2J`
- move cursor   `\033[1;1H`

### Debugging

sending messages to stderr

```go
fmt.Fprintf(os.Stderr, "snake after: %v\n", g.snake)
```

send stderr to file

```sh
go run . 2> demo.txt
```

## References

[[00] I will teach you how to make games in your terminal 👾](https://www.youtube.com/watch?v=Fyf-Wh4ckDk)
