package cpumodule

import (
	. "E6502/memorymodule"
	. "E6502/utils"
	"testing"
)

func Test_INS_STX_ZP(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0x42
	var address Byte = 0x42
	cpu.X = val
	memory.WB(0xFFFC, STX_ZP)
	memory.WB(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(3, &memory)

	ValidateStore(success, cycles, memory.RB(Word(address)), val, cpu, cpuCopy, t)
}

func Test_INS_STX_ZY(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	cpu.Y = 0x0F
	var val Byte = 0x42
	var address Byte = 0x80
	cpu.X = val
	memory.WB(0xFFFC, STX_ZY)
	memory.WB(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(4, &memory)

	ValidateStore(success, cycles, memory.RB(Word(address+cpu.Y)), val, cpu, cpuCopy, t)
}

func Test_INS_STX_AB(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0x42
	var address Word = 0x32F0
	cpu.X = val
	memory.WB(0xFFFC, STX_AB)
	memory.WW(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(4, &memory)

	ValidateStore(success, cycles, memory.RB(address), val, cpu, cpuCopy, t)
}
