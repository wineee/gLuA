package state

type luaStack struct {
	slots []luaValue // slots用来存放值
	top   int       // 记录栈顶索引
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top: 0,
	}
}

// 检查栈的空闲空间是否还可以容纳至少n个值，如果不满足这个条件，则扩容
func (self *luaStack) check (n int) {
	free := len(self.slots) - self.top;
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
}

func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}
	self.slots[self.top] = val
	self.top ++
}

func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil
	return val
}

// 把索引转换成绝对索引
func (self *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx+self.top+1
}

// 判断索引是否有效
func (self *luaStack) isValid(idx int) bool {
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx <= self.top
}

func (self *luaStack) get(idx int) luaValue {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx-1]
	}
	return nil
}

func (self *luaStack) set(idx int, val luaValue) {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		self.slots[absIdx-1] = val
	} else {
		panic("invalid index!")
	}
}

func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
