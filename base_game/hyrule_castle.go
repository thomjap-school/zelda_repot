package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type Character struct {
	Name   string
	HP     int
	MaxHP  int
	STR    int
}

func (c *Character) Attack(target *Character) {
	target.HP -= c.STR
	if target.HP < 0 {
		target.HP = 0
	}
}

func (c *Character) Heal() {
	healAmount := c.MaxHP / 2
	c.HP += healAmount
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
}

func (c *Character) IsAlive() bool {
	return c.HP > 0
}

func fight(player *Character, enemy *Character) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n⚔️ A wild %s appears!\n\n", enemy.Name)

	for player.IsAlive() && enemy.IsAlive() {

		fmt.Printf("%s: %d/%d HP\n", player.Name, player.HP, player.MaxHP)
		fmt.Printf("%s: %d/%d HP\n\n", enemy.Name, enemy.HP, enemy.MaxHP)

		fmt.Print("Choose action (Attack / Heal): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "attack" {
			player.Attack(enemy)
			fmt.Printf("%s attacks %s for %d damage!\n", player.Name, enemy.Name, player.STR)
		} else if input == "heal" {
			player.Heal()
			fmt.Println("You healed!")
		} else {
			fmt.Println("Invalid action. You lose your turn!")
		}

		if enemy.IsAlive() {
			enemy.Attack(player)
			fmt.Printf("%s attacks for %d damage!\n", enemy.Name, enemy.STR)
		}

		fmt.Println("-----------------------")
	}

	return player.IsAlive()
}

func main() {
	fmt.Println("Welcome to Hyrule Castle!\n")

	player := Character{
		Name:  "Link",
		HP:    60,
		MaxHP: 60,
		STR:   15,
	}

	for floor := 1; floor <= 10; floor++ {

		fmt.Printf("🏰 Floor %d\n\n", floor)

		var enemy Character

		if floor == 10 {
			enemy = Character{
				Name:  "Ganon",
				HP:    150,
				MaxHP: 150,
				STR:   20,
			}
		} else {
			enemy = Character{
				Name:  "Bokoblin",
				HP:    30,
				MaxHP: 30,
				STR:   5,
			}
		}

		survived := fight(&player, &enemy)

		if !survived {
			fmt.Println("💀 You have been defeated...")
			return
		}

		fmt.Printf("✅ You cleared floor %d!\n\n", floor)
	}

	fmt.Println("🎉 Congratulations! You defeated Ganon and cleared Hyrule Castle!")
}
