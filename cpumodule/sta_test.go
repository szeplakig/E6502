package cpumodule

import (
	. "E6502/memorymodule"
	. "E6502/utils"
	"testing"
)

func Test_INS_STA_ZP(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0xF0
	var address Byte = 0x42
	cpu.A = val
	memory.WB(0xFFFC, STA_ZP)
	memory.WB(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(3, &memory)

	ValidateStore(success, cycles, memory.RB(Word(address)), val, cpu, cpuCopy, t)
}
