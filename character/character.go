package character
import (
	"fmt"
	"os"
	"time"
)

type Player struct {
	Name string
	Level int
	Health int
}

type Hollow struct {
	Health int
}

func(std *Hollow) Create_Hollow() {
	std.Health = 30
}

func Opening (you Player){
	fmt.Println("Welcome to GoSouls!!!!!!!!")
	fmt.Println("Your Character Name is:", you.Name)
	fmt.Println("Your Character Level is:", you.Level)
	fmt.Println("Your Character Health Pool is: ", you.Health)
}

func GameOver (){
	fmt.Println("YOU DIED")
	os.Exit(0)
}

func Combat (you Player, enemy Hollow){
	for you.Health > 0{
		
		for x := 5; x > 0; x--{
			var input string
			fmt.Scanln(&input)
			if input == "x" {
				break
			}
			fmt.Println("Enemy Attacking")
			fmt.Println("Press X to block or Z to attack")
			time.Sleep(1 * time.Second)
			you.Health = you.Health - 1
		}
		enemy.Health = enemy.Health -2
		fmt.Println("Player Health:", you.Health)
		fmt.Println("Enemy Health:", enemy.Health)
		if you.Health <= 0{
			GameOver()
		} else if enemy.Health <= 0 {
			fmt.Println("You beat the enemy!!!!")
			break
		}
	}
	

}