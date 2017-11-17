package main

import (
	"fmt"
	"gogame/game"
)

func createPlayer(name string, live int) {
	actor := game.NewActor(name)
	actor.Live = live
	fmt.Println("go-game:", actor.Id, actor.Name, actor.Live)
}

func main() {
	createPlayer("Player ONE", 100)
	createPlayer("Player TWO", 200)
}
