package game

import "fmt"

const (
	levelUpXP = 30
)

type Level struct {
	Level     int `json:"level"`
	XP        int `json:"xp"`
	LevelUpXP int `json:"level_up_xp"`
}

func (l *Level) String() string {
	return fmt.Sprintf("Level: %v [%v/%v]", l.Level, l.XP, l.LevelUpXP)
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