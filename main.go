package main

import (
	"fmt"
	"io/ioutil"
)

// Integer Register
var IntRegNames = map[uint32]string{
	0:  "zero",
	1:  "ra",
	2:  "sp",
	3:  "gp",
	4:  "tp",
	5:  "t0",
	6:  "t1",
	7:  "t2",
	8:  "s0",
	9:  "s1",
	10: "a0",
	11: "a1",
	12: "a2",
	13: "a3",
	14: "a4",
	15: "a5",
	16: "a6",
	17: "a7",
	18: "s2",
	19: "s3",
	20: "s4",
	21: "s5",
	22: "s6",
	23: "s7",
	24: "s8",
	25: "s9",
	26: "s10",
	27: "s11",
	28: "t3",
	29: "t4",
	30: "t5",
	31: "t6",
}

type IntReg struct {
	data [32]uint32
}

func (reg IntReg) Read(index uint32) uint32 {
	return reg.data[index]
}

func (reg IntReg) Write(index uint32, value uint32) {
	if index == 0 {
		return
	}
	reg.data[index] = value
}

// CPU core
type Core struct {
	pc     uint32
	nextPc uint32
	reg    IntReg
}

// op interface
type Op interface {
	Execute(core *Core)
}

// Unknown op
type UnknownOp struct{}

func (op UnknownOp) Execute(core *Core) {}

func (op UnknownOp) String() string {
	return "unknown"
}

// LUI
type LUI struct {
	rd  uint32
	imm uint32
}

func (op LUI) Execute(core *Core) {
	core.reg.Write(op.rd, op.imm)
}

func (op LUI) String() string {
	return fmt.Sprint("lui ", IntRegNames[op.rd], ",", op.imm)
}

// Emulation logic
func pick(data uint32, lsb uint32, width uint32) uint32 {
	return (data >> lsb) & ((1 << width) - 1)
}

func decode(insn uint32) Op {
	funct7 := pick(insn, 0, 7)

	if funct7 == 0b0110111 {
		rd := pick(insn, 7, 5)
		imm := pick(insn, 12, 20)
		return LUI{rd: rd, imm: imm}
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
