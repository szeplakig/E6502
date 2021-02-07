package memory

// The Memory structure of the 6502
type Memory struct {
	Data []Byte
}

// Reset the memory
func (memory *Memory) Reset() {
	memory.Data = make([]Byte, 0xFFFF)
}

// Create a new memory
func NewMemory() Memory {
	memory := Memory{}
	memory.Reset()
	return memory
}

// Write a Byte to the memory at address
func (memory *Memory) WB(address Word, value Byte) {
	memory.Data[address%0xFFFF] = value
}

// Write a Word to the memory at address (big endian)
func (memory *Memory) WW(address Word, value Word) {
	memory.Data[address%0xFFFF] = Byte(value)
	memory.Data[(address+1)%0xFFFF] = Byte(value >> 8)
}

// Read a Byte from memory at address
func (memory *Memory) RB(address Word) Byte {
	return memory.Data[address%0xFFFF]
}

// Read a Word from memory at address
func (memory *Memory) RW(address Word) Word {
	return Word(memory.Data[address%0xFFFF]) | Word(memory.Data[(address+1)%0xFFFF])<<8
}
