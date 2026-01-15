# Learn Go

This is a personal repository for learning the Go programming language. It contains various examples, exercises, and notes to help me understand Go's syntax, features, and best practices.

## 1. The Module Name

In Node, you name a project `learn-go` in `package.json`, but it lives at `github.com/user/learn-go`. The name and address are separate.

**In Go, the name IS the address.**

When you run:

```bash
go mod init github.com/ccrsxx/learn-go
```

You are telling Go: **"This URL is my Global ID."**

- **Node/Python:** Centralized Registry (NPM/PyPI). You ask for a name, the registry gives you the code.
- **Go:** Decentralized. The import string `github.com/user/repo` tells Go exactly where to go on the internet to find the code. It needs to be a valid URL and Git repository in order to fetch it remotely.

## 2. Absolute vs. Relative Imports

This is the biggest muscle-memory change.

### Node.js (Relative)

Node allows (and encourages) relative imports. The file system dictates the relationship.

```typescript
// "Go up one folder, then into utils"
import { validate } from '../utils/validation';
```

### Go (Absolute)

Go forbids relative imports like `../`. You must use the **Full Module Path** defined in `go.mod`.

```go
// "Start from the Project Root, then go to internal/utils"
import "github.com/ccrsxx/learn-go/internal/utils"
```

### The Pitfall of Wrong Module Names

If you initialize your module as `go mod init example/test`:

1. It works locally because Go treats it as a namespace.
1. **It fails remotely** because `example/test` is not a valid URL.
1. If you push to GitHub later, you have to **Find & Replace** every single import string in your entire codebase to match the real URL.
1. **Rule:** Always name your module the actual URL where it will live (e.g., `github.com/username/project`).

## 3. How `go.mod` "Routes" Imports

If I import my own code:

```go
import "github.com/ccrsxx/learn-go/internal/utils"
```

Why doesn't Go try to download this from the internet?

**The Magic:** `go.mod` acts as a local router.

1. Compiler sees the import path.
1. It checks `go.mod`: `module github.com/ccrsxx/learn-go`.
1. **Match\!** It realizes "This import starts with my username. I don't need to go to the internet; I will look in the local folder structure."

## 4. Exports: Capitalization vs. Keywords

- **Node/TS:** You explicitly decide what to export using the `export` keyword.
- **Go:** You decide visibility using **Capital Letters**.

<!-- end list -->

```go
func MyFunction() {} // Exported (Public) - Visible to other packages
func myFunction() {} // Unexported (Private) - Hidden inside this package
```

**Usage:**

```go
import "math"

math.Sqrt(4) // Works (Public)
math.sqrt(4) // Fails (Private/Undefined)
```

## 5. Dependency Management

Go solves "Dependency Hell" differently than `node_modules` or `venv`.

- **Node:** Nested `node_modules` (Duplication).
- **Python:** Global environment (Conflicts require `venv`).
- **Go:**
  1. **Global Cache:** Downloads version `v1.0` and `v2.0` to a central folder on your disk.
  1. **Per-Project Linking:** Your `go.mod` picks exactly which version from the cache to use.
  1. **Major Versions:** Go treats `v2` as a **different import path** (`github.com/lib/foo/v2`). This allows your app to use `v1` and `v2` of the same library simultaneously without conflict.

## 6. The "Package Main" Rule: Executable vs. Library

In Node.js, any file can be an entry point (`node utils.js` works fine). In Go, the entry point is strict.

### The Golden Rule

A Go program will **only compile as an executable** if:

1. The package name is **`package main`**.
1. It contains a **`func main()`**.

### The "One Package Per Folder" Constraint

Go does not allow you to mix package names in the same directory.

- **Node:** You can have `index.js` (app) and `helper.js` (module) in the same folder.
- **Go:** A folder can only be **one thing**:
  - **EITHER** an Executable (`package main`)
  - **OR** a Library (`package utils`)

## 7. The "Unused Import" Trap

Go is extremely strict about code hygiene compared to interpreted languages.

- **Node/Python:** You can leave `import fs` or `import os` at the top of the file while you code.
- **Go:** **Compile-Time Error.** The code will not run if an import is unused.

**The Debugging Workaround:**
If you need to comment out code but keep the import, use the **Blank Identifier (`_`)**:

```go
import (
    "fmt"
    _ "log/slog" // Compiler ignores that this is unused
)
```

## 8. Advanced Printing: Replacing Template Literals

In Node.js, you use backticks for string interpolation. Go uses **Printf verbs**.

### 1. Simple Printing (`Printf`)

Unlike `Println`, `Printf` does **not** add a new line automatically. You must add `\n`.

```go
// Node: console.log(`Hello ${name}, count: ${count}`)
fmt.Printf("Hello %s, count: %d\n", name, count)
```

