package state

import "fmt"
import . "luago/api"

// 把给定Lua类型转换成对应的字符串表
func (self *luaState) TypeName(tp LuaType) string {
	switch tp {
	case LUA_TNONE:
		return "no value"
	case LUA_TNIL:
		return "nil"
	case LUA_TBOOLEAN:
		return "boolean"
	case LUA_TNUMBER:
		return "number"
	case LUA_TSTRING:
		return "string"
	case LUA_TTABLE:
		return "table"
	case LUA_TFUNCTION:
		return "function"
	case LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}

// 根据索引返回值的类型，如果索引无效，则返回LUA_TNONE
func (self *luaState) Type(idx int) LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == LUA_TNONE
}

func (self *luaState) IsNil(idx int) bool {
	return self.Type(idx) == LUA_TNIL
}

func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= LUA_TNIL
}

func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == LUA_TBOOLEAN
}

func (self *luaState) IsTable(idx int) bool {
	return self.Type(idx) == LUA_TTABLE
}

func (self *luaState) IsFunction(idx int) bool {
	return self.Type(idx) == LUA_TFUNCTION
}

func (self *luaState) IsThread(idx int) bool {
	return self.Type(idx) == LUA_TTHREAD
}

// 否是字符串（或是数字)
func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

// 是否是（或者可以转换为）数字类型
func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

// 
func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

// 如果值不是数字类型并且也没办法转换成数字类型，前者只是简单地返回0，后者则会报告转换是否成功
func (self *luaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

func (self *luaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	// https://blog.csdn.net/nevergk/article/details/79174733
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

// 如果值是字符串，则返回该字符串。如果值是数字，则将值转换为字符串（注意会修改栈），然后返回字符串。否则，返回空字符串
func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)

	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x) // https://www.cnblogs.com/jcjc/p/12290915.html
		self.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}
