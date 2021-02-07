package CPU

import (
	"E6502/Memory"
)

func (cpu *CPU) SetLDYFlags() {
	cpu.Z = cpu.Y == 0
	cpu.N = (cpu.Y >> 7) == 1
}

func (cpu *CPU) H_LDY_IM(cycles *int, memory *Memory.Memory) {
	value := cpu.FetchBytePC(cycles, memory)
	cpu.WriteY(cycles, value)
	cpu.SetLDYFlags()
}

func (cpu *CPU) H_LDY_ZP(cycles *int, memory *Memory.Memory) {
	address := cpu.ZeroPageAddressing(cycles, memory)
	cpu.LoadY(cycles, memory, address)
	cpu.SetLDYFlags()
}

func (cpu *CPU) H_LDY_ZX(cycles *int, memory *Memory.Memory) {
	address := cpu.ZeroPageXAddressing(cycles, memory)
	cpu.LoadY(cycles, memory, address)
	cpu.SetLDYFlags()
}

func (cpu *CPU) H_LDY_AB(cycles *int, memory *Memory.Memory) {
	address := cpu.AbsoluteAddressing(cycles, memory)
	cpu.LoadY(cycles, memory, address)
	cpu.SetLDYFlags()
}

func (cpu *CPU) H_LDY_AX(cycles *int, memory *Memory.Memory) {
	address := cpu.AbsoluteXAddressing(cycles, memory)
	cpu.LoadY(cycles, memory, address)
	cpu.SetLDYFlags()
}
