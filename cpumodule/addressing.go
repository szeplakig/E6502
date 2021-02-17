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
	offsetAddress := Word(address) + Word(cpu.X)
	*cycles--
	return offsetAddress
}

func (cpu *CPU) ZeroPageYAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := cpu.fetchBytePC(cycles, mem)
	offsetAddress := Word(address) + Word(cpu.Y)
	*cycles--
	return offsetAddress
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
	offsetAddress := add(cycles, address, Word(cpu.X))
	return offsetAddress
}

func (cpu *CPU) AbsoluteXAddressingLong(cycles *int, mem *memorymodule.Memory) Word {
	lower := cpu.fetchBytePC(cycles, mem)
	upper := cpu.fetchBytePC(cycles, mem)
	address := Word(upper)<<8 | Word(lower)
	offsetAddress := address + Word(cpu.X)
	*cycles--
	return offsetAddress
}

func (cpu *CPU) AbsoluteYAddressing(cycles *int, mem *memorymodule.Memory) Word {
	lower := cpu.fetchBytePC(cycles, mem)
	upper := cpu.fetchBytePC(cycles, mem)
	address := Word(upper)<<8 | Word(lower)
	offsetAddress := add(cycles, address, Word(cpu.Y))
	return offsetAddress
}

func (cpu *CPU) AbsoluteYAddressingLong(cycles *int, mem *memorymodule.Memory) Word {
	lower := cpu.fetchBytePC(cycles, mem)
	upper := cpu.fetchBytePC(cycles, mem)
	address := Word(upper)<<8 | Word(lower)
	offsetAddress := address + Word(cpu.Y)
	*cycles--
	return offsetAddress
}

func (cpu *CPU) IndirectXAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := cpu.fetchBytePC(cycles, mem)
	offsetAddress := Word(Byte(address + cpu.X))
	*cycles--
	indirectAddress := fetchWord(cycles, mem, offsetAddress)
	return indirectAddress
}

func (cpu *CPU) IndirectYAddressing(cycles *int, mem *memorymodule.Memory) Word {
	address := Word(cpu.fetchBytePC(cycles, mem))
	indirectAddress := fetchWord(cycles, mem, address)
	offsetAddress := add(cycles, indirectAddress, Word(cpu.Y))
	return offsetAddress
}

func (cpu *CPU) IndirectYAddressingLong(cycles *int, mem *memorymodule.Memory) Word {
	address := Word(cpu.fetchBytePC(cycles, mem))
	indirectAddress := fetchWord(cycles, mem, address)
	offsetAddress := indirectAddress + Word(cpu.Y)
	*cycles--
	return offsetAddress
}
