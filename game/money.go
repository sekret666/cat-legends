package game

import (
	"fmt"
	"strconv"
	"strings"
)

type Money struct {
	Money int `bson:"money"`
}

func (m *Money) String() string {
	return fmt.Sprintf("Money: %v", m.Money)
}

func (m *Money) GetStringMap() map[string]string {
	_m := make(map[string]string)
	_m["%money%"] = strconv.Itoa(m.Money)
	return _m
}

func (m *Money) ReplaceInString(text string) string {
	for k, v := range m.GetStringMap() {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
