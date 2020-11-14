package game

import "fmt"

type Health struct {
	Health    int `bson:"health"`
	MaxHealth int `bson:"maxHealth"`
}

func (h *Health) String() string {
	return fmt.Sprintf("Health: [%v/%v]", h.Health, h.MaxHealth)
}

