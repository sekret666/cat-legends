package game

import "fmt"

type BaseEntity struct {
	Level  Level  `bson:"level"`
	Health Health `bson:"health"`
	Mana   Mana   `bson:"mana"`
}

func (b *BaseEntity) String() string {
	return fmt.Sprintf("Entity: %v %v %v", &b.Level, &b.Health, &b.Mana)
}