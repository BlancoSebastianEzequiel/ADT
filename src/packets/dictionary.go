package packets

type Dictionary struct {
	Root Node
}

func (dictionary Dictionary) NewDicc() (newDictionary Dictionary) {
	root := Node{}
	root.Init("", 0, &Leave{}, &Leave{})
	newDictionary = Dictionary{Root: root}
	return
}

func (dictionary *Dictionary) Put(key string) () {
	dictionary.Root.Add(key)
	return
}

func (dictionary Dictionary) Get(key string) (value int) {
	value = dictionary.Root.Get(key)
	return
}

func (dictionary Dictionary) Equals(anotherDictionary Dictionary) (value bool) {
	value = dictionary.Root.IsEqual(&anotherDictionary.Root)
	return
}

func (dictionary Dictionary) Print() string {
	return dictionary.Root.Print()
}