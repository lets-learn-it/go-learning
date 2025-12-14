# Concurrency

**OS scheduler**: Preemptive multitasking while **Go schedule**: Cooperative multitasking.
**Linux thread**: 2MB memory while **Go routine**: 2KB memory.

## Concurrency is hard

### Race Conditions

- occurs when two or more operations must execute in correct order but its not.
- **data race**: one concurrent operation attempts to read a variable while at some undermined time another concurrent operation is attempting to write to the same variable.

### Atomicity

- Within the context it is operating, it is indivisible or uninterruptible.

### Memory access synchronization

- **critical section**: section of program that needs exclusive access to a shared resource.

### Deadlocks, livelocks & starvation



## References

[[00] https://www.udemy.com/course/working-with-concurrency-in-go-golang/](https://www.udemy.com/course/working-with-concurrency-in-go-golang/)
