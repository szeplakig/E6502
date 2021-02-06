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

func Test_INS_LDA_IM(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0b10000000
	memory.WB(0xFFFC, LDA_IM)
	memory.WB(0xFFFD, val)

	cpuCopy := cpu
	cpu.Execute(2, &memory)

	if cpu.A != val {
		t.Error("LDA IM not loading next byte into Accumulator!")
	}
	if ((val>>7) == 1 && !cpu.N) || ((val>>7) == 0 && cpu.N) {
		t.Error("LDA IM not setting negative flag correctly!")
	}
	if (val == 0 && !cpu.Z) || (val != 0 && cpu.Z) {
		t.Error("LDA IM not setting zero flag correctly!")
	}
	if !VerifyFlagsUnchanged(cpu, cpuCopy, ZN) {
		t.Error("LDA IM not setting the flags correctly!")
	}
}

func Test_INS_LDA_ZP(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x0
	memory.WB(0x0000, val)
	memory.WB(0xFFFC, LDA_ZP)
	memory.WB(0xFFFD, 0x00)

	cpuCopy := cpu
	cpu.Execute(3, &memory)

	if cpu.A != val {
		t.Error("LDA ZP not loading next byte into Accumulator!")
	}
	if ((val>>7) == 1 && !cpu.N) || ((val>>7) == 0 && cpu.N) {
		t.Error("LDA ZP not setting negative flag correctly!")
	}
	if (val == 0 && !cpu.Z) || (val != 0 && cpu.Z) {
		t.Error("LDA ZP not setting zero flag correctly!")
	}
	if !VerifyFlagsUnchanged(cpu, cpuCopy, ZN) {
		t.Error("LDA ZP not setting the flags correctly!")
	}
}

func Test_INS_LDA_ZX(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	cpu.X = 0x0F
	memory.WB(0x008F, val)
	memory.WB(0xFFFC, LDA_ZX)
	memory.WB(0xFFFD, 0x80)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	if cpu.A != val {
		t.Error("LDA ZX not loading next byte into Accumulator!")
	}
	if ((val>>7) == 1 && !cpu.N) || ((val>>7) == 0 && cpu.N) {
		t.Error("LDA ZX not setting negative flag correctly!")
	}
	if (val == 0 && !cpu.Z) || (val != 0 && cpu.Z) {
		t.Error("LDA ZX not setting zero flag correctly!")
	}
	if !VerifyFlagsUnchanged(cpu, cpuCopy, ZN) {
		t.Error("LDA ZX not setting the flags correctly!")
	}
}

func Test_INS_LDA_AB(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	memory.WB(0x4224, val)
	memory.WB(0xFFFC, LDA_AB)
	memory.WW(0xFFFD, 0x4224)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	if cpu.A != val {
		t.Error("LDA AB not loading next byte into Accumulator!")
	}
	if ((val>>7) == 1 && !cpu.N) || ((val>>7) == 0 && cpu.N) {
		t.Error("LDA AB not setting negative flag correctly!")
	}
	if (val == 0 && !cpu.Z) || (val != 0 && cpu.Z) {
		t.Error("LDA AB not setting zero flag correctly!")
	}
	if !VerifyFlagsUnchanged(cpu, cpuCopy, ZN) {
		t.Error("LDA AB not setting the flags correctly!")
	}
}

func Test_INS_LDA_AX(t *testing.T) {
	cpu := NewCPU()
	memory := Memory.NewMemory()
	var val Byte = 0x42
	cpu.X = 0x92
	memory.WB(0x2092, val)
	memory.WB(0xFFFC, LDA_AX)
	memory.WW(0xFFFD, 0x2000)

	cpuCopy := cpu
	cpu.Execute(4, &memory)

	if cpu.A != val {
		t.Error("LDA AX not loading next byte into Accumulator!")
	}
	if ((val>>7) == 1 && !cpu.N) || ((val>>7) == 0 && cpu.N) {
		t.Error("LDA AX not setting negative flag correctly!")
	}
	if (val == 0 && !cpu.Z) || (val != 0 && cpu.Z) {
		t.Error("LDA AX not setting zero flag correctly!")
	}
	if !VerifyFlagsUnchanged(cpu, cpuCopy, ZN) {
		t.Error("LDA AX not setting the flags correctly!")
	}
}
