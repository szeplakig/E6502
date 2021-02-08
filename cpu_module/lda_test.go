package cpu_module

import (
	"E6502/memory_module"
	"testing"
)

func Test_INS_LDA_IM(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	memory.WB(0xFFFC, LDA_IM)
	memory.WB(0xFFFD, val)

	cpuCopy := cpu
	success, _ := cpu.Execute(2, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_ZP(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	memory.WB(0x0000, val)
	memory.WB(0xFFFC, LDA_ZP)
	memory.WB(0xFFFD, 0x00)

	cpuCopy := cpu
	success, _ := cpu.Execute(3, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_ZX(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.X = 0x0F
	memory.WB(0x008F, val)
	memory.WB(0xFFFC, LDA_ZX)
	memory.WB(0xFFFD, 0x80)

	cpuCopy := cpu

	success, _ := cpu.Execute(4, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AB(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	memory.WB(0x4224, val)
	memory.WB(0xFFFC, LDA_AB)
	memory.WW(0xFFFD, 0x4224)

	cpuCopy := cpu
	success, _ := cpu.Execute(4, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AX(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.X = 0x92
	memory.WB(0x2092, val)
	memory.WB(0xFFFC, LDA_AX)
	memory.WW(0xFFFD, 0x2000)

	cpuCopy := cpu
	success, _ := cpu.Execute(4, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AX_CROSSES_PAGE_BOUNDARY(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.X = 0x2
	memory.WB(0x20E0, val)
	memory.WB(0xFFFC, LDA_AX)
	memory.WW(0xFFFD, 0x20DE)

	cpuCopy := cpu
	success, cycles := cpu.Execute(5, &memory)

	if cycles > 0 {
		t.Error("LDA AX should take one more cycle if the value crosses page boundary.")
	}

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AY(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.Y = 0x92
	memory.WB(0x2092, val)
	memory.WB(0xFFFC, LDA_AY)
	memory.WW(0xFFFD, 0x2000)

	cpuCopy := cpu
	success, _ := cpu.Execute(4, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AY_CROSSES_PAGE_BOUNDARY(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.Y = 0x2
	memory.WB(0x20E0, val)
	memory.WB(0xFFFC, LDA_AY)
	memory.WW(0xFFFD, 0x20DE)

	cpuCopy := cpu
	success, cycles := cpu.Execute(5, &memory)

	if cycles > 0 {
		t.Error("LDA AY should take one more cycle if the value crosses page boundary.")
	}

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_IX(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.X = 0x04
	memory.WW(0x0006, 0x0042)
	memory.WB(0x0042, val)
	memory.WB(0xFFFC, LDA_IX)
	memory.WB(0xFFFD, 0x02)

	cpuCopy := cpu
	success, _ := cpu.Execute(6, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_IY(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.Y = 0x04
	memory.WW(0x0002, 0x8000)
	memory.WB(0x8004, val)
	memory.WB(0xFFFC, LDA_IY)
	memory.WB(0xFFFD, 0x02)

	cpuCopy := cpu
	success, _ := cpu.Execute(5, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_IY_CROSSES_PAGE_BOUNDARY(t *testing.T) {
	cpu := NewCPU()
	memory := memory_module.NewMemory()
	var val Byte = 0xF0
	cpu.Y = 0xFE

	memory.WB(0xFFFC, LDA_IY) // Indirect indexed
	memory.WB(0xFFFD, 0x02)   // Will read from zero page 0x02  and 0x03
	memory.WW(0x0002, 0x80FE) // Will read the addr 0x80FE and add Y to it = 81FC
	memory.WB(0x81FC, val)

	cpuCopy := cpu
	success, _ := cpu.Execute(6, &memory)

	ValidateLoad(success, cpu.A, val, cpu, cpuCopy, t)
}
