package dto

type CreateAccountDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateAccountDto struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
