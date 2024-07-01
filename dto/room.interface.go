package dto

type TRoom struct {
	Name  string   `json:"name" default:"Untitled"`
	Users []string `json:"users"`
	Chat  []TChat  `json:"chat"`
	TDefaultValue
}

type TChat struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Time   string `json:"time"`
	IsMe   bool   `json:"isMe"`
	TDefaultValue
}

type TRoomResponse struct {
	Message string
	Status  int
	Data    TRoom
}
