package packets

type NodeInterface interface {
	GetKey() string
	GetValue() int
	GetLeft() (left NodeInterface)
	GetRight() (right NodeInterface)
	IsLeave() (val bool)
	Add(key string) bool
	Get(key string) int
	IsEqual(anotherNode NodeInterface) bool
	Print() string
}