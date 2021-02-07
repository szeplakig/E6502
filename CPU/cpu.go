package CPU

import (
	"E6502/Memory"
)

type CPU struct {
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
	cpu := CPU{}
	cpu.Reset()
	return cpu
}

func Add(cycles *int, a Word, b Word) Word {
	if (a%0xFF)+(b%0xFF) >= 0xFF {
		*cycles--
	}
	return a + b
}

func (cpu *CPU) WriteA(cycles *int, value Byte) {
	cpu.A = value
}

func (cpu *CPU) WriteX(cycles *int, value Byte) {
	cpu.X = value
}

func (cpu *CPU) WriteY(cycles *int, value Byte) {
	cpu.Y = value
}

func (cpu *CPU) LoadA(cycles *int, memory *Memory.Memory, address Word) {
	value := cpu.FetchByte(cycles, memory, address)
	cpu.WriteA(cycles, value)
}

func (cpu *CPU) LoadX(cycles *int, memory *Memory.Memory, address Word) {
	value := cpu.FetchByte(cycles, memory, address)
	cpu.WriteX(cycles, value)
}

func (cpu *CPU) LoadY(cycles *int, memory *Memory.Memory, address Word) {
	value := cpu.FetchByte(cycles, memory, address)
	cpu.WriteY(cycles, value)
}

func (cpu *CPU) FetchBytePC(cycles *int, memory *Memory.Memory) Byte {
	value := cpu.FetchByte(cycles, memory, cpu.PC)
	cpu.PC++
	return value
}

func (cpu *CPU) FetchByte(cycles *int, memory *Memory.Memory, address Word) Byte {
	value := memory.RB(address)
	*cycles--
	return value
}

func (cpu *CPU) FetchWord(cycles *int, memory *Memory.Memory, address Word) Word {
	value := memory.RW(address)
	*cycles -= 2
	return value
}

func RegisterLoaderFactory(addressor func(*int, *Memory.Memory) Word, loader func(*int, *Memory.Memory, Word), flaghandler func()) func(cycles *int, memory *Memory.Memory) {
	return func(cycles *int, memory *Memory.Memory) {
		address := addressor(cycles, memory)
		loader(cycles, memory, address)
		flaghandler()
	}
}

func LoadFlagLoaderFactory(register *Byte, Z *bool, N *bool) func() {
	return func() {
		*Z = *register == 0
		*N = (*register >> 7) == 1
	}
}

func (cpu *CPU) Execute(cycles int, memory *Memory.Memory) (bool, int) {
	AFlagLoader := LoadFlagLoaderFactory(&cpu.A, &cpu.Z, &cpu.N)
	XFlagLoader := LoadFlagLoaderFactory(&cpu.X, &cpu.Z, &cpu.N)
	YFlagLoader := LoadFlagLoaderFactory(&cpu.Y, &cpu.Z, &cpu.N)

	instruction_map := map[Byte]func(*int, *Memory.Memory){
		LDA_IM: RegisterLoaderFactory(cpu.ImmediateAddressing, cpu.LoadA, AFlagLoader),
		LDA_ZP: RegisterLoaderFactory(cpu.ZeroPageAddressing, cpu.LoadA, AFlagLoader),
		LDA_ZX: RegisterLoaderFactory(cpu.ZeroPageXAddressing, cpu.LoadA, AFlagLoader),
		LDA_AB: RegisterLoaderFactory(cpu.AbsoluteAddressing, cpu.LoadA, AFlagLoader),
		LDA_AX: RegisterLoaderFactory(cpu.AbsoluteXAddressing, cpu.LoadA, AFlagLoader),
		LDA_AY: RegisterLoaderFactory(cpu.AbsoluteYAddressing, cpu.LoadA, AFlagLoader),
		LDA_IX: RegisterLoaderFactory(cpu.IndirectXAddressing, cpu.LoadA, AFlagLoader),
		LDA_IY: RegisterLoaderFactory(cpu.IndirectYAddressing, cpu.LoadA, AFlagLoader),

		LDX_IM: RegisterLoaderFactory(cpu.ImmediateAddressing, cpu.LoadX, XFlagLoader),
		LDX_ZP: RegisterLoaderFactory(cpu.ZeroPageAddressing, cpu.LoadX, XFlagLoader),
		LDX_ZY: RegisterLoaderFactory(cpu.ZeroPageYAddressing, cpu.LoadX, XFlagLoader),
		LDX_AB: RegisterLoaderFactory(cpu.AbsoluteAddressing, cpu.LoadX, XFlagLoader),
		LDX_AY: RegisterLoaderFactory(cpu.AbsoluteYAddressing, cpu.LoadX, XFlagLoader),

		LDY_IM: RegisterLoaderFactory(cpu.ImmediateAddressing, cpu.LoadY, YFlagLoader),
		LDY_ZP: RegisterLoaderFactory(cpu.ZeroPageAddressing, cpu.LoadY, YFlagLoader),
		LDY_ZX: RegisterLoaderFactory(cpu.ZeroPageXAddressing, cpu.LoadY, YFlagLoader),
		LDY_AB: RegisterLoaderFactory(cpu.AbsoluteAddressing, cpu.LoadY, YFlagLoader),
		LDY_AX: RegisterLoaderFactory(cpu.AbsoluteXAddressing, cpu.LoadY, YFlagLoader),
	}

	for cycles > 0 {
		next_ins := cpu.FetchBytePC(&cycles, memory)
		if handler, ok := instruction_map[next_ins]; ok {
			handler(&cycles, memory)
		} else {
			return false, cycles
		}
	}
	//fmt.Println(cycles)
	return cycles == 0, cycles
}
