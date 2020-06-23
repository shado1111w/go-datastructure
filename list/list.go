package list

type List interface {
	Init() bool
	GetElem(i int) (e int, ok bool)
	LocateElem(e int) (index int)
	Insert(i, e int) (ok bool)
	Delete(i int) (ok bool)
}