package character

import (
	"reflect"
	"testing"
)

func TestRaceCreation(t *testing.T) {
	race := CustomRace("Test", Stats{}, make([]Feature, 0), make([]string, 0))
	expected_name := "Test"
	if race.Name != expected_name {
		t.Errorf("Name is diffrent from expected name. expected %s got %s", expected_name, race.Name)
	}
}

func TestDwarf(t *testing.T) {
	race := *Dwarf()
	expected := Race{Name: "Dwarf", Modifiers: Stats{Strength: 2, Constitution: 3},
		Traits: []Feature{{Name: "Dwarven skin", Description: "Gives +5 defence against physical attacks."}}, Languages: []string{}}

	if !reflect.DeepEqual(race, expected) {
		t.Errorf("Race is diffrent from expected. expected %v got %v", expected, race)
	}
}

func TestOrc(t *testing.T) {
	race := *Orc()
	expected := Race{Name: "Orc", Modifiers: Stats{Strength: 3, Constitution: 2}, Traits: []Feature{{Name: "Brutal Strength", Description: "Gives +10 to attack damage with melee weapons."}}}
	if !reflect.DeepEqual(race, expected) {
		t.Errorf("Race is diffrent from expected. expected %v got %v", expected, race)
	}
}

func TestHuman(t *testing.T) {
	race := *Human()
	modifier := Stats{Charisma: 2, Wisdom: 1, Strength: 1, Dexterity: 1}
	features := make([]Feature, 0)
	expected := Race{Name: "Human", Modifiers: modifier, Traits: features}
	if !reflect.DeepEqual(race, expected) {
		t.Errorf("Race is diffrent from expected. expected %v got %v", expected, race)
	}
}
