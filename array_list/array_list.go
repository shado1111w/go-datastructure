package array_list

const (
	maxsize = 200
)

type ArrayList struct {
	length int
	elem   []int
}

// 1.初始化
func (a *ArrayList) Init() bool {
	a.elem = make([]int, maxsize)
	return true
}

// 2.取值
func (a *ArrayList) GetElem(i int) (e int, ok bool) {
	if i < 1 || i > a.length {
		ok = false
		return
	}

	e = a.elem[i-1]
	ok = true

	return
}

// 3.查找
func (a *ArrayList) LocateElem(e int) (index int) {
	for i := 0; i < a.length; i++ {
		if a.elem[i] == e {
			index = i + 1
			return
		}
	}

	return
}

// 4.插入
func (a *ArrayList) Insert(i, e int) (ok bool) {
	if i < 1 || i > maxsize {
		ok = false
		return
	}

	if a.length == maxsize {
		ok = false
		return
	}

	// for j := a.Length; j >= i; j-- {
	// 	a.Elem[j] = a.Elem[j-1]
	// }
	copy(a.elem[i:], a.elem[i-1:])

	a.elem[i-1] = e
	a.length++
	ok = true

	return
}

// 5.删除
func (a *ArrayList) Delete(i int) (ok bool) {
	if i < 1 || i > a.length {
		ok = false
		return
	}

	for j := i; j <= a.length; j++ {
		a.elem[j-1] = a.elem[j]
	}
	a.length--
	ok = true

	return
}
