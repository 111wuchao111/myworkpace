package models

type Article struct {
	Id int `json:"id"`
	//ClassId      int         `json:"class_id"`
	Class        *ArticleClass `orm:"rel(fk)" json:"class_info"`
	Title        string        `json:"title"`
	SubTitle     string        `json:"sub_title"`
	Icon         string        `json:"icon"`
	Content      string        `json:"content"`
	User         *User         `orm:"rel(fk)" json:"user_info"`
	ViewCount    string        `json:"view_count"`
	CommentCount string        `json:"comment_count"`
	ZanCount     string        `json:"zan_count"`
	CreateTime   string        `json:"create_time"`
	UpdateTime   string        `json:"update_time"`
}