### 2. Indexed Arguments (Reusing Variables)

You can reuse arguments without passing them multiple times using `[n]`.

```go
user := "Alice"
role := "Admin"

// Use user (1st arg) twice
fmt.Printf("%[1]s is a %[2]s. Goodbye %[1]s.\n", user, role)
```

### 3. Debugging Structs (`%+v`)

This is the closest thing to `JSON.stringify` for debugging objects.

- `%v`: Prints values `{Alice 101}` (Hard to read)
- `%+v`: Prints fields `{Name:Alice ID:101}` (**Use this\!**)

### 4. Using slog for Structured Logging

Go's `slog` package provides structured logging capabilities, allowing you to log messages with key-value pairs for better context.

```go
import "log/slog"

// User login user="Alice" id=101 active=true
slog.Info("User login", "user", "Alice", "id", 101, "active", true)
```

## 9. Array on Go with slices

Here's an array on Go compared to JavaScript arrays.

| Operation  | JavaScript               | Go                          |
| :--------- | :----------------------- | :-------------------------- |
| **Create** | `const arr = ["a", "b"]` | `arr := []string{"a", "b"}` |
| **Length** | `arr.length`             | `len(arr)`                  |
| **Push**   | `arr.push("c")`          | `arr = append(arr, "c")`    |
| **Get**    | `arr[0]`                 | `arr[0]`                    |
| **Slice**  | `arr.slice(1, 3)`        | `arr[1:3]`                  |

## 10. Most used fmt functions

Go has many printing functions, but these 4 cover 95% of use cases.

| Function          | Output     | Adds Newline? | Use Case               | Node.js Equivalent      |
| :---------------- | :--------- | :------------ | :--------------------- | :---------------------- |
| **`fmt.Println`** | Console    | ✅ Yes        | Simple logging         | `console.log(val)`      |
| **`fmt.Printf`**  | Console    | ❌ No         | Formatting variables   | `console.log(template)` |
| **`fmt.Sprintf`** | **String** | ❌ No         | Format to **variable** | `const s = template`    |
| **`fmt.Errorf`**  | **Error**  | ❌ No         | Wrap/Create **Error**  | `new Error(template)`   |

## 11. Way to declare variables

Go offers several ways to declare variables, each suited for different scenarios. Here are the most common methods:

| Keyword     | Where to use?         | Use Case                                                                                      |
| :---------- | :-------------------- | :-------------------------------------------------------------------------------------------- |
| **`const`** | Top level & Functions | Values that **never change** and are known at compile time.                                   |
| **`var`**   | Top level & Functions | When you need an **empty** variable (Zero Value) to fill later or you want to override later. |
| **`:=`**    | **Inside Functions**  | 99% of your code. Creating new variables with values.                                         |

## 12. Types of Error Checks on Go

There are two main types of error checks in Go:

| Method          | Question                   | Result            | Use Case                                                |
| :-------------- | :------------------------- | :---------------- | :------------------------------------------------------ |
| **`errors.Is`** | "Is it **Equal** to X?"    | `bool`            | Checking Sentinel Errors (`ErrNotFound`)                |
| **`errors.As`** | "Can it be **Cast** to X?" | `bool` + **Data** | Reading fields (Status Code, ID) from the error Structs |

## 13. Types of Loops on Go

Go has only one type of loop, the `for` loop, which can be used in various ways to achieve different looping behaviors. Here are the main types of loops you can create using `for`:

| Loop Type           | Syntax Example                            | Use Case                                      |
| :------------------ | :---------------------------------------- | :-------------------------------------------- |
| **Range Loop**      | `for index, value := range collection {}` | Iterate over arrays, slices, maps, or strings |
| **Traditional For** | `for i := 0; i < 10; i++ {}`              | Standard counting loop                        |
| **While Loop**      | `for condition {}`                        | Loop while a condition is true                |

You can use `break` and `continue` statements within these loops to control the flow as needed.

## 14. Testing in Go

Go has a built-in testing framework that allows you to write and run tests easily. Here are the main components of testing in Go:

- File Naming: Test files should be named with the `_test.go` suffix (e.g., `greetings_test.go`).
- Test Functions: Test functions should start with `Test` to be runnable by the `go test` command.
- Assertions: Use the `t.Errorf` method to report failures in tests.
- Running Tests: Use the `go test` command to run all tests in the package. with `-v` for verbose output.

## 15. Installing Go Binaries

Go allows you to install command-line tools (binaries) directly from source code repositories. Here's how you can do it:

```bash
go install example.com/user/tool
```

It'll download, compile, and place the binary in your `$GOPATH/bin` or `$HOME/go/bin` directory. Then you can run the tool from anywhere in your terminal.

## 16. `fmt.Printf` Cheat Sheet: All the Verbs

