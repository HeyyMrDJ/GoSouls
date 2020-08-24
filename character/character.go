package character
import (
	"fmt"
	"os"
	"time"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"math/rand"
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
	color.Red("Enenmy Hollow attacking")
	color.Red("ENTERING COMBAT")
	fmt.Println("Press X to block or Z to attack")

	for randnum := rand.Intn(10); you.Health > 0;{
		time.Sleep(time.Second * time.Duration(randnum))
		color.Red("Enemy Attacking")

		you.Health = you.Health - 1
		GetInput()
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

func GetInput(){
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if string(char) == "z"{
			color.Yellow("ATTACKING")
			break
		}
		if string(char) == "x"{
			color.Blue("BLOCKING")
			break
		}
        if key == keyboard.KeyEsc {
			break
		}
		if key == keyboard.KeyCtrlC {
			break
		}
	}
}