package main

import (
	"fmt"
	"os"
)

type Machine struct {
	memory    [4096]uint8
	registers [16]uint8
	stack     [16]uint16

	i  uint16
	pc uint16
	sp uint8
	dt uint8
	st uint8

	opcode uint16
}

func createMachine() *Machine {
	m := new(Machine)
	m.pc = 0

	return m
}

func (m *Machine) loadRom(path string) {
	file, _ := os.Open(path)
	file.Read(m.memory[0:])
}

func (m *Machine) memdump() {
	fmt.Print(m.memory)
}

func (m *Machine) fetchOpcodeAt(location uint16) uint16 {
	var opcode uint16
	hi := m.memory[location]
	lo := m.memory[location+1]

	fmt.Printf("hi: %x\n", hi)
	fmt.Printf("lo: %x\n", lo)

	opcode = uint16(hi)<<8 | uint16(lo)
	m.opcode = opcode
	fmt.Printf("OPCODE: %x\n", opcode)

	return opcode
}

func (m *Machine) fetchOpcode() uint16 {
	return m.fetchOpcodeAt(m.pc)
}

func (m *Machine) handleOpcode(opcode uint16) {
	// 0x0 Cases
	switch opcode {
	case 0x00E0:
		m.opCls()
	case 0x00EE:
		m.opCls()
	}

	// 0x1 -> 0xD (not 0x8) Cases
	switch opcode & 0xF000 {
	case 0x1000:
		m.opJmp()
	case 0x2000:
		m.opRet()
	case 0x3000:
		m.opSe()
	case 0x4000:
		m.opSne()
		return
	case 0x5000:
		m.opSeReg()
		return
	case 0x6000:
		m.opLdVx()
		return
	case 0x7000:
		m.opAdd()
		return
	case 0x9000:
		m.opSneReg()
		return
	case 0xA000:
		m.opLdI()
		return
	case 0xB000:
		m.opJmpV0()
		return
	case 0xC000:
		m.opRnd()
		return
	case 0xD000:
		m.opDrw()
		return
	}

	// 0x8nn0 -> 0x8nnE Cases
	switch opcode & 0x000F {
	case 0x0:
		m.opLdReg()
		return
	case 0x1:
		m.opOr()
		return
	case 0x2:
		m.opAnd()
		return
	case 0x3:
		m.opXor()
		return
	case 0x4:
		m.opAddReg()
		return
	case 0x5:
		m.opSub()
		return
	case 0x6:
		m.opShr()
		return
	case 0x7:
		m.opSubn()
		return
	case 0xE:
		m.opShl()
		return
	}

	// ExAA Cases
	switch opcode & 0xF0FF {
	case 0xE09E:
		m.opSkp()
		return
	case 0xE0A1:
		m.opSknp()
		return
	}

	// FxAA Cases
	if opcode&0xF000 == 0xF000 {
		switch opcode & 0xF0FF {
		case 0xF007:
			m.opLdVxDT()
			return
		case 0xF00A:
			m.opLdVxK()
			return
		case 0xF015:
			m.opLdDTVx()
			return
		case 0xF018:
			m.opLdSTVx()
			return
		case 0xF01E:
			m.opAddIVx()
			return
		case 0xF029:
			m.opLdFVx()
			return
		case 0xF033:
			m.opLdBVx()
			return
		case 0xF055:
			m.opLdDMemVx()
			return
		case 0xF065:
			m.opLdVxMem()
			return
		}
	}
}
