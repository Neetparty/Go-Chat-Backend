package dto

type TRoom struct {
	Name      string   `json:"name" default:"Untitled"`
	Users     []string `json:"users"`
	Chat      []TChat  `json:"chat"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	IsDeleted bool     `json:"isDeleted"`
}

type TChat struct {
	Sender    string `json:"sender"`
	Text      string `json:"text"`
	Time      string `json:"time"`
	IsMe      bool   `json:"isMe"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	IsDeleted bool   `json:"isDeleted"`
}

type TRoomResponse struct {
	Message string
	Status  int
	Data    TRoom
}
