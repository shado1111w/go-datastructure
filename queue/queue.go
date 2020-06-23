package queue

const (
	maxsize = 5
)

type SqQueue struct {
	front int
	rear  int
	data  []int
}

type QNode struct {
	data int
	next *QNode
}

type LinkQueue struct {
	front *QNode
	rear  *QNode
}

// 1.初始化
func (s *SqQueue) Init() bool {
	s.data = make([]int, maxsize)
	s.rear = 0
	s.front = s.rear

	return true
}

// 2.求队列长度
func (s *SqQueue) Length() int {
	return (s.rear - s.front + maxsize) % maxsize
}

// 3.入队
func (s *SqQueue) EnQueue(e int) bool {
	if (s.rear+1)%maxsize == s.front {
		return false
	}

	s.data[s.rear] = e
	s.rear = (s.rear + 1) % maxsize

	return true
}

// 4.出队
func (s *SqQueue) DeQueue() (int, bool) {
	if s.rear == s.front {
		return 0, false
	}

	e := s.data[s.front]
	s.front = (s.front + 1) % maxsize

	return e, true
}

// 5.取队头元素
func (s *SqQueue) GetHead() (int, bool) {
	if s.rear == s.front {
		return 0, false
	}

	e := s.data[s.front]

	return e, true
}

// 1.初始化
func (l *LinkQueue) Init() bool {
	q := new(QNode)
	l.front = q
	l.rear = q

	l.front.next = nil

	return true
}

// 2.入队
func (l *LinkQueue) EnQueue(e int) bool {
	q := new(QNode)
	q.data = e
	q.next = nil

	l.rear.next = q
	l.rear = q

	return true
}

// 3.出队
func (l *LinkQueue) DeQueue() (int, bool) {
	if l.rear == l.front {
		return 0, false
	}

	q := l.front.next
	e := q.data
	l.front.next = q.next

	// 当最后一个元素被删，队尾指针重新指向头结点
	if l.rear == q {
		l.rear = l.front
	}

	return e, true
}

// 4.取队头元素
func (l *LinkQueue) GetHead() (int, bool) {
	if l.rear == l.front {
		return 0, false
	}

	return l.front.next.data, true
}