When using `fmt.Printf`, you need to use specific "verbs" (placeholders) for different data types.

### General (Use these 90% of the time)

| Verb      | Description                                    | Example Output                    |
| :-------- | :--------------------------------------------- | :-------------------------------- |
| **`%v`**  | **Default format** (works for any type)        | `Hello` / `123` / `[1 2 3]`       |
| **`%+v`** | **Struct with fields** (Crucial for debugging) | `{Name:Alice Age:30}`             |
| **`%#v`** | **Go Syntax** (Code to recreate the value)     | `main.User{Name:"Alice", Age:30}` |
| **`%T`**  | **Type** of the value                          | `int` / `string` / `main.User`    |
| **`%%`**  | Literal percent sign                           | `%`                               |

### Integers (Numbers)

| Verb     | Description                      | Example (`123`) |
| :------- | :------------------------------- | :-------------- |
| **`%d`** | **Decimal** (Base 10) - Standard | `123`           |
| **`%b`** | Binary (Base 2)                  | `1111011`       |
| **`%x`** | Hexadecimal (Base 16)            | `7b`            |
| **`%o`** | Octal (Base 8)                   | `173`           |

### Floats (Decimals)

| Verb       | Description                              | Example (`123.456`) |
| :--------- | :--------------------------------------- | :------------------ |
| **`%f`**   | Decimal (default precision)              | `123.456000`        |
| **`%.2f`** | **Limited Precision** (2 decimal places) | `123.46`            |
| **`%e`**   | Scientific Notation                      | `1.234560e+02`      |

### Strings & Booleans

| Verb     | Description                                 | Example           |
| :------- | :------------------------------------------ | :---------------- |
| **`%s`** | **String** (Standard)                       | `Hello`           |
| **`%q`** | **Quoted String** (Great for empty strings) | `"Hello"` or `""` |
| **`%t`** | **Boolean** (true/false)                    | `true`            |

### Padding & Alignment (Making Tables)

| Verb       | Description                        | Example (`"Go"`) |
| :--------- | :--------------------------------- | :--------------- |
| **`%5s`**  | Right-align (pad left with spaces) | `Go`             |
| **`%-5s`** | Left-align (pad right with spaces) | `Go`             |
| **`%05d`** | Pad number with Zeros (`12`)       | `00012`          |

## 16. Naked Returns (The "Magic" Return)

Go allows "Naked Returns" where you define the return variables in the function signature.

```go
func split(sum int) (x, y int) { // x and y are initialized to 0 here
    x = sum * 4 / 9
    y = sum - x
    return // Automatically returns x and y
}

```

It is best idea to not used it at all, because it can make code less readable.

## 17. Variable Shadowing

In Go, you can declare a new variable with the same name as an existing one within a new scope (block). This is called **Shadowing**.

```go
func main() {
    n := 10
    if n > 5 {
        n := 0 // This is a NEW variable, valid only inside this {} block
        fmt.Println(n) // Prints 0
    }
    fmt.Println(n) // Prints 10 (Original variable is untouched)
}
```

**Use Cases:**

1. **Type Conversion:** `data := []byte(data)` (Transforming a variable while keeping the name).
1. **Safety:** protecting an outer variable from being mutated inside a loop or goroutine (pre-Go 1.22).

## 18. Truthy and Falsy Values

Go does not have "truthy" or "falsy" values like JavaScript or Python. Every condition must explicitly evaluate to a boolean (`true` or `false`).

| Type          | Node/Python Style | Go Style (Explicit)                                                     |
| :------------ | :---------------- | :---------------------------------------------------------------------- |
| **String**    | `if (str)`        | `if str != ""` (or `len(str) > 0`)                                      |
| **Integer**   | `if (num)`        | `if num != 0`                                                           |
| **Boolean**   | `if (bool)`       | `if bool` (Same\!)                                                      |
| **Pointer**   | `if (ptr)`        | `if ptr != nil`                                                         |
| **Slice/Map** | `if (arr)`        | `if len(arr) > 0` (Check contents) or `if arr != nil` (Check existence) |
| **Error**     | `if (err)`        | `if err != nil`                                                         |

## 16. If with a Short Statement

Go allows you to execute a short statement **before** the condition in an `if` block. It's available only within the scope of that `if` block and any associated `else` blocks.

### The Syntax

`if statement; condition { ... }`

```go
import "math"

func pow(x, n, lim float64) float64 {
    // "v" is initialized right here
    if v := math.Pow(x, n); v < lim {
        return v
    } else  {
      // "v" is also available here
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // "v" is NOT available here. It died at the closing brace } above.
    return lim
}
```

## 17. Switch Statements: Safe by Default

Go's `switch` statement works differently than JavaScript's or C's. It is safer and more powerful.

### 1. Automatic Break (No Fallthrough)

