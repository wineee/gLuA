package state

import "luago/binchunk"

type luaState struct {
	stack *luaStack

	proto *binchunk.Prototype // 函数原型，这样就可以从中提取指令或者常量
	pc    int // 程序计数器
}

func New() *luaState {
	return &luaState{
		stack: newLuaStack(20),
		proto: proto,
		pc:    0, // 虚拟机肯定是从第1条指令开始执行
	}
}

