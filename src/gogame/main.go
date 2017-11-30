package main

import (
	//"bufio"
	"fmt"
	"gogame/game"
	//"os"
	"bytes"
	"gopkg.in/validator.v2"
	"strings"
)

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

func createPlayer(name string, live int, desc string) game.Actor {
	actor := game.NewActor(name)
	actor.Live = live
	actor.Desc = desc
	fmt.Printf("go-game: %d : %s : %d - %s\n", actor.Id, actor.Name, actor.Live, actor.Desc)
	return actor
}

func roll(d game.Dice) {
	fmt.Printf("Roll dice with %d faces: %d\n", d.Faces, d.Roll())
}

func createWeapon(name string, faces int, damage int) game.Weapon {
	weapon := game.NewWeapon(name, faces, damage)
	fmt.Printf("weapon: %s : %dD%d\n", weapon.Name, weapon.Dice.Faces, damage)
	return weapon
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	//name, _ := reader.ReadString('\n')
	name := "PLAYER"
	player := createPlayer(strings.TrimSpace(name), 100, "I am player number one")
	player.Weapon = createWeapon("Weapon", 4, 1)
	player.Attack = 2
	player.Defense = 5
	fmt.Printf("%s stats Life: %d, Attack: %d, Defense: %d\n",
		player.Name, player.Live, player.Attack, player.Defense)
	fmt.Printf("%s rolls dice with %d\n", player.Name, player.Hit())
	fmt.Println("Weapon damage for ", player.Weapon.Damage())

	checkValidateModule()
}