In Node.js, if you forget `break`, code "falls through" to the next case.
In Go, **`break` is automatic**. Code stops after a match.

```go
name := "Emilia"
switch name {
case "Emilia":
    fmt.Println("Best Girl")
    // STOPS HERE. Does not run the next case.
case "Rem":
    fmt.Println("Who?")
}
```

### 1. The `fallthrough` Keyword

If you _want_ the behavior of running the next case (ignoring the next condition\!), you must explicitly ask for it.

```go
switch n {
case 1:
    fmt.Println("One")
    fallthrough // Forcefully runs "case 2" code without checking!
case 2:
    fmt.Println("Two")
}
```

### 1. Multiple Cases & Logic

You can list multiple values with commas, or use a switch without a variable to replace `if/else` chains.

```go
// Multiple matches
case "Saturday", "Sunday":
    fmt.Println("Weekend")

// Logic replacement (cleaner than if/else if)
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Morning")
case t.Hour() < 17:
    fmt.Println("Afternoon")
default:
    fmt.Println("Evening")
}
```

## 18. Defer: The Cleanup Scheduler

Go uses defer instead of try/finally. It schedules a function call to run immediately before the surrounding function returns.

### 1. Basic Usage

```go
func main() {
defer fmt.Println("World") // Runs LAST
fmt.Println("Hello") // Runs FIRST
}
// Output: Hello World
```

### 1. The Real Use Case (Cleanup)

Open a resource and immediately defer closing it. This guarantees cleanup even if errors occur later.

```go
file, err := os.Open("data.txt")
if err != nil { return err }

// ⚠️ CRITICAL: Only defer AFTER checking for error.
// If os.Open fails, 'file' is nil. Calling nil.Close() would panic.
defer file.Close()

// Read file...
// Parse lines...
// Return result...
```

### 1. LIFO Order (Stack)

Multiple defers are executed in Last-In, First-Out order (like a stack of plates).

```go
defer fmt.Println("First") // Runs 3rd
defer fmt.Println("Second") // Runs 2nd
defer fmt.Println("Third") // Runs 1st
// Output: Third, Second, First
```

## 19. Pointers: References to Values

Pointers in Go are variables that store the memory address of another variable. They allow you to reference and manipulate the original value directly.

### 1. Declaring Pointers

Create a pointer using the `*` operator and get the address of a variable using the `&` operator.

```go
var p *int        // Declare a pointer to an int
i := 42           // An integer variable
p = &i           // Assign the address of i to p
```

### 2. Dereferencing Pointers

Basically how you get or set the value at the address a pointer is pointing to.

```go
fmt.Println(*p) // Dereference p to get the value of i (prints 42)
*p = 21         // Change the value at the address p points to
fmt.Println(i)  // Now i is 21
```

## 20. Structs: Public vs. Private Fields

Just like functions, the visibility of struct fields is controlled entirely by **Capitalization**.

| Case          | Visibility               | Description                                                    |
| :------------ | :----------------------- | :------------------------------------------------------------- |
| **Uppercase** | **Exported (Public)**    | Visible to **other packages**. Required for JSON/XML encoding. |
| **Lowercase** | **Unexported (Private)** | Visible **ONLY** inside the same package.                      |

## 21. Handling Optional Values (The `nil` vs. `0` Dilemma)

In JavaScript/Node, you often rely on values being "Falsy" (`0`, `null`, `undefined`, `false`) to check if data exists. In Go, this concept **does not exist**.

### The Problem: Zero Values

In Go, every variable **always** has a value. If you don't assign one, it gets the **Zero Value**.

| Type     | Zero Value | "Does it exist?"                          |
| :------- | :--------- | :---------------------------------------- |
| `int`    | `0`        | You don't know (Is it 0, or missing?)     |
| `bool`   | `false`    | You don't know (Is it false, or missing?) |
| `string` | `""`       | You don't know (Is it empty, or missing?) |

**The "Baby Age" Bug:**
If you use a plain `int` for Age, you cannot distinguish between a "Newborn" (0) and "User didn't tell us" (Default 0).

### The Solution: Pointers (`*int`)

To represent "Optional" or "Missing" data for primitives (numbers, booleans), you must use a **Pointer**.

- **`nil`** = Data is missing.
- **`ptr`** = Data exists (even if it is `0`).

```go
type User struct {
    Name string
    Age  *int // Pointer allows 'nil'
}
```

### The "Clumsy Dance" of Assignment

Because Go is strict, you cannot just pass a raw number into a pointer slot. You also cannot take the address of a literal number (`&10` is an error). You must perform a 2-step process:

**JavaScript:**

```javascript
const u = { age: 10 }; // Easy
```

**Go:**

```go
// Step 1: Create a variable to hold the value
myAge := 10

// Step 2: Point to that variable
u := User{
    Age: &myAge,
}
```

