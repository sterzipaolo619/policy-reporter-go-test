package fsm

import (
	"testing"
)

func TestNewFiniteAutomaton(t *testing.T) {
	states := []State{"S0", "S1", "S2"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S1", "S2"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		if currentState == "S0" && symbol == "0" {
			return "S0"
		}
		if currentState == "S0" && symbol == "1" {
			return "S1"
		}
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	if fa == nil {
		t.Fatal("Expected non-nil FiniteAutomaton")
	}

	if len(fa.States) != 3 {
		t.Errorf("Expected 3 states, got %d", len(fa.States))
	}

	if len(fa.Alphabet) != 2 {
		t.Errorf("Expected 2 symbols in alphabet, got %d", len(fa.Alphabet))
	}

	if fa.InitialState != "S0" {
		t.Errorf("Expected initial state S0, got %s", fa.InitialState)
	}
}

func TestProcessInput_ValidInput(t *testing.T) {
	states := []State{"S0", "S1", "S2"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S1", "S2"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		switch currentState {
		case "S0":
			switch symbol {
			case "0":
				return "S0"
			case "1":
				return "S1"
			}
		case "S1":
			return "S1"
		}
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	tests := []struct {
		input         string
		expectedState State
	}{
		{"0", "S0"},
		{"1", "S1"},
		{"00", "S0"},
		{"01", "S1"},
		{"10", "S1"},
		{"11", "S1"},
		{"000", "S0"},
		{"001", "S1"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			finalState, err := fa.ProcessInput(test.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if finalState != test.expectedState {
				t.Errorf("Expected final state %s, got %s", test.expectedState, finalState)
			}
		})
	}
}

func TestProcessInput_InvalidInput(t *testing.T) {
	states := []State{"S0", "S1"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S1"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	invalidInputs := []string{"2", "a", "01a", "a01", "012"}

	for _, input := range invalidInputs {
		t.Run(input, func(t *testing.T) {
			_, err := fa.ProcessInput(input)
			if err == nil {
				t.Errorf("Expected error for invalid input '%s', but got none", input)
			}
		})
	}
}

func TestIsValidSymbol(t *testing.T) {
	states := []State{"S0", "S1"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S1"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	validSymbols := []Symbol{"0", "1"}
	for _, symbol := range validSymbols {
		if !fa.isValidSymbol(symbol) {
			t.Errorf("Symbol '%s' should be valid", symbol)
		}
	}

	invalidSymbols := []Symbol{"2", "a", "A", " "}
	for _, symbol := range invalidSymbols {
		if fa.isValidSymbol(symbol) {
			t.Errorf("Symbol '%s' should be invalid", symbol)
		}
	}
}

func TestIsAcceptingState(t *testing.T) {
	states := []State{"S0", "S1", "S2"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S2"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	acceptingStatesList := []State{"S0", "S2"}
	for _, state := range acceptingStatesList {
		if !fa.IsAcceptingState(state) {
			t.Errorf("State '%s' should be accepting", state)
		}
	}

	nonAcceptingStates := []State{"S1"}
	for _, state := range nonAcceptingStates {
		if fa.IsAcceptingState(state) {
			t.Errorf("State '%s' should not be accepting", state)
		}
	}
}

func TestGetStates(t *testing.T) {
	states := []State{"S0", "S1", "S2"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S1", "S2"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	retrievedStates := fa.GetStates()
	if len(retrievedStates) != len(states) {
		t.Errorf("Expected %d states, got %d", len(states), len(retrievedStates))
	}

	for i, state := range states {
		if retrievedStates[i] != state {
			t.Errorf("Expected state %s at position %d, got %s", state, i, retrievedStates[i])
		}
	}
}

func TestGetAlphabet(t *testing.T) {
	states := []State{"S0", "S1"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S1"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	retrievedAlphabet := fa.GetAlphabet()
	if len(retrievedAlphabet) != len(alphabet) {
		t.Errorf("Expected %d symbols in alphabet, got %d", len(alphabet), len(retrievedAlphabet))
	}

	for i, symbol := range alphabet {
		if retrievedAlphabet[i] != symbol {
			t.Errorf("Expected symbol %s at position %d, got %s", symbol, i, retrievedAlphabet[i])
		}
	}
}

func TestGetInitialState(t *testing.T) {
	states := []State{"S0", "S1"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S1"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	retrievedInitialState := fa.GetInitialState()
	if retrievedInitialState != initialState {
		t.Errorf("Expected initial state %s, got %s", initialState, retrievedInitialState)
	}
}

func TestGetAcceptingStates(t *testing.T) {
	states := []State{"S0", "S1", "S2"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S2"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	retrievedAcceptingStates := fa.GetAcceptingStates()
	if len(retrievedAcceptingStates) != len(acceptingStates) {
		t.Errorf("Expected %d accepting states, got %d", len(acceptingStates), len(retrievedAcceptingStates))
	}

	for i, state := range acceptingStates {
		if retrievedAcceptingStates[i] != state {
			t.Errorf("Expected accepting state %s at position %d, got %s", state, i, retrievedAcceptingStates[i])
		}
	}
}

func TestString(t *testing.T) {
	states := []State{"S0", "S1"}
	alphabet := []Symbol{"0", "1"}
	initialState := State("S0")
	acceptingStates := []State{"S0", "S1"}

	transitionFunction := func(currentState State, symbol Symbol) State {
		return currentState
	}

	fa := NewFiniteAutomaton(states, alphabet, initialState, acceptingStates, transitionFunction)

	str := fa.String()
	if str == "" {
		t.Error("Expected non-empty string representation")
	}

	expectedComponents := []string{"Finite Automaton:", "States:", "Alphabet:", "Initial State:", "Accepting States:"}
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
