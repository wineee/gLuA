package binchunk

// 魔数
const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSZIET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

// 常量表, 每个常量都以1字节tag开头
const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

type binaryChunk struct {
	header // 头部
	sizeUpvalues  byte // 主函数upvalue数量
	mainFunc     *Prototype  // 主函数原型
}

type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

// 函数原型主要包含函数基本信息、指令表、常量表、upvalue表、子函数原型表以及调试信息
type Prototype struct {
	Source          string
	// 源文件名,只有在主函数原型里,该字段才真正有值
	LineDefined     uint32
	LastLineDefined uint32
	// 起止行号,如果是普通的函数,起止行号都应该大于0；如果是主函数，则起止行号都是0
	NumParams       byte
	// 固定参数个数,这里的固定参数,是相对于变长参数（Vararg）而言的
	IsVararg        byte // 是否是vararg函数
	MaxStackSize    byte
	// 运行函数所必要的寄存器数量, Lua虚拟机在执行函数时，真正使用的其实是一种栈结构，这种栈结构除了可以进行常规地推入和弹出操作以外，还可以按索引访问，所以可以用来模拟寄存器。
	Code           []uint32
	// 指令表,每条指令占4个字节
	Constants      []interface{}
	// 常量表,用于存放Lua代码里出现的字面量，包括nil、布尔值、整数、浮点数和字符串五种
	Upvalues       []Upvalue
	// 该表的每个元素占用2个字节
	Protos         []*Prototype //子函数原型表
	LineInfo       []uint32
	// 行号表,行号表中的行号和指令表中的指令一一对应，分别记录每条指令在源代码中对应的行号
	LocVars        []LocVar
	// 局部变量表，用于记录局部变量名，表中每个元素都包含变量名（按字符串类型存储）和起止指令索引（按cint类型存储）
	UpvalueNames   []string
	// upvalue名列表,该列表中的元素和前面Upvalue表中的元素一一对应
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}

// 用于解析二进制chunk
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader() // 检验头部
	reader.readByte() // skip the size of Upvalue
	return reader.readProto("")  // 读取原型
}
