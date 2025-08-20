package main

import (
	"bufio"
	"fmt"
	"fsm-modulo-three/modthree"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Finite State Machine Modulo Three Implementation ===")
	fmt.Println("This program demonstrates the FSM library and mod-three functionality.")
	fmt.Println()

	fsm := modthree.NewModThreeFSM()

	fmt.Println("FSM Configuration:")
	fmt.Println(fsm.String())
	fmt.Println()

	fmt.Println("=== Example Cases ===")
	examples := []string{"1101", "1110", "1111", "110", "1010"}

	for _, example := range examples {
		result, err := fsm.ModThree(example)
		if err != nil {
			fmt.Printf("Error processing '%s': %v\n", example, err)
			continue
		}

		fmt.Printf("Input: %s (decimal: %d) -> Remainder: %d (Final State: %s)\n",
			result.Input, result.DecimalValue, result.Remainder, result.FinalState)
	}
	fmt.Println()

	fmt.Println("=== Interactive Mode ===")
	fmt.Println("Enter binary strings to compute their modulo three remainder.")
	fmt.Println("Enter 'quit' to exit.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter binary string: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "quit" || input == "exit" {
			break
		}

		if input == "" {
			continue
		}

		result, err := fsm.ModThree(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("Result: %s (decimal: %d) %% 3 = %d (Final State: %s)\n",
			result.Input, result.DecimalValue, result.Remainder, result.FinalState)
		fmt.Println()
	}

	fmt.Println("Goodbye!")
}
