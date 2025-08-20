# Finite State Machine Modulo Three Implementation

This project implements an advanced Finite State Machine (FSM) library in Go, with the modulo three calculation as a concrete example. The implementation follows the abstract definition of finite automata and provides a reusable API for other developers.

## Project Structure

```
fsm-modulo-three/
├── fsm/                    # Core FSM library
│   ├── fsm.go             # Main FSM implementation
│   └── fsm_test.go        # FSM unit tests
├── modthree/              # Mod-three specific implementation
│   ├── modthree.go        # Mod-three FSM implementation
│   └── modthree_test.go   # Mod-three unit tests
├── cmd/                   # Application entry point
│   └── main.go           # Interactive demo application
├── go.mod                 # Go module file
└── README.md             # This file
```

## Features

### Core FSM Library (`fsm` package)
- **Abstract FSM Implementation**: Implements the 5-tuple (Q,Σ,q0,F,δ) definition
- **Flexible API**: Designed for extensibility and reuse by other developers
- **Input Validation**: Validates input symbols against the defined alphabet
- **Comprehensive Testing**: Full unit test coverage with edge cases

### Mod-Three Implementation (`modthree` package)
- **State Transition Logic**: Implements the exact state diagram from the exercise
- **Result Verification**: Cross-validates FSM results with traditional modulo arithmetic
- **Rich Output**: Provides detailed results including final state, remainder, and decimal conversion
- **Error Handling**: Comprehensive input validation and error reporting

## Installation and Setup

### Prerequisites
- Go 1.21 or later

### Setup Instructions
1. Clone or download the project
2. Navigate to the project directory
3. Run the tests to verify everything works:
   ```bash
   go test ./...
   ```
4. Build and run the demo application:
   ```bash
   go run cmd/main.go
   ```

## Usage Examples

### Using the FSM Library

```go
package main

import (
    "fmt"
    "fsm-modulo-three/fsm"
)

func main() {
    // Define states
    states := []fsm.State{"S0", "S1", "S2"}
    
    // Define alphabet
    alphabet := []fsm.Symbol{"0", "1"}
    
    // Define transition function
    transitionFunction := func(currentState fsm.State, symbol fsm.Symbol) fsm.State {
        // Your transition logic here
        return currentState
    }
    
    // Create FSM
    automaton := fsm.NewFiniteAutomaton(
        states,
        alphabet,
        "S0", // initial state
        states, // all states are accepting
        transitionFunction,
    )
    
    // Process input
    finalState, err := automaton.ProcessInput("1101")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Final state: %s\n", finalState)
}
```

### Using the Mod-Three Implementation

```go
package main

import (
    "fmt"
    "fsm-modulo-three/modthree"
)

func main() {
    // Create mod-three FSM
    fsm := modthree.NewModThreeFSM()
    
    // Calculate modulo three
    result, err := fsm.ModThree("1101")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Input: %s\n", result.Input)
    fmt.Printf("Decimal Value: %d\n", result.DecimalValue)
    fmt.Printf("Remainder: %d\n", result.Remainder)
    fmt.Printf("Final State: %s\n", result.FinalState)
}
```

## FSM State Transition Diagram

The mod-three FSM implements the following state transitions:

```
Initial State: S0

S0 --0--> S0
S0 --1--> S1
S1 --0--> S2
S1 --1--> S0
S2 --0--> S1
S2 --1--> S2

Final State Mapping:
S0 -> Remainder 0
S1 -> Remainder 1
S2 -> Remainder 2
```

## Example Walkthrough

For input string "110":
1. Initial state = S0, Input = 1, result state = S1
2. Current state = S1, Input = 1, result state = S0
3. Current state = S0, Input = 0, result state = S0
4. No more input
5. Final state = S0, Output = 0

For input string "1010":
1. Initial state = S0, Input = 1, result state = S1
2. Current state = S1, Input = 0, result state = S2
3. Current state = S2, Input = 1, result state = S2
4. Current state = S2, Input = 0, result state = S1
5. No more input
6. Final state = S1, Output = 1

## Testing

The project includes comprehensive unit tests covering:

### FSM Library Tests
- FSM creation and configuration
- Input processing with valid and invalid inputs
- Symbol validation
- State management
- Error handling

### Mod-Three Tests
- Example cases from the exercise
- Comprehensive range of binary numbers (0-31)
- Edge cases (empty strings, leading zeros, large numbers)
- Invalid input handling
- State-to-remainder mapping verification

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./fsm
go test ./modthree
```

## Design Decisions and Assumptions

### 1. **Extensibility Focus**
- The FSM library is designed as a reusable component
- Clear separation between the abstract FSM and the concrete mod-three implementation
- Generic API that can be used for other FSM applications

### 2. **Error Handling**
- Comprehensive input validation
- Descriptive error messages with context
- Graceful handling of edge cases

### 3. **Verification Strategy**
- Cross-validation between FSM results and traditional modulo arithmetic
- This ensures correctness while demonstrating the FSM approach

### 4. **API Design**
- Clean, intuitive interfaces
- Immutable FSM configuration
- Rich result objects with detailed information

### 5. **Testing Strategy**
- Unit tests for all public methods
- Edge case coverage
- Invalid input testing
- Comprehensive example validation

### 6. **Performance Considerations**
- Efficient state transitions using switch statements
- Minimal memory allocation
- O(n) time complexity where n is input length

## Assumptions Made

1. **Input Format**: Binary strings containing only '0' and '1' characters
2. **Processing Order**: Most significant bit first (left-to-right processing)
3. **State Representation**: String-based state names for clarity
4. **Error Handling**: Fail-fast approach with descriptive error messages
5. **Verification**: Traditional modulo arithmetic is used to verify FSM correctness

## Future Enhancements

The FSM library is designed to support future enhancements:

1. **Additional FSM Types**: Deterministic, non-deterministic, and pushdown automata
2. **Visualization**: State diagram generation
3. **Optimization**: Compilation to optimized state machines
4. **Serialization**: Save/load FSM configurations
5. **Performance Monitoring**: State transition tracking and analysis

## Troubleshooting

### Common Issues

1. **Import Errors**: Ensure you're in the correct directory and the module path is correct
2. **Test Failures**: Run `go mod tidy` to ensure dependencies are properly resolved
3. **Build Errors**: Verify Go version compatibility (requires Go 1.21+)

### Getting Help

If you encounter issues:
1. Check that all tests pass: `go test ./...`
2. Verify Go version: `go version`
3. Clean and rebuild: `go clean && go build ./...`

## License

This project is provided as a technical exercise implementation. Please refer to your organization's licensing policies for usage guidelines.

