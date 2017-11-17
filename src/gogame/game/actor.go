package game

type Actor struct {
	Object
	Live int
}

func NewActor(name string) Actor {
	return Actor{Object: NewObject(name)}
}
