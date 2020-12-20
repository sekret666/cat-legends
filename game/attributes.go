package game

import (
	"fmt"
	"strconv"
	"strings"
)

type Attributes struct {
	Endurance    int `bson:"endurance"`
	Agility      int `bson:"agility"`
	Intelligence int `bson:"intelligence"`
	Strength     int `bson:"strength"`
	Luck         int `bson:"luck"`
}

func (a *Attributes) String() string {
	return fmt.Sprintf("Attributes: Endurance %v Agility %v Intelligence %v Strength %v Luck %v", a.Endurance, a.Agility, a.Intelligence, a.Strength, a.Luck)
}

func (a *Attributes) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%end%"] = strconv.Itoa(a.Endurance)
	m["%agi%"] = strconv.Itoa(a.Agility)
	m["%int%"] = strconv.Itoa(a.Intelligence)
	m["%str%"] = strconv.Itoa(a.Strength)
	m["%luc%"] = strconv.Itoa(a.Luck)
	return m
}

func (a *Attributes) ReplaceInString(text string) string {
	for k, v := range a.GetStringMap() {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
