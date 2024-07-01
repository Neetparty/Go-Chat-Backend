package dto

type TLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TLoginResponse struct {
	Message string
	Status  int
	Data    TRegister
}
