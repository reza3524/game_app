package entity

type Game struct {
	Id          uint
	Score       uint
	CategoryId  uint
	WinnerId    uint
	QuestionsId []uint
	Users       []User
}
