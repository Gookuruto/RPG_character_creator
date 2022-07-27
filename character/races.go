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
		return Human(), nil
	case "Orc":
		return Orc(), nil
	case "Elven":
		return Elven(), nil
	case "Half-orc":
		return HalfOrc(), nil
	case "Half-elven":
		return HalfElven(), nil
	default:
		return nil, fmt.Errorf("Race %s not exist", raceName)
	}
}

func Dwarf() *Race {
	var moifiers Stats = Stats{Strength: 2, Constitution: 3}
	features := make([]Feature, 0)
	features = append(features, Feature{Name: "Dwarven skin", Description: "Gives +5 defence against physical attacks."})
	race := &Race{Name: "Dwarf", Modifiers: moifiers, Traits: features, Languages: make([]string, 0)}

	return race
}

func Orc() *Race {
	modifier := Stats{Strength: 3, Constitution: 2}
	features := make([]Feature, 0)
	features = append(features, Feature{Name: "Brutal Strength", Description: "Gives +10 to attack damage with melee weapons."})
	return &Race{Name: "Orc", Modifiers: modifier, Traits: features}
}

func HalfOrc() *Race {
	modifier := Stats{Strength: 2, Constitution: 2, Dexterity: 1}
	features := make([]Feature, 0)
	return &Race{Name: "Half-Orc", Modifiers: modifier, Traits: features}

}

func Elven() *Race {
	modifier := Stats{Dexterity: 2, Wisdom: 2, Inteligence: 1}
	features := make([]Feature, 2)
	features[0] = Feature{Name: "Hawk Eye", Description: "Increase accuracy with range weapons by 10"}
	features[1] = Feature{Name: "Precise Eye", Description: "Increase critical chance by 10% with ranged weapons and spells."}
	return &Race{Name: "Elven", Modifiers: modifier, Traits: features}
}

func HalfElven() *Race {
	modifier := Stats{Charisma: 2, Dexterity: 2, Wisdom: 1}
	features := make([]Feature, 0)
	return &Race{Name: "Half-Elven", Modifiers: modifier, Traits: features}
}

func Human() *Race {
	modifier := Stats{Charisma: 2, Wisdom: 1, Strength: 1, Dexterity: 1}
	features := make([]Feature, 0)
	return &Race{Name: "Human", Modifiers: modifier, Traits: features}
}

func (race *Race) SaveToJson(path string) {
	json, err := json.MarshalIndent(&race, "", " ")
	if err != nil {
		panic("Cannot marshall race to json.")
	}
	_ = ioutil.WriteFile(path, json, 0644)
}
