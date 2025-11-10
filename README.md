# Go Key-Value Store

A robust key-value store with multiple backend options and interceptor support, built to master Go's interfaces, error handling, and concurrency patterns.

---

## 🎯 Project Goal

This project demonstrates a clean interface-based architecture in Go, showcasing how to build extensible systems with proper separation of concerns. It has evolved from a simple learning exercise into a flexible component that can adapt to different storage needs.

---

## 🧠 Core Concepts

This implementation demonstrates several key Go patterns for building extensible software:

- **Interface-Based Design:** The `Storage` interface allows for multiple backend implementations (JSON, XML, in-memory) that can be swapped without changing application code.
- **Composition over Inheritance:** The persistent stores embed the in-memory implementation to reuse core functionality.
- **Interceptor Pattern:** Cross-cutting concerns like logging are handled through interceptors that can observe storage operations.
- **Concurrency Safety:** All access to the in-memory map is protected by a `sync.RWMutex`, allowing for multiple concurrent readers or a single exclusive writer.
- **Factory Pattern:** The `New` function creates appropriate storage instances based on configuration.

---

## 🚀 Setup

1. **Prerequisites:** Ensure you have Go (version 1.19 or later) installed.
2. **Clone the repository:**
   ```bash
   git clone https://github.com/EronAlves1996/go-key-value-store.git
   cd go-key-value-store
   ```

---

## 📖 Usage

Running the application will execute the test logic in `main.go`, which demonstrates setting and getting values with different storage backends.

```bash
go run main.go
```

This will:

1. Create a storage file in the project directory (JSON or XML based on configuration).
2. Set a few key-value pairs.
3. Retrieve and print values to the console.
4. Log operations through the interceptor.

### Storage Backends

The key-value store supports three storage backends:

- **In-Memory Storage:** Non-persistent storage that exists only during runtime.
- **JSON Storage:** Persistent storage using JSON format.
- **XML Storage:** Persistent storage using XML format.

### Interceptors

Interceptors allow you to observe and potentially modify storage operations. The example includes a logging interceptor that prints details about each operation.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
