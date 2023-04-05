package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		tm   *TuringMachine = nil
		file *os.File       = nil
	)

	// Best 1-state Busy-Beaver, score = 1
	fmt.Println("\n1-state busy beaver:")
	tm = newTuringMachine(1)
	file, _ = os.Open("machines/1-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm.tape)

	// Best 2-state Busy-Beaver, score = 4
	fmt.Println("2-state busy beaver:")
	tm = newTuringMachine(2)
	file, _ = os.Open("machines/2-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm.tape)

	// Best 3-state Busy-Beaver, score = 6
	fmt.Println("3-state busy beaver:")
	tm = newTuringMachine(3)
	file, _ = os.Open("machines/3-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm.tape)
}
