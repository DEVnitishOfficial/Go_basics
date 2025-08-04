

# GoModDemo Explained : How to access files which lies in different packages

This project demonstrates how to set up and work with a **multi-file, multi-package Go project** using **Go Modules**.

---

## 📁 Project Structure

```
GoModDemo/
├── go.mod               # Module definition
├── main.go              # Entry point of the application
└── app/
    └── complex.go       # Contains the 'Complex' struct and related logic
```

---

## 🚀 Getting Started with Go Modules

### 1️⃣ Initialize the Module

Navigate to your project root and run:

```bash
go mod init GoModDemo
```

This creates a `go.mod` file which declares the **module path** (e.g., `GoModDemo`).
This path becomes the **import root** for your project’s internal packages.

---

### 2️⃣ Organize Code into Packages

* Place related logic in subdirectories (e.g., `app/`).
* Each subdirectory should declare its own package at the top of the `.go` files, like:

```go
package app
```

This allows Go to treat each folder as a separate **package**.

---

### 3️⃣ Import Local Packages

In `main.go`, import your local packages using the **module path** from `go.mod`, like this:

```go
import "GoModDemo/app"
```

Go resolves `GoModDemo/app` by looking for the `app` package inside the module directory.

✅ **Note:** Always use the exact module path defined in your `go.mod`.

---

### 4️⃣ Build and Run the Project

Always build or run your code from the **module root directory** (where `go.mod` is located):

```bash
go run .
# or
go run main.go
```

This ensures Go correctly resolves all local package imports.

---

## 🔍 How Go Resolves Paths and Dependencies

* **Module Path:** Defined in `go.mod` (e.g., `GoModDemo`) — used as the root path for imports.
* **Local Imports:** Use `module_path/package_name` to import your own code.
* **External Dependencies:** Automatically added to `go.mod` and `go.sum` when used, or when you run:

```bash
go mod tidy
```

---

## 💡 Example Summary

* `main.go` imports the local `app` package as `GoModDemo/app`.
* The `Complex` struct and related functions (like `NewComplexNumber`, `Add`, `Multiply`, etc.) are defined in `app/complex.go`.
* You can create, operate on, and display complex numbers in your `main.go`.




# Next Goal : How to acces files which present in the same package

suppose you have a file name `lib.go`, how will you run this file?

To run this file alogn with main.go we can use the following command:

```bash
go run main.go lib.go
```

or

```bash
go run *.go // use this command only in the macos or linux
```

for windows you can use the following command:

```bash
go run .
```

