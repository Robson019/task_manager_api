package request

type LoginDTO struct {
	Email    string `json:"email" example:"robson@gmail.com"`
	Password string `json:"password" example:"Test1234!"`
}

type RefreshDTO struct {
	RefreshToken string `json:"refresh_token"`
}
