# Learn Go

## Go vs. Node/Python: The Import & Module Paradigm Shift

Coming from Node.js (TypeScript) or Python, the way Go handles imports and modules feels like a "Culture Shock." This document summarizes the key differences I learned during my time learning Go.

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

### Why "example/test" is a trap

If you initialize your module as `go mod init example/test`:

1. It works locally because Go treats it as a namespace.
2. **It fails remotely** because `example/test` is not a valid URL.
3. If you push to GitHub later, you have to **Find & Replace** every single import string in your entire codebase to match the real URL.
4. **Rule:** Always name your module the actual URL where it will live (e.g., `github.com/username/project`).

## 3. How `go.mod` "Routes" Imports

If I import my own code:

```go
import "github.com/ccrsxx/learn-go/src/getting-started/greetings"
```

Why doesn't Go try to download this from the internet?

**The Magic:** `go.mod` acts as a local router.

1. Compiler sees the import path.
2. It checks `go.mod`: `module github.com/ccrsxx/learn-go`.
3. **Match\!** It realizes "This import starts with my username. I don't need to go to the internet; I will look in the local folder structure."

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
  2. **Per-Project Linking:** Your `go.mod` picks exactly which version from the cache to use.
  3. **Major Versions:** Go treats `v2` as a **different import path** (`github.com/lib/foo/v2`). This allows your app to use `v1` and `v2` of the same library simultaneously without conflict.

## 6. The "Package Main" Rule: Executable vs. Library

In Node.js, any file can be an entry point (`node utils.js` works fine). In Go, the entry point is strict.

### The Golden Rule

A Go program will **only compile as an executable** if:

1. The package name is **`package main`**.
2. It contains a **`func main()`**.

### The "One Package Per Folder" Constraint

Go does not allow you to mix package names in the same directory.

- **Node:** You can have `index.js` (app) and `helper.js` (module) in the same folder.
- **Go:** A folder can only be **one thing**:
  - **EITHER** an Executable (`package main`)
  - **OR** a Library (`package utils`)

**Common Project Structure to solve this:**

```text
/my-project
  /cmd/api       -> package main (The entry point)
  /internal/db   -> package db   (The library code)
```

## 7. String Formatting: Replacing Template Literals

In Node.js, you use backticks for string interpolation. Go uses **Printf verbs** (similar to C or Python's old `%` formatting).

### The Syntax Shift

- **Node.js:** `` `User ${name} has ID ${id}` ``
- **Go:** `fmt.Printf("User %s has ID %d\n", name, id)`

  ### 1. Simple Printing to Console

  Unlike `Println`, `Printf` does **not** add a new line automatically. You must add `\n`.

  ```go
  name := "Alice"
  count := 5

  // Node: console.log(`Hello ${name}, count: ${count}`)
  fmt.Printf("Hello %s, count: %d\n", name, count)
  ```

  ### 2. Better Formatting Control with Explicit Argument Indexes

  Using argument indexes, you can reuse arguments without passing them multiple times.

  ```go
  user := "Alice"
  role := "Admin"

  // "Alice is a Admin. Goodbye Alice."
  // We use user (1st arg) twice without passing it twice!
  fmt.Printf("%[1]s is a %[2]s. Goodbye %[1]s.\n", user, role)
  ```

  ### 3. Using slog for Structured Logging

  Go's `slog` package allows for structured logging, which is more powerful than simple string formatting. You can log key-value pairs that can be easily parsed and analyzed.

  ```go
  import (
    "log/slog"
  )

  // Example of structured logging
  // User login user="Alice" id=101 active=true
  slog.Info("User login", "user", "Alice", "id", 101, "active", true)
  ```
