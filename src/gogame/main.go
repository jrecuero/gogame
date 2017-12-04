package main

import (
	//"bufio"
	"fmt"
	"gogame/game"
	//"os"
	"bytes"
	"strings"

	"gopkg.in/validator.v2"
)

// User defined structure.
type User struct {
	FirstName string `validate:"min=3,max=40"`
	LastName  string `validate:"min=3, max=8"`
	Age       int    `validate:"min=18"`
	Zip       string `validate:"nonzero"`
	Salute    func(salutation string) string
}

func checkValidateModule() {
	ur := User{
		FirstName: "ok",
		LastName:  "Recuero Arias",
		Age:       10,
		Salute: func(salutation string) string {
			var buffer bytes.Buffer
			buffer.WriteString(salutation)
			buffer.WriteString(" world!")
			return buffer.String()
		},
	}
	if errs := validator.Validate(ur); errs != nil {
		err := errs.(validator.ErrorMap)
		for f, e := range err {
			fmt.Printf("\t - %s (%v)\n", f, e)
		}
	} else {
		fmt.Println("User said: ", ur.Salute("Hello"))
	}
}

func roll(d game.Dice) {
	fmt.Printf("Roll dice with %d faces: %d\n", d.Faces, d.Roll())
}

func createPlayer(actor *game.Actor) error {
	actor.Live = 100
	actor.Attack = 5
	actor.Defense = 10
	actor.Desc = "I am the main player"
	weapon, _ := game.NewWeapon("Sword", 4, 1)
	actor.Weapon = weapon
	return nil
}

func createEnemy(actor *game.Actor) error {
	actor.Live = 100
	actor.Attack = 2
	actor.Defense = 5
	actor.Desc = "I am the main enemy"
	weapon, _ := game.NewWeapon("Axe", 3, 1)
	actor.Weapon = weapon
	return nil
}

func getName() string {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter name: ")
	//name, _ := reader.ReadString('\n')
	name := "PLAYER"
	return strings.TrimSpace(name)
}

func isAHit(orig *game.Actor, target *game.Actor) bool {
	//return orig.Hit() > target.Defense
	hit := orig.Hit()
	defense := target.Defense
	fmt.Printf("%s hit %s : %d vs %d\n", orig.Name, target.Name, hit, defense)
	return hit > defense
}

func main() {
	player, errs := game.NewActor("PLAYER", createPlayer)
	if errs != nil {
		panic(errs)
	}
	enemy, errs := game.NewActor("ENEMY", createEnemy)
	if errs != nil {
		panic(errs)
	}
	player.Print()
	enemy.Print()
	if isAHit(player, enemy) {
		damage := player.Weapon.Damage()
		fmt.Printf("%s damage %s for %d\n", player.Name, enemy.Name, damage)
	}

	//checkValidateModule()
}
