package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	tapeSize = 20
	alphSize = 2
	L        = false
	R        = true
)

// Alphabet is {0, 1} so we
// represent symbols as booleans.
type Input struct {
	symbol bool
	state  string
}

type Output struct {
	symbol    bool
	state     string
	direction bool
}

type Instructions map[Input]Output

func (i Instructions) String() string {
	out := ""
	for input := range i {
		output := i[input]
		out += fmt.Sprintf("state %s, symb %t -> symb %t, dir %t, state %s\n",
			input.state, input.symbol, output.symbol,
			output.direction, output.state)
	}
	return out
}

type tape [tapeSize]bool

type TuringMachine struct {
	tape         tape
	position     int
	state        string
	states       int
	instructions Instructions
}

func (t tape) String() string {
	out := ""
	for i := range t {
		if t[i] {
			out += "1 "
		} else {
			out += "0 "
		}
	}
	out += "\n"
	return out
}

func newTuringMachine(states int) *TuringMachine {
	return &TuringMachine{
		position:     tapeSize / 2,
		states:       states,
		state:        "0",
		instructions: make(Instructions, states*alphSize),
	}
}

func (tm *TuringMachine) setInstructions(file *os.File) {
	reader := bufio.NewReader(file)

	for i := 0; i < tm.states*alphSize; i++ {
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1]
		arr := strings.Split(line, ",")

		input := Input{state: arr[0], symbol: (arr[1] == "1")}
		output := Output{symbol: arr[2] == "1", direction: arr[3] == "R", state: arr[4]}

		tm.instructions[input] = output
	}
}

func (tm *TuringMachine) step() {
	sym := tm.tape[tm.position]
	input := Input{state: tm.state, symbol: sym}
	output := tm.instructions[input]

	// Set symbol on the tape
	tm.tape[tm.position] = output.symbol
	// Move left or right
	if output.direction {
		tm.position++
	} else {
		tm.position--
	}
	// Change state
	tm.state = output.state
}

func (tm *TuringMachine) run() {
	for tm.state != "H" {
		tm.step()
	}
}

/*
STATE ; SYMBOL ;  NSYMBOL; DIRECTION ; NSTATE
*/

func main() {
	var (
		tm   *TuringMachine
		file *os.File
	)

	// Best 1-state Busy-Beaver, score = 1
	fmt.Println("Best 1-state busy beaver:")
	tm = newTuringMachine(1)
	file, _ = os.Open("1-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm.tape)

	// Best 2-state Busy-Beaver, score = 4
	fmt.Println("Best 2-state busy beaver:")
	tm = newTuringMachine(2)
	file, _ = os.Open("2-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm.tape)

	// Best 3-state Busy-Beaver, score = 6
	fmt.Println("Best 3-state busy beaver:")
	tm = newTuringMachine(3)
	file, _ = os.Open("3-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm.tape)
}
