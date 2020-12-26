package items

import (
	"strings"
)

type Accessory struct {
	Effect string `bson:"effect"`
}

func (a *Accessory) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%effect%"] = a.Effect
	return m
}

func (a *Accessory) DefaultPattern() string {
	return `<code>Ефект: %effect%</code>`
}

func (a *Accessory) Info(pattern string) string {
	for k, v := range a.GetStringMap() {
		pattern = strings.ReplaceAll(pattern, k, v)
	}
	return pattern
}
