// 杂项

package state

// Len()方法访问指定索引处的值，取其长度，然后推入栈顶
func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else {
		panic("length error!")
	}
}

// Concat()方法从栈顶弹出n个值，对这些值进行拼接，然后把结果推入栈顶
/*
   "a"
   "b" -> "ba" 
 */
func (self *luaState) Concat(n int) {
	if n == 0 {  // 如果n是0，不弹出任何值，直接往栈顶推入一个空字符串即可
		self.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++ {
			if self.IsString(-1) && self.IsString(-2) {
				s1, s2 := self.ToString(-2), self.ToString(-1)
				self.stack.pop();
				self.stack.pop();
				self.stack.push(s1 + s2)
			} else {
				panic("concatenation error!")
			}
		}
	}
	//  如果n==1， do nothing
}
