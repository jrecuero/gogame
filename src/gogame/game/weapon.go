package game

type Weapon struct {
	Object
	Dice
	damage int
}

func NewWeapon(name string, faces int, damage_ int) Weapon {
	return Weapon{Object: NewObject(name), Dice: NewDice(faces), damage: damage_}
}

func (w Weapon) Damage() int {
	return w.Dice.Roll() + w.damage
}
