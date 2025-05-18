# TRUMP Programming Language

A tremendous programming language inspired by Donald Trump's speaking style. BELIEVE ME, it's going to be YUGE!

## Overview

TRUMP (Tremendously Radical Universal Multi-Purpose Programming) is a programming language that combines traditional programming concepts with the unique linguistic style of Donald Trump. It's designed to be both functional and entertaining, featuring Trump-themed syntax and error messages.

## Features

- **Trump-style syntax**: Write code that sounds like Trump's speeches
- **Variable declarations with YUGE and TREMENDOUS**
- **Conditional statements with BUILD WALL IF/ELSE**
- **Loop constructs with MAKE DEALS WHILE and MAKE AMERICA GREAT AGAIN FOR**
- **Output with TWEET, RALLY, and EXECUTIVE_ORDER**
- **Specialized built-in functions like TREMENDOUS_SORT and AMERICA_FIRST**
- **Trump-themed error messages**: When code fails, it fails in the most tremendous way

## Installation

### Prerequisites

- Go 1.22 or later

### Building from Source

```bash
git clone https://github.com/AndrewDonelson/trumplang.git
cd trumplang
make
```

### Running the Examples

```bash
# Run the "Hello, World!" example
./build/trumpc run examples/hello/main.trump

# Run the Fibonacci example
./build/trumpc run examples/fibonacci/main.trump
```

## Language Reference

### Variable Declaration

```
YUGE variableName = value;      // For regular variables
TREMENDOUS bigVariable = value; // For particularly important variables
```

### Functions

```
YUGE FUNCTION functionName(param1, param2) RATED "10/10" {
    // Function body
    RETURN value;
}
```

### Conditional Statements

```
BUILD WALL IF (condition) {
    // Code to execute if condition is true
} ELSE {
    // Code to execute if condition is false
}
```

### Loops

```
// While loop
MAKE DEALS WHILE (condition) {
    // Loop body
}

// For loop
MAKE AMERICA GREAT AGAIN FOR (initialization; condition; update) {
    // Loop body
}
```

### Output Statements

```
TWEET expression;               // Regular output
RALLY expression;               // Emphasized output with Trump-style flourishes
EXECUTIVE_ORDER expression;     // Error/warning output
```

### Built-in Functions

- `TREMENDOUS_SORT(array)` - Sorts an array (with a small chance of being unpredictable)
- `AMERICA_FIRST(array)` - Prioritizes the value 45 in an array
- `DEAL(value1, value2)` - Swaps two values
- `len(string/array)` - Returns the length 
- `first(array)` - Returns the first element
- `last(array)` - Returns the last element
- `rest(array)` - Returns all elements except the first
- `push(array, element)` - Adds an element to an array

### Data Types

- Integers
- Floats
- Strings
- Booleans (`WINNING` for true, `LOSER` for false)
- Arrays
- Functions

### Comments

```
// Regular comment

FAKE NEWS: Trump-style comment for longer explanations
```

## Command-Line Interface

TRUMP comes with a command-line interface (CLI) for various operations:

```bash
# Build a TRUMP program
trumpc build file.trump

# Run a TRUMP program directly
trumpc run file.trump

# Create a new TRUMP project
trumpc create myproject

# Check a program for errors without compiling
trumpc inspect file.trump
```

## Examples

### Hello World

```
// file: examples/hello/main.trump
// description: Hello World example in TRUMP language

FAKE NEWS: The simplest TRUMP program ever written!
FAKE NEWS: This will be TREMENDOUS! BELIEVE ME!

// Say hello to the world, Trump style
TWEET "Hello, World! It's going to be TREMENDOUS!";

// Build a terrific message
TREMENDOUS message = "MAKE PROGRAMMING GREAT AGAIN!";
RALLY message;
```

### Fibonacci Sequence

See the [Fibonacci example](examples/fibonacci/main.trump) for a more complex program.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This project is created for entertainment and educational purposes. It's meant to be a humorous take on programming languages and is not intended to make any political statement.