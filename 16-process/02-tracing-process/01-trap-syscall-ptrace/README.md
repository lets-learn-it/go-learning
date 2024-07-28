# Trap syscalls using ptrace

## Running

```sh
# build
go build

# run
./trap-syscall-ptrace pwd
```

## References

[[00] Strace in 60 lines of Go by Liz Rice - medium.com](https://medium.com/hackernoon/strace-in-60-lines-of-go-b4b76e3ecd64)

[[01] Linux System Call Table for x86 64 - blog.rchapman.org](https://blog.rchapman.org/posts/Linux_System_Call_Table_for_x86_64/)

[[02] google/gvisor - github.com](https://github.com/google/gvisor/blob/master/pkg/sentry/platform/ptrace/subprocess.go)

[[03] Part 1: DIY debugger in Golang - poonai.github.io](https://poonai.github.io/posts/how-debuggers-works-part1/)
