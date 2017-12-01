package game

import (
	"math/rand"
	"time"

	validator "gopkg.in/validator.v2"
)

// Dice represents any dice in the game.
type Dice struct {
	Faces int `validate:"nonzero"`
}

// NewDice function creates a new dice.
func NewDice(faces int) Dice {
	dice := Dice{
		Faces: faces,
	}
	if errs := validator.Validate(dice); errs != nil {
		panic(errs)
	}
	return dice
}

// Roll function roll a dice.
func (d Dice) Roll() int {
	return rand.Intn(d.Faces) + 1
}

// NewSeed generate a new random seed.
func NewSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func init() {
	NewSeed()
}
