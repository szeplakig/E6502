package cpumodule

import (
	"E6502/memorymodule"
	. "E6502/utils"
)

type CPU struct {
	// 6502 cpu
	PC Word // Program Counter
	SP Byte // Stack Pointer

	A Byte // Accumulator
	X Byte // Index Register X
	Y Byte // Index Register Y

	C bool // Carry Flag
	Z bool // Zero Flag
	I bool // Interrupt Disable
	D bool // Decimal Mode
	B bool // Break Command
	V bool // Overflow Flag
	N bool // Negative Flag
}

func (cpu *CPU) Reset() {
	// Reset the cpu
	cpu.PC = 0xFFFC
	cpu.SP = 0xFF

	cpu.A = 0x0
	cpu.X = 0x0
	cpu.Y = 0x0

	cpu.C = false
	cpu.Z = false
	cpu.I = false
	cpu.D = false
	cpu.B = false
	cpu.V = false
	cpu.N = false
}

func NewCPU() CPU {
	// Create a new cpu with default values
	cpu := CPU{}
	cpu.Reset()
	return cpu
}

func add(cycles *int, a Word, b Word) Word {
	if (a%0xFF)+(b%0xFF) >= 0xFF {
		*cycles--
	}
	return a + b
}

/*
func (cpu *CPU) writeA(cycles *int, value Byte) {
	cpu.A = value
}

func (cpu *CPU) writeX(cycles *int, value Byte) {
	cpu.X = value
}

func (cpu *CPU) writeY(cycles *int, value Byte) {
	cpu.Y = value
}

func (cpu *CPU) loadA(cycles *int, mem *memorymodule.Memory, address Word) {
	value := fetchByte(cycles, mem, address)
	cpu.writeA(cycles, value)
}

func (cpu *CPU) loadX(cycles *int, mem *memorymodule.Memory, address Word) {
	value := fetchByte(cycles, mem, address)
	cpu.writeX(cycles, value)
}

func (cpu *CPU) loadY(cycles *int, mem *memorymodule.Memory, address Word) {
	value := fetchByte(cycles, mem, address)
	cpu.writeY(cycles, value)
} */

func (cpu *CPU) fetchBytePC(cycles *int, mem *memorymodule.Memory) Byte {
	value := fetchByte(cycles, mem, cpu.PC)
	cpu.PC++
	return value
}

func fetchByte(cycles *int, mem *memorymodule.Memory, address Word) Byte {
	value := mem.RB(address)
	*cycles--
	return value
}

func fetchWord(cycles *int, mem *memorymodule.Memory, address Word) Word {
	value := mem.RW(address)
	*cycles -= 2
	return value
}

func writeByte(cycles *int, mem *memorymodule.Memory, address Word, value Byte) {
	mem.WB(address, value)
	*cycles--
}

/* func writeWord(cycles *int, mem *memorymodule.Memory, address Word, value Word) {
	mem.WW(address, value)
	*cycles -= 2
} */

func registerLoaderFactory(addressor func(*int, *memorymodule.Memory) Word, registerAddress *Byte, flaghandler func()) func(*int, *memorymodule.Memory) {
	return func(cycles *int, mem *memorymodule.Memory) {
		address := addressor(cycles, mem)
		value := fetchByte(cycles, mem, address)
		*registerAddress = value
		flaghandler()
	}
}

func registerStorerFactory(addressor func(*int, *memorymodule.Memory) Word, registerAddress *Byte) func(*int, *memorymodule.Memory) {
	return func(cycles *int, mem *memorymodule.Memory) {
		address := addressor(cycles, mem)
		value := *registerAddress
		writeByte(cycles, mem, address, value)
	}
}

func loadFlagLoaderFactory(register *Byte, Z *bool, N *bool) func() {
	return func() {
		*Z = *register == 0
		*N = (*register >> 7) == 1
	}
}

