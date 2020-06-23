package array_list

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	list := new(ArrayList)
	// 初始化
	ok := list.Init()
	if !ok {
		t.Error("初始化失败")
	}
	// 赋值
	ok = list.Insert(1, 1)
	if ok {
		t.Logf("length:%d,elem:%v\n", list.length, list.elem)
	}
	ok = list.Insert(2, 2)
	if ok {
		t.Logf("length:%d,elem:%v\n", list.length, list.elem)
	}
	ok = list.Insert(3, 3)
	if ok {
		t.Logf("length:%d,elem:%v\n", list.length, list.elem)
	}
	ok = list.Insert(4, 4)
	if ok {
		t.Logf("length:%d,elem:%v\n", list.length, list.elem)
	}
	ok = list.Insert(5, 5)
	if ok {
		t.Logf("length:%d,elem:%v\n", list.length, list.elem)
	}

	// 取值
	e, ok := list.GetElem(3)
	if ok {
		t.Logf("index:3,elem:%d\n", e)
	}

	// 查找
	index := list.LocateElem(3)
	t.Logf("elem:3,index:%d\n", index)

	// 删除
	t.Logf("length:%d,list:%v\n", list.length, list.elem)
	ok = list.Delete(3)
	if ok {
		t.Logf("length:%d,list:%v\n", list.length, list.elem)
	}
}
