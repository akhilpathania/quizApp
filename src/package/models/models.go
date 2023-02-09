package models

import (
	_ "fmt"
)

type Score struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	UserID   uint   `json:"userid"`
	QuizID   uint   `json:"quizid"`
	GenreID  uint   `json:"genreid"`
	Value    uint   `json:"value"`
	Attempts uint   `json:"attempts"`
}

type User struct {
	ID         uint    `json:"id"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	FirstName  string  `json:"firstname"`
	LastName   string  `json:"lastname"`
	Scores     []Score `json:"scores"`
	TotalScore uint    `json:"totalscore"`
	Admin      bool    `json:"admin"`
}
