package game

import "fmt"

type Mana struct {
	Mana    int `bson:"mana"`
	MaxMana int `bson:"maxMana"`
}

func (m *Mana) String() string {
	return fmt.Sprintf("Mana: [%v/%v]", m.Mana, m.MaxMana)
}