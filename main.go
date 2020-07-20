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

	// character data files
	humanFile  = "data/humans.json"
	zombieFile = "data/zombies.json"
)

func main() {
	humans := charactersFromFile(humanFile)
	zombies := charactersFromFile(zombieFile)

	blue.Printf("zombiego\n--------\n\n")
	player := chooseCharacter(humans)
	time.Sleep(1000 * time.Millisecond)

	for i := 0; i < len(zombies); i++ {
		zombie := zombies[i]
		survived := fightLoop(player, zombie)
		if !survived {
			red.Printf("You didn't make it :/\n")
			break
		}

		if survived && i == (len(zombies)-1) {
			green.Printf("You made it!\n")
		}

	}
}
