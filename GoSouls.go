package main

import (
	"character/character"
)

func main(){
	you := character.Player{Name: "Ben", Level: 1, Health: 10}
	hollow := character.Hollow{}
	hollow.Create_Hollow()
	character.Opening(you)
	character.Combat(you, hollow)
}