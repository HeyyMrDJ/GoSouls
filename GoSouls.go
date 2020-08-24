package main

import (
	//"character/character"
	"github.com/eiannone/keyboard"
	"fmt"
)

func main(){
	//you := character.Player{Name: "Ben", Level: 1, Health: 10}
	//hollow := character.Hollow{}
	//hollow.Create_Hollow()
	//character.Opening(you)
	//fmt.Println("Enemy Hollow Attacking!")
	//character.Combat(you, hollow)
	gameloop()

}

func gameloop() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if string(char) == "z"{
			fmt.Println("ATTACKING")
		}
		if string(char) == "x"{
			fmt.Println("BLOCKING")
		}
        if key == keyboard.KeyEsc {
			break
		}
	}
}