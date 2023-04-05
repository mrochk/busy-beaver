package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	tapeSize     = 51 // The tape has a fixed size.
	alphabetSize = 2  // Alphabet = {0, 1}
	diverging    = 20 // We consider machine as diverging after this n of iters.
)

type Input struct {
	symbol bool
	state  string
}

type Output struct {
	symbol    bool
	state     string
	direction bool
}

type tape [tapeSize]bool

type Instructions map[Input]Output

type TuringMachine struct {
	tape         tape
	position     int
	state        string
	states       int
	instructions Instructions
}

func newTuringMachine(states int) *TuringMachine {
	return &TuringMachine{
		position:     tapeSize / 2,
		states:       states,
		state:        "0",
		instructions: make(Instructions, states*alphabetSize),
	}
}

func (tm *TuringMachine) setInstructions(file *os.File) {
	reader := bufio.NewReader(file)
	for i := 0; i < tm.states*alphabetSize; i++ {
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1]
		arr := strings.Split(line, ",")
		input := Input{state: arr[0], symbol: (arr[1] == "1")}
		output := Output{symbol: arr[2] == "1", direction: arr[3] == "R", state: arr[4]}
		tm.instructions[input] = output
	}
}

func (tm *TuringMachine) step() {
	var (
		symbol = tm.tape[tm.position]
		input  = Input{state: tm.state, symbol: symbol}
		output = tm.instructions[input]
	)
	tm.tape[tm.position] = output.symbol
	if output.direction {
		tm.position++
	} else {
		tm.position--
	}
	tm.state = output.state
}

func (tm *TuringMachine) run() {
	for count := 0; tm.state != "H"; count++ {
		if count >= diverging {
			fmt.Println("Machine is probably diverging.")
			return
		}
		tm.step()
	}
}

func (t tape) getScore() int {
	score := 0
	for i := range t {
		if t[i] {
			score++
		}
	}
	return score
}

func (i Instructions) String() string {
	out := ""
	for input := range i {
		output := i[input]
		out += fmt.Sprintf("state %s, symb %t -> symb %t, dir %t, state %s\n",
			input.state, input.symbol, output.symbol, output.direction, output.state)
	}
	return out
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
	out += "\nScore = " + fmt.Sprint(t.getScore()) + "\n\n"
	return out
}
