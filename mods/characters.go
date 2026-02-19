package main

import (
	"encoding/json"
	"io/ioutil"
)

type Character struct {
	ID          int
	Name        string
	HP          int
	MaxHP       int
	MP          int
	MaxMP       int
	STR         int
	INT         int
	DEF         int
	RES         int
	SPD         int
	LUCK        int
	Class       int
	Race        int
	Rarity      int
	IsDefending bool
}

type Class struct {
	ID         int
	Name       string
	Strengths  []int
	Weaknesses []int
}

type Race struct {
	ID       int
	Name     string
	Strength []int
	Weakness []int
}

// Load JSON
func loadCharacters(file string) ([]Character, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var chars []Character
	err = json.Unmarshal(data, &chars)
	if err != nil {
		return nil, err
	}

	for i := range chars {
		chars[i].MaxHP = chars[i].HP
		chars[i].MaxMP = chars[i].MP
	}

	return chars, nil
}

func loadClasses(file string) ([]Class, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var classes []Class
	err = json.Unmarshal(data, &classes)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func loadRaces(file string) ([]Race, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var races []Race
	err = json.Unmarshal(data, &races)
	if err != nil {
		return nil, err
	}
	return races, nil
}
