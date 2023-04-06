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

	// 1-state busy beaver, score = 1
	fmt.Println("\n1-state busy beaver:")
	tm = newTuringMachine(1)
	file, _ = os.Open("machines/best-1-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm)

	// Diverging 1-state Turing machine.
	fmt.Println("\n1-state diverging:")
	tm = newTuringMachine(1)
	file, _ = os.Open("machines/div-1-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm)

	// 2-state busy beaver, score = 4
	fmt.Println("2-state busy beaver:")
	tm = newTuringMachine(2)
	file, _ = os.Open("machines/best-2-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm)

	// 3-state busy beaver, score = 6
	fmt.Println("3-state busy beaver:")
	tm = newTuringMachine(3)
	file, _ = os.Open("machines/best-3-state.txt")
	tm.setInstructions(file)
	tm.run()
	fmt.Print(tm)
}
