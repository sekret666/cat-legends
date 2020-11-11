package game

import "fmt"

type Health struct {
	Health    int `json:"health"`
	MaxHealth int `json:"max_health"`
}

func (h *Health) String() string {
	return fmt.Sprintf("Health: [%v/%v]", h.Health, h.MaxHealth)
}

