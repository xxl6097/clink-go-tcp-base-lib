package setutil

// 定义一个 Set 结构体
type Set map[string]bool

// 向 Set 中添加元素
func (s Set) Add(item string) {
	s[item] = true
}

// 从 Set 中删除元素
func (s Set) Remove(item string) {
	delete(s, item)
}

// 检查 Set 中是否存在某个元素
func (s Set) Contains(item string) bool {
	_, exists := s[item]
	return exists
}

// 获取 Set 中的所有元素
func (s Set) Elements() []string {
	elements := make([]string, 0, len(s))
	for element := range s {
		elements = append(elements, element)
	}
	return elements
}
