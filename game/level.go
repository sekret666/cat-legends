package game

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	levelUpXP = 30
)

type Level struct {
	Level     int `bson:"level"`
	XP        int `bson:"xp"`
	LevelUpXP int `bson:"levelUpXP"`
}

func (l *Level) String() string {
	return fmt.Sprintf("Level: %v [%v/%v]", l.Level, l.XP, l.LevelUpXP)
}

func (l *Level) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%lvl%"] = strconv.Itoa(l.Level)
	m["%XP%"] = strconv.Itoa(l.XP)
	m["%lvlUpXP%"] = strconv.Itoa(l.LevelUpXP)
	return m
}

func (l *Level) ReplaceInString(text string) string {
	for k, v := range l.GetStringMap() {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}

/*

   FIX

func (l *Level)LevelUp()  {
	for l.XP >= l.LevelUpXP {
		l.Level++
		l.XP -= l.LevelUpXP
		l.LevelUpXP =  (2 * l.Level - 1) * levelUpXP
	}
}

*/
