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
	memory.Data[address] = uint8(value)
	memory.Data[address+1] = uint8(value>>8)
}

func (memory *Memory) RB(address Word) Byte {
	return memory.Data[address]
}

func (memory *Memory) RW(address Word) Word {
	return uint16(memory.Data[address]) | uint16(memory.Data[address+1])<<8
}
