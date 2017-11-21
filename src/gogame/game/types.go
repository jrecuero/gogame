package game

var id int = 0

type Object struct {
	Id   int
	Name string
	Desc string
}

func NewObject(name string) Object {
	id++
	return Object{Id: id, Name: name}
}
