package CPU

import (
	"E6502/Memory"
)

func (cpu *CPU) SetLDAFlags() {
	cpu.Z = cpu.A == 0
	cpu.N = (cpu.A >> 7) == 1
}

func (cpu *CPU) H_LDA_IM(cycles *int, memory *Memory.Memory) {
	value := cpu.FetchBytePC(cycles, memory)
	cpu.WriteA(cycles, value)
	cpu.SetLDAFlags()
}

func (cpu *CPU) H_LDA_ZP(cycles *int, memory *Memory.Memory) {
	address := cpu.ZeroPageAddressing(cycles, memory)
	cpu.LoadA(cycles, memory, address)
	cpu.SetLDAFlags()
}

func (cpu *CPU) H_LDA_ZX(cycles *int, memory *Memory.Memory) {
	address := cpu.ZeroPageXAddressing(cycles, memory)
	cpu.LoadA(cycles, memory, address)
	cpu.SetLDAFlags()
}

func (cpu *CPU) H_LDA_AB(cycles *int, memory *Memory.Memory) {
	address := cpu.AbsoluteAddressing(cycles, memory)
	cpu.LoadA(cycles, memory, address)
	cpu.SetLDAFlags()
}

func (cpu *CPU) H_LDA_AX(cycles *int, memory *Memory.Memory) {
	address := cpu.AbsoluteXAddressing(cycles, memory)
	cpu.LoadA(cycles, memory, address)
	cpu.SetLDAFlags()
}

func (cpu *CPU) H_LDA_AY(cycles *int, memory *Memory.Memory) {
	address := cpu.AbsoluteYAddressing(cycles, memory)
	cpu.LoadA(cycles, memory, address)
	cpu.SetLDAFlags()
}

func (cpu *CPU) H_LDA_IX(cycles *int, memory *Memory.Memory) {
	address := cpu.IdirectXAddressing(cycles, memory)
	cpu.LoadA(cycles, memory, address)
	cpu.SetLDAFlags()
}

func (cpu *CPU) H_LDA_IY(cycles *int, memory *Memory.Memory) {
	address := cpu.IndirectYAdressing(cycles, memory)
	cpu.LoadA(cycles, memory, address)
	cpu.SetLDAFlags()
}
