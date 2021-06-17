package state

import (
	"math"
	. "luago/api"
	"luago/number"
)

var (
	iadd = func(a, b int64) int64 { return a+b }
	fadd  = func(a, b float64) float64 { return a + b }

	isub  = func(a, b int64) int64 { return a - b }
	fsub  = func(a, b float64) float64 { return a - b }

	imul  = func(a, b int64) int64 { return a * b }
	fmul  = func(a, b float64) float64 { return a * b }

	imod = number.IMod
	fmod = number.FMod

	pow = math.Pow

	div = func(a, b float64) float64 { return a / b }

	// 整除
	iidiv = number.IFloorDiv
	fidiv = number.FFloorDiv

	band  = func(a, b int64) int64 { return a & b }
	bor   = func(a, b int64) int64 { return a | b }
	bxor  = func(a, b int64) int64 { return a ^ b }
	
	shl   = number.ShiftLeft
	shr   = number.ShiftRight

	// 使用接收两个参数且返回一个值的函数来统一表示Lua运算符，一元运算符简单忽略第二个参数就可以
	iunm  = func(a, _ int64) int64 { return -a }
	funm  = func(a, _ float64) float64 { return -a }

	bnot  = func(a, _ int64) int64 { return ^a }
)

type operator struct {
	integerFunc func(int64, int64) int64
	floatFunc func(float64, float64) float64
}

// 注意要和前面定义的Lua运算码常量顺序一致
var operators = []operator {
	//{integerFunc floatFunc}
	operator{iadd, fadd}, // 成员是函数
	operator{isub, fsub},
	operator{imul, fmul},
	operator{imod, fmod},
	operator{nil, pow},
	operator{nil, div},
	operator{iidiv, fidiv},
	operator{band, nil},
	operator{bor, nil},
	operator{bxor, nil},
	operator{shl, nil},
	operator{shr, nil},
	operator{iunm, funm},
	operator{bnot, nil},
}


// 先根据情况从Lua栈里弹出一到两个操作数，然后按索引取出相应的operator实例，最后调用_arith()函数执行计算
func (self *luaState) Arith(op ArithOp) {
	var a, b luaValue
	b = self.stack.pop()
	if op != LUA_OPUNM && op != LUA_OPBNOT {
		a = self.stack.pop()
	} else {
		a = b
	}

	operator := operators[op] // 顺序
	if result := _arith(a, b, operator); result != nil {
		self.stack.push(result)
	} else {
		panic("arithmetic error!")
	}
}

func _arith(a, b luaValue, op operator) luaValue {
	if op.floatFunc == nil { 
		if x, ok := convertToInteger(a); ok {
			if y, ok := convertToInteger(b); ok {
				return op.integerFunc(x, y)
			}
		}
	} else { 
		if op.integerFunc != nil { 
			if x, ok := a.(int64); ok {
				if y, ok := b.(int64); ok {
					return op.integerFunc(x, y)
				}
			}
		}
		if x, ok := convertToFloat(a); ok {
			if y, ok := convertToFloat(b); ok {
				return op.floatFunc(x, y)
			}
		}
	}
	return nil
}
