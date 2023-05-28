package utils

import "regexp"

// ternary conditional operator 三元表达式
// condition ? trueVal : falseVal
// 如果condition为true,那么返回trueVal，如果a为false，那么返回falseVal。
func TCO(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

// IsValid ...
func IsValid(str string) bool {
	reg := regexp.MustCompile(`[\\\\/:*?\"<>|]`)
	if reg == nil {
		return false
	}
	result := reg.FindAllString(str, -1)
	return len(result) <= 0
}
