package modthree

import (
	"fsm-modulo-three/fsm"
	"strconv"
	"testing"
)

func TestNewModThreeFSM(t *testing.T) {
	fsm := NewModThreeFSM()

	if fsm == nil {
		t.Fatal("Expected non-nil ModThreeFSM")
	}

	automaton := fsm.GetAutomaton()
	if automaton == nil {
		t.Fatal("Expected non-nil automaton")
	}

	states := automaton.GetStates()
	if len(states) != 3 {
		t.Errorf("Expected 3 states, got %d", len(states))
	}

	expectedStates := []string{"S0", "S1", "S2"}
	for i, state := range expectedStates {
		if string(states[i]) != state {
			t.Errorf("Expected state %s at position %d, got %s", state, i, states[i])
		}
	}

	alphabet := automaton.GetAlphabet()
	if len(alphabet) != 2 {
		t.Errorf("Expected 2 symbols in alphabet, got %d", len(alphabet))
	}

	expectedAlphabet := []string{"0", "1"}
	for i, symbol := range expectedAlphabet {
		if string(alphabet[i]) != symbol {
			t.Errorf("Expected symbol %s at position %d, got %s", symbol, i, alphabet[i])
		}
	}

	initialState := automaton.GetInitialState()
	if string(initialState) != "S0" {
		t.Errorf("Expected initial state S0, got %s", initialState)
	}
}

func TestModThree_ExampleCases(t *testing.T) {
	fsm := NewModThreeFSM()

	tests := []struct {
		input             string
		expectedRemainder int
		description       string
	}{
		{"1101", 1, "thirteen (1101 = 13, 13 % 3 = 1)"},
		{"1110", 2, "fourteen (1110 = 14, 14 % 3 = 2)"},
		{"1111", 0, "fifteen (1111 = 15, 15 % 3 = 0)"},
		{"110", 0, "six (110 = 6, 6 % 3 = 0)"},
		{"1010", 1, "ten (1010 = 10, 10 % 3 = 1)"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := fsm.ModThree(test.input)
			if err != nil {
				t.Errorf("Unexpected error for input '%s': %v", test.input, err)
				return
			}

			if result.Remainder != test.expectedRemainder {
				t.Errorf("For input '%s' (%s): expected remainder %d, got %d",
					test.input, test.description, test.expectedRemainder, result.Remainder)
			}

			expectedDecimal, _ := strconv.ParseInt(test.input, 2, 64)
			if result.DecimalValue != int(expectedDecimal) {
				t.Errorf("For input '%s': expected decimal value %d, got %d",
					test.input, expectedDecimal, result.DecimalValue)
			}
		})
	}
}

func TestModThree_ComprehensiveTest(t *testing.T) {
	fsm := NewModThreeFSM()

	testCases := []struct {
		binary            string
		decimal           int
		expectedRemainder int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"10", 2, 2},
		{"11", 3, 0},
		{"100", 4, 1},
		{"101", 5, 2},
		{"110", 6, 0},
		{"111", 7, 1},
		{"1000", 8, 2},
		{"1001", 9, 0},
		{"1010", 10, 1},
		{"1011", 11, 2},
		{"1100", 12, 0},
		{"1101", 13, 1},
		{"1110", 14, 2},
		{"1111", 15, 0},
		{"10000", 16, 1},
		{"10001", 17, 2},
		{"10010", 18, 0},
		{"10011", 19, 1},
		{"10100", 20, 2},
		{"10101", 21, 0},
		{"10110", 22, 1},
		{"10111", 23, 2},
		{"11000", 24, 0},
		{"11001", 25, 1},
		{"11010", 26, 2},
		{"11011", 27, 0},
		{"11100", 28, 1},
		{"11101", 29, 2},
		{"11110", 30, 0},
		{"11111", 31, 1},
	}

	for _, testCase := range testCases {
		t.Run(testCase.binary, func(t *testing.T) {
			result, err := fsm.ModThree(testCase.binary)
			if err != nil {
				t.Errorf("Unexpected error for input '%s': %v", testCase.binary, err)
				return
			}

			if result.Remainder != testCase.expectedRemainder {
				t.Errorf("For input '%s' (decimal %d): expected remainder %d, got %d",
					testCase.binary, testCase.decimal, testCase.expectedRemainder, result.Remainder)
			}

			if result.DecimalValue != testCase.decimal {
				t.Errorf("For input '%s': expected decimal value %d, got %d",
					testCase.binary, testCase.decimal, result.DecimalValue)
			}

			expectedState := fsm.stateToRemainder(result.FinalState)
			if expectedState != testCase.expectedRemainder {
				t.Errorf("For input '%s': final state %s should map to remainder %d, but got %d",
					testCase.binary, result.FinalState, testCase.expectedRemainder, expectedState)
			}
		})
	}
}

