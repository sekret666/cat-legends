package game

import (
	"fmt"
	"strconv"
	"strings"
)

type Health struct {
	Health    int `bson:"health"`
	MaxHealth int `bson:"maxHealth"`
}

func (h *Health) String() string {
	return fmt.Sprintf("Health: [%v/%v]", h.Health, h.MaxHealth)
}

func (h *Health) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%HP%"] = strconv.Itoa(h.Health)
	m["%maxHP%"] = strconv.Itoa(h.MaxHealth)
	return m
}

func (h *Health) ReplaceInString(text string) string {
	for k, v := range h.GetStringMap() {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
