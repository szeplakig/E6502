package cpu

const (
	NOP = 0xEA

	LDA_IM = 0xA9
	LDA_ZP = 0xA5
	LDA_ZX = 0xB5
	LDA_AB = 0xAD
	LDA_AX = 0xBD
	LDA_AY = 0xB9
	LDA_IX = 0xA1
	LDA_IY = 0xB1

	LDX_IM = 0xA2
	LDX_ZP = 0xA6
	LDX_ZY = 0xB6
	LDX_AB = 0xAE
	LDX_AY = 0xBE

	LDY_IM = 0xA0
	LDY_ZP = 0xA4
	LDY_ZX = 0xB4
	LDY_AB = 0xAC
	LDY_AX = 0xBC
)
