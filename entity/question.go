package entity

import "game/enumertion"

type Question struct {
	Id              uint
	Text            string
	Score           int
	CategoryId      uint
	CorrectAnswerId uint
	Difficulty      enumertion.QuestionDifficulty
	Answers         []Answer
}
