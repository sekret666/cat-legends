package items

import (
	"strconv"
	"strings"
)

type Clothing struct {
	Defence int `bson:"defence"`
}

func (c *Clothing) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%defence%"] = strconv.Itoa(c.Defence)
	return m
}

func (c *Clothing) DefaultPattern() string {
	return `
<code>Захист: %defence%</code>
`
}

func (c *Clothing) Info(pattern string) string {
	for k, v := range c.GetStringMap() {
		pattern = strings.ReplaceAll(pattern, k, v)
	}
	return pattern
}
