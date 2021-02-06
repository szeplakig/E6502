package CPU

import (
	"E6502/Memory"
	"fmt"
)

type Byte = uint8
type Word = uint16

type Instruction = Byte

const (
	LDA_IM = 0xA9
	LDA_ZP = 0xA5
	LDA_ZX = 0xB5
	LDA_AB = 0xAD
	LDA_AX = 0xBD
	LDA_AY = 0xB9
	LDA_IX = 0xA1
	LDA_IY = 0xB1
)

type CPU struct {
	PC Word // Program Counter
	SP Byte // Stack Pointer

	A Byte // Accumulator
	X Byte // Index Register X
	Y Byte // Index Register Y

	C bool // Carry Flag
	Z bool // Zero Flag
	I bool // Interrupt Disable
	D bool // Decimal Mode
	B bool // Break Command
	V bool // Overflow Flag
	N bool // Negative Flag
}

func (cpu *CPU) Reset() {
	cpu.PC = 0xFFFC
	cpu.SP = 0xFF

	cpu.A = 0x0
	cpu.X = 0x0
	cpu.Y = 0x0

	cpu.C = false
	cpu.Z = false
	cpu.I = false
	cpu.D = false
	cpu.B = false
	cpu.V = false
	cpu.N = false
}

func NewCPU() CPU {
	cpu := CPU{}
	cpu.Reset()
	return cpu
}

func (cpu *CPU) FetchBytePC(cycles *int, memory *Memory.Memory) Byte {
	value := cpu.FetchByte(cycles, memory, cpu.PC)
	cpu.PC++
	return value
}

func (cpu *CPU) FetchByte(cycles *int, memory *Memory.Memory, address Word) Byte {
	value := memory.RB(address)
	*cycles--
	return value
}

func (cpu *CPU) Execute(cycles int, memory *Memory.Memory) {
	instruction_map := map[Byte]func(*int, *Memory.Memory){
		LDA_IM: cpu.H_LDA_IM,
		LDA_ZP: cpu.H_LDA_ZP,
		LDA_ZX: cpu.H_LDA_ZX,
		LDA_AB: cpu.H_LDA_AB,
	}

	for cycles > 0 {
		next_ins := cpu.FetchBytePC(&cycles, memory)
		if handler, ok := instruction_map[next_ins]; ok {
			handler(&cycles, memory)
		} else {
			fmt.Print("Unknown insturction: ", next_ins)
			return
		}
	}
}

func (cpu *CPU) H_LDA_IM(cycles *int, memory *Memory.Memory) {
	value := cpu.FetchBytePC(cycles, memory)
	cpu.A = value
	*cycles--
	if cpu.A == 0 {
		cpu.Z = true
	}
	if (cpu.A >> 7) == 1 {
		cpu.N = true
	}
}

func (cpu *CPU) H_LDA_ZP(cycles *int, memory *Memory.Memory) {
	address := cpu.FetchBytePC(cycles, memory)
	value := cpu.FetchByte(cycles, memory, Word(address))
	cpu.A = value
	*cycles--
	if cpu.A == 0 {
		cpu.Z = true
	}
	if (cpu.A >> 7) == 1 {
		cpu.N = true
	}
}

func (cpu *CPU) H_LDA_ZX(cycles *int, memory *Memory.Memory) {
	address := cpu.FetchBytePC(cycles, memory)
	address += cpu.X
	*cycles--
	value := cpu.FetchByte(cycles, memory, Word(address))
	cpu.A = value
	*cycles--
	if cpu.A == 0 {
		cpu.Z = true
	}
	if (cpu.A >> 7) == 1 {
		cpu.N = true
	}
}

func (cpu *CPU) H_LDA_AB(cycles *int, memory *Memory.Memory) {
	byte1 := cpu.FetchBytePC(cycles, memory)
	byte2 := cpu.FetchBytePC(cycles, memory)
	address := uint16(byte1) | uint16(byte2)<<8
	address += cpu.X
	value := cpu.FetchByte(cycles, memory, address)
	cpu.A = value
	*cycles--
	if cpu.A == 0 {
		cpu.Z = true
	}
	if (cpu.A >> 7) == 1 {
		cpu.N = true
	}
}
