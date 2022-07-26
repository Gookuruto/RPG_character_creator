package character

type Class struct {
	Name        string
	Modifiers   Stats
	Description string
	Features    []Feature
	Skills      []Skill
}

func CustomClass(name, desciption string, modifiers Stats, features []Feature) *Class {
	return &Class{Name: name, Description: desciption, Modifiers: modifiers, Features: features}

}
