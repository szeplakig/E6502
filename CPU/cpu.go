package CPU

import (
	"E6502/Memory"
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

func Add(cycles *int, a Word, b Word) Word {
	if (a%0xFF)+(b%0xFF) >= 0xFF {
		*cycles--
	}
	return a + b
}

func (cpu *CPU) WriteA(cycles *int, value Byte) {
	cpu.A = value
}

func (cpu *CPU) WriteX(cycles *int, value Byte) {
	cpu.X = value
}

func (cpu *CPU) WriteY(cycles *int, value Byte) {
	cpu.Y = value
}

func (cpu *CPU) LoadA(cycles *int, memory *Memory.Memory, address Word) {
	value := cpu.FetchByte(cycles, memory, address)
	cpu.WriteA(cycles, value)
}

func (cpu *CPU) LoadX(cycles *int, memory *Memory.Memory, address Word) {
	value := cpu.FetchByte(cycles, memory, address)
	cpu.WriteX(cycles, value)
}

func (cpu *CPU) LoadY(cycles *int, memory *Memory.Memory, address Word) {
	value := cpu.FetchByte(cycles, memory, address)
	cpu.WriteY(cycles, value)
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

func (cpu *CPU) FetchWord(cycles *int, memory *Memory.Memory, address Word) Word {
	value := memory.RW(address)
	*cycles -= 2
	return value
}

func (cpu *CPU) Execute(cycles int, memory *Memory.Memory) (bool, int) {
	instruction_map := map[Byte]func(*int, *Memory.Memory){
		LDA_IM: cpu.H_LDA_IM,
		LDA_ZP: cpu.H_LDA_ZP,
		LDA_ZX: cpu.H_LDA_ZX,
		LDA_AB: cpu.H_LDA_AB,
		LDA_AX: cpu.H_LDA_AX,
		LDA_AY: cpu.H_LDA_AY,
		LDA_IX: cpu.H_LDA_IX,
		LDA_IY: cpu.H_LDA_IY,

		LDX_IM: cpu.H_LDX_IM,
		LDX_ZP: cpu.H_LDX_ZP,
		LDX_ZY: cpu.H_LDX_ZY,
		LDX_AB: cpu.H_LDX_AB,
		LDX_AY: cpu.H_LDX_AY,

		LDY_IM: cpu.H_LDY_IM,
		LDY_ZP: cpu.H_LDY_ZP,
		LDY_ZX: cpu.H_LDY_ZX,
		LDY_AB: cpu.H_LDY_AB,
		LDY_AX: cpu.H_LDY_AX,
	}

	for cycles > 0 {
		next_ins := cpu.FetchBytePC(&cycles, memory)
		if handler, ok := instruction_map[next_ins]; ok {
			handler(&cycles, memory)
		} else {
			return false, cycles
		}
	}
	//fmt.Println(cycles)
	return cycles == 0, cycles
}
