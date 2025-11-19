# Go Key-Value Store

A robust key-value store with multiple backend options and interceptor support, built to master Go's interfaces, structured error handling, and concurrency patterns.

---

## 🎯 Project Goal

This project demonstrates a clean interface-based architecture in Go, showcasing how to build extensible systems with proper separation of concerns. It has evolved from a simple learning exercise into a flexible component that can adapt to different storage needs and demonstrates modern Go 1.13+ error handling patterns.

---

## 🧠 Core Concepts

This implementation demonstrates several key Go patterns for building extensible software:

- **Interface-Based Design:** Use the `Storage` interface to swap between JSON, XML, and in-memory backends without changing application code.
- **Composition over Inheritance:** Persistent stores embed the in-memory implementation to reuse core functionality.
- **Structured Error Handling:** Custom error types are used with `errors.Is` and `errors.As` for precise error checking and handling.
- **Interceptor Pattern:** Cross-cutting concerns like logging are handled through interceptors that can observe storage operations.
- **Concurrency Safety:** All access to the in-memory map is protected by a `sync.RWMutex`, allowing for multiple concurrent readers or a single exclusive writer.
- **Factory Pattern:** The `New` function creates appropriate storage instances based on configuration.

---

## 🚀 Setup

1.  **Prerequisites:** Ensure you have Go (version 1.19 or later) installed.
2.  **Clone the repository:**
    ```bash
    git clone https://github.com/EronAlves1996/go-key-value-store.git
    cd go-key-value-store
    ```

---

## 📖 Usage

Running the application will execute the test logic in `main.go`, which demonstrates setting and getting values with different storage backends and handling potential errors.

```bash
go run main.go
```

This will:

1.  Create a storage file in the project directory (JSON or XML based on configuration).
2.  Set a few valid key-value pairs.
3.  Attempt to set a value with an invalid key and handle the resulting custom error.
4.  Retrieve and print values to the console.
5.  Log operations through the interceptor.

### Error Handling Example

The application includes an example of robust error handling. When trying to set a value with an empty key, the store returns a custom `InvalidKeyError`. The main function then uses `errors.Is` to check the error type and `errors.As` to safely extract the specific error details.

```go
// From main.go
if err := kv.Set("", "some value"); err != nil {
    // Check if the error is of a specific type (or wraps it)
    if errors.Is(err, &internal_errors.InvalidKeyError{}) {
        fmt.Println("Caught an invalid key error!")
    }

    // Extract the custom error to access its fields
    var invalidKeyErr *internal_errors.InvalidKeyError
    if errors.As(err, &invalidKeyErr) {
        fmt.Printf("The problematic key was: '%s'\n", invalidKeyErr.Key)
    }
}
```

### Storage Backends

The key-value store supports three storage backends:

- **In-Memory Storage:** Non-persistent storage that exists only during runtime.
- **JSON Storage:** Persistent storage using JSON format.
- **XML Storage:** Persistent storage using XML format.

### Interceptors

Interceptors allow you to observe and potentially modify storage operations. The example includes a logging interceptor that prints details about each operation.

---

## 🗺️ Roadmap

Potential future improvements to further explore Go's patterns:

- **Error Wrapping:** Implement error wrapping with the `%w` verb in `fmt.Errorf` to add more context to errors while preserving the original cause.
- **Configuration Management:** Add support for selecting the storage backend via command-line flags or a configuration file, instead of hardcoding it.
- **Expanded Interceptor Support:** Implement additional interceptors for cross-cutting concerns like metrics collection (e.g., Prometheus), rate limiting, or caching.
- **Comprehensive Testing:** Add a dedicated `_test.go` file with table-driven tests to ensure the reliability of all storage backends and error conditions.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
