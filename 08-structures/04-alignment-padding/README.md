# Alignment & Padding

- Golang defines the concept of byte alignment in the structure struct. Appropriate adjustment of the field order can save a lot of memory.
- Alignment speeds up memory access by generating code that requires one instruction each to read & write a memory location. without alignment, we may face a situation where the processor will have to use two or more instructions to access data located between addresses that are multiples of the size of the machine word.
- To align with machine word size, compiler adds padding in struct fields.

| type    | bytes |
|---------|-------|
| bool    | 1     |
| byte    | 1     |
| int     | 8     |
| float32 | 4     |
| float64 | 8     |
| string  | 16    |
| slice   | 24    |
| map     | 8     |

- use **aligncheck** tool to optimize struct size. (https://gitlab.com/opennota/check).