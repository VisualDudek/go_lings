| Question | Answer |
|----------|--------|
| Every Go executable file must contain which two things at minimum? | The {{c1::package main}} declaration and a {{c1::func main()}} function |
| What is the role of the {{c1::func main()}} function in Go? | It is the entry point where program execution starts when the executable runs |
| What does the {{c1::package main}} declaration mean? | This file belongs to the main package, and when compiled, will produce an executable |
| What is the syntax for importing a single package in Go? | {{c1::import "packagename"}} |
| How do you import multiple packages in Go? | Using a factored import block:<br>{{c1::import (<br>&nbsp;&nbsp;"package1"<br>&nbsp;&nbsp;"package2"<br>)}} |
| What is a standard library package in Go and give an example? | Standard library packages ship with Go. Example: {{c1::fmt}} for formatted I/O or {{c1::math}} for math functions |
| How would you import the {{c1::Intn}} function from the {{c1::math/rand}} package? | {{c1::import "math/rand"}} then call it as {{c1::rand.Intn(100)}} |
| In Go, what determines whether an identifier (function, constant, variable) is exported? | Whether it begins with an {{c1::uppercase}} letter (exported) or {{c1::lowercase}} letter (unexported/private) |
| Why is {{c1::math.pi}} not accessible but {{c1::math.Pi}} is? | {{c1::math.pi}} starts with lowercase (unexported), while {{c1::math.Pi}} starts with uppercase (exported) |
| Can you access an unexported identifier from another package? | No, unexported identifiers are private to their package |
| In Go, what is the basic syntax for defining a function that adds two integers? | {{c1::func add(a int, b int) int { return a + b }}} |
| When parameters in a Go function share the same type, how can you shorten the declaration? | List the type once after all parameter names: {{c1::func add(a, b int) int}} |
| What does a function return type declaration look like in Go? | The type is listed after the parameter list: {{c1::func add(a, b int) int}} |
| Can Go functions return multiple values? | Yes, functions can return multiple values |
| How do you declare a function that returns two string values? | Use the syntax: {{c1::func swap(x, y string) (string, string)}} |
| How do you capture multiple return values when calling a function? | Use the short variable declaration with comma separation: {{c1::a, b := swap("hello", "world")}} |
| What is a named return value in Go? | A return value with a declared name and type, automatically initialized to its zero value |
| What is the advantage of using named return values? | You can use a bare {{c1::return}} statement (without listing values) since the named values are already in scope |
| In the {{c1::split(sum int) (x, y int)}} example, what do {{c1::x}} and {{c1::y}} represent? | Named return values: {{c1::x}} and {{c1::y}} are initialized to zero and are automatically returned when you call {{c1::return}} |
| Is {{c1::return}} without values idiomatic in Go? | It's allowed when using named return values but can reduce code clarity; usually used only in short functions |
| What is the zero value in Go and why does it matter? | The default value assigned to a declared variable when no initial value is specified (0 for int, false for bool, "" for string, etc.) |
| What is the output of this code: {{c1::var i int; fmt.Println(i)}}? | {{c1::0}} (the zero value of int) |
| What is the difference between {{c1::var i int}} and {{c1::var i int = 5}}? | The first declares {{c1::i}} with zero value (0), the second declares and initializes it to 5 |
| How do you declare and initialize multiple variables of the same type in Go? | {{c1::var a, b, c int = 1, 2, 3}} or declare them on separate lines |
| How do you declare variables of different types with initialization? | {{c1::var a int = 1; var b string = "hello"}} or use the block syntax:<br>{{c1::var (<br>&nbsp;&nbsp;a int = 1<br>&nbsp;&nbsp;b string = "hello"<br>)}} |
| In the declaration {{c1::var a = 1}}, what type will {{c1::a}} have? | {{c1::int}} (inferred from the initial value) |
| In the declaration {{c1::var a, b = 1, "hello"}}, what types will they have? | {{c1::a}} will be {{c1::int}} and {{c1::b}} will be {{c1::string}}, inferred from initial values |
| What is the short variable declaration syntax in Go? | Using {{c1:::=}} operator, for example: {{c1::k := 3}} |
| Where can you use the short variable declaration operator {{c1:::=}}? | Only inside functions (not at package level) |
| Why can't you use {{c1::x := 5}} at the package level? | The {{c1:::=}} short declaration operator is only valid inside functions; package-level variables must use {{c1::var}} |
| What would happen if you tried {{c1::x := 5}} outside a function? | You would get a syntax error |
| What are the basic numeric types in Go? | Integer types ({{c1::int8}}, {{c1::int16}}, {{c1::int32}}, {{c1::int64}}, {{c1::uint}}, etc.), floating-point types ({{c1::float32}}, {{c1::float64}}), and complex types ({{c1::complex64}}, {{c1::complex128}}) |
| What is {{c1::uint64}} and what does the {{c1::u}} prefix mean? | {{c1::uint64}} is an unsigned 64-bit integer; the {{c1::u}} means it can only hold non-negative values |
| What is a complex number type in Go and give an example? | {{c1::complex128}} is a complex number type; example: {{c1::z := 5 + 12i}} |
| How do you calculate the square root of a negative number in Go? | Use {{c1::cmplx.Sqrt()}} from the {{c1::math/cmplx}} package, which handles complex numbers |
| What does the {{c1::<<}} operator do in Go? | It is the left bit shift operator; {{c1::1<<64}} shifts 1 left by 64 bits |
| How would you express the maximum value for a {{c1::uint64}}? | Using bit shift: {{c1::1<<64 - 1}} (all bits set to 1) |
| What is the purpose of using {{c1::1<<64 - 1}} instead of a magic number? | It makes the intent clear (maximum value for a 64-bit unsigned integer) and is self-documenting |
| What does the {{c1::%T}} format verb do in {{c1::fmt.Printf()}}? | Prints the type of the value |
| What does the {{c1::%v}} format verb do in {{c1::fmt.Printf()}}? | Prints the value in a default format |
| How would you print both the type and value of a variable {{c1::x}}? | {{c1::fmt.Printf("Type: %T Value: %v\n", x, x)}} |
| What is the difference between {{c1::var c bool}} and {{c1::c := true}}? | {{c1::var c bool}} declares {{c1::c}} at package level with zero value false; {{c1::c := true}} is short declaration inside a function, must be inside a function |
| Why might forgetting to capitalize a function name cause issues? | An unexported (lowercase) function cannot be called from another package; only exported (capitalized) functions are accessible across packages |
| What happens if you try to use {{c1:::=}} at package level? | You get a syntax error; {{c1:::=}} is only valid inside functions |
| What is shadowing and why is it problematic? | Declaring a variable with the same name in an inner scope hides the outer scope variable; can cause bugs if unintended |
| In the hello world program, what does {{c1::fmt.Println()}} do? | Prints its arguments to standard output, automatically adding a newline at the end |
| What is the output of {{c1::fmt.Println(3, 5)}}? | {{c1::3 5}} (arguments separated by spaces with a newline) |
