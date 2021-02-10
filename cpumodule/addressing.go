package cpumodule

import (
	"E6502/memorymodule"
	. "E6502/utils"
)

func (cpu *CPU) ImmediateAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := cpu.PC
	return address
}

func (cpu *CPU) ZeroPageAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := Word(cpu.fetchBytePC(cycles, mem))
	return address
}

func (cpu *CPU) ZeroPageXAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := cpu.fetchBytePC(cycles, mem)
	offset_address := Word(address) + Word(cpu.X)
	*cycles--
	return offset_address
}

func (cpu *CPU) ZeroPageYAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := cpu.fetchBytePC(cycles, mem)
	offset_address := Word(address) + Word(cpu.Y)
	*cycles--
	return offset_address
}

func (cpu *CPU) AbsoluteAddressing(cycles *int, mem *memorymodule.Memory) Word {
	lower := cpu.fetchBytePC(cycles, mem)
	upper := cpu.fetchBytePC(cycles, mem)
	address := Word(upper)<<8 | Word(lower)
	return address
}

func (cpu *CPU) AbsoluteXAddressing(cycles *int, mem *memorymodule.Memory) Word {
	lower := cpu.fetchBytePC(cycles, mem)
	upper := cpu.fetchBytePC(cycles, mem)
	address := Word(upper)<<8 | Word(lower)
	offset_address := add(cycles, address, Word(cpu.X))
	return offset_address
}

func (cpu *CPU) AbsoluteYAddressing(cycles *int, mem *memorymodule.Memory) Word {
	lower := cpu.fetchBytePC(cycles, mem)
	upper := cpu.fetchBytePC(cycles, mem)
	address := Word(upper)<<8 | Word(lower)
	offset_address := add(cycles, address, Word(cpu.Y))
	return offset_address
}

func (cpu *CPU) IndirectXAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := cpu.fetchBytePC(cycles, mem)
	offset_address := Word(address + cpu.X)
	*cycles--
	indirect_address := fetchWord(cycles, mem, offset_address)
	return indirect_address
}

func (cpu *CPU) IndirectYAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := Word(cpu.fetchBytePC(cycles, mem))
	indirect_address := fetchWord(cycles, mem, address)
	offset_address := add(cycles, indirect_address, Word(cpu.Y))
	return offset_address
}
