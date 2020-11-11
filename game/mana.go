package game

import "fmt"

type Mana struct {
	Mana    int `json:"mana"`
	MaxMana int `json:"max_mana"`
}

func (m *Mana) String() string {
	return fmt.Sprintf("Mana: [%v/%v]", m.Mana, m.MaxMana)
}