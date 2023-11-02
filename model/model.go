package model

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Phone   int     `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}

type Error struct {
	Message string
}
