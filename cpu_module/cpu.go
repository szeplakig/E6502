package cpu_module

import "E6502/memory_module"

type CPU struct { // 6502 cpu
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

// Reset the cpu
func (cpu *CPU) Reset() {
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

// Create a new cpu with default values
func NewCPU() CPU {
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

func (cpu *CPU) writeA(cycles *int, value Byte) {
	cpu.A = value
}

func (cpu *CPU) writeX(cycles *int, value Byte) {
	cpu.X = value
}

func (cpu *CPU) writeY(cycles *int, value Byte) {
	cpu.Y = value
}

func (cpu *CPU) loadA(cycles *int, mem *memory_module.Memory, address Word) {
	value := cpu.fetchByte(cycles, mem, address)
	cpu.writeA(cycles, value)
}

func (cpu *CPU) loadX(cycles *int, mem *memory_module.Memory, address Word) {
	value := cpu.fetchByte(cycles, mem, address)
	cpu.writeX(cycles, value)
}

func (cpu *CPU) loadY(cycles *int, mem *memory_module.Memory, address Word) {
	value := cpu.fetchByte(cycles, mem, address)
	cpu.writeY(cycles, value)
}

func (cpu *CPU) fetchBytePC(cycles *int, mem *memory_module.Memory) Byte {
	value := cpu.fetchByte(cycles, mem, cpu.PC)
	cpu.PC++
	return value
}

func (cpu *CPU) fetchByte(cycles *int, mem *memory_module.Memory, address Word) Byte {
	value := mem.RB(address)
	*cycles--
	return value
}

func (cpu *CPU) fetchWord(cycles *int, mem *memory_module.Memory, address Word) Word {
	value := mem.RW(address)
	*cycles -= 2
	return value
}

func registerLoaderFactory(addressor func(*int, *memory_module.Memory) Word, loader func(*int, *memory_module.Memory, Word), flaghandler func()) func(*int, *memory_module.Memory) {
	return func(cycles *int, mem *memory_module.Memory) {
		address := addressor(cycles, mem)
		loader(cycles, mem, address)
		flaghandler()
	}
}

func loadFlagLoaderFactory(register *Byte, Z *bool, N *bool) func() {
	return func() {
		*Z = *register == 0
		*N = (*register >> 7) == 1
	}
}

// Execute n cycles using the memory
func (cpu *CPU) Execute(cycles int, mem *memory_module.Memory) (bool, int) {
	AFlagLoader := loadFlagLoaderFactory(&cpu.A, &cpu.Z, &cpu.N)
	XFlagLoader := loadFlagLoaderFactory(&cpu.X, &cpu.Z, &cpu.N)
	YFlagLoader := loadFlagLoaderFactory(&cpu.Y, &cpu.Z, &cpu.N)

	instruction_map := map[Byte]func(*int, *memory_module.Memory){
		LDA_IM: registerLoaderFactory(cpu.ImmediateAddressing, cpu.loadA, AFlagLoader),
		LDA_ZP: registerLoaderFactory(cpu.ZeroPageAddressing, cpu.loadA, AFlagLoader),
		LDA_ZX: registerLoaderFactory(cpu.ZeroPageXAddressing, cpu.loadA, AFlagLoader),
		LDA_AB: registerLoaderFactory(cpu.AbsoluteAddressing, cpu.loadA, AFlagLoader),
		LDA_AX: registerLoaderFactory(cpu.AbsoluteXAddressing, cpu.loadA, AFlagLoader),
		LDA_AY: registerLoaderFactory(cpu.AbsoluteYAddressing, cpu.loadA, AFlagLoader),
		LDA_IX: registerLoaderFactory(cpu.IndirectXAddressing, cpu.loadA, AFlagLoader),
		LDA_IY: registerLoaderFactory(cpu.IndirectYAddressing, cpu.loadA, AFlagLoader),

		LDX_IM: registerLoaderFactory(cpu.ImmediateAddressing, cpu.loadX, XFlagLoader),
		LDX_ZP: registerLoaderFactory(cpu.ZeroPageAddressing, cpu.loadX, XFlagLoader),
		LDX_ZY: registerLoaderFactory(cpu.ZeroPageYAddressing, cpu.loadX, XFlagLoader),
		LDX_AB: registerLoaderFactory(cpu.AbsoluteAddressing, cpu.loadX, XFlagLoader),
		LDX_AY: registerLoaderFactory(cpu.AbsoluteYAddressing, cpu.loadX, XFlagLoader),

		LDY_IM: registerLoaderFactory(cpu.ImmediateAddressing, cpu.loadY, YFlagLoader),
		LDY_ZP: registerLoaderFactory(cpu.ZeroPageAddressing, cpu.loadY, YFlagLoader),
		LDY_ZX: registerLoaderFactory(cpu.ZeroPageXAddressing, cpu.loadY, YFlagLoader),
		LDY_AB: registerLoaderFactory(cpu.AbsoluteAddressing, cpu.loadY, YFlagLoader),
		LDY_AX: registerLoaderFactory(cpu.AbsoluteXAddressing, cpu.loadY, YFlagLoader),
	}

	for cycles > 0 {
		next_ins := cpu.fetchBytePC(&cycles, mem)
		if handler, ok := instruction_map[next_ins]; ok {
			handler(&cycles, mem)
		} else {
			return false, cycles
		}
	}
	//fmt.Println(cycles)
	return cycles == 0, cycles
}
