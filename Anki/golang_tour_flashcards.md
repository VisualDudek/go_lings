# Go Fundamentals - Flashcards

| Question | Answer |
|----------|--------|
| What is the entry point function in every Go executable? | The `func main()` function in the `main` package |
| What two elements must every executable Go program contain? | 1. A package main declaration<br>2. A func main() function |
| How do you import a single package in Go? | `import "package_name"` or import in a block with multiple packages |
| What is the difference between a single import and factored imports? | Single: `import "fmt"` on one line<br>Factored: Multiple imports in () block:<br>`import (`<br>`"fmt"`<br>`"math"`<br>`)` |
| In Go, what determines if a name is exported from a package? | Whether it starts with a capital letter (exported) or lowercase (unexported) |
| What will happen if you try to access `math.pi` instead of `math.Pi`? | Compilation error - Go is case-sensitive and only capitalized names are exported from packages |
| How do you call an exported function from an imported package? | Use dot notation: `PackageName.ExportedFunction()`<br>Example: `math.Sqrt(16)` |
| What is the syntax for a Go function that takes two parameters? | `func functionName(param1 type, param2 type) returnType { }` |
| How can you shorten parameter type declarations when consecutive parameters share the same type? | Instead of `func add(a int, b int)`, write `func add(a, b int)` |
| What is the syntax for a Go function that returns multiple values? | List return types in parentheses: `func swap(x, y string) (string, string) { return y, x }` |
| How do you capture multiple return values from a function call? | Use the short declaration with comma separation:<br>`a, b := functionReturningTwo()` |
| What is a naked return in Go? | A return statement without values in a function with named return values<br>Example: `func split(sum int) (x, y int) { x = 5; y = 3; return }` |
| What is the benefit of named return values in Go? | They serve as documentation and can be returned without explicitly listing them (naked return) |
| What is the zero value in Go? | The default value given to variables when declared without initialization<br>Examples: 0 for int, false for bool, empty string "" for string |
| Write code that declares an integer variable without initialization and prints its value | `var i int`<br>`fmt.Println(i)` // prints 0 |
| How do you declare multiple variables with initial values in Go? | `var a, b int = 1, 2`<br>or `var name type = value, name2 type = value2` |
| Can you use the short variable declaration `:=` outside of a function? | No, `:=` can only be used inside functions. Use `var` at package level. |
| What does `:=` do in Go and when can it be used? | It declares and initializes a variable with inferred type. It can only be used inside functions. |
| What is the syntax for declaring variables with inferred types in a factored block? | `var (`<br>&nbsp;&nbsp;&nbsp;&nbsp;`ToBe bool = false`<br>&nbsp;&nbsp;&nbsp;&nbsp;`MaxInt uint64 = 1<<64 - 1`<br>`)` |
| What are the basic numeric types in Go? | int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128 |
| How do you create a complex number in Go? | Use imaginary notation: `z := 1 + 2i` or use `cmplx` package functions |
| What does `fmt.Printf("Type: %T Value: %v\n", value)` print? | Both the type (%T) and value (%v) of a variable |
| What is the package structure required for a Go program to be executable? | `package main` declaration at the top of the file with a `func main()` entry point |
| What happens if you try to use `import "math/rand"` and then call `math.Intn()`? | Compilation error - you must use `rand.Intn()` because the package name is `rand`, not `math` |
| In the context of Go packages, what does "nested package path" mean? | A subpackage accessed via a path like "math/rand" where `math` is a collection and `rand` is the actual package |
| Predict output: `var x, y int`<br>`fmt.Println(x, y)` | `0 0` (zero values for uninitialized int variables) |
| Predict output: `var a, b, c = 1, "hello", true`<br>`fmt.Printf("%T %T %T\n", a, b, c)` | `int string bool` (types inferred from assignment) |
| Predict output: `func swap(a, b int) (int, int) { return b, a }`<br>`x, y := swap(3, 5)`<br>`fmt.Println(x, y)` | `5 3` |
| Write a function that returns two values and uses a naked return statement | `func split(n int) (x, y int) {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`x = n / 2`<br>&nbsp;&nbsp;&nbsp;&nbsp;`y = n - x`<br>&nbsp;&nbsp;&nbsp;&nbsp;`return`<br>`}` |
| What is the main difference between `var i int` (at package level) and `i := 5` in terms of scope and usage? | `var` declares at package level and gets zero value; `:=` declares at function level with inferred type from initialization |
| Why would you use `var` blocks in Go instead of separate var declarations? | For better organization and readability when declaring multiple related variables at package level |
| What is the syntax for explicit type conversion in Go? | `T(v)` where T is the target type and v is the value<br>Example: `float64(3)` or `uint(f)` |
| Why does Go require explicit type conversion between types? | To prevent implicit coercion bugs and make type changes visible and intentional in code |
| What happens if you try to assign a float value to an int variable without conversion in Go? | Compilation error - Go does not allow implicit type conversion |
| Predict output: `var x int = 3`<br>`var f float64 = float64(x)`<br>`fmt.Println(f)` | `3` (the int is converted to float64, which prints as 3) |
| What is the difference between variables declared with `var` and constants declared with `const`? | Variables can change at runtime; constants have values fixed at compile time and cannot be modified |
| Can you use the short declaration operator `:=` to declare constants in Go? | No, constants must be declared with `const` keyword; `:=` cannot be used for constants |
| What is the scope of a constant in Go? | Constants can be package-level (accessible throughout the package) or block-scoped (inside a function) |
| Predict output: `const Pi = 3.14`<br>`func main() {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`const Greeting = "Hi"`<br>&nbsp;&nbsp;&nbsp;&nbsp;`fmt.Println(Pi, Greeting)`<br>`}` | `3.14 Hi` (both package-level and function-level constants are accessible) |
| What are "untyped constants" in Go? | Constants without explicit type annotations that can hold very large values and adapt to the context in which they are used |
| Why can untyped constants hold very large values like 1<<100 without overflow? | Because they exist at compile time with arbitrary precision and only take on a concrete type when assigned to a variable or used in a context requiring a specific type |
| Predict output: `const Big = 1 << 100`<br>`fmt.Println(Big)` | Compilation error - Big has no concrete type and cannot be printed without explicit conversion or assignment to a typed variable |
| What is `iota` in Go and what does it do? | `iota` is a built-in constant that starts at 0 and increments by 1 for each const declaration in a const block |
| How do you create enumeration values in Go using `iota`? | Use `iota` in a const block where each constant automatically increments:<br>`const (`<br>&nbsp;&nbsp;&nbsp;&nbsp;`Red = iota   // 0`<br>&nbsp;&nbsp;&nbsp;&nbsp;`Green        // 1`<br>&nbsp;&nbsp;&nbsp;&nbsp;`Blue         // 2`<br>`)` |
| Can you use expressions with `iota` in Go? | Yes, you can perform arithmetic on iota:<br>`const (`<br>&nbsp;&nbsp;&nbsp;&nbsp;`StatusOK = iota + 200 // 200`<br>&nbsp;&nbsp;&nbsp;&nbsp;`StatusCreated          // 201`<br>`)` |
| What is a bit shift operator in Go and what does `1 << 8` mean? | `<<` is the left bit shift operator; `1 << 8` shifts 1 left by 8 bits, resulting in 256 (2^8) |
| What does `1 << 16 - 1` equal and what does it represent? | 65535, which is the maximum value for a uint16 (2^16 - 1) |
| What is the purpose of bit flags in Go? | To represent multiple boolean states using a single integer, where each bit represents a different flag |
| Predict output: `const (`<br>&nbsp;&nbsp;&nbsp;&nbsp;`READ = 1 << iota   // ?`<br>&nbsp;&nbsp;&nbsp;&nbsp;`WRITE              // ?`<br>&nbsp;&nbsp;&nbsp;&nbsp;`EXECUTE            // ?`<br>`)` | READ = 1 (0b001), WRITE = 2 (0b010), EXECUTE = 4 (0b100) |
| How do you check if a bit flag is set in a Go variable? | Use the bitwise AND operator `&`:<br>`if userPerms&READ != 0 { /* can read */ }` |
| How do you combine multiple bit flags in Go? | Use the bitwise OR operator `\|`:<br>`userPerms := READ \| EXECUTE` |
| What is the syntax for complex number literals in Go? | Use the form `real + imagi` (e.g., `0.8 + 0.5i`, `3 + 4i`) |
| What package should you import to work with complex numbers in Go? | The `math/cmplx` package provides functions like `cmplx.Sqrt()` for complex arithmetic |
| When Go infers the type of a numeric constant like `3.14`, what type does it become? | `float64` |
| When Go infers the type of a numeric constant like `42`, what type does it become? | `int` |
| When Go infers the type of a numeric constant like `0.8 + 0.5i`, what type does it become? | `complex128` |
| What is the key advantage of untyped constants maintaining full precision? | You can perform arithmetic on very large numbers at compile time without loss of precision, and the result adapts to the target type when assigned |
| Predict output: `const HighPrecision = 1e20`<br>`const LowPrecision = HighPrecision / 1e10`<br>`fmt.Println(LowPrecision)` | `1e+10` (the division is computed with full precision at compile time) |
| In Go, if you have `const x = 42` and later use `f := float64(x)`, what happens? | The constant 42 (which is untyped) is implicitly converted to float64 at the assignment; this is allowed because untyped constants adapt to context |
| Why does Go allow implicit conversion of untyped constants but not variables? | Untyped constants are compile-time values with arbitrary precision, so the conversion is safe and verified at compile time; variables require explicit conversion to prevent runtime errors |
| Can you mix typed and untyped constants in arithmetic in Go? | Yes, untyped constants will take on the type of the typed operand in the expression, limited by that type's range |
| What is the syntax for a traditional for loop in Go? | `for init; condition; post { }`<br>Example: `for i := 0; i < 10; i++ { }` |
| Are the init, condition, and post statements required in a Go for loop? | No, any of them can be omitted. Example: `for i := 0; i < 10` (no post) or `for i < 10` (no init or post) |
| How does the modern range-based for loop work in Go, and what does `for i := range 10` do? | It iterates from 0 to 9 (10 iterations total). Introduced in Go 1.22, it's equivalent to `for i := 0; i < 10; i++` |
| What is the difference between `for i := 0; i < 10; i++` and `for i := range 10`? | Traditional form requires explicit init, condition, and post; range form is shorter and only requires the upper limit. Both iterate from 0 to 9. |
| What is a while-like loop in Go and how do you write one? | Use `for` with only a condition: `for condition { }` without init or post. Example: `for i < 10 { i++ }` |
| Can you have a for loop with no init, condition, or post statements in Go? | Yes, `for { }` creates an infinite loop that runs forever until break or return |
| What keyword would you use to exit an infinite loop in Go? | Use `break` to exit the loop |
| In a traditional for loop, when are the init, condition, and post statements executed? | init runs once before the loop starts; condition is checked before each iteration; post runs after each iteration |
| Predict output: `for i := 0; i < 3; i++ { fmt.Println(i) }` | `0`<br>`1`<br>`2` |
| Predict output: `for i := range 3 { fmt.Println(i) }` | `0`<br>`1`<br>`2` |
| Predict output: `i := 0; for i < 3 { fmt.Println(i); i++ }` | `0`<br>`1`<br>`2` |
| Predict output: `for { fmt.Println("infinite"); break }` | `infinite` (loop breaks immediately on first iteration) |
| Predict output: `sum := 0; for i := 0; i < 3; i++ { sum += i }; fmt.Println(sum)` | `3` (0 + 1 + 2 = 3) |
| What is the scope of a variable declared in the init statement of a for loop? | The variable is scoped to the for loop only and is not accessible after the loop ends |
| Does `i` declared in `for i := 0; i < 10; i++` exist after the loop? | No, the variable is scoped to the for loop only |
| What happens if you try to modify the loop variable in the post statement? | It executes normally; the post statement modifies the loop variable after each iteration. Example: `for i := 0; i < 10; i++` increments i by 1 each iteration |
| How do you create a loop that counts from 10 down to 1? | `for i := 10; i >= 1; i-- { fmt.Println(i) }` |
| Can you have multiple statements in the init or post part of a for loop? | Yes, separate them with commas: `for i, j := 0, 10; i < j; i, j = i+1, j-1 { }` |
| What is an if statement in Go and what is the basic syntax? | It evaluates a condition and executes a block if true: `if condition { }` |
| Does an if statement in Go require parentheses around the condition? | No, parentheses are not allowed. Syntax: `if condition { }` (no parentheses) |
| What is the syntax for an if statement with a short variable declaration? | `if varName := expression; condition { }`<br>Example: `if v := computeValue(); v < limit { }` |
| What is the scope of a variable declared in the short declaration of an if statement? | The variable is scoped to the if block only and is not accessible after the if block ends |
| Predict output: `if v := 5; v > 3 { fmt.Println("yes") } else { fmt.Println("no") }` | `yes` |
| Predict output: `if v := 2; v > 3 { fmt.Println("yes") } else { fmt.Println("no") }` | `no` |
| Predict output: `if v := 5; v > 3 { fmt.Println(v) }; fmt.Println(v)` | Compilation error - v is scoped to the if block only |
| How do you write an if-else statement in Go? | `if condition { } else { }`<br>The else block executes if the condition is false |
| Can you declare a variable in the if statement and use it in the else block? | Yes, a variable declared in the if short declaration is accessible in both the if and else blocks |
| Predict output: `if v := 5; v > 3 { fmt.Println("big") } else { fmt.Println(v) }` | `big` (else block doesn't execute because condition is true) |
| Predict output: `if v := 2; v > 3 { fmt.Println("big") } else { fmt.Println(v) }` | `2` (else block executes and prints v) |
| Can you chain multiple else if conditions in Go? | Yes: `if cond1 { } else if cond2 { } else if cond3 { } else { }` |
| What is a switch statement in Go? | A statement that evaluates an expression and executes one of several code blocks based on matching cases |
| What is the syntax for a basic switch statement in Go? | `switch expr { case val1: ... case val2: ... default: ... }` |
| Does Go require a `break` statement between cases in a switch? | No, Go does not fall through between cases by default; each case is separate. The break is implicit. |
| What is the difference between Go's switch and switch statements in other languages like C or Java? | Go does not fall through between cases by default; cases are independent. Other languages require explicit `break` to prevent fallthrough. |
| How do you write a switch statement with a short variable declaration? | `switch varName := expression; varName { case val1: ... case val2: ... }`<br>Example: `switch os := runtime.GOOS; os { case "linux": ... }` |
| What is a switch statement without a tag? | A switch with no expression where each case is a boolean expression: `switch { case condition1: ... case condition2: ... }` |
| When would you use a switch without a tag instead of if-else? | When you have multiple boolean conditions and want cleaner syntax than chained if-else-if statements |
| Predict output: `switch 2 { case 1: fmt.Println("one") case 2: fmt.Println("two") default: fmt.Println("other") }` | `two` |
| Predict output: `switch 4 { case 1: fmt.Println("one") case 2: fmt.Println("two") default: fmt.Println("other") }` | `other` |
| Predict output: `x := 5; switch { case x < 0: fmt.Println("negative") case x == 0: fmt.Println("zero") default: fmt.Println("positive") }` | `positive` |
| Predict output: `t := time.Now().Hour(); switch { case t < 12: println("morning") case t < 17: println("afternoon") default: println("evening") }` | Depends on current time; if 10:00 AM: `morning`; if 3:00 PM: `afternoon`; if 8:00 PM: `evening` |
| Can you use complex expressions in case statements? | Yes, case expressions are evaluated and compared to the switch expression: `case today + 1` or `case expr1 + expr2` |
| What does a `default` case do in a switch statement? | It executes if no other cases match. It's optional and serves as a fallback. |
| What is the `defer` statement in Go? | A keyword that schedules a function call to run after the surrounding function returns |
| What is the basic syntax of a defer statement? | `defer functionCall()`<br>Example: `defer fmt.Println("cleanup")` |
| When does a deferred function call execute? | After the surrounding function returns but before it actually exits and the control returns to the caller |
| Predict output: `fmt.Println("a"); defer fmt.Println("c"); fmt.Println("b")` | `a`<br>`b`<br>`c` (defer executes after all other statements in the function) |
| Predict output: `defer fmt.Println("2nd"); fmt.Println("1st")` | `1st`<br>`2nd` (fmt.Println("1st") executes first, then deferred statement) |
| If you have multiple defer statements, in what order do they execute? | In LIFO (Last-In-First-Out) order, like a stack. The last defer statement is the first to execute. |
| Predict output: `defer fmt.Println("1"); defer fmt.Println("2"); defer fmt.Println("3")` | `3`<br>`2`<br>`1` |
| What happens when you defer a function call inside a for loop? | All deferred calls are pushed onto a stack and execute in LIFO order after the loop completes and the function returns |
| Predict output: `for i := 0; i < 3; i++ { defer fmt.Println(i) }; fmt.Println("done")` | `done`<br>`2`<br>`1`<br>`0` (defers execute in LIFO order after the loop) |
| Predict output: `for i := range 3 { defer fmt.Println(i) }; fmt.Println("end")` | `end`<br>`2`<br>`1`<br>`0` |
| Why is the defer order LIFO when deferring in a loop? | Because each iteration pushes a new defer onto the stack; the last iteration's defer is on top of the stack and executes first |
| What is a common use case for the defer statement? | Resource cleanup, such as closing files or database connections: `defer file.Close()` or `defer db.Close()` |
| Can you defer multiple different function calls in the same function? | Yes, all deferred calls will execute in LIFO order: `defer a(); defer b(); defer c()` executes c, then b, then a |
| Can the defer statement defer a method call or only a function? | You can defer both function calls and method calls: `defer fmt.Println()` or `defer obj.Method()` |
| What is the key takeaway about defer execution timing? | Deferred functions execute in LIFO order after the surrounding function's other statements complete but before the function returns |
| Can you defer an anonymous function in Go? | Yes, you can defer an anonymous function (closure): `defer func() { fmt.Println("cleanup") }()` |
| Write code that demonstrates multiple defers executing in LIFO order | `func main() {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`defer fmt.Println("first")`<br>&nbsp;&nbsp;&nbsp;&nbsp;`defer fmt.Println("second")`<br>&nbsp;&nbsp;&nbsp;&nbsp;`fmt.Println("main")`<br>`}`<br>Output: `main` then `second` then `first` |
| Write code that accumulates values in a for loop using a traditional for statement | `sum := 0`<br>`for i := 0; i < 10; i++ {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`sum += i`<br>`}`<br>`fmt.Println(sum) // 45` |
| Write code that demonstrates variable scoping in an if statement with short declaration | `if v := getValue(); v < limit {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`fmt.Println(v) // v is scoped to if block`<br>`}`<br>`// fmt.Println(v) would cause compilation error` |
| Write code that demonstrates a switch statement with a short variable declaration | `switch os := runtime.GOOS; os {`<br>`case "linux":`<br>&nbsp;&nbsp;&nbsp;&nbsp;`fmt.Println("Linux")`<br>`case "darwin":`<br>&nbsp;&nbsp;&nbsp;&nbsp;`fmt.Println("macOS")`<br>`default:`<br>&nbsp;&nbsp;&nbsp;&nbsp;`fmt.Println("Other")`<br>`}` |
| Write code that demonstrates a switch without a tag using boolean expressions | `t := time.Now().Hour()`<br>`switch {`<br>`case t < 12:`<br>&nbsp;&nbsp;&nbsp;&nbsp;`println("morning")`<br>`case t < 17:`<br>&nbsp;&nbsp;&nbsp;&nbsp;`println("afternoon")`<br>`default:`<br>&nbsp;&nbsp;&nbsp;&nbsp;`println("evening")`<br>`}` |
| Write a function that uses defer to demonstrate cleanup behavior | `func readFile(filename string) string {`<br>&nbsp;&nbsp;&nbsp;&nbsp;`file, _ := os.Open(filename)`<br>&nbsp;&nbsp;&nbsp;&nbsp;`defer file.Close()`<br>&nbsp;&nbsp;&nbsp;&nbsp;`// read from file`<br>&nbsp;&nbsp;&nbsp;&nbsp;`return "content"`<br>`}` |
| What is the key difference between `for i := 0; i < 10; i++` and `for i := range 10` in terms of readability? | The range form is more concise and modern (Go 1.22+), while traditional form is more explicit about init/condition/post |
| When iterating from 0 to 9, which form is preferred in modern Go? | `for i := range 10` is preferred in Go 1.22+; `for i := 0; i < 10; i++` is still valid and common in older code |
| How do you skip to the next iteration of a for loop? | Use the `continue` keyword: `for i := 0; i < 10; i++ { if i == 5 { continue }; fmt.Println(i) }` |
| How do you exit a for loop early? | Use the `break` keyword: `for i := 0; i < 10; i++ { if i == 5 { break }; fmt.Println(i) }` |
| What happens if you break out of a for loop with deferred statements? | The deferred statements still execute before the function returns |
| Predict output: `for i := 0; i < 5; i++ { defer fmt.Println(i); if i == 2 { break } }; fmt.Println("done")` | `done`<br>`2`<br>`1`<br>`0` |
| Can you use a defer statement inside an if block? | Yes, the defer still applies to the entire function: `if cond { defer fmt.Println("cleanup") }` |
| In a switch statement, does the default case have to be the last case? | No, the default case can appear anywhere, but it conventionally appears last for readability |
| Can you have a switch statement with no cases and no default? | Yes, it's valid syntax but does nothing: `switch expr { }` |
