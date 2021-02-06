package CPU

import (
	"E6502/Memory"
	"testing"
)

const ZN Byte = 0b1011110
const ALL Byte = 0b1111111

func VerifyFlagsUnchanged(cpu CPU, cpuCopy CPU, mask Byte) bool {
	if (mask&0b10000000) == 0 && (cpu.C != cpuCopy.C) {
		return false
	}
	if (mask&0b1000000) == 0 && (cpu.Z != cpuCopy.Z) {
		return false
	}
	if (mask&0b100000) == 0 && (cpu.I != cpuCopy.I) {
		return false
	}
	if (mask&0b10000) == 0 && (cpu.D != cpuCopy.D) {
		return false
	}
	if (mask&0b1000) == 0 && (cpu.B != cpuCopy.B) {
		return false
	}
	if (mask&0b100) == 0 && (cpu.V != cpuCopy.V) {
		return false
	}
	if (mask&0b10) == 0 && (cpu.N != cpuCopy.N) {
		return false
	}
	return true
}

func ValidateLoad(register Byte, value Byte, cpu CPU, cpuCopy CPU, t *testing.T) {
	if register != value {
		t.Error("LDA IM not loading next byte into Accumulator!")
	}
	if ((value>>7) == 1 && !cpu.N) || ((value>>7) == 0 && cpu.N) {
		t.Error("LDA IM not setting negative flag correctly!")
	}
	if (value == 0 && !cpu.Z) || (value != 0 && cpu.Z) {
		t.Error("LDA IM not setting zero flag correctly!")
	}
	if !VerifyFlagsUnchanged(cpu, cpuCopy, ZN) {
		t.Error("LDA IM not setting the flags correctly!")
	}
}

func Test_EXECUTION_RETURN_WITH_UNKNOWN_INSTRUCTION(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	memory.WB(0xFFFC, 0x00)

	cpuCopy := cpu
	success, _ := cpu.Execute(1, &memory)

	if success {
		t.Error("Execution should fail with an unknown instruction!")
	}

	if !VerifyFlagsUnchanged(cpu, cpuCopy, ALL) {
		t.Error("No instructions executed should leave all flags unchanged!")
	}
}

func Test_INS_LDA_IM_ZERO(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x0
	memory.WB(0xFFFC, LDA_IM)
	memory.WB(0xFFFD, val)

	cpuCopy := cpu
	cpu.Execute(2, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_IM_NEGATIVE(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	memory.WB(0xFFFC, LDA_IM)
	memory.WB(0xFFFD, val)

	cpuCopy := cpu
	cpu.Execute(2, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_ZP_ZERO(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x0
	memory.WB(0x0000, val)
	memory.WB(0xFFFC, LDA_ZP)
	memory.WB(0xFFFD, 0x00)

	cpuCopy := cpu
	cpu.Execute(3, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_ZP_NEGATIVE(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	memory.WB(0x0000, val)
	memory.WB(0xFFFC, LDA_ZP)
	memory.WB(0xFFFD, 0x00)

	cpuCopy := cpu
	cpu.Execute(3, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_ZX_ZERO(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x0
	cpu.X = 0x0F
	memory.WB(0x008F, val)
	memory.WB(0xFFFC, LDA_ZX)
	memory.WB(0xFFFD, 0x80)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_ZX_NEGATIVE(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	cpu.X = 0x0F
	memory.WB(0x008F, val)
	memory.WB(0xFFFC, LDA_ZX)
	memory.WB(0xFFFD, 0x80)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AB_ZERO(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x0
	memory.WB(0x4224, val)
	memory.WB(0xFFFC, LDA_AB)
	memory.WW(0xFFFD, 0x4224)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AB_NEGATIVE(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	memory.WB(0x4224, val)
	memory.WB(0xFFFC, LDA_AB)
	memory.WW(0xFFFD, 0x4224)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AX_ZERO(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x0
	cpu.X = 0x92
	memory.WB(0x2092, val)
	memory.WB(0xFFFC, LDA_AX)
	memory.WW(0xFFFD, 0x2000)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AX_NEGATIVE(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	cpu.X = 0x92
	memory.WB(0x2092, val)
	memory.WB(0xFFFC, LDA_AX)
	memory.WW(0xFFFD, 0x2000)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}

func Test_INS_LDA_AX_CROSSES_PAGE_BOUNDARY(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	cpu.X = 0x92
	memory.WB(0x2092, val)
	memory.WB(0xFFFC, LDA_AX)
	memory.WW(0xFFFD, 0x2000)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	ValidateLoad(cpu.A, val, cpu, cpuCopy, t)
}
