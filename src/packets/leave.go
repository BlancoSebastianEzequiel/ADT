package packets

type Leave struct {}

func (leave Leave) GetKey() string {
	return ""
}
func (leave Leave) GetValue() int {
	return -1
}

func (leave Leave) GetLeft() (left NodeInterface) {
	left = nil
	return
}

func (leave Leave) GetRight() (right NodeInterface) {
	right = nil
	return
}

func (leave Leave) IsLeave() (val bool) {
	val = true
	return
}

func (leave *Leave) Add(key string) (couldAdd bool) {
	couldAdd = false
	return
}

func (leave *Leave) Get(key string) (value int) {
	value = -1
	return
}

func (leave Leave) IsEqual(anotherNode NodeInterface) (val bool) {
	val = anotherNode.IsLeave()
	return
}

func (leave Leave) Print() string {
	return "leave"
}