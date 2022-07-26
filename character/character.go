package character

type Character struct {
	Name       string
	Hp         uint64
	Resource   uint64
	Statistics Stats
	Story      string
	Class      Class
	Race       Race
	Skills     []Skill
	ArmorClass string
}
