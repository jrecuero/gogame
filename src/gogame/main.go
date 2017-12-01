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
	actor.Attack = 2
	actor.Defense = 5
	actor.Desc = "I am player number one"
	weapon := game.NewWeapon("Sword", 4, 1)
	actor.Weapon = &weapon
	return nil
}

func getName() string {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter name: ")
	//name, _ := reader.ReadString('\n')
	name := "PLAYER"
	return strings.TrimSpace(name)
}

func main() {
	player := game.NewActor(getName(), createPlayer)
	player.Print()
	fmt.Printf("%s rolls dice with %d\n", player.Name, player.Hit())
	fmt.Println("Weapon damage for ", player.Weapon.Damage())

	//checkValidateModule()
}
