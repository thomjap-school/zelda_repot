package main

import "fmt"

// -------------------------
// Options de combat : Escape & Protect
// -------------------------

func Escape(player *Character) {
    fmt.Println(player.Name, "fled the fight! (repercussions à coder)")
}

func Protect(player *Character) {
    player.IsDefending = true
    fmt.Println(player.Name, "is protecting! Damage will be halved this turn.")
}
