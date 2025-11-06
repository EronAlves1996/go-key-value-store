# Go Key-Value Store

A robust in-memory key-value store with a persistent JSON backend, built to master Go's error handling and concurrency patterns from the ground up.

---

## 🎯 Project Goal

This project is a practical learning exercise that has evolved into a fully-functional component. The primary objective is to implement a concrete application showcasing Go's error handling philosophy, progressing from basic checks to robust, atomic operations. It serves as a study in building reliable concurrent systems.

---

## 🧠 Core Concepts

This implementation demonstrates several key Go patterns for building reliable software:

- **Atomic Operations:** The `Set` method is atomic. If persisting the data to disk fails, the in-memory state is automatically rolled back to its original value, preventing partial updates.
- **Error Wrapping:** Errors are wrapped with context using `fmt.Errorf`, which preserves the original error message and provides richer context for debugging.
- **Concurrency Safety:** All access to the in-memory map is protected by a `sync.RWMutex`, allowing for multiple concurrent readers or a single exclusive writer.
- **Persistent JSON Backend:** The state of the store is saved to a `store.json` file, ensuring data survives application restarts.

---

## 🚀 Setup

1.  **Prerequisites:** Ensure you have Go (version 1.25.3 or later) installed.
2.  **Clone the repository:**
    ```bash
    git clone https://github.com/EronAlves1996/go-key-value-store.git
    cd go-key-value-store
    ```
    ```

    ```

---

## 📖 Usage

Running the application will execute the test logic in `main.go`, which demonstrates setting and getting values.

```bash
go run main.go
```

This will:

1.  Create a `store.json` file in the project directory.
2.  Set a few key-value pairs.
3.  Retrieve and print a value to the console.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
