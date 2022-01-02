package main

import (
	"fmt"
	"io/ioutil"
)

// op interface
type Op interface {
	Execute()
}

// unknown op
type UnknownOp struct{}

func (op UnknownOp) Execute() {}

func (op UnknownOp) String() string {
	return "unknown"
}

// LUI
type LUI struct {
	rd  uint32
	imm uint32
}

func (op LUI) Execute() {}

func (op LUI) String() string {
	return "LUI"
}

// decode
func decode(insn uint32) Op {
	funct7 := insn & 0x7f

	if funct7 == 0b0110111 {
		rd := (insn >> 7) & 0x1f
		return LUI{rd: rd}
	}

	return UnknownOp{}
}

func main() {
	bytes, err := ioutil.ReadFile("./rafi-prebuilt-binary/riscv-tests/isa/rv32ui-p-add.bin")
	if err != nil {
		panic(err)
	}

	for i := 0; i+4 <= len(bytes); i += 4 {
		insn := uint32(bytes[i+3])<<24 |
			uint32(bytes[i+2])<<16 |
			uint32(bytes[i+1])<<8 |
			uint32(bytes[i])
		op := decode(insn)
		fmt.Printf("%08x %s\n", insn, op)
	}
}
