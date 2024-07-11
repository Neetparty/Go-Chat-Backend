package dto

type TRegister struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Image     string `json:"image"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	IsDeleted bool   `json:"isDeleted"`
}

type TRegisterResponse struct {
	Message string
	Status  int
	Data    TRegister
}
