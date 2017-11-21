package game

var dice = NewDice(20)

type Actor struct {
	Object
	Live    int
	Attack  int
	Defense int
	Weapon
}

func NewActor(name string) Actor {
	return Actor{Object: NewObject(name)}
}

func (a Actor) Hit() int {
	return dice.Roll() + a.Attack
}
