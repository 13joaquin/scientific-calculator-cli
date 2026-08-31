# 🧮 Scientific Calculator CLI

An educational scientific calculator written in Go, designed to demonstrate and teach **clean architecture**, **modular design**, **testing practices**, and **terminal user interface (TUI) development**.

## 📚 Purpose

This project is a learning resource for developers who want to understand:

- **Clean Architecture**: Separation of concerns with organized layers (domain, application, infrastructure)
- **Modular Design**: Building extensible, reusable components
- **Testing Best Practices**: Unit tests, integration tests, and test-driven development (TDD)
- **Terminal User Interface (TUI)**: Creating interactive CLI applications with Go
- **Go Development Patterns**: Idiomatic Go code organization and conventions

## ✨ Features

- Basic arithmetic operations (+, -, *, /, %)
- Scientific functions (trigonometry, logarithms, exponentials)
- Extensible operation system for easy feature addition
- Clean, modular codebase following Go best practices
- Comprehensive test coverage
- Interactive terminal interface

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/13joaquin/scientific-calculator-cli.git
cd scientific-calculator-cli

# Build the project
go build -o calculator

# Run the calculator
./calculator
```

### Usage

```bash
./calculator
```

The calculator will open an interactive terminal interface where you can:
- Enter mathematical expressions
- Use standard operators: `+`, `-`, `*`, `/`, `%`
- Apply scientific functions: `sin()`, `cos()`, `tan()`, `log()`, `exp()`, etc.
- View calculation history
- Exit with `quit` or `exit`

## 📁 Project Structure

```
scientific-calculator-cli/
├── main.go              # Application entry point
├── internal/            # Private application code
│   ├── domain/         # Core business logic
│   ├── application/    # Use cases and orchestration
│   ├── infrastructure/ # External dependencies (I/O, storage)
│   └── ui/             # Terminal user interface
├── go.mod              # Go module definition
├── README.md           # This file
└── LICENSE             # MIT License
```

## 🏗️ Architecture

This project follows **Clean Architecture** principles:

- **Domain Layer**: Core calculator logic, operation definitions
- **Application Layer**: Calculator orchestration, operation coordination
- **Infrastructure Layer**: I/O, parsing, TUI rendering
- **Presentation Layer**: User interface and input handling

This separation enables:
- Easy testing of business logic without UI dependencies
- Simple component reuse across different interfaces
- Clear responsibility boundaries
- Framework/library independence

## 🧪 Testing

The project includes comprehensive tests demonstrating:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

## 🛠️ Extensibility

Adding new operations is straightforward:

1. Define the operation in the domain layer
2. Register it in the application layer
3. Add tests
4. Update the UI (if needed)

See `internal/domain` for examples.

## 📝 Learning Resources

This project is ideal for:

- **Go Beginners**: Understanding Go project structure and idioms
- **Clean Architecture Students**: Practical implementation of architectural principles
- **TUI Development**: Building interactive CLI applications
- **Test-Driven Development**: Exploring testing patterns and practices

## 🤝 Contributing

This is an educational project. Contributions that enhance learning value are welcome:

- Improved documentation
- Additional test examples
- New scientific functions
- Better architectural patterns
- Performance optimizations

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎯 Roadmap

- [ ] Advanced mathematical functions (matrices, complex numbers)
- [ ] Command-line argument parsing for batch operations
- [ ] Configuration file support
- [ ] Expression history and variables
- [ ] Multiple precision arithmetic

## 👤 Author

**Joaquín** - [GitHub Profile](https://github.com/13joaquin)

---

**Happy Learning!** 🚀 Feel free to explore, modify, and learn from this codebase.
