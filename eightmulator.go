package main

import (
	"fmt"
)

func main() {
	fmt.Println("Eightmulator Initialized")

	machine := createMachine()
	machine.loadRom("./PONG")

	for {
		opcode := machine.fetchOpcode()
		machine.handleOpcode(opcode)
	}
}
