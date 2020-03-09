package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Prompts for character choice to start the game
func chooseCharacter() (player Character) {
	fmt.Printf("Choose your character:\n")

	for i := 0; i < len(Humans); i++ {
		Humans[i].Display()
	}

	fmt.Print("--> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch input := scanner.Text(); input {
		case "r":
			player = Humans[0]
		case "d":
			player = Humans[1]
		case "c":
			player = Humans[2]
		default:
			fmt.Printf("Invalid option. Choose again.\n--> ")
		}

		if player.Name != "" {
			blue.Printf("%s chosen!\n\n", player.Name)
			break
		}
	}

	return player
}

// Prompts for a characters attack choice
func chooseAttack(player Character) (attack Attack) {
	fmt.Printf("Choose your attack:\n")
	player.DisplayAttacks()

	fmt.Print("--> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		for i := 0; i < len(player.Attacks); i++ {
			if input == strconv.Itoa(i) {
				attack = player.Attacks[i]
			}
		}

		if attack.Name != "" {
			break
		} else {
			fmt.Printf("Invalid option. Choose again.\n--> ")
		}
	}

	return attack
}

// Handles the fight sequence when the player encounters
// a zombie
func fightLoop(player Character, zombie Character) bool {
	fmt.Printf("Zombie spotted!\n")
	zombie.Display()

	for zombie.isAlive() {
		// zombie attack
		var zd int
		za := zombie.Attacks[0] // only one attack, for now

		fmt.Printf("Zombie is attacking...\n")
		time.Sleep(1000 * time.Millisecond)

		if za.wasSuccessful() {
			zd = za.Damage
			player.Hp = (player.Hp - zd)
			red.Printf("%s successful for %d damage! (%s: %d HP remaining)\n", za.Name, zd, player.Name, player.Hp)
		} else {
			fmt.Printf("%s missed!\n", za.Name)
		}
		time.Sleep(1000 * time.Millisecond)

		if !player.isAlive() {
			return false
		}

		// player attack
		var pd int
		pa := chooseAttack(player)
		time.Sleep(1000 * time.Millisecond)

		if pa.wasSuccessful() {
			pd = pa.Damage
			zombie.Hp = (zombie.Hp - pd)
			green.Printf("%s successful for %d damage! (%s: %d HP remaining)\n", pa.Name, pd, zombie.Name, zombie.Hp)
		} else {
			fmt.Printf("%s missed!\n", pa.Name)
		}
		time.Sleep(1000 * time.Millisecond)

	}
	green.Printf("Zombie %s defeated!\n\n", zombie.Name)
	time.Sleep(1000 * time.Millisecond)

	return true
}