### Checking the Value (Safety)

You trade **Write Convenience** for **Read Safety**. You never accidentally treat `null` as `0`.

```go
if u.Age != nil {
    // We know the user provided an age!
    // Now we extract the value using '*'
    fmt.Println("Real Age:", *u.Age)
} else {
    fmt.Println("User declined to answer")
}
```

**Summary:**

- Use **`int`** when the value is required and `0` is a valid number.
- Use **`*int`** when you strictly need to know the difference between "Zero" and "Nothing."

## 22. Arrays vs. Slices: The "Fixed" vs. "Dynamic"

Coming from JavaScript/Python, you usually only have "Lists." Go splits this into two concepts.

### 1. Arrays (`[Size]Type`)

A solid block of memory with a **Fixed Size**.

- **Syntax:** `[3]int{1, 2, 3}` (Number in brackets).
- **Behavior:** It is a **Value Type**. Assigning it to a new variable **copies the whole array**.
- **Use Case:** Rarely used directly (mostly for memory optimization).

### 2. Slices (`[]Type`)

A "Window" looking at an underlying array.

- **Syntax:** `[]int{1, 2, 3}` (Empty brackets).
- **Behavior:** It is a **Reference Type** (sort of). Assigning it copies the **Pointer**, not the data.
- **Use Case:** Used 99% of the time in Go.

| Operation      | Array `[3]int`        | Slice `[]int`            |
| :------------- | :-------------------- | :----------------------- |
| **Resize**     | ❌ Impossible         | ✅ Dynamic (`append`)    |
| **Assignment** | 🐢 Copies Data (Slow) | 🐇 Copies Pointer (Fast) |

## 23. Slice Internals: Length vs. Capacity

A slice is just a small struct with 3 fields: **Pointer**, **Length**, and **Capacity**.

```go
s := []int{2, 3, 5, 7, 11, 13}
s = s[:0] // Length is 0, but Data is still there!
s = s[:4] // We can "resurrect" the data because Cap is still 6.
```

### The Rules of Slicing

1. **Slicing the End (`s[:4]`)**:
   - **Reversible.** You are just "closing the curtain."
   - **Capacity** stays the same.
2. **Slicing the Start (`s[2:]`)**:
   - **Destructive.** You move the pointer forward.
   - **Capacity shrinks.** You cannot look back to the left.

## 24. The "Nil" Slice (No Crash!)

In Node/Python, `null` lists cause crashes. In Go, a `nil` slice is useful.

```go
var s []int // s is nil (Pointer is nil)
```

- **It behaves like an empty list:**
  - `len(s)` returns `0` (Safe).
  - `for range s` runs 0 times (Safe).
  - `append(s, 1)` works (Safe - automatically creates the array).

**Note on Printing:** `fmt.Println(s)` will print `[]` (a convenient lie). To see if it's truly nil, use `fmt.Printf("%#v", s)` which prints `[]int(nil)`.

## 25. Maps: The "Read Safe, Write Panic" Trap

Maps in Go are similar to Objects/Dictionaries in other languages, but they have a unique relationship with `nil`.

- **Declaration:** `var m map[string]int` (Defaults to `nil`).
- **Initialization:** `m = make(map[string]int)` (Ready to use).

### The Golden Rule of Maps

Just like Slices, a `nil` map has no memory allocated. However, the behavior differs for reading vs. writing.

| Operation  | On `nil` Map                 | On Initialized Map |
| :--------- | :--------------------------- | :----------------- |
| **Read**   | ✅ Safe (Returns Zero Value) | ✅ Returns value   |
| **Write**  | ❌ **PANIC** (Crash)         | ✅ Works           |
| **Delete** | ✅ Safe (No-op)              | ✅ Works           |

**The "Write" Crash:**

```go
var m map[string]int // m is nil
// fmt.Println(m["key"]) // ✅ Prints 0 (Safe!)
// m["key"] = 10         // ❌ PANIC! assignment to entry in nil map
```

**Fix:** Always use `make()` or a literal `map[string]int{}` before writing.

## 26. Cheat Sheet: Zero Values vs. Nil

In Go, every variable has a default value the moment it is declared. It is critical to know which ones start as "Empty Boxes" (Zero Value) and which ones start as "Missing Boxes" (Nil).

### 1. Value Types (Default to "Zero Value")

These exist immediately. You can use them safely right away.

| Type                | Default Value        | Example        |
| :------------------ | :------------------- | :------------- |
| **Integer / Float** | `0`                  | `var i int`    |
| **Boolean**         | `false`              | `var b bool`   |
| **String**          | `""` (Empty string)  | `var s string` |
| **Array**           | `[0, 0...]`          | `var a [3]int` |
| **Struct**          | `{}` (Fields zeroed) | `var u User`   |

