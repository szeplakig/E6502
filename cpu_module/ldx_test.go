package cpu_module

import (
	"E6502/memory_module"
	"testing"
)

func Test_INS_LDX_IM(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	memory.WB(0xFFFC, LDX_IM)
	memory.WB(0xFFFD, val)

	cpuCopy := cpu
	success, _ := cpu.Execute(2, &memory)

	ValidateLoad(success, cpu.X, val, cpu, cpuCopy, t)
}

func Test_INS_LDX_ZP(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	memory.WB(0x0000, val)
	memory.WB(0xFFFC, LDX_ZP)
	memory.WB(0xFFFD, 0x00)

	cpuCopy := cpu
	success, _ := cpu.Execute(3, &memory)

	ValidateLoad(success, cpu.X, val, cpu, cpuCopy, t)
}

func Test_INS_LDX_ZY(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.Y = 0x0F
	memory.WB(0x008F, val)
	memory.WB(0xFFFC, LDX_ZY)
	memory.WB(0xFFFD, 0x80)

	cpuCopy := cpu

	success, _ := cpu.Execute(4, &memory)

	ValidateLoad(success, cpu.X, val, cpu, cpuCopy, t)
}

func Test_INS_LDX_AB(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	memory.WB(0x4224, val)
	memory.WB(0xFFFC, LDX_AB)
	memory.WW(0xFFFD, 0x4224)

	cpuCopy := cpu
	success, _ := cpu.Execute(4, &memory)

	ValidateLoad(success, cpu.X, val, cpu, cpuCopy, t)
}

func Test_INS_LDX_AY(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.Y = 0x92
	memory.WB(0x2092, val)
	memory.WB(0xFFFC, LDX_AY)
	memory.WW(0xFFFD, 0x2000)

	cpuCopy := cpu
	success, _ := cpu.Execute(4, &memory)

	ValidateLoad(success, cpu.X, val, cpu, cpuCopy, t)
}

func Test_INS_LDX_AY_CROSSES_PAGE_BOUNDARY(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.Y = 0x2
	memory.WB(0x20E0, val)
	memory.WB(0xFFFC, LDX_AY)
	memory.WW(0xFFFD, 0x20DE)

	cpuCopy := cpu
	success, cycles := cpu.Execute(5, &memory)

	if cycles > 0 {
		t.Error("LDX AY should take one more cycle if the value crosses page boundary.")
	}

	ValidateLoad(success, cpu.X, val, cpu, cpuCopy, t)
}
