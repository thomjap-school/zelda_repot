package main

import (
	"math/rand"
)

func pickRandomCharacter(chars []Character) Character {
	if len(chars) == 0 {
		panic("empty character slice")
	}

	total := 0
	for _, c := range chars {
		total += c.Rarity
	}
	r := rand.Intn(total)
	for _, c := range chars {
		r -= c.Rarity
		if r < 0 {
			return c
		}
	}
	return chars[0]
}

func getClassByID(classes []Class, id int) Class {
	for _, c := range classes {
		if c.ID == id {
			return c
		}
	}
	return Class{}
}

func getRaceByID(races []Race, id int) Race {
	for _, r := range races {
		if r.ID == id {
			return r
		}
	}
	return Race{}
}

// Enemy scaling
func scaleEnemy(enemy Character, floor int) Character {
	multiplier := 1.0 + float64(floor-1)*0.1
	enemy.HP = int(float64(enemy.HP) * multiplier)
	enemy.MaxHP = enemy.HP
	enemy.MP = int(float64(enemy.MP) * multiplier)
	enemy.MaxMP = enemy.MP
	enemy.STR = int(float64(enemy.STR) * multiplier)
	enemy.INT = int(float64(enemy.INT) * multiplier)
	enemy.DEF = int(float64(enemy.DEF) * multiplier)
	enemy.RES = int(float64(enemy.RES) * multiplier)
	enemy.SPD = int(float64(enemy.SPD) * multiplier)
	return enemy
}
