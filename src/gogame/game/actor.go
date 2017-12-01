package game

import (
	"fmt"

	validator "gopkg.in/validator.v2"
)

var dice = NewDice(20)

// Actor represents any player in the game.
type Actor struct {
	Object
	Live    int
	Attack  int `validate:"nonzero"`
	Defense int `validate:"nonzero"`
	*Weapon
}

// NewActor function creates a new Actor.
func NewActor(name string, options ...func(*Actor) error) Actor {
	actor := Actor{Object: NewObject(name)}
	for _, option := range options {
		if err := option(&actor); err != nil {
			panic(err)
		}
	}
	if errs := validator.Validate(actor); errs != nil {
		panic(errs)
	}
	return actor
}

// Hit function rolls a hit dice.
func (a Actor) Hit() int {
	return dice.Roll() + a.Attack
}

// Print function prints on display Actor information.
func (a Actor) Print() {
	fmt.Printf("%s stats:\n\t- Life: %d\n\t- Attack: %d\n\t- Defense: %d\n",
		a.Name, a.Live, a.Attack, a.Defense)
	if a.Weapon != nil {
		a.Weapon.Print()
	}
}
