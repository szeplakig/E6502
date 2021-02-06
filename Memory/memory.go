package Memory

type Byte = uint8
type Word = uint16

type Memory struct {
	Data []Byte
}

func (memory *Memory) Reset() {
	memory.Data = make([]Byte, 0xFFFF)
}

func NewMemory() Memory {
	memory := Memory{}
	memory.Reset()
	return memory
}

func (memory *Memory) WB(address Word, value Byte) {
	memory.Data[address] = value
}

func (memory *Memory) WW(address Word, value Word) {
	memory.Data[address] = Byte(value)
	memory.Data[address+1] = Byte(value >> 8)
}

func (memory *Memory) RB(address Word) Byte {
	return memory.Data[address]
}

func (memory *Memory) RW(address Word) Word {
	return Word(memory.Data[address]) | Word(memory.Data[address+1])<<8
}
