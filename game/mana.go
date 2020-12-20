package game

import (
	"fmt"
	"strconv"
	"strings"
)

type Mana struct {
	Mana    int `bson:"mana"`
	MaxMana int `bson:"maxMana"`
}

func (m *Mana) String() string {
	return fmt.Sprintf("Mana: [%v/%v]", m.Mana, m.MaxMana)
}

func (m *Mana) GetStringMap() map[string]string {
	_m := make(map[string]string)
	_m["%MP%"] = strconv.Itoa(m.Mana)
	_m["%maxMP%"] = strconv.Itoa(m.MaxMana)
	return _m
}

func (m *Mana) ReplaceInString(text string) string {
	for k, v := range m.GetStringMap() {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