### 2. Reference Types (Default to `nil`)

These are just pointers/labels. They don't point to anything yet.

| Type          | Default Value | Danger Zone ⚠️                         |
| :------------ | :------------ | :------------------------------------- |
| **Slice**     | `nil`         | Safe to read/append.                   |
| **Map**       | `nil`         | **CRASH** if you write.                |
| **Pointer**   | `nil`         | **CRASH** if you dereference (`*p`).   |
| **Interface** | `nil`         | **CRASH** if you call a method.        |
| **Function**  | `nil`         | **CRASH** if you call it.              |
| **Channel**   | `nil`         | **BLOCKS** forever (deadlock) if used. |

**Mental Model:**

- **Zero Value:** "I have a box, but it's empty."
- **Nil:** "I have a label, but no box."

Here is a clean, formatted section ready to copy-paste directly into your documentation.

## 27. Reference Types & Nullability

In Go, **Reference Types** (Slices, Maps, Channels, Interfaces) are inherently nullable. You generally **should not** use a pointer (e.g., `*[]string` or `*map[string]int`) to make them optional.

### Why? (Under the Hood)

Reference types are already "descriptors" or pointers to underlying data structures. Adding a pointer to them creates unnecessary indirection (a pointer to a pointer).

- **Maps & Channels:** act as pure pointers. Their zero value is `nil`.
- **Slices:** act as a tiny struct containing a pointer to an array. If that internal pointer is missing, the slice is considered `nil`.

### The "Do I Need a Pointer?" Cheat Sheet

| Type Category  | Examples               | Zero Value   | Needs `*` for Nullable? | Best Practice                                                     |
| -------------- | ---------------------- | ------------ | ----------------------- | ----------------------------------------------------------------- |
| **Primitives** | `int`, `float`, `bool` | `0`, `false` | **YES**                 | Use `*int` to distinguish `0` from `nil`.                         |
| **Structs**    | `User`, `Config`       | `{}` (Empty) | **YES**                 | Use `*User` to avoid creating empty objects.                      |
| **Strings**    | `string`               | `""`         | **YES**                 | Use `*string` only if `""` is a valid value different from `nil`. |
| **Slices**     | `[]T`                  | `nil`        | **NO** 🚫               | `*[]T` is redundant and awkward to use.                           |
| **Maps**       | `map[K]V`              | `nil`        | **NO** 🚫               | `*map` is a code smell.                                           |

## 28. File & Directory Naming Conventions

Go is opinionated about naming. Here is the standard way to name your files and folders.

| Type                | Convention                    | Example                          | Why?                                                                            |
| :------------------ | :---------------------------- | :------------------------------- | :------------------------------------------------------------------------------ |
| **Source Files**    | `snake_case`                  | `user_profile.go`, `api_test.go` | Required for `_test.go` and `_windows.go` build tags to work.                   |
| **Package Folders** | `snake_case` (or single word) | `json_parser/`, `handlers/`      | Package name must match folder name. `kebab-case` is invalid in `package name`. |
| **Project Root**    | `kebab-case`                  | `github.com/user/learn-go`       | Standard for repositories.                                                      |

## 29. Methods are just Functions (Under the Hood)

Mathematically, a method is just a function where the first argument (the "receiver") is moved to a special position before the function name.

- **Syntax:** `func (v Vertex) Abs()`
- **Reality:** The compiler sees `func Vertex_Abs(v Vertex)`.
- **Why use them?** They enable "Dot Syntax" (`v.Abs()`) and are required to satisfy Interfaces.

## 30. Pointer Receiver vs. Value Receiver

When defining a method, you must choose how the data is passed:

- **Pointer Receiver `(v *Vertex)`:** Use this if you need to **modify** the data or if the struct is **large** (to avoid copying). This is the most common default.
- **Value Receiver `(v Vertex)`:** Use this for **read-only** behavior on small, primitive-like data (e.g., Coordinates, Time). You get a copy, so changes inside the method are lost.

## 31. The "Consistency" Rule (Google Style)

If **any** method on a struct needs a pointer receiver (for mutation), then **ALL** methods on that struct should use pointer receivers.

- **Why?** It prevents confusion and ensures the type always behaves like a reference.
- **Tip:** When in doubt, default to **Pointer (`*`)**.

## 32. Method Syntactic Sugar (The "Auto-Fix")

Go methods are flexible, whereas standard functions are strict.

- **Auto-Reference:** If you call a Pointer Method on a Value Variable (`v.Scale()`), Go automatically injects `&v`.
- **Auto-Dereference:** If you call a Value Method on a Pointer Variable (`p.Abs()`), Go automatically injects `*p`.
- **Contrast:** Standalone functions (e.g., `ScaleFunc(v)`) will **crash** if the types don't match exactly.

