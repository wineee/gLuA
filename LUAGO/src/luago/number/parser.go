package number

// 用于将字符串解析为整数和浮点数

import "strconv"

// 第一个返回值是解析后的数，第二个返回值说明解析是否成功
func ParseInteger(str string) (int64, bool) {
	i, err := strconv.ParseInt(str, 10, 64)
	return i, err == nil
}

func ParseFloat(str string) (float64, bool) {
	f, err := strconv.ParseFloat(str, 64)
	return f, err == nil
}
