package api

type LuaVM interface {
	LuaState  // 扩展了LuaState接口，增加了5个方法
	PC() int // 返回当前 PC （测试用的）
	AddPC(n int)  // 修改PC, 实现跳转指令
	Fetch() uint32 // 取出当前指令，PC+1
	GetConst(idx int) // 将常量推入栈顶 
	GetRK(rk int) // 从常量表里提取常量或者从栈里提取值，然后推入栈顶
}