## 33. Extending Non-Struct Types

You can attach methods to almost any type, not just structs, by creating a **Type Definition**.

- **Example:** `type MyFloat float64` allows you to write `func (f MyFloat) Abs()`.
- **Use Case:** Adding logic to basic values, like formatted printing for status codes (Enums).

## 34. The "Type Ownership" Rule

You can only define a method for a type that is defined in the **same package**.

- **Allowed:** Defining `MyUser` in `main` and adding methods to it.
- **Forbidden:** You cannot add methods to standard types (like `int`) or types from other libraries (like `time.Time`) directly. You must "wrap" or "alias" them first.

## 35. Interfaces: Implicit Implementation

In Java or TypeScript, you explicitly say `class User implements Stringer`. In Go, you don't.

**If you have the methods, you are the interface.**

- **Definition:** An interface is just a set of method signatures.
- **Implementation:** A type implements an interface by implementing its methods. There is no `implements` keyword.

```go
type Abser interface {
    Abs() float64
}

type MyFloat float64

// This method means MyFloat AUTOMATICALLY implements Abser
func (f MyFloat) Abs() float64 {
    return float64(f)
}
```

## 36. The Interface "Tuple" & The Nil Trap

This is the most dangerous concept for beginners. An interface value is effectively a tuple of **(Type, Value)**.

### The "Two Nils" Problem

There is a huge difference between an **Empty Interface** and an **Interface holding a Nil Pointer**.

| Scenario            | Code       | Internal Tuple `(Type, Value)` | Call `i.Method()` | Why?                                                        |
| ------------------- | ---------- | ------------------------------ | ----------------- | ----------------------------------------------------------- |
| **1. Unassigned**   | `var i I`  | `(nil, nil)`                   | **💥 CRASH**      | No Type info. Go doesn't know which function to call.       |
| **2. Assigned Nil** | `var t *T` | `(*T, nil)`                    | **✅ WORKS**      | Type is `*T`. Go finds the function and passes `nil` to it. |

### Why Scenario 2 Works (The "Manual on the Wall")

In Go, methods are not attached to the object (like JS). They are attached to the **Type**.
When you call `i.Method()` in Scenario 2:

1. Go sees the type `*T`.
2. Go looks up the "Instruction Manual" (code) for `*T`.
3. Go runs the function, passing `nil` as the argument.

**Rule:** It is safe to call a method on a `nil` pointer, as long as the method itself handles the `nil` safely (e.g., `if t == nil { return }`).

## 37. The Empty Interface (`any`)

The interface with zero methods is called the **Empty Interface**: `interface{}`.
Since Go 1.18, it has an alias: **`any`**.

- **Logic:** Every type has at least zero methods. Therefore, `any` can hold **anything**.
- **The Trap:** It works like `any` in TypeScript, but worse. Once you put data in, **you lose all type info**. You cannot call methods or do math until you pull it back out.

```go
var i any = "hello"
// i.len() // ❌ Error! 'any' has no methods.
```

**Use Case:** Only use `any` when you truly don't know the data structure (e.g., `fmt.Println`, JSON parsing).

## 38. Type Assertions: Getting Data Out

To get the value back out of an interface, you must use a **Type Assertion**.

`t := i.(T)`

### The "Comma OK" Idiom (Safety First)

This is a special syntax that changes behavior based on how many variables you assign.

| Syntax          | Code                  | Behavior on Mismatch                       |
| --------------- | --------------------- | ------------------------------------------ |
| **1 Variable**  | `s := i.(string)`     | **💥 PANIC!** (App crashes immediately)    |
| **2 Variables** | `s, ok := i.(string)` | **✅ SAFE.** `ok` is `false`, `s` is `""`. |

### The "Magic" of Safety (Why it doesn't crash)

When you use the 2-variable form, Go **guarantees** that the value variable (`s`) will be populated with the **Zero Value** of that type if the assertion fails.

- **If Match:** `s` = "Actual String", `ok` = `true`
- **If Fail:** `s` = `""` (Empty String), `ok` = `false`

Because the variable is always valid (never "undefined"), the program continues safely without panic.

```go
var i any = 123

s, ok := i.(string)
if !ok {
    // 's' is now safely "", so we don't crash.
    fmt.Println("This is not a string!")
}
```

## 39. Type Switches

If you need to check multiple types, don't chain `if/else` assertions. Use a **Type Switch**.

It gives you a special variable `v` that **morphs** into the correct type inside each case.

