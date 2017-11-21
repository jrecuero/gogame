package game

import (
	"math/rand"
	"time"
)

type Dice struct {
	Faces int
}

func NewDice(faces int) Dice {
	return Dice{Faces: faces}
}

func (d Dice) Roll() int {
	return rand.Intn(d.Faces) + 1
}

func NewSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func init() {
	NewSeed()
}
