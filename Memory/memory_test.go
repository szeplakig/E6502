package Memory

import (
	"testing"
)


func TestWriteByteAndReadByteReversible(t *testing.T) {
	memory := NewMemory()
	var address Word = 0x8000
	var value Byte = 0x42
	memory.WB(address, value)
	var read_value = memory.RB(address)
	if value != read_value {
		t.Error("Written and read values do no match!")
	}
}

func TestWriteWordAndReadWordReversible(t *testing.T) {
	memory := NewMemory()
	var address Word = 0x8000
	var value Word = 0x4284
	memory.WW(address, value)
	var read_value = memory.RW(address)
	if value != read_value {
		t.Error("Written and read values do no match!")
	}
}
