package main

import (
	"time"

	"github.com/fatih/color"
)

var (
	// Re-usable colorized output
	blue   = color.New(color.FgBlue).Add(color.Bold)
	red    = color.New(color.FgRed)
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
)

func main() {
	blue.Printf("zombiego\n--------\n\n")
	player := chooseCharacter()
	time.Sleep(1000 * time.Millisecond)

	for i := 0; i < len(Zombies); i++ {
		zombie := Zombies[i]
		survived := fightLoop(player, zombie)
		if !survived {
			red.Printf("You didn't make it :/\n")
			break
		}

		if survived && i == (len(Zombies)-1) {
			green.Printf("You made it!\n")
		}

	}
}
