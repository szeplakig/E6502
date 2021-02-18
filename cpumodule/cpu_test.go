package cpumodule

import (
	. "E6502/memorymodule"
	. "E6502/utils"
	"reflect"
	"testing"
)

func Test_EXECUTION_RETURN_WITH_UNKNOWN_INSTRUCTION(t *testing.T) {
	cpu := NewCPU()
	mem := NewMemory()
	mem.WB(0xFFFC, 0x00)

	cpuCopy := cpu
	success, _ := cpu.Execute(1, &mem)

	if success {
		t.Error("Execution should fail with an unknown instruction!")
	}

	if !VerifyFlagsUnchanged(cpu, cpuCopy, ALL) {
		t.Error("No instructions executed should leave all flags unchanged!")
	}
}

func Test_NOP_SHOULD_ONLY_CHANGE_THE_PC(t *testing.T) {
	cpu := NewCPU()
	mem := NewMemory()
	mem.WB(0xFFFC, NOP)

	memCopy := mem
	cpuCopy := cpu
	cpu.Execute(2, &mem)
	cpu.PC--

	if !reflect.DeepEqual(cpu, cpuCopy) || !reflect.DeepEqual(mem, memCopy) {
		t.Error("NOP should change nothing!")
	}
}

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

func ValidateLoad(success bool, cycles int, valueInRegister Byte, valueToLoad Byte, cpu CPU, cpuBeforeOp CPU, t *testing.T) {
	if !success {
		t.Errorf("Loading value into register did not take the exact cycles to get performed! Cycles taken: %d", cycles)
	}
	if valueInRegister != valueToLoad {
		t.Errorf("Operation did not load the correct byte into the register! Register: %d, Value: %d", valueInRegister, valueToLoad)
	}
	if ((valueToLoad>>7) == 1 && !cpu.N) || ((valueToLoad>>7) == 0 && cpu.N) {
		t.Error("Operation did not set the negative flag correctly!")
	}
	if (valueToLoad == 0 && !cpu.Z) || (valueToLoad != 0 && cpu.Z) {
		t.Error("Operation did not set the zero flag correctly!")
	}
	if !VerifyFlagsUnchanged(cpu, cpuBeforeOp, ZN) {
		t.Error("After loading register the flags were not set correctly!")
	}
}

func ValidateStore(success bool, cycles int, valueInMemory Byte, valueToStore Byte, cpu CPU, cpuBeforeOp CPU, t *testing.T) {
	if !success {
		t.Errorf("Storing the value from the register did not take the exact cycles to get performed! Cycles taken: %d", cycles)
	}

	if valueInMemory != valueToStore {
		t.Errorf("Operation did not store the correct byte into memory! In memory: %d, Value to store: %d", valueInMemory, valueToStore)
	}

	if !VerifyFlagsUnchanged(cpu, cpuBeforeOp, ALL) {
		t.Error("The operation changed some flags!")
	}
}
