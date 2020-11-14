package game

import "fmt"

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