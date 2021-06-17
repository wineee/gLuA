package number

import "math"

//加、减、乘、除和取反运算符可以直接映射为Go语言里相应的运算符
// 乘方运算符（Go语言里没有乘方运算符）、整除运算符（Go语言整除运算仅适用于整数，并且是直接截断结果而非向下取整）和取模运算符（原因和整除运算符类似）却不能直接映射

func IFloorDiv(a, b int64) int64 {
	// 向下取整
	if (a>0 && b>0) || (a<0 && b<0) || a%b==0 {
		return a / b
	}
	return a/b - 1
}

func FFloorDiv(a, b float64) float64 {
	return math.Floor(a / b)
}

// a%b = a - a//b*b
func IMod(a, b int64) int64 {
	return a - IFloorDiv(a, b) * b
}

func FMod(a, b float64) float64 {
	return a - math.Floor(a/b)*b;
}

// 按位与、按位或、异或、按位取反运算符可以直接映射为Go语言里相应的运算符
// 但是位移运算需要稍微处理一下

func ShiftLeft(a, n int64) int64 {
	if n >= 0 {
		return a << uint64(n)
	} else {
		return ShiftRight(a, -n)
	}
}

// 无符号右移（空缺补0）
func ShiftRight(a, n int64) int64 {
	if n >= 0 {
		return int64(uint64(a) >> uint64(n))
	} else {
		return ShiftLeft(a, -n)
	}
}

/*
利用短路的三目运算符   
> a = 9
> b = 888
> c = a>b and a or b
> c
888
*/

// 自动类型转换

func FloatToInteger(f float64) (int64, bool) {
	i := int64(f)
	return i, float64(i) == f
}
