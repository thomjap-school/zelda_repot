package main

import (
	"fmt"
	"strings"
	"math/rand"
)

func fight(player *Character, enemy *Character, classes []Class, races []Race) bool {
	fmt.Printf("\n⚔️ %sA wild %s appears!%s\n\n", Red+Bold, enemy.Name, Reset)

	playerClass := getClassByID(classes, player.Class)
	playerRace := getRaceByID(races, player.Race)
	enemyClass := getClassByID(classes, enemy.Class)
	enemyRace := getRaceByID(races, enemy.Race)

	for player.HP > 0 && enemy.HP > 0 {
		// Affichage HP avec couleurs
		fmt.Printf("%s%s%s: %s%d/%d HP%s\n", Cyan+Bold, player.Name, Reset, Green, player.HP, player.MaxHP, Reset)
		fmt.Printf("%s%s%s: %s%d/%d HP%s\n\n", Red+Bold, enemy.Name, Reset, Red, enemy.HP, enemy.MaxHP, Reset)

		fmt.Printf("Choose action: %s1.Attack%s / %s2.Heal%s / %s3.Protect%s: ",
			Red+Bold, Reset, Green+Bold, Reset, Blue+Bold, Reset)

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// Reset defending each turn
		player.IsDefending = false

		switch input {
		case "attack","1":
			damage, hitType := calculateDamage(player, enemy, playerClass, playerRace, enemyClass, enemyRace)
			enemy.HP -= damage
			if enemy.HP < 0 {
				enemy.HP = 0
			}
			// couleur critique ou dégâts
			msgColor := Red
			if strings.Contains(hitType, "Critical") {
				msgColor = Yellow + Bold
			}
			fmt.Println(msgColor + hitType + fmt.Sprintf("%s attacks %s for %d damage!", player.Name, enemy.Name, damage) + Reset)

		case "heal","2":
			heal := player.MaxHP / 2
			player.HP += heal
			if player.HP > player.MaxHP {
				player.HP = player.MaxHP
			}
			fmt.Println(Green + "💚 You healed!" + Reset)

		case "protect","3":
			player.IsDefending = true
			fmt.Println(Cyan + "🛡️ You are protecting! Damage will be halved this turn." + Reset)

		default:
			fmt.Println(Red + "Invalid action! You lose your turn!" + Reset)
		}

		// Enemy turn
		if enemy.HP > 0 {
			damage, hitType := calculateDamage(enemy, player, enemyClass, enemyRace, playerClass, playerRace)
			if player.IsDefending {
				damage /= 2
			}
			player.HP -= damage
			if player.HP < 0 {
				player.HP = 0
			}
			msgColor := Red
			if strings.Contains(hitType, "Critical") {
				msgColor = Yellow + Bold
			}
			fmt.Println(msgColor + hitType + fmt.Sprintf("%s attacks %s for %d damage!", enemy.Name, player.Name, damage) + Reset)
		}

		fmt.Println("-----------------------")
	}

	if player.HP > 0 {
		fmt.Println(Green + Bold + "✅ You cleared the floor!" + Reset)
	} else {
		fmt.Println(Red + Bold + "💀 You have been defeated..." + Reset)
	}

	return player.HP > 0
}

func calculateDamage(attacker, defender *Character, attackerClass Class, attackerRace Race, defenderClass Class, defenderRace Race) (int, string) {
	damage := attacker.STR
	hitType := ""

	if rand.Intn(100) < attacker.LUCK {
		damage *= 2
		hitType += "Critical Strike! "
	}

	for _, s := range attackerClass.Strengths {
		if s == defender.Class {
			damage *= 2
			hitType += "Crushing hit! "
		}
	}
	for _, w := range attackerClass.Weaknesses {
		if w == defender.Class {
			damage /= 2
			hitType += "Glancing hit! "
		}
	}

	for _, s := range attackerRace.Strength {
		if s == defender.Race {
			damage *= 2
			if hitType == "" {
				hitType += "Crushing hit! "
			}
		}
	}
	for _, w := range attackerRace.Weakness {
		if w == defender.Race {
			damage /= 2
			if hitType == "" {
				hitType += "Glancing hit! "
			}
		}
	}

	damage -= defender.DEF
	if damage < 0 {
		damage = 0
	}

	return damage, hitType
}
