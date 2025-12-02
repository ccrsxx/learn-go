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
