package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// chooseCharacter prompts for a character choice to start the game
func chooseCharacter(humans []Character) (player Character) {
	fmt.Println("Choose your character:")

	for _, human := range humans {
		yellow.Print(human.Display())
	}

	fmt.Print("--> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		for _, h := range humans {
			if input == h.Alias {
				player = h
			}
		}

		if player.Name != "" {
			blue.Printf("%s chosen!\n\n", player.Name)
			break
		}
		fmt.Printf("Invalid option. Choose again.\n--> ")
	}

	return player
}

// chooseAttach prompts for a characters attack choice
func (player *Character) chooseAttack() (attack Attack) {
	fmt.Println("Choose your attack:")
	yellow.Print(player.DisplayAttacks())

	fmt.Print("--> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		for i, a := range player.Attacks {
			if input == strconv.Itoa(i) {
				attack = a
			}
		}

		if attack.Name != "" {
			break
		}
		fmt.Printf("Invalid option. Choose again.\n--> ")
	}

	return attack
}

// Fight handles the fight sequence when the player encounters a zombie
func (player *Character) Fight(zombie Character) bool {
	fmt.Println("Zombie spotted!")
	yellow.Print(zombie.Display())

	for zombie.isAlive() {
		// zombie attack
		fmt.Println("Zombie is attacking...")
		pauseScreen()

		za := zombie.Attacks[0] // only one attack, for now
		if za.wasSuccessful() {
			player.Hp = (player.Hp - za.Damage)
			red.Printf("%s successful for %d damage! (%s: %d HP remaining)\n", za.Name, za.Damage, player.Name, player.Hp)
		} else {
			fmt.Printf("%s missed!\n", za.Name)
		}
		pauseScreen()

		if !player.isAlive() {
			return false
		}

		// player attack
		pa := player.chooseAttack()
		pauseScreen()

		if pa.wasSuccessful() {
			zombie.Hp = (zombie.Hp - pa.Damage)
			green.Printf("%s successful for %d damage! (%s: %d HP remaining)\n", pa.Name, pa.Damage, zombie.Name, zombie.Hp)
		} else {
			fmt.Printf("%s missed!\n", pa.Name)
		}
		pauseScreen()

	}

	green.Printf("Zombie %s defeated!\n\n", zombie.Name)
	pauseScreen()

	return true
}

// pauseScreen is a helper to sleep for a configured amount of time to
// make the scrolling text easier to follow
func pauseScreen() {
	time.Sleep(pauseDuration)
}
