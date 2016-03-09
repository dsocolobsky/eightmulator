package main

import (
	"fmt"
)

// 0x0 Cases

func (m *Machine) opCls() {

}

func (m *Machine) opRet() {

}

// 0x1 -> 0x7 Cases

func (m *Machine) opJmp() {
	m.pc = m.opcode & 0x0FFF
	fmt.Printf("JMP %d\n", m.pc)
}

func (m *Machine) opCall() {

}

func (m *Machine) opSe() {

}

func (m *Machine) opSne() {

}

func (m *Machine) opSeReg() {

}

func (m *Machine) opLdVx() {
    reg := (m.opcode & 0x0F00) >> 8 
    val := m.opcode & 0x00FF
    
    m.registers[reg] = uint8(val)
    
    m.pc += 2
}

func (m *Machine) opAdd() {

}

// 0x8nn0 -> 0x8nnE

func (m *Machine) opLdReg() {

}

func (m *Machine) opOr() {

}

func (m *Machine) opAnd() {

}

func (m *Machine) opXor() {

}

func (m *Machine) opAddReg() {

}

func (m *Machine) opSub() {

}

func (m *Machine) opShr() {

}

func (m *Machine) opSubn() {

}

func (m *Machine) opShl() {

}

// 0x9 -> 0xD Cases

func (m *Machine) opSneReg() {

}

func (m *Machine) opLdI() {

}

func (m *Machine) opJmpV0() {

}

func (m *Machine) opRnd() {

}

func (m *Machine) opDrw() {

}

// 0xE Cases

func (m *Machine) opSkp() {

}

func (m *Machine) opSknp() {

}

// 0xF Cases

func (m *Machine) opLdVxDT() {

}

func (m *Machine) opLdVxK() {

}

func (m *Machine) opLdDTVx() {

}

func (m *Machine) opLdSTVx() {

}

func (m *Machine) opAddIVx() {

}

func (m *Machine) opLdFVx() {

}

func (m *Machine) opLdBVx() {

}

func (m *Machine) opLdDMemVx() {

}

func (m *Machine) opLdVxMem() {

}
