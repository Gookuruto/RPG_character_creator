package character

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Race struct {
	Name      string
	Modifiers Stats
	Traits    []Feature
	Languages []string
}

func CustomRace(name string, modifiers Stats, traits []Feature, languages []string) *Race {
	return &Race{Name: name, Modifiers: modifiers, Traits: traits, Languages: languages}
}

func CreateDefinedRace(raceName string) (*Race, error) {
	switch raceName {
	case "Dwarf":
		return Dwarf(), nil
	case "Human":
	case "Orc":
		return Orc(), nil
	case "Elven":
	case "Half-orc":
	case "Half-elven":
	default:
		return nil, fmt.Errorf("Race %s not exist", raceName)
	}
	return nil, nil
}

func Dwarf() *Race {
	var moifiers Stats = Stats{Strength: 2, Constitution: 3}
	features := make([]Feature, 0)
	features = append(features, Feature{Name: "Dwarven skin", Description: "Gives +5 defence against physical attacks."})
	race := &Race{Name: "Dwarf", Modifiers: moifiers, Traits: features,Languages: make([]string, 0)}

	return race
}

func Orc() *Race {
	modifier := Stats{Strength: 3, Constitution: 2}
	features := make([]Feature, 0)
	features = append(features, Feature{Name: "Brutal Strength", Description: "Gives +10 to attack damage with melee weapons."})
	return &Race{Modifiers: modifier, Traits: features}
}

func (race *Race) SaveToJson(path string) {
	json, err := json.MarshalIndent(&race, "", " ")
	if err != nil {
		panic("Cannot marshall race to json.")
	}
	_ = ioutil.WriteFile(path, json, 0644)
}
