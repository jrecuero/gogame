package game

var id int = 0

type Objectable interface {
	getName() string
}

type Object struct {
	Id   int
	Name string
}

func NewObject(name string) Object {
	id++
	return Object{id, name}
}
