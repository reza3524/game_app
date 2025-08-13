package entity

type Round struct {
	Id           uint
	GameId       uint
	QuestionId   uint
	RoundDetails []RoundDetail
}
