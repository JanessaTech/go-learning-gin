package model

type Entry struct {
	UserId  string `json:"userID"`
	Content string `json:"content" binding:"required"`
}
