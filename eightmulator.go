package main

import (
	"fmt"
)

func main() {
	fmt.Println("Eightmulator Initialized")

	machine := createMachine()
	machine.loadRom("./INVADERS")

	for {
		opcode := machine.fetchOpcode()
		machine.handleOpcode(opcode)
	}
}
