package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ApplyCrit calculates if a critical strike occurs
func ApplyCrit(attacker *Character, baseDamage int) (int, string) {
	hitType := ""
	if rand.Intn(100) < attacker.LUCK { // utiliser le champ exact de ton struct
		baseDamage *= 2
		hitType = "Critical Strike! "
	}
	return baseDamage, hitType
}

// ApplyDef reduces damage by DEF
func ApplyDef(defender *Character, damage int) int {
	damage -= defender.DEF
	if damage < 0 {
		damage = 0
	}
	return damage
}

// ApplyRaceClassMultipliers returns damage multipliers based on race/class
func ApplyRaceClassMultipliers(damage int, attackerClass, defenderClass Class, attackerRace, defenderRace Race) (int, string) {
	hitType := ""

	// Class multiplier
	for _, strong := range attackerClass.Strengths {
		if strong == defenderClass.ID {
			damage *= 2
			hitType = "Crushing hit! "
		}
	}
	for _, weak := range attackerClass.Weaknesses {
		if weak == defenderClass.ID {
			damage /= 2
			hitType = "Glancing hit! "
		}
	}

	// Race multiplier
	for _, strong := range attackerRace.Strength {
		if strong == defenderRace.ID {
			damage *= 2
			if hitType == "" {
				hitType = "Crushing hit! "
			}
		}
	}
	for _, weak := range attackerRace.Weakness {
		if weak == defenderRace.ID {
			damage /= 2
			if hitType == "" {
				hitType = "Glancing hit! "
			}
		}
	}

	return damage, hitType
}
