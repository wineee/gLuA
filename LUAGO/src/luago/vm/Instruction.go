package vm

type Instruction uint32

const MAXARG_Bx = 1<<18 - 1       // 262143
const MAXARG_sBx = MAXARG_Bx >> 1 // 131071

// Opcode()方法从指令中提取操作码
func (self Instruction) Opcode() int {
	return int(self & 0x3F)
}

// ABC()方法从iABC模式指令中提取参数
func (self Instruction) ABC() (a, b, c int) {
	a = int(self >> 6 & 0xFF)
	c = int(self >> 14 & 0x1FF)
	b = int(self >> 23 & 0x1FF)
	return
}

// ABx()方法从iABx模式指令中提取参数
func (self Instruction) ABx() (a, bx int) {
	a = int(self >> 6 & 0xFF);
	bx = int(self >> 14);
	return
}

// AsBx()方法从iAsBx模式指令中提取参数
func (self Instruction) AsBx() (a, sbx int) {
	a, bx := self.ABx()
	return a, bx - MAXARG_sBx // 偏移二进制码
}

// Ax()方法从iAx模式指令中提取参数
func (self Instruction) Ax() int {
	return int(self >> 6)
}


// 这4个方法分别返回指令的操作码名字、编码模式、操作数B的使用模式以及操作数C的使用模式
func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}
