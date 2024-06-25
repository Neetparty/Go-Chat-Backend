package dto

type TRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TRegisterResponse struct {
	Message string
	Status  int
	Data    TRegister
}
