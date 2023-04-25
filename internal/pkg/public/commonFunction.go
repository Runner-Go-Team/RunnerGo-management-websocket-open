package public

import "sort"

// GetStringNum 获取字符串字符个数
func GetStringNum(stringData string) int {
	num := 0
	for range stringData {
		num++
	}
	return num
}

// StringInSlice 判断目标字符串是否在切片中存在
func StringInSlice(s string, slice []string) bool {
	i := sort.SearchStrings(slice, s)
	if i < len(slice) && slice[i] == s {
		return true
	}
	return false
}
