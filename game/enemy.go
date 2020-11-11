package game

type Enemy struct {
	BaseEntity
}

func NewEnemy() *Enemy {
	return &Enemy{}
}
