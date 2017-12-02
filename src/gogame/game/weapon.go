package game

import "fmt"

// Weapon represents any weapon in the game.
type Weapon struct {
	Object
	Dice
	damage int
}

// NewWeapon function creates a new weapon.
func NewWeapon(name string, faces int, damag int) (*Weapon, error) {
	dice, errs := NewDice(faces)
	if errs != nil {
		return nil, errs
	}
	obj, errs := NewObject(name)
	if errs != nil {
		return nil, errs
	}
	weapon := Weapon{
		Object: *obj,
		Dice:   *dice,
		damage: damag,
	}
	return &weapon, nil
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
