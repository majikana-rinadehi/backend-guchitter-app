package model

// アバター
type Avatar struct {
	AvatarId   int    `json:"avatarId" example:"1234567890" gorm:"primaryKey"`
	AvatarName string `json:"avatarName" example:"Nino"`
	AvatarText string `json:"avatarText" example:"なのよ"`
	ImageUrl   string `json:"imageUrl" example:"https://hoge.com/fuga"`
	Color      string `json:"color" example:"#f6f6f6"`
}