```go
func do(i any) {
    switch v := i.(type) {
    case int:
        // v is an INT here. We can do math.
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        // v is a STRING here. We can check length.
        fmt.Printf("String length: %v\n", len(v))
    default:
        // v is still ANY (interface{}) here.
        // We know nothing about it except it's not an int or string.
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

**Quirk:** In the `default` case, `v` maintains the original value but remains type `interface{}`. You cannot use type-specific methods on it yet.

## 40. Stringers: The `toString()` of Go

In JavaScript, every object inherits `.toString()`. In Go, you opt-in by implementing the `fmt.Stringer` interface.

```go
type Stringer interface {
    String() string
}
```

If you define a `String()` method, `fmt.Println` will automatically call it instead of printing the raw struct.

### The Infinite Loop Trap ⚠️

When implementing `String()`, do not pass the object itself back into `fmt.Sprintf` with `%v`, or it will call `String()` again forever.

**Bad:**

```go
func (p Person) String() string {
    // 💥 CRASH: Sprintf calls p.String(), which calls Sprintf...
    return fmt.Sprintf("Person: %v", p)
}
```

**Good:**

```go
func (p Person) String() string {
    // Safe: We only print specific fields
    return fmt.Sprintf("Person: %s (%d)", p.Name, p.Age)
}
```

## 41. Errors: Values, Not Exceptions

This is the biggest culture shock coming from TypeScript.

- **Node/TS:** You `throw` an Error (Grenade) 💣. It blows up the stack until a `catch` block catches it.
- **Go:** You `return` an Error (Rock) 🪨. The function hands it to you as a normal return value. You must check it immediately.

### Extracting Error Data (Type Assertion)

Since `error` is just an interface, it's a "black box." To see if it contains custom data (like a timestamp or code), you must cast it.

**TypeScript:**

```typescript
try {
  run();
} catch (err) {
  if (err instanceof MyError) {
    console.log(err.when);
  }
}
```

**Go:**

```go
err := run()
if err != nil {
    // Check: "Is this error actually a *MyError pointer?"
    if e, ok := err.(*MyError); ok {
        fmt.Println(e.When) // ✅ Access fields safely
    }
}
```

## 42. Readers: The "Bucket" Metaphor

The `io.Reader` interface is fundamental to how Go handles I/O (Files, HTTP, Streams).

```go
Read(p []byte) (n int, err error)
```

### The difference from `fs.readFile`

- **Node (`fs.readFile`):** Loads the **entire** 50GB file into RAM, then gives it to you.
- **Go (`Reader`):** You give it a small "Bucket" (e.g., 32KB). It fills the bucket, hands it back, and waits for you to ask for more.

**The Loop Logic:**

1. You hand the Reader an empty slice `p` (The Bucket).
2. Reader fills it with data.
3. Reader tells you `n` (how many bytes it poured).
4. You process those `n` bytes.
5. Repeat until `err == io.EOF`.

## 43. The Middleware Pattern (Rot13)

A common Go pattern is wrapping one Reader inside another to modify the stream on the fly. This is how **Gzip** (Compression) and **TLS** (Encryption) work.

**The Flow:**
`Your Code` <— `Rot13Reader` <— `StringsReader`

1. **`io.Copy`** (The Truck Driver): Asks `Rot13Reader` for data.
2. **`Rot13Reader`**: Asks `StringsReader` for data.
3. **`StringsReader`**: "Here is encrypted text: 'Lbh...'".
4. **`Rot13Reader`**: "Hold on." (Modifies the bytes in place: 'L' -> 'Y'). "Okay, here is 'You...'".
5. **`io.Copy`**: Delivers the clean data to `os.Stdout`.

**Key Concept:** `io.Copy` never looks at the data. It just moves bytes. The transformation happens in the middle.

## 44. Images: Paint by Numbers

In Go, an Image isn't necessarily a file in memory. It is any type that can answer three questions:

1. **`Bounds()`**: How big is the canvas?
2. **`ColorModel()`**: What colors do you use?
3. **`At(x, y)`**: **"What color is at this pixel?"**

This allows you to create **Procedural Images** (like Fractals) that take up zero RAM because the pixels are calculated using math (e.g., `x ^ y`) only when they are requested.

## 45. Design Pattern Quirk: Methods vs. Config Objects

Coming from JavaScript, you might be tempted to pass "Configuration Objects" with functions inside them.

**The JS Pattern (Flexible but Memory Heavy):**

```javascript
// A new closure is created for every instance!
const myImage = {
  at: (x, y) => 'blue'
};
```

**The Go Pattern (Strict but Efficient):**

```go
// The method code exists ONCE in memory.
// All 10,000 instances share the same instruction set.
func (i Image) At(x, y int) color.Color { ... }
```

### The Rule of Thumb

- **Use Structs with Function Fields** (The JS way) for **Configuration** or **One-off logic** (e.g., `http.Server` error handler).
- **Use Interfaces & Methods** (The Go way) for **Core Data Types** (e.g., `Image`, `Reader`, `User`). This separates **Data** (Struct) from **Behavior** (Methods).
