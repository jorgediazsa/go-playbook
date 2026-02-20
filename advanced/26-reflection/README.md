# 26 - Reflection (`reflect`)

Go is statically typed. Reflection allows a program to inspect and manipulate its own types and variables at runtime. 
The standard library uses it heavily for things like `encoding/json` and `fmt.Printf`.

---

## 1. `reflect.Type` and `reflect.Value`

- `reflect.TypeOf(x)` gives you metadata: the name, the kind (Struct, String, Slice), and struct tags.
- `reflect.ValueOf(x)` gives you the actual runtime value, allowing you to read it or mutate it.

### The Production Danger 
Reflection comes with three enormous penalties:
1. **Loss of Compile-Time Safety:** If you try to call `v.SetString("hello")` on an `int`, it will cause a structural `panic` at runtime, crashing your app.
2. **Performance Devastation:** Reflection defies compile-time optimizations like inlining. It is orders of magnitude slower than native code.
3. **Obscure Code:** Heavy reflection creates code that is impossible to navigate with an IDE.

---

## 2. When to Use It (And When NOT To)

**When to use it:**
- Building generic decoders like `json.Unmarshal`.
- Validating structs dynamically via Struct Tags (`validate:"required"`).

**When NOT to use it:**
- If an `interface` works, use an interface.
- If Generics (`[T any]`) solve the problem mathematically at compile time, use Generics.

---

## 3. How to Mutate Values

To change a value via reflection, you MUST pass a pointer to it.
`v := reflect.ValueOf(&myStruct).Elem()`
If you do not call `.Elem()` on a pointer, you cannot use `.Set()`, and the app will panic.

---

## Exercises

- `ex01_validator.go`
- `ex02_copier.go`
- `ex03_refactor.go`
