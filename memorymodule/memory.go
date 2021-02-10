package memorymodule

import (
	. "E6502/utils"
)

type Memory struct { // The Memory structure of the 6502
	Data []Byte
}

// Reset the memory
func (mem *Memory) Reset() {
	mem.Data = make([]Byte, 0xFFFF)
}

// Create a new memory
func NewMemory() Memory {
	memory := Memory{}
	memory.Reset()
	return memory
}

// Write a Byte to the memory at address
func (mem *Memory) WB(address Word, value Byte) {
	mem.Data[address%0xFFFF] = value
}

// Write a Word to the memory at address (big endian)
func (mem *Memory) WW(address Word, value Word) {
	mem.Data[address%0xFFFF] = Byte(value)
	mem.Data[(address+1)%0xFFFF] = Byte(value >> 8)
}

// Read a Byte from memory at address
func (mem *Memory) RB(address Word) Byte {
	return mem.Data[address%0xFFFF]
}

// Read a Word from memory at address
func (mem *Memory) RW(address Word) Word {
	return Word(mem.Data[address%0xFFFF]) | Word(mem.Data[(address+1)%0xFFFF])<<8
}
