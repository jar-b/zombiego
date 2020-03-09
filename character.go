package main

import (
	"fmt"
	"math/rand"
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

	// Seed that changes to produce truly random numbers each iteration
	seed = rand.NewSource(time.Now().UnixNano())
	// A random number generator for determining attack success
	randgen = rand.New(seed)
)

// Verifies if the character is alive, based on remaining Hp
func (c *Character) isAlive() bool {
	if c.Hp > 0 {
		return true
	}
	return false
}

// Displays the characters attributes, including Hp and attacks
func (c *Character) Display() {
	yellow.Printf("[%s] %s (%d HP)\n  Attacks:\n", c.Alias, c.Name, c.Hp)
	c.DisplayAttacks()
	fmt.Println()
}

// Displays the characters attacks options
func (c *Character) DisplayAttacks() {
	for i := 0; i < len(c.Attacks); i++ {
		a := c.Attacks[i]
		yellow.Printf("  [%d] %s\t(%d Damage, %d Accuracy)\n", i, a.Name, a.Damage, a.Accuracy)
	}
}

// Determines if an attack is successful by generating a random
// number (0-100) and comparing with the Accuracy attribute for the
// chosen attack
func (a *Attack) wasSuccessful() bool {
	if randgen.Intn(100) <= a.Accuracy {
		return true
	}
	return false
}
