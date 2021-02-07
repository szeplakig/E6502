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

func ValidateLoad(success bool, register Byte, value Byte, cpu CPU, cpuCopy CPU, t *testing.T) {
	if !success {
		t.Error("Register did not take the exact cycles to get performed!")
	}
	if register != value {
		t.Errorf("Register not loading correct byte into Accumulator! Register: %d, value: %d", register, value)
	}
	if ((value>>7) == 1 && !cpu.N) || ((value>>7) == 0 && cpu.N) {
		t.Error("Register not setting negative flag correctly!")
	}
	if (value == 0 && !cpu.Z) || (value != 0 && cpu.Z) {
		t.Error("Register not setting zero flag correctly!")
	}
	if !VerifyFlagsUnchanged(cpu, cpuCopy, ZN) {
		t.Error("Register not setting the flags correctly!")
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
