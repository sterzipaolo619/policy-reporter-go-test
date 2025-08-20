package modthree

import (
	"fmt"
	"fsm-modulo-three/fsm"
	"strconv"
)

type ModThreeResult struct {
	Input        string
	FinalState   fsm.State
	Remainder    int
	BinaryValue  int
	DecimalValue int
}

type ModThreeFSM struct {
	automaton *fsm.FiniteAutomaton
}

func NewModThreeFSM() *ModThreeFSM {
	states := []fsm.State{"S0", "S1", "S2"}

	alphabet := []fsm.Symbol{"0", "1"}

	initialState := fsm.State("S0")

	acceptingStates := []fsm.State{"S0", "S1", "S2"}

	transitionFunction := func(currentState fsm.State, symbol fsm.Symbol) fsm.State {
		switch currentState {
		case "S0":
			switch symbol {
			case "0":
				return "S0"
			case "1":
				return "S1"
			}
		case "S1":
			switch symbol {
			case "0":
				return "S2"
			case "1":
				return "S0"
			}
		case "S2":
			switch symbol {
			case "0":
				return "S1"
			case "1":
				return "S2"
			}
		}
		return currentState
	}

	automaton := fsm.NewFiniteAutomaton(
		states,
		alphabet,
		initialState,
		acceptingStates,
		transitionFunction,
	)

	return &ModThreeFSM{
		automaton: automaton,
	}
}

func (m *ModThreeFSM) ModThree(input string) (*ModThreeResult, error) {
	if err := m.validateInput(input); err != nil {
		return nil, err
	}

	finalState, err := m.automaton.ProcessInput(input)
	if err != nil {
		return nil, fmt.Errorf("FSM processing error: %w", err)
	}

	remainder := m.stateToRemainder(finalState)

	binaryValue, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse binary string: %w", err)
	}

	expectedRemainder := int(binaryValue % 3)
	if remainder != expectedRemainder {
		return nil, fmt.Errorf("FSM result mismatch: got %d, expected %d", remainder, expectedRemainder)
	}

	return &ModThreeResult{
		Input:        input,
		FinalState:   finalState,
		Remainder:    remainder,
		BinaryValue:  int(binaryValue),
		DecimalValue: int(binaryValue),
	}, nil
}

func (m *ModThreeFSM) validateInput(input string) error {
	if input == "" {
		return fmt.Errorf("input string cannot be empty")
	}

	for i, char := range input {
		if char != '0' && char != '1' {
			return fmt.Errorf("invalid character '%c' at position %d: only '0' and '1' are allowed", char, i)
		}
	}

	return nil
}

func (m *ModThreeFSM) stateToRemainder(state fsm.State) int {
	switch state {
	case "S0":
		return 0
	case "S1":
		return 1
	case "S2":
		return 2
	default:
		return -1
	}
}

func (m *ModThreeFSM) GetAutomaton() *fsm.FiniteAutomaton {
	return m.automaton
}

func (m *ModThreeFSM) String() string {
	return fmt.Sprintf("ModThree FSM:\n%s", m.automaton.String())
}
