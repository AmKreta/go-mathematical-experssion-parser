# Golang Expression Parser

A mathematical expression parser and evaluator written in Go. This project implements a complete interpreter pipeline: lexical analysis, parsing, and evaluation using the visitor pattern.

## Features

- **Mathematical Operations**: Addition, subtraction, multiplication, division, modulo, integer division, and exponentiation
- **Unary Operators**: Factorial (`!`), trigonometric functions (sin, cos, tan, cot, sec, cosec), and logarithm
- **Constants**: Built-in support for `pi` (π)
- **Grouping**: Supports parentheses `()`, square brackets `[]`, and curly braces `{}`
- **REPL**: Interactive Read-Eval-Print Loop for testing expressions
- **Visitor Pattern**: Clean separation of concerns using the visitor pattern for AST traversal

## Project Structure

```
experssion-parser/
├── ast/                    # Abstract Syntax Tree nodes
│   ├── ast.go             # Core AST interfaces (Node, Visitor, Visitable)
│   ├── binaryOperator.go  # Binary operator node (+, -, *, /, %, //, ^)
│   ├── number.go          # Number literal node
│   ├── unaryOperator.go   # Unary operator node (!, sin, cos, etc.)
│   └── type.go            # Type definitions (ComputeSize)
│
├── lexer/                  # Lexical analyzer (tokenizer)
│   ├── lexer.go           # Main lexer implementation
│   └── factory.go         # Lexer factory functions
│
├── parser/                 # Parser (syntax analyzer)
│   └── parser.go          # Recursive descent parser
│
├── interpreter/            # Expression evaluator
│   └── interpreter.go     # Visitor implementation for evaluation
│
├── token/                  # Token definitions
│   ├── token.go           # Token structure
│   └── types.go           # Token type constants
│
├── repl/                   # Read-Eval-Print Loop
│   ├── repl.go            # Interactive REPL interface
│   └── eval.go            # Expression evaluation wrapper
│
├── util/                   # Utility functions
│   └── factorial.go       # Factorial calculation
│
├── visitor/                # Visitor pattern (type aliases for backward compatibility)
│
├── main.go                 # Application entry point
└── go.mod                  # Go module definition
```

## Architecture

The project follows a classic interpreter architecture:

1. **Lexer** (`lexer/`): Converts input strings into a stream of tokens
2. **Parser** (`parser/`): Builds an Abstract Syntax Tree (AST) from tokens using recursive descent parsing
3. **AST** (`ast/`): Represents the parsed expression as a tree structure
4. **Interpreter** (`interpreter/`): Traverses the AST using the visitor pattern to evaluate expressions

### Grammar

```
expression -> additive_expression
additive_expression -> multiplicative_expression (('+' | '-') multiplicative_expression)*
multiplicative_expression -> unary_expression (('*' | '/' | '%' | '//' | '^') unary_expression)*
unary_expression -> factor (('!' | 'sin' | 'cos' | 'tan' | 'cot' | 'sec' | 'cosec' | 'log') factor)*
prefix_unary_expression -> ('sin' | 'cos' | 'tan' | 'cot' | 'sec' | 'cosec' | 'log')(factor | prefix_unary_expression | suffix_unary_expression)*
suffix_unary_expression -> factor!
factor -> number | 'pi' | '(' expression ')' | '[' expression ']' | '{' expression '}'
```

## Usage

### Prerequisites

- Go 1.23.2 or later

### Building and Running

1. **Clone or navigate to the project directory:**
   ```bash
   cd experssion-parser
   ```

2. **Build the project:**
   ```bash
   go build
   ```

3. **Run the REPL:**
   ```bash
   ./experssion-parser
   # On Windows:
   experssion-parser.exe
   ```

4. **Enter expressions at the prompt:**
   ```
   >>> enter expression: 2 + 3 * 4
   result:  14
   
   >>> enter expression: sin(pi/2)
   result:  1
   
   >>> enter expression: 5!
   result:  120
   
   >>> enter expression: exit
   ```

### Supported Operations

#### Binary Operators
- `+` : Addition
- `-` : Subtraction
- `*` : Multiplication
- `/` : Division
- `%` : Modulo
- `//` : Integer division
- `^` : Exponentiation

#### Unary Operators
- `!` : Factorial (suffix)
- `sin()` : Sine function
- `cos()` : Cosine function
- `tan()` : Tangent function
- `cot()` : Cotangent function
- `sec()` : Secant function
- `cosec()` : Cosecant function
- `log()` : Natural logarithm

#### Constants
- `pi` : Mathematical constant π (approximately 3.14159...)

#### Grouping
- `()` : Parentheses
- `[]` : Square brackets
- `{}` : Curly braces

### Example Expressions

```bash
# Basic arithmetic
>>> 10 + 20 * 2
result: 50

# With parentheses
>>> (10 + 20) * 2
result: 60

# Trigonometric functions
>>> sin(pi/2)
result: 1

>>> cos(0)
result: 1

# Factorial
>>> 5!
result: 120

# Exponentiation
>>> 2^8
result: 256

# Complex expressions
>>> sin(pi/4) * cos(pi/4)
result: 0.5

# Logarithm
>>> log(2.71828)
result: 0.999999327347282
```

### Code Structure Notes

- The **visitor pattern** is used to separate AST structure from operations (evaluation, pretty-printing, etc.)
- The `Visitor` and `Visitable` interfaces are defined in the `ast` package to avoid import cycles
- The `visitor` package provides type aliases for backward compatibility
- All AST nodes implement the `Node` interface and support the visitor pattern via `AcceptVisitor()`

## License

This project is part of a learning exercise in building interpreters and parsers in Go.

