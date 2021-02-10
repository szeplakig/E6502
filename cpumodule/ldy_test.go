package cpumodule

import (
	. "E6502/memorymodule"
	. "E6502/utils"
	"testing"
)

func Test_INS_LDY_IM(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0xF0
	memory.WB(0xFFFC, LDY_IM)
	memory.WB(0xFFFD, val)

	cpuCopy := cpu
	success, cycles := cpu.Execute(2, &memory)

	ValidateLoad(success, cycles, cpu.Y, val, cpu, cpuCopy, t)
}

func Test_INS_LDY_ZP(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0xF0
	memory.WB(0x0000, val)
	memory.WB(0xFFFC, LDY_ZP)
	memory.WB(0xFFFD, 0x00)

	cpuCopy := cpu
	success, cycles := cpu.Execute(3, &memory)

	ValidateLoad(success, cycles, cpu.Y, val, cpu, cpuCopy, t)
}

func Test_INS_LDY_ZX(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0xF0
	cpu.X = 0x0F
	memory.WB(0x008F, val)
	memory.WB(0xFFFC, LDY_ZX)
	memory.WB(0xFFFD, 0x80)

	cpuCopy := cpu
	success, cycles := cpu.Execute(4, &memory)

	ValidateLoad(success, cycles, cpu.Y, val, cpu, cpuCopy, t)
}

func Test_INS_LDY_AB(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0xF0
	memory.WB(0x4224, val)
	memory.WB(0xFFFC, LDY_AB)
	memory.WW(0xFFFD, 0x4224)

	cpuCopy := cpu
	success, cycles := cpu.Execute(4, &memory)

	ValidateLoad(success, cycles, cpu.Y, val, cpu, cpuCopy, t)
}

func Test_INS_LDY_AX(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0xF0
	cpu.X = 0x92
	memory.WB(0x2092, val)
	memory.WB(0xFFFC, LDY_AX)
	memory.WW(0xFFFD, 0x2000)

	cpuCopy := cpu
	success, cycles := cpu.Execute(4, &memory)

	ValidateLoad(success, cycles, cpu.Y, val, cpu, cpuCopy, t)
}

func Test_INS_LDY_AX_CROSSES_PAGE_BOUNDARY(t *testing.T) {
	cpu := NewCPU()
	memory := NewMemory()
	var val Byte = 0xF0
	cpu.X = 0x2
	memory.WB(0x20E0, val)
	memory.WB(0xFFFC, LDY_AX)
	memory.WW(0xFFFD, 0x20DE)

	cpuCopy := cpu
	success, cycles := cpu.Execute(5, &memory)

	if cycles > 0 {
		t.Error("LDY AY should take one more cycle if the value crosses page boundary.")
	}

	ValidateLoad(success, cycles, cpu.Y, val, cpu, cpuCopy, t)
}
