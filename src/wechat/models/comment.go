package models

type Comments struct {
	Id           int    `json:"id"`
	ArticleId    int    `json:"article_id"`
	UserOpenid   string `json:"user_openid"`
	UserHeadicon string `json:"user_headicon"`
	UserNickname string `json:"user_nickname"`
	ToOpenid     string `json:"to_openid"`
	ToNickname   string `json:"to_nickname"`
	Content      string `json:"content"`
	CreateTime   string `json:"create_time"`
}

func (u *Comments) TableName() string {
	// db table name
	return "article_comments"
}
