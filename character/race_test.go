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
		Traits: []Feature{{Name: "Dwarven skin", Description: "Gives +5 defence against physical attacks."}}}

	if reflect.DeepEqual(race, expected) {
		t.Errorf("Name is diffrent from expected name. expected %v got %v", expected, race)
	}
}
