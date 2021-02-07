package cpu

import "E6502/memory"

func (cpu *CPU) ImmediateAddressing(cycles *int, memory *memory.Memory) Word {
	address := cpu.PC
	return address
}

func (cpu *CPU) ZeroPageAddressing(cycles *int, memory *memory.Memory) Word {
	address := Word(cpu.fetchBytePC(cycles, memory))
	return address
}

func (cpu *CPU) ZeroPageXAddressing(cycles *int, memory *memory.Memory) Word {
	address := cpu.fetchBytePC(cycles, memory)
	offset_address := Word(address) + Word(cpu.X)
	*cycles--
	return offset_address
}

func (cpu *CPU) ZeroPageYAddressing(cycles *int, memory *memory.Memory) Word {
	address := cpu.fetchBytePC(cycles, memory)
	offset_address := Word(address) + Word(cpu.Y)
	*cycles--
	return offset_address
}

func (cpu *CPU) AbsoluteAddressing(cycles *int, memory *memory.Memory) Word {
	lower := cpu.fetchBytePC(cycles, memory)
	upper := cpu.fetchBytePC(cycles, memory)
	address := Word(upper)<<8 | Word(lower)
	return address
}

func (cpu *CPU) AbsoluteXAddressing(cycles *int, memory *memory.Memory) Word {
	lower := cpu.fetchBytePC(cycles, memory)
	upper := cpu.fetchBytePC(cycles, memory)
	address := Word(upper)<<8 | Word(lower)
	offset_address := add(cycles, address, Word(cpu.X))
	return offset_address
}

func (cpu *CPU) AbsoluteYAddressing(cycles *int, memory *memory.Memory) Word {
	lower := cpu.fetchBytePC(cycles, memory)
	upper := cpu.fetchBytePC(cycles, memory)
	address := Word(upper)<<8 | Word(lower)
	offset_address := add(cycles, address, Word(cpu.Y))
	return offset_address
}

func (cpu *CPU) IndirectXAddressing(cycles *int, memory *memory.Memory) Word {
	address := cpu.fetchBytePC(cycles, memory)
	offset_address := Word(address + cpu.X)
	*cycles--
	indirect_address := cpu.fetchWord(cycles, memory, offset_address)
	return indirect_address
}

func (cpu *CPU) IndirectYAddressing(cycles *int, memory *memory.Memory) Word {
	address := Word(cpu.fetchBytePC(cycles, memory))
	indirect_address := cpu.fetchWord(cycles, memory, address)
	offset_address := add(cycles, indirect_address, Word(cpu.Y))
	return offset_address
}
