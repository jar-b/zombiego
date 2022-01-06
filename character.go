package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
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

var (
	// seed is a seed that changes to produce truly random numbers each iteration
	seed = rand.NewSource(time.Now().UnixNano())

	// randgen is a random number generator for determining attack success
	randgen = rand.New(seed)
)

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

// Displays prints the character's attributes, including Hp and attacks
func (c *Character) Display() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("[%s] %s (%d HP)\n  Attacks:\n%s\n", c.Alias, c.Name, c.Hp, c.DisplayAttacks()))
	return b.String()
}

// DisplayAttacks prints the character's attacks options
func (c *Character) DisplayAttacks() string {
	var b strings.Builder
	for i, a := range c.Attacks {
		b.WriteString(fmt.Sprintf("  [%d] %s\t(%d Damage, %d Accuracy)\n", i, a.Name, a.Damage, a.Accuracy))
	}
	return b.String()
}

// isAlive verifies if the character is alive based on remaining Hp
func (c *Character) isAlive() bool {
	return c.Hp > 0
}

// wasSuccessful determines if an attack is successful by generating a random
// number (0-100) and comparing with the Accuracy attribute for the chosen attack
func (a *Attack) wasSuccessful() bool {
	return randgen.Intn(100) <= a.Accuracy
}
