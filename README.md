# Thaibaht-converter

This project is a Go-based solution for converting decimal currency values into Thai Baht text format. It is designed with a focus on accuracy, reusability, and professional project structure, intended for integration into financial services.

## Features
- **Accurate Conversion**: Converts `shopspring/decimal` values to Thai text without floating-point errors.
- **Standard Compliant**: Correctly handles "บาทถ้วน" (Baht Thuan) for whole numbers and "สตางค์" (Satang) for decimals.
- **Production Ready**: Organizes logic into an internal package to prevent unintended external access and ensure maintainability.

## Project Structure
The project follows a standard Go layout:
- `/internal/convert`: Contains the core conversion logic (`convertThaibaht.go`).
- `main.go`: Entry point demonstrating the converter with example cases.
- `Dockerfile`: Configuration for containerized execution.
- `go.mod` & `go.sum`: Dependency management.

## Prerequisites
- **Go**: Version 1.22 or higher.
- **Docker**: Optional, for running the application in a container.

## How to Run

### Using Local Go
1. Clone the repository and navigate to the project folder.
2. Run the application:
```bash
go run main.go