package main

import (
	"fmt"
	"io/ioutil"
)

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
		fmt.Printf("%08x\n", insn)
	}
}
