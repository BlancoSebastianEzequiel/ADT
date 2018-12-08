package packets

import "strconv"

type Node struct {
	key string
	value int
	left NodeInterface
	right NodeInterface
}

func (node *Node) Init(key string, value int, left NodeInterface, right NodeInterface) {
	node.key = key
	node.value = value
	node.left = left
	node.right = right
}

func (node Node) GetKey() string {
	return node.key
}

func (node Node) GetValue() int {
	return node.value
}

func (node Node) GetLeft() (left NodeInterface) {
	left = node.left
	return
}

func (node Node) GetRight() (right NodeInterface) {
	right = node.right
	return
}

func (node Node) IsLeave() (val bool) {
	val = false
	return
}

func (node *Node) AddLeft(leftNode NodeInterface) {
	node.left = leftNode
}

func (node *Node) AddRight(rightNode NodeInterface) {
	node.right = rightNode
}

func (node *Node) swap(key string, isLeft bool) {
	var newLeave = Node{}
	newLeave.Init(key, 1, &Leave{}, &Leave{})
	if isLeft {
		node.AddLeft(&newLeave)
		return
	}
	node.AddRight(&newLeave)
}

func (node *Node) Add(key string) (couldAdd bool) {
	if node.key == "" {
		node.key = key
		node.value = 1
		return
	}
	if node.key == key {
		node.value += 1
		couldAdd = true
		return
	}
	if node.key > key {
		couldAdd = node.left.Add(key)
		if !couldAdd {
			node.swap(key, true)
		}
		return
	}
	couldAdd = node.right.Add(key)
	if !couldAdd {
		node.swap(key, false)
	}
	return
}

func (node *Node) Get(key string) (value int) {
	if node.key == key {
		value = node.value
		return
	}
	if node.key > key {
		value = node.left.Get(key)
		return
	}
	value = node.right.Get(key)
	return
}

func (node Node) IsEqual(anotherNode NodeInterface) (val bool) {
	if anotherNode.IsLeave() {
		return false
	}
	if (node.key != anotherNode.GetKey()) || (node.value != anotherNode.GetValue()) {
		return false
	}
	var leftEquals = node.left.IsEqual(anotherNode.GetLeft())
	var rightEquals = node.right.IsEqual(anotherNode.GetRight())
	val = leftEquals && rightEquals
	return
}

func (node Node) Print() string {
	msg := "key: " + node.key + " - value: " + strconv.Itoa(node.value) + "\n"
	msg += "\tLeft:" + node.left.Print() + "\n"
	msg += "\tRight:" + node.right.Print() + "\n"
	return msg
}