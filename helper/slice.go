package helper

func RemoveUnordered(s []string, i int) []string {
	s[i] = s[len(s)-1]  // 用最后一个元素覆盖目标
	return s[:len(s)-1] // 截断最后一位
}
