package main

import (
	"fmt"

	"github.com/mrochk/busy-beaver/tm"
)

func main() {
	fmt.Println("\n1-state busy beaver:")
	m := tm.NewTuringMachine(1)
	m.Run("machines/best-1-state.txt")

	fmt.Println("\n1-state diverging:")
	m = tm.NewTuringMachine(1)
	m.Run("machines/div-1-state.txt")

	fmt.Println("2-state busy beaver:")
	m = tm.NewTuringMachine(2)
	m.Run("machines/best-2-state.txt")

	fmt.Println("3-state busy beaver:")
	m = tm.NewTuringMachine(3)
	m.Run("machines/best-3-state.txt")
}
