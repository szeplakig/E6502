package cpumodule

import (
	. "E6502/memorymodule"
	. "E6502/utils"
	"testing"
)

func Test_INS_STA_ZP(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0x42
	var address Byte = 0x42
	cpu.A = val
	memory.WB(0xFFFC, STA_ZP)
	memory.WB(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(3, &memory)

	ValidateStore(success, cycles, memory.RB(Word(address)), val, cpu, cpuCopy, t)
}

func Test_INS_STA_ZX(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	cpu.X = 0x0F
	var val Byte = 0x42
	var address Byte = 0x80
	cpu.A = val
	memory.WB(0xFFFC, STA_ZX)
	memory.WB(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(4, &memory)

	ValidateStore(success, cycles, memory.RB(Word(address+cpu.X)), val, cpu, cpuCopy, t)
}

func Test_INS_STA_AB(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0x42
	var address Word = 0x32F0
	cpu.A = val
	memory.WB(0xFFFC, STA_AB)
	memory.WW(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(4, &memory)

	ValidateStore(success, cycles, memory.RB(address), val, cpu, cpuCopy, t)
}

func Test_INS_STA_AX(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0x42
	var address Word = 0x2000
	cpu.X = 0x92
	cpu.A = val
	memory.WB(0xFFFC, STA_AX)
	memory.WW(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(5, &memory)
	ValidateStore(success, cycles, memory.RB(address+Word(cpu.X)), val, cpu, cpuCopy, t)
}

func Test_INS_STA_AY(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0x42
	var address Word = 0x2000
	cpu.Y = 0x92
	cpu.A = val
	memory.WB(0xFFFC, STA_AY)
	memory.WW(0xFFFD, address)

	cpuCopy := cpu
	success, cycles := cpu.Execute(5, &memory)
	ValidateStore(success, cycles, memory.RB(address+Word(cpu.Y)), val, cpu, cpuCopy, t)
}

func Test_INS_STA_IX(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()

	var val Byte = 0x42
	var address Word = 0x4242
	cpu.A = val
	cpu.X = 0xFF
	memory.WW(0x0005, address)
	memory.WB(0xFFFC, STA_IX)
	memory.WB(0xFFFD, 0x06)

	cpuCopy := cpu
	success, cycles := cpu.Execute(6, &memory)
	ValidateStore(success, cycles, memory.RB(address), val, cpu, cpuCopy, t)
}

func Test_INS_STA_IY(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()

	var val Byte = 0x42
	var address Word = 0x4242
	cpu.A = val
	cpu.Y = 0x04
	memory.WW(0x0006, address)
	memory.WB(0xFFFC, STA_IY)
	memory.WB(0xFFFD, 0x06)

	cpuCopy := cpu
	success, cycles := cpu.Execute(6, &memory)
	ValidateStore(success, cycles, memory.RB(address+Word(cpu.Y)), val, cpu, cpuCopy, t)
}
