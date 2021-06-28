package state

func (self *luaState) PC() int {
	return self.pc
}

func (self *luaState) AddPC(n int) {
	self.pc += n
}

func (self *luaState) Fetch() uint32 {
	i := self.proto.Code[self.pc]
	self.pc += 1
	return i
}

// 据索引从函数原型的常量表里取出一个常量值，然后把它推入栈顶
func (self *luaState) GetConst(idx int) {
	c := self.proto.Constants[idx]
	self.stack.push(c)
}

// 根据情况调用GetConst()方法把某个常量推入栈顶，
// 或者调用PushValue()方法把某个索引处的栈值推入栈顶

/*
   传递给GetRK()方法的参数实际上是iABC模式指令里的OpArgK类型参数。
   这种类型的参数一共占9个比特。
   如果最高位是1，那么参数里存放的是常量表索引，把最高位去掉就可以得到索引值
   否则最高位是0，参数里存放的就是寄存器索引值。
 */
func (self *luaState) GetRK(rk int) {
	if rk > 0xFF {
		self.GetConst(rk & 0xFF)
	} else {
		// 寄存器索引从0开始，而Lua API里的栈索引是从1开始
		self.PushValue(rk+1)
	}
}