func TestModThree_EdgeCases(t *testing.T) {
	fsm := NewModThreeFSM()

	edgeCases := []struct {
		input             string
		expectedRemainder int
		description       string
	}{
		{"0", 0, "zero"},
		{"1", 1, "one"},
		{"00", 0, "zero with leading zero"},
		{"01", 1, "one with leading zero"},
		{"000", 0, "zero with multiple leading zeros"},
		{"001", 1, "one with multiple leading zeros"},
		{"1000000", 1, "large power of two"},
		{"1111111", 1, "large number of ones"},
	}

	for _, testCase := range edgeCases {
		t.Run(testCase.input, func(t *testing.T) {
			result, err := fsm.ModThree(testCase.input)
			if err != nil {
				t.Errorf("Unexpected error for input '%s' (%s): %v", testCase.input, testCase.description, err)
				return
			}

			if result.Remainder != testCase.expectedRemainder {
				t.Errorf("For input '%s' (%s): expected remainder %d, got %d",
					testCase.input, testCase.description, testCase.expectedRemainder, result.Remainder)
			}
		})
	}
}

func TestModThree_InvalidInput(t *testing.T) {
	fsm := NewModThreeFSM()

	invalidInputs := []struct {
		input       string
		description string
	}{
		{"", "empty string"},
		{"2", "invalid digit"},
		{"a", "letter"},
		{"01a", "mixed valid and invalid"},
		{"a01", "invalid at start"},
		{"01a", "invalid at end"},
		{"012", "multiple invalid digits"},
		{"abc", "all letters"},
		{" ", "space"},
		{"01 ", "space at end"},
		{" 01", "space at start"},
	}

	for _, testCase := range invalidInputs {
		t.Run(testCase.input, func(t *testing.T) {
			_, err := fsm.ModThree(testCase.input)
			if err == nil {
				t.Errorf("Expected error for invalid input '%s' (%s), but got none",
					testCase.input, testCase.description)
			}
		})
	}
}

func TestStateToRemainder(t *testing.T) {
	modThreeFSM := NewModThreeFSM()

	tests := []struct {
		state             fsm.State
		expectedRemainder int
	}{
		{"S0", 0},
		{"S1", 1},
		{"S2", 2},
		{"invalid", -1},
	}

	for _, test := range tests {
		t.Run(string(test.state), func(t *testing.T) {
			remainder := modThreeFSM.stateToRemainder(test.state)
			if remainder != test.expectedRemainder {
				t.Errorf("For state %s: expected remainder %d, got %d",
					test.state, test.expectedRemainder, remainder)
			}
		})
	}
}

func TestValidateInput(t *testing.T) {
	fsm := NewModThreeFSM()

	validInputs := []string{"0", "1", "00", "01", "10", "11", "000", "001", "010", "011", "100", "101", "110", "111"}
	for _, input := range validInputs {
		t.Run("valid_"+input, func(t *testing.T) {
			err := fsm.validateInput(input)
			if err != nil {
				t.Errorf("Expected no error for valid input '%s', but got: %v", input, err)
			}
		})
	}

	invalidInputs := []string{"", "2", "a", "01a", "a01", "012", "abc", " ", "01 ", " 01"}
	for _, input := range invalidInputs {
		t.Run("invalid_"+input, func(t *testing.T) {
			err := fsm.validateInput(input)
			if err == nil {
				t.Errorf("Expected error for invalid input '%s', but got none", input)
			}
		})
	}
}

func TestModThreeResult_StringRepresentation(t *testing.T) {
	fsm := NewModThreeFSM()

	result, err := fsm.ModThree("1101")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.Input != "1101" {
		t.Errorf("Expected input '1101', got '%s'", result.Input)
	}

	if result.FinalState != "S1" {
		t.Errorf("Expected final state 'S1', got '%s'", result.FinalState)
	}

	if result.Remainder != 1 {
		t.Errorf("Expected remainder 1, got %d", result.Remainder)
	}

	if result.DecimalValue != 13 {
		t.Errorf("Expected decimal value 13, got %d", result.DecimalValue)
	}

	if result.BinaryValue != 13 {
		t.Errorf("Expected binary value 13, got %d", result.BinaryValue)
	}
}

func TestString(t *testing.T) {
	fsm := NewModThreeFSM()

	str := fsm.String()
	if str == "" {
		t.Error("Expected non-empty string representation")
	}

	expectedComponents := []string{"ModThree FSM:", "Finite Automaton:", "States:", "Alphabet:", "Initial State:", "Accepting States:"}
	for _, component := range expectedComponents {
		if !contains(str, component) {
			t.Errorf("String representation should contain '%s'", component)
		}
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			contains(s[1:len(s)-1], substr)))
}
