package model

// ぐち
type Complaint struct {
	// ID, CreatedAt, UpdatedAt, DeletedAt が付与される
	// => Error 1054: Unknown column 'created_at' in 'field list'のエラー
	// gorm.Model
	// タグ`gorm:"primaryKey"`を付与。goの構造体は、複数のタグがある場合は半角スペースで区切って記載
	ComplaintId   int    `json:"complaintId" example:"56" gorm:"primaryKey"`
	ComplaintText string `json:"complaintText" example:"勘弁してくれ!"`
	AvatarId      int    `json:"avatarId" example:"1"`
}
