package main

import (
	"os"
	"time"

	"github.com/fatih/color"
)

const (
	// character data files
	humanFile  = "data/humans.json"
	zombieFile = "data/zombies.json"

	// pauseDuration is the amount of time the screen will pause between actions
	pauseDuration = 1000 * time.Millisecond
)

var (
	// Re-usable colorized output
	blue   = color.New(color.FgBlue).Add(color.Bold)
	red    = color.New(color.FgRed)
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
)

func main() {
	humans := charactersFromFile(humanFile)
	zombies := charactersFromFile(zombieFile)

	blue.Printf("zombiego\n--------\n\n")
	player := chooseCharacter(humans)
	pauseScreen()

	for _, zombie := range zombies {
		survived := player.Fight(zombie)
		if !survived {
			red.Println("You didn't make it :/")
			os.Exit(0)
		}
	}

	green.Println("You made it!")
}
