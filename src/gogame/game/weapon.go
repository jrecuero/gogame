package game

import "fmt"

// Weapon represents any weapon in the game.
type Weapon struct {
	Object
	Dice
	damage int
}

// NewWeapon function creates a new weapon.
func NewWeapon(name string, faces int, damag int) Weapon {
	return Weapon{
		Object: NewObject(name),
		Dice:   NewDice(faces),
		damage: damag,
	}
}

// Damage function rolls a damage dice.
func (w Weapon) Damage() int {
	return w.Dice.Roll() + w.damage
}

// Print function prints on display Weapon information.
func (w Weapon) Print() {
	fmt.Printf("\t- %s stats:\n\t\t- Dice faces: %d\n\t\t- Damage: %d\n",
		w.Name, w.Faces, w.damage)
}
