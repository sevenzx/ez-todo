package util

import "github.com/bytedance/sonic"

// ToString 转成字符串
func ToString(v interface{}) string {
	s, _ := sonic.MarshalString(v)
	return s
}
