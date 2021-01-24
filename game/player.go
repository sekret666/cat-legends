package game

import (
	"CatLegends/utils"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Player struct {
	BaseEntity `bson:"base"`
	ChatId     int64     `bson:"chatId"`
	Inventory  Inventory `bson:"inventory"`
}

func NewPlayer() *Player {
	return &Player{}
}

func InitPlayer(chatId int64) *Player {
	p := NewPlayer()
	p.ChatId = chatId

	p.Level = Level{
		Level:     1,
		XP:        0,
		LevelUpXP: 10,
	}

	p.Health = Health{
		Health:    30,
		MaxHealth: 30,
	}

	p.Mana = Mana{
		Mana:    0,
		MaxMana: 0,
	}

	return p
}

func GetPlayerById(chatId int64) (*Player, bool) {
	db := utils.GetDB()

	p := NewPlayer()
	if err := db.Players.FindOne(db.Ctx, bson.M{"chatId": chatId}).Decode(&p); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false
		} else {
			log.Error(err)
			return nil, false
		}
	} else {
		return p, true
	}
}
