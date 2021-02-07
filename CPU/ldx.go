package CPU

import (
	"E6502/Memory"
)

func (cpu *CPU) SetLDXFlags() {
	cpu.Z = cpu.X == 0
	cpu.N = (cpu.X >> 7) == 1
}

func (cpu *CPU) H_LDX_IM(cycles *int, memory *Memory.Memory) {
	value := cpu.FetchBytePC(cycles, memory)
	cpu.WriteX(cycles, value)
	cpu.SetLDXFlags()
}

func (cpu *CPU) H_LDX_ZP(cycles *int, memory *Memory.Memory) {
	address := cpu.ZeroPageAddressing(cycles, memory)
	cpu.LoadX(cycles, memory, address)
	cpu.SetLDXFlags()
}

func (cpu *CPU) H_LDX_ZY(cycles *int, memory *Memory.Memory) {
	address := cpu.ZeroPageYAddressing(cycles, memory)
	cpu.LoadX(cycles, memory, address)
	cpu.SetLDXFlags()
}

func (cpu *CPU) H_LDX_AB(cycles *int, memory *Memory.Memory) {
	address := cpu.AbsoluteAddressing(cycles, memory)
	cpu.LoadX(cycles, memory, address)
	cpu.SetLDXFlags()
}

func (cpu *CPU) H_LDX_AY(cycles *int, memory *Memory.Memory) {
	address := cpu.AbsoluteYAddressing(cycles, memory)
	cpu.LoadX(cycles, memory, address)
	cpu.SetLDXFlags()
}
