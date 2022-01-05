package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// Character contains attributes about an individual (human or zombie)
type Character struct {
	Name    string   `json:"name"`
	Alias   string   `json:"alias"`
	Hp      int      `json:"hp"`
	Attacks []Attack `json:"attacks"`
}

// Attack contains attributes for a single attack
type Attack struct {
	Name     string `json:"name"`
	Accuracy int    `json:"accuracy"`
	Damage   int    `json:"damage"`
}

// seed is a seed that changes to produce truly random numbers each iteration
var seed = rand.NewSource(time.Now().UnixNano())

// randgen is a random number generator for determining attack success
var randgen = rand.New(seed)

// charactersFromFile reads a list of characters from the specified JSON file
func charactersFromFile(file string) []Character {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file %s", file)

	}

	var characters []Character
	err = json.Unmarshal(content, &characters)
	if err != nil {
		log.Fatalf("Error unmarshalling file %s", file)
	}

	return characters
}

// isAlive verifies if the character is alive based on remaining Hp
func (c *Character) isAlive() bool {
	return c.Hp > 0
}

// Displays prints the character's attributes, including Hp and attacks
func (c *Character) Display() {
	yellow.Printf("[%s] %s (%d HP)\n  Attacks:\n", c.Alias, c.Name, c.Hp)
	c.DisplayAttacks()
	fmt.Println()
}

// DisplayAttacks prints the character's attacks options
func (c *Character) DisplayAttacks() {
	for i, a := range c.Attacks {
		yellow.Printf("  [%d] %s\t(%d Damage, %d Accuracy)\n", i, a.Name, a.Damage, a.Accuracy)
	}
}

// wasSuccessful determines if an attack is successful by generating a random
// number (0-100) and comparing with the Accuracy attribute for the chosen attack
func (a *Attack) wasSuccessful() bool {
	return randgen.Intn(100) <= a.Accuracy
}
