package model

// ぐち
type Complaint struct {
	ComplaintText string `json:"complaintText" example:"勘弁してくれ!"`
	AvatarId      int    `json:"avatarId" example:"1"`
}
