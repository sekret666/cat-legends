package items

import (
	"strconv"
	"strings"
)

type Weapon struct {
	Damage    int  `bson:"damage"`
	OneHanded bool `bson:"oneHanded"`
}

func (w *Weapon) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%damage%"] = strconv.Itoa(w.Damage)
	if w.OneHanded {
		m["%oneHanded%"] = "так"
	} else {
		m["%oneHanded%"] = "ні"
	}
	return m
}

func (w *Weapon) DefaultPattern() string {
	return `<code>Ушкодження: %damage%
Одноручна:  %oneHanded%</code>`
}

func (w *Weapon) Info(pattern string) string {
	for k, v := range w.GetStringMap() {
		pattern = strings.ReplaceAll(pattern, k, v)
	}
	return pattern
}
