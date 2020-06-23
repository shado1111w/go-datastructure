package stack

const (
	maxsize = 100
)

type SqStack struct {
	base      int // 栈底索引
	top       int // 栈顶索引
	stacksize int
	data      []interface{}
}

type StackNode struct {
	data interface{}
	next *StackNode
}

type IntStack []interface{}

// 1.初始化
func (s *SqStack) Init() bool {
	s.data = make([]interface{}, maxsize)
	s.base = 0
	s.top = s.base
	s.stacksize = maxsize
	return true
}

// 2.入栈
func (s *SqStack) Push(e interface{}) bool {
	if s.top-s.base == s.stacksize {
		return false
	}

	s.data[s.top] = e
	s.top++
	return true
}

// 3.出栈
func (s *SqStack) Pop() (e interface{}, ok bool) {
	if s.top == s.base {
		return 0, false
	}

	s.top--
	e = s.data[s.top]

	return e, true
}

// 4.取栈顶元素
func (s *SqStack) GetTop() (e interface{}, ok bool) {
	if s.top == s.base {
		return 0, false
	}

	e = s.data[s.top-1]

	return e, true
}

// 1.初始化
func (s *StackNode) Init() bool {
	s = nil
	return true
}

// 2.入栈
func (s *StackNode) Push(e interface{}) bool {
	var sTem StackNode = *s
	p := StackNode{
		data: e,
		next: &sTem,
	}
	*s = p

	return true
}

// 3.出栈
func (s *StackNode) Pop() (e interface{}, ok bool) {
	if s == nil {
		return 0, false
	}

	e = s.data
	*s = *s.next

	return e, true
}

// 4.取栈顶元素
func (s *StackNode) GetTop() (e interface{}, ok bool) {
	if s == nil {
		return 0, false
	}

	return s.data, true
}

// 1.初始化
func (i *IntStack) Init() bool {
	i = nil
	return true
}

// 2.入栈
func (i *IntStack) Push(e interface{}) bool {
	*i = append(*i, e)
	return true
}

// 3.出栈
func (i *IntStack) Pop() (e interface{}, ok bool) {
	if i.IsEmpty() {
		return 0, false
	}

	index := len(*i) - 1
	e = (*i)[index]
	*i = (*i)[:index]
	return e, true
}

// 取栈顶元素
func (i *IntStack) GetTop() (e interface{}, ok bool) {
	if i.IsEmpty() {
		return 0, false
	}

	index := len(*i) - 1
	return (*i)[index], true
}

// 4.是否为空
func (i *IntStack) IsEmpty() bool {
	return len(*i) == 0
}
