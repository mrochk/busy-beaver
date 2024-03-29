package tm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	tapeSize     = 31 // The tape has a fixed size.
	alphabetSize = 2  // Alphabet = {0, 1}
	diverging    = (tapeSize + 1) / 2
)

type TuringMachine struct {
	tape         tape
	position     int
	state        string
	states       int
	instructions Instructions
}

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

type tape [tapeSize]bool

func NewTuringMachine(states int) *TuringMachine {
	return &TuringMachine{
		position:     tapeSize / 2,
		states:       states,
		state:        "0",
		instructions: make(Instructions, states*alphabetSize),
	}
}

func (tm *TuringMachine) Run(filename string) {
	file, _ := os.Open(filename)
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm)
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

func (tm *TuringMachine) run() {
	for count := 0; tm.state != "H"; count++ {
		if count >= diverging {
			fmt.Fprintln(os.Stderr, "ERROR: machine is divergent")
			return
		}
		tm.step()
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

func (tm TuringMachine) String() string {
	out := "["
	for i := range tm.tape {
		if tm.tape[i] {
			out += "1 "
		} else {
			out += "0 "
		}
	}
	out = out[:len(out)-1] + "]\n"
	for i := 0; i < tm.position; i++ {
		out += "  "
	}
	out += " ^\n"
	return out + "Score: " + fmt.Sprint(tm.tape.getScore()) + "\n\n"
}
