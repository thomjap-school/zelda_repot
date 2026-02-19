package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(Cyan+Bold+"Welcome to Hyrule Castle!"+Reset, "\n")

	// Charger JSON avec gestion d'erreurs
	players, err := loadCharacters("../json/players.json")
	if err != nil {
		fmt.Println(Red + "Erreur lors du chargement des joueurs :" + err.Error() + Reset)
		return
	}
	enemies, err := loadCharacters("../json/enemies.json")
	if err != nil {
		fmt.Println(Red + "Erreur lors du chargement des ennemis :" + err.Error() + Reset)
		return
	}
	bosses, err := loadCharacters("../json/bosses.json")
	if err != nil {
		fmt.Println(Red + "Erreur lors du chargement des boss :" + err.Error() + Reset)
		return
	}
	classes, err := loadClasses("../json/classes.json")
	if err != nil {
		fmt.Println(Red + "Erreur lors du chargement des classes :" + err.Error() + Reset)
		return
	}
	races, err := loadRaces("../json/races.json")
	if err != nil {
		fmt.Println(Red + "Erreur lors du chargement des races :" + err.Error() + Reset)
		return
	}

	player := pickRandomCharacter(players)

	// ======================
	// MENU DE DÉMARRAGE
	// ======================
	for {
		fmt.Print(Yellow + "Voulez-vous entrer dans Hyrule Castle ? (yes/no) " + Reset)
		var choice string
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(Red + "Erreur de lecture. Veuillez réessayer." + Reset)
			continue
		}

		choice = strings.ToLower(strings.TrimSpace(choice))
		if choice == "yes" || choice == "y" {
			fmt.Println(Green + "Très bien, l'aventure commence !" + Reset + "\n")
			break
		} else if choice == "no" || choice == "n" {
			fmt.Println(Red + "D'accord, peut-être une autre fois..." + Reset)
			return
		} else {
			fmt.Println(Red + "Réponse invalide. Tapez 'yes' ou 'no'." + Reset)
		}
	}

	// ======================
	// Boucle des étages
	// ======================
	totalFloors := 10
	for floor := 1; floor <= totalFloors; floor++ {
		fmt.Printf("%s🏰 Floor %d%s\n\n", Cyan+Bold, floor, Reset)

		var enemy Character
		if floor == totalFloors {
			enemy = pickRandomCharacter(bosses)
		} else {
			enemy = pickRandomCharacter(enemies)
		}

		survived := fight(&player, &enemy, classes, races)
		if !survived {
			fmt.Println(Red + "💀 Vous avez été vaincu..." + Reset)
			return
		}

		fmt.Printf(Green+"✅ Vous avez terminé le floor %d!\n\n"+Reset, floor)
	}

	fmt.Println(Green + Bold + "🎉 Félicitations ! Vous avez vaincu tous les ennemis et conquis Hyrule Castle !" + Reset)
}
