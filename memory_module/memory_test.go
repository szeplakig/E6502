package memory_module

import (
	"testing"
)

func TestWriteByteAndReadByteReversible(t *testing.T) {
	memory := NewMemory()
	var address Word = 0x8000
	var value Byte = 0x42
	memory.WB(address, value)
	var readValue = memory.RB(address)
	if value != readValue {
		t.Error("Written and read values do no match!")
	}
}

func TestWriteWordAndReadWordReversible(t *testing.T) {
	memory := NewMemory()
	var address Word = 0x8000
	var value Word = 0x4284
	memory.WW(address, value)
	var readValue = memory.RW(address)
	if value != readValue {
		t.Error("Written and read values do no match!")
	}
}
