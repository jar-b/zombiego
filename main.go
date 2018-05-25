package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type (
	Character struct {
		Name    string
		Alias   string
		Hp      int
		Attacks []Attack
	}

	Attack struct {
		Name     string
		Accuracy int
		Damage   int
	}
)

var (
	// Seed that changes to produce truly random numbers each iteration
	seed = rand.NewSource(time.Now().UnixNano())
	// A random number generator for determining attack success
	randgen = rand.New(seed)

	// Player character options
	Humans = []Character{
		{
			Name:  "Rick",
			Alias: "r",
			Hp:    100,
			Attacks: []Attack{
				{Name: "Punch", Accuracy: 50, Damage: 15},
				{Name: "Revolver", Accuracy: 20, Damage: 60},
			},
		},
		{
			Name:  "Daryl",
			Alias: "d",
			Hp:    100,
			Attacks: []Attack{
				{Name: "Punch", Accuracy: 50, Damage: 15},
				{Name: "Crossbow", Accuracy: 25, Damage: 50},
			},
		},
		{
			Name:  "Carl",
			Alias: "c",
			Hp:    50,
			Attacks: []Attack{
				{Name: "Punch", Accuracy: 40, Damage: 10},
				{Name: "Cower", Accuracy: 100, Damage: 5},
			},
		},
	}

	// Zombie opponents
	Zombies = []Character{
		{
			Name:  "Shane",
			Alias: "s",
			Hp:    10,
			Attacks: []Attack{
				{Name: "Bite", Accuracy: 50, Damage: 10},
			},
		},
		{
			Name:  "Glenn",
			Alias: "g",
			Hp:    20,
			Attacks: []Attack{
				{Name: "Bite", Accuracy: 40, Damage: 20},
			},
		},
	}
)

// Verifies if the character is alive, based on remaining Hp
func (c *Character) isAlive() bool {
	if c.Hp > 0 {
		return true
	} else {
		return false
	}
}

// Displays the characters attributes, including Hp and attacks
func (c *Character) Display() {
	fmt.Printf("[%s] %s (%d HP)\n  Attacks:\n", c.Alias, c.Name, c.Hp)
	c.DisplayAttacks()
	fmt.Println()
}

// Displays the characters attacks options
func (c *Character) DisplayAttacks() {
	for i := 0; i < len(c.Attacks); i++ {
		a := c.Attacks[i]
		fmt.Printf("  [%d] %s\t(%d Damage, %d Accuracy)\n", i, a.Name, a.Damage, a.Accuracy)
	}
}

// Determines if an attack is successful by generating a random
// number (0-100) and comparing with the Accuracy attribute for the
// chosen attack
func (a *Attack) wasSuccessful() bool {
	if randgen.Intn(100) <= a.Accuracy {
		return true
	} else {
		return false
	}
}

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
			fmt.Printf("%s chosen!\n\n", player.Name)
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
			fmt.Printf("%s successful for %d damage! (%s: %d HP remaining)\n", za.Name, zd, player.Name, player.Hp)
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
			fmt.Printf("%s successful for %d damage! (%s: %d HP remaining)\n", pa.Name, pd, zombie.Name, zombie.Hp)
		} else {
			fmt.Printf("%s missed!\n", pa.Name)
		}
		time.Sleep(1000 * time.Millisecond)

	}
	fmt.Printf("Zombie %s defeated!\n\n", zombie.Name)
	time.Sleep(1000 * time.Millisecond)

	return true
}

func main() {
	fmt.Printf("zombiego\n--------\n\n")
	player := chooseCharacter()
	time.Sleep(1000 * time.Millisecond)

	for i := 0; i < len(Zombies); i++ {
		zombie := Zombies[i]
		survived := fightLoop(player, zombie)
		if !survived {
			fmt.Printf("You didn't make it :/\n")
			break
		}

		if survived && i == (len(Zombies)-1) {
			fmt.Printf("You made it!\n")
		}

	}
}
