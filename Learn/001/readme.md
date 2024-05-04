**README**

This document serves as a guide to understanding variables, zero values, blank identifiers, declaration, assignment, initialization, and the concepts of value and type in the Go programming language.

### Variables, Zero Values, Blank Identifier

In Go, variables are declared using the `var` keyword, which assigns them a zero value if no initial value is provided. The zero value depends on the type of the variable. Blank identifiers (`_`) can be used to discard values returned by functions or to import packages for side effects.

### Declaration, Assignment, Initialization

- **Declaration**: Specifying the type of a variable.
- **Assignment**: Giving a variable a value.
- **Initialization**: Declaration combined with assignment.

#### General Guidelines

- Use `var` when you need the zero value.
- Utilize the short declaration operator (`:=`) for concise variable declaration and initialization.
- Use `var` when specificity is required, such as declaring variables with explicit types.

#### Multiple Initializations

You can initialize multiple variables in a single statement.

### Value and Type

- Go is statically typed, meaning variables have fixed types determined at compile time.
- Variables hold values of specific types, ensuring type safety.
- Unused variables are not allowed by the compiler to prevent code pollution.

### Zero Values

In Go, variables are initialized with zero values if no explicit value is assigned. Common zero values include:

- `false` for boolean.
- `0` for integers.
- `0.0` for floats.
- `""` for strings.
- `nil` for pointers, functions, interfaces, slices, channels, and maps.

Reference: Todd McLeod - Learn To Code Go - Page 15
