package list

import (
	"testing"

	_ "shado1111w.com/algorithm/array_list"
	"shado1111w.com/algorithm/linked_list"
)

func TestList(t *testing.T) {
	var list List
	// list = new(array_list.ArrayList)
	list = new(linked_list.LNode)

	// 初始化
	ok := list.Init()
	if !ok {
		t.Error("初始化失败")
	}

	// 赋值
	list.Insert(1, 1)
	list.Insert(2, 2)
	list.Insert(3, 3)
	list.Insert(4, 4)
	list.Insert(5, 5)

	// 取值
	e, ok := list.GetElem(3)
	if ok {
		t.Logf("index:3,elem:%d\n", e)
	}

	// 查找
	index := list.LocateElem(3)
	t.Logf("elem:3,index:%d\n", index)

	// 删除
	list.Delete(3)
}
