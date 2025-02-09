# Learning Data Structures in Go

## Pointers

```go
var i int = 42
var p *int = &i // p is a pointer to i
fmt.Println(*p) // by doing *p we are dereferencing the pointer
```

- Pointers help in passing the memory address of a variable to a function. This can be useful when you want to modify the value of a variable in a function.

## Hash Table

- Chaining for collision resolution using linked list

```go
type Bucket[K comparable, V any] struct {
    head *BucketNode[K, V]  // Linked list head
}

type BucketNode[K comparable, V any] struct {
    key K
    value V
    next *BucketNode[K, V]  // Next node in chain
}
```

- Generic types for flexibility
Standard hash table operations
Modern Go features like generics and the maphash package

### Understanding the Bitwise Operation in Hash Function

The hash function uses a bitwise AND operation (&) to calculate the bucket index. Here's a detailed explanation:

How It Works
Size Power of 2

Hash table size is typically a power of 2 (e.g., 16, 32, 64)
When we subtract 1 from a power of 2, we get a number with all bits set to 1 up to that position
Example for size 16:

> Bitwise AND Operation

The & operator performs a bitwise AND between:
The hash value (a large uint64 number)
(size-1) converted to uint64
Example:

```
16     = 0001 0000
16 - 1 = 0000 1111
```

```
hashValue   = 1101 1100 1011 0110
size-1      = 0000 0000 0000 1111
result      = 0000 0000 0000 0110 (6)
```


Benefits
Fast Operation: Bitwise operations are very efficient
Even Distribution: Ensures index is always within array bounds
Modulo Alternative: Equivalent to hashValue % size when size is power of 2, but faster
```
Bitwise AND: ~1 CPU cycle
Modulus: ~10-20 CPU cycles
```

# Stack & Queue

We can use slices to implement stacks and queues.

## Stacks
for stacks appending to the slice is the same as pushing to the stack
for pops we can use the built-in function `len` to get the length of the slice and then use `len-1` to get the last element of the slice
> and update the slice by slicing it from the beginning to the end - 1

## Queues
for queues we can use the built-in function `len` to get the length of the slice and then use `len-1` to get the last element of the slice
> and update the slice by slicing it from beginning + 1 to the end

## References

[Pointers in GO](https://youtu.be/2XEQsJLsLN0?si=MRGwlgKO38sbS5rc)