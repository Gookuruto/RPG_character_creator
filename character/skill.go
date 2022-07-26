package character

type ISkill interface {
	Use(target Character)
	GetType() string
	GetDescription() string
}

type Skill struct {
	Name            string
	Description     string
	MainStat        string
	Cost            float64
	Is_cost_percont bool
	RequiredLevel   int
	isBought        bool
}

func NewSkill(name, desc, mainStat, succesFormula string, cost float64, is_percent bool) *Skill {
	return &Skill{Name: name, Description: desc, MainStat: mainStat, Cost: cost, Is_cost_percont: is_percent}

}
