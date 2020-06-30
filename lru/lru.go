package lru

type LRUCache struct {
	Cap    int
	Bucket map[int]*Node
	Head   *Node
	Tail   *Node
}

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Cap:    capacity,
		Bucket: make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if item, ok := this.Bucket[key]; ok {
		this.RefreshNode(item)
		return item.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key, value int) {
	if this.Bucket == nil {
		capacity := 100
		this.Cap = capacity
		this.Bucket = make(map[int]*Node, capacity)
	}

	if item, ok := this.Bucket[key]; !ok {
		if len(this.Bucket) >= this.Cap {
			this.RemoveNode(this.Head)
		}

		node := &Node{
			Key:  key,
			Val:  value,
			Pre:  nil,
			Next: nil,
		}
		this.AddNode(node)
		this.Bucket[key] = node
	} else {
		item.Val = value
		this.RefreshNode(item)
	}
}

func (this *LRUCache) RefreshNode(node *Node) {
	if this.Tail == node {
		return
	}

	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) AddNode(node *Node) {
	if this.Tail != nil {
		this.Tail.Next = node
		node.Pre = this.Tail
		node.Next = nil
	}

	this.Tail = node

	if this.Head == nil {
		this.Head = node
		this.Head.Next = nil
	}

	this.Bucket[node.Key] = node
}

func (this *LRUCache) RemoveNode(node *Node) {
	if node == this.Tail {
		this.Tail = this.Tail.Pre
		this.Tail.Next = nil
	} else if node == this.Head {
		this.Head = this.Head.Next
		this.Head.Pre = nil
	} else {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}

	delete(this.Bucket, node.Key)
}
