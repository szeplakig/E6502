package CPU

import (
	"E6502/Memory"
)

func (cpu *CPU) ZeroPageAddressing(cycles *int, memory *Memory.Memory) Word {
	address := Word(cpu.FetchBytePC(cycles, memory))
	return address
}

func (cpu *CPU) ZeroPageXAddressing(cycles *int, memory *Memory.Memory) Word {
	address := cpu.FetchBytePC(cycles, memory)
	offset_address := Word(address) + Word(cpu.X)
	*cycles--
	return offset_address
}

func (cpu *CPU) ZeroPageYAddressing(cycles *int, memory *Memory.Memory) Word {
	address := cpu.FetchBytePC(cycles, memory)
	offset_address := Word(address) + Word(cpu.Y)
	*cycles--
	return offset_address
}

func (cpu *CPU) AbsoluteAddressing(cycles *int, memory *Memory.Memory) Word {
	lower := cpu.FetchBytePC(cycles, memory)
	upper := cpu.FetchBytePC(cycles, memory)
	address := Word(upper)<<8 | Word(lower)
	return address
}

func (cpu *CPU) AbsoluteXAddressing(cycles *int, memory *Memory.Memory) Word {
	lower := cpu.FetchBytePC(cycles, memory)
	upper := cpu.FetchBytePC(cycles, memory)
	address := Word(upper)<<8 | Word(lower)
	offset_address := Add(cycles, address, Word(cpu.X))
	return offset_address
}

func (cpu *CPU) AbsoluteYAddressing(cycles *int, memory *Memory.Memory) Word {
	lower := cpu.FetchBytePC(cycles, memory)
	upper := cpu.FetchBytePC(cycles, memory)
	address := Word(upper)<<8 | Word(lower)
	offset_address := Add(cycles, address, Word(cpu.Y))
	return offset_address
}

func (cpu *CPU) IdirectXAddressing(cycles *int, memory *Memory.Memory) Word {
	address := cpu.FetchBytePC(cycles, memory)
	offset_address := Word(address + cpu.X)
	*cycles--
	indirect_address := cpu.FetchWord(cycles, memory, offset_address)
	return indirect_address
}

func (cpu *CPU) IndirectYAdressing(cycles *int, memory *Memory.Memory) Word {
	address := Word(cpu.FetchBytePC(cycles, memory))
	indirect_address := cpu.FetchWord(cycles, memory, address)
	offset_address := Add(cycles, indirect_address, Word(cpu.Y))
	return offset_address
}
