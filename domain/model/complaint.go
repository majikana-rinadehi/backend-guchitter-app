package model

// ぐち
type Complaint struct {
	ComplaintId   int
	ComplaintText string `json:"complaintText" example:"勘弁してくれ!"`
	AvatarId      string `json:"avatarId" example:"1"`
}
