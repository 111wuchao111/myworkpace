package models

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	HeadIcon string `json:"head_icon"`
}
