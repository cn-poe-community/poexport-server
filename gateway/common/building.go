package common

type Building struct {
	Items         map[string]any `json:"items"`
	PassiveSkills map[string]any `json:"passiveSkills"`
}

func NewBuilding() *Building {
	return &Building{
		Items:         make(map[string]any),
		PassiveSkills: make(map[string]any),
	}
}
