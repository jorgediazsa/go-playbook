# 03 - Functions

This module covers the subtle execution timing and state capture behaviors of Go functions.

## 1. Named Returns and Defer

Go allows you to name the return values of a function. This creates zero-initialized variables in the function scope that are automatically returned when `return` (naked or not) is called.

### The Production Danger

A `defer` block executes *after* the `return` statement is evaluated, but *before* the function actually hands control back to the caller. This means a deferred closure can observe and **modify** named return values before they are emitted.

```go
// Dangerous or clever? (It returns 2!)
func getCount() (count int) {
    defer func() {
        count++
    }()
    return 1
}
```
This pattern is heavily used in production for modifying returned `error` values (e.g., decorating an error with trace contexts if an error occurred).

---

## 2. Closure Variable Capture

When an anonymous function (closure) references variables defined outside its body, it captures those variables by **reference**, not by value.

### The Production Danger

If the closure executes later (e.g., in a background goroutine or via a callback), it sees the *current* state of the variable at execution time, not the state at the time the closure was defined. 
This behaves similarly to loop-variable capture but happens in normal procedural logic when state mutates after registering a callback.

### Idiomatic Solution

Explicitly pass variables as arguments to the closure if you want to capture them by value (copying them at registration time).

---

## 3. Variadic Functions and Slice Unpacking

Variadic functions (`func sum(nums ...int)`) accept zero or more arguments of a type. Inside the function, `nums` is a slice `[]int`.

### The Production Danger

If you already have a slice and want to pass it to a variadic function, you must use the spread operator (`...`). If you modify the slice *inside* the variadic function, you are modifying the backing array of the original slice passed in by the caller!

```go
func modify(nums ...int) {
    nums[0] = 999
}

func main() {
    a := []int{1, 2, 3}
    modify(a...) // Modifies 'a'!
}
```

---

## Exercises

- `ex01_named_returns.go`
- `ex02_closures.go`
- `ex03_variadic.go`
