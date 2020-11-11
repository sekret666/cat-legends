package game

import "fmt"

type BaseEntity struct {
	Level  Level  `json:"level"`
	Health Health `json:"health"`
	Mana   Mana   `json:"mana"`
}

func (b *BaseEntity) String() string {
	return fmt.Sprintf("Entity: %v %v %v", &b.Level, &b.Health, &b.Mana)
}