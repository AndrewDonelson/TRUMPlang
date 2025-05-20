# TRUMP Programming Language

A tremendously great programming language - the best language, believe me!

## Overview

TRUMP is a satirical programming language inspired by the rhetoric of former U.S. President Donald Trump. It combines humor with actual programming language features to create a fully functional, if unconventional, programming experience.

## Installation

### Prerequisites

- Go 1.22 or higher

### Building from Source

```bash
# Clone the repository
git clone https://github.com/AndrewDonelson/trumplang.git
cd trumplang

# Build the compiler
go build -o trumpc ./cmd/trumpc

# Add to PATH (optional)
# On Linux/macOS:
sudo cp trumpc /usr/local/bin/
# On Windows: Add the directory to your PATH environment variable
```

## Getting Started

### Creating a New Project

```bash
./trumpc create MyFirstProject
```

This creates a new directory with a template `main.trump` file and a README.

### Running a TRUMP Program

```bash
# Run a project directory (looks for main.trump)
./trumpc run MyFirstProject

# Or run a specific file
./trumpc run MyFirstProject/main.trump
```

### Inspecting a TRUMP Program

```bash
./trumpc inspect MyFirstProject/main.trump
```

This shows detailed information about the tokens and syntax of your program.

### Building a TRUMP Program

```bash
./trumpc build MyFirstProject/main.trump
```

This compiles the program to a `.djt` file.

## Language Syntax

### Variables

Variables are declared with `YUGE` or `TREMENDOUS`:

```
YUGE name = "Trump";
TREMENDOUS number = 45;
```

### Functions

Functions are defined with the `FUNCTION` keyword and can have ratings:

```
YUGE FUNCTION greet(name) RATED 10/10 {
    TWEET "Hello, " + name + "! You're doing a FANTASTIC job!";
    RETURN "Greeted " + name;
}
```

### Control Flow

#### If Statements

```
BUILD WALL IF (value == 45) {
    RALLY "That's the BEST number, BELIEVE ME!";
} ELSE {
    TWEET "Not a great number. SAD!";
}
```

#### While Loops

```
YUGE counter = 0;
MAKE DEALS WHILE (counter < 3) {
    TWEET "Making deal #" + counter;
    counter = counter + 1;
}
```

#### For Loops

```
MAKE AMERICA GREAT AGAIN FOR (YUGE i = 0; i < 3; i = i + 1) {
    TWEET "Making America Great Again: " + i;
}
```

### Output

```
TWEET "Normal output";                   // Standard output
RALLY "Emphasized output";               // Emphasized output (ALL CAPS with random Trump-like emphasis)
EXECUTIVE_ORDER "Warning or error";      // Error output
```

### Data Types

- Integers: `45`
- Floats: `3.14`
- Strings: `"Make Programming Great Again"`
- Arrays: `[1, 2, 3, 45]`
- Booleans: `WINNING` (true) and `LOSER` (false)

### Comments

```
// Single-line comment

/* 
 * Multi-line comment
 * The best comments, believe me!
 */
```

### Built-in Functions

- `TREMENDOUS_SORT(array)` - Sorts an array (with a twist)
- `AMERICA_FIRST(array)` - Prioritizes certain elements in an array
- Standard functions: `len`, `first`, `last`, `rest`, `push`

## Examples

### Hello World

```
TWEET "Hello, World! It's going to be TREMENDOUS!";
```

### Calculating Factorial

```
YUGE FUNCTION factorial(n) RATED 10/10 {
    BUILD WALL IF (n <= 1) {
        RETURN 1;
    } ELSE {
        RETURN n * factorial(n - 1);
    }
}

TWEET "Factorial of 5 is: " + factorial(5);
```

## Contributing

Contributions to make TRUMP even greater are welcome! Please feel free to submit pull requests or report issues.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This project is satirical and meant for humor and educational purposes only. It is not affiliated with or endorsed by Donald Trump or any political entity.

---

Made with love to bring humor to programming. Let's Make Programming Great Again!