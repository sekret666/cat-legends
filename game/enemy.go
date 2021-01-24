package game

import (
	"CatLegends/utils"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type EscapeStatus int

const (
	CanEscape EscapeStatus = iota
	CantEscape
)

type Enemy struct {
	BaseEntity
	ChatId       int64        `bson:"chatId"`
	Name         string       `bson:"name"`
	Description  string       `bson:"description"`
	EscapeStatus EscapeStatus `bson:"escapeStatus"`
}

func NewEnemy() *Enemy {
	return &Enemy{}
}

func GenerateEnemy(chatId int64, level int) *Enemy {
	e := NewEnemy()
	e.ChatId = chatId

	e.Name = "Test Enemy"
	e.Description = "Enemy Description"

	e.Level.Level = level

	return e
}

func GetEnemyById(chatId int64) (*Enemy, bool) {
	db := utils.GetDB()

	e := NewEnemy()
	if err := db.Enemies.FindOne(db.Ctx, bson.M{"chatId": chatId}).Decode(&e); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false
		} else {
			log.Error(err)
			return nil, false
		}
	} else {
		return e, true
	}
}

func (e *Enemy) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%name%"] = e.Name
	m["%description%"] = e.Description
	return m
}

func (e *Enemy) ReplaceInString(text string) string {
	for k, v := range e.GetStringMap() {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