// Execute n cycles using the memory
func (cpu *CPU) Execute(cycles int, mem *memorymodule.Memory) (bool, int) {
	AFlagLoader := loadFlagLoaderFactory(&cpu.A, &cpu.Z, &cpu.N)
	XFlagLoader := loadFlagLoaderFactory(&cpu.X, &cpu.Z, &cpu.N)
	YFlagLoader := loadFlagLoaderFactory(&cpu.Y, &cpu.Z, &cpu.N)

	instructionMap := map[Byte]func(*int, *memorymodule.Memory){
		NOP:    func(cycles *int, mem *memorymodule.Memory) { *cycles-- },
		LDA_IM: registerLoaderFactory(cpu.ImmediateAddressing, &cpu.A, AFlagLoader),
		LDA_ZP: registerLoaderFactory(cpu.ZeroPageAddressing, &cpu.A, AFlagLoader),
		LDA_ZX: registerLoaderFactory(cpu.ZeroPageXAddressing, &cpu.A, AFlagLoader),
		LDA_AB: registerLoaderFactory(cpu.AbsoluteAddressing, &cpu.A, AFlagLoader),
		LDA_AX: registerLoaderFactory(cpu.AbsoluteXAddressing, &cpu.A, AFlagLoader),
		LDA_AY: registerLoaderFactory(cpu.AbsoluteYAddressing, &cpu.A, AFlagLoader),
		LDA_IX: registerLoaderFactory(cpu.IndirectXAddressing, &cpu.A, AFlagLoader),
		LDA_IY: registerLoaderFactory(cpu.IndirectYAddressing, &cpu.A, AFlagLoader),

		LDX_IM: registerLoaderFactory(cpu.ImmediateAddressing, &cpu.X, XFlagLoader),
		LDX_ZP: registerLoaderFactory(cpu.ZeroPageAddressing, &cpu.X, XFlagLoader),
		LDX_ZY: registerLoaderFactory(cpu.ZeroPageYAddressing, &cpu.X, XFlagLoader),
		LDX_AB: registerLoaderFactory(cpu.AbsoluteAddressing, &cpu.X, XFlagLoader),
		LDX_AY: registerLoaderFactory(cpu.AbsoluteYAddressing, &cpu.X, XFlagLoader),

		LDY_IM: registerLoaderFactory(cpu.ImmediateAddressing, &cpu.Y, YFlagLoader),
		LDY_ZP: registerLoaderFactory(cpu.ZeroPageAddressing, &cpu.Y, YFlagLoader),
		LDY_ZX: registerLoaderFactory(cpu.ZeroPageXAddressing, &cpu.Y, YFlagLoader),
		LDY_AB: registerLoaderFactory(cpu.AbsoluteAddressing, &cpu.Y, YFlagLoader),
		LDY_AX: registerLoaderFactory(cpu.AbsoluteXAddressing, &cpu.Y, YFlagLoader),

		STA_ZP: registerStorerFactory(cpu.ZeroPageAddressing, &cpu.A),
		STA_ZX: registerStorerFactory(cpu.ZeroPageXAddressing, &cpu.A),
		STA_AB: registerStorerFactory(cpu.AbsoluteAddressing, &cpu.A),
		STA_AX: registerStorerFactory(cpu.AbsoluteXAddressingLong, &cpu.A),
		STA_AY: registerStorerFactory(cpu.AbsoluteYAddressingLong, &cpu.A),
		STA_IX: registerStorerFactory(cpu.IndirectXAddressing, &cpu.A),
		STA_IY: registerStorerFactory(cpu.IndirectYAddressingLong, &cpu.A),

		STX_ZP: registerStorerFactory(cpu.ZeroPageAddressing, &cpu.X),
		STX_ZY: registerStorerFactory(cpu.ZeroPageYAddressing, &cpu.X),
		STX_AB: registerStorerFactory(cpu.AbsoluteAddressing, &cpu.X),

		STY_ZP: registerStorerFactory(cpu.ZeroPageAddressing, &cpu.Y),
		STY_ZX: registerStorerFactory(cpu.ZeroPageXAddressing, &cpu.Y),
		STY_AB: registerStorerFactory(cpu.AbsoluteAddressing, &cpu.Y),
	}

	for cycles > 0 {
		nextIns := cpu.fetchBytePC(&cycles, mem)
		if handler, ok := instructionMap[nextIns]; ok {
			handler(&cycles, mem)
		} else {
			return false, cycles
		}
	}
	//fmt.Println(cycles)
	return cycles == 0, cycles
}
