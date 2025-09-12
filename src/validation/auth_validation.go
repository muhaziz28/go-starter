package validation

type Register struct {
	Name     string `json:"name" validate:"required,max=50" example:"fake name"`
	Email    string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
	Password string `json:"password" validate:"required,min=8,max=20,password" example:"password1"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
	Password string `json:"password" validate:"required,min=8,max=20,password" example:"password1"`
}

type Logout struct {
	RefreshToken string `json:"refresh_token" validate:"required,max=255"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required,max=255"`
}

type ForgotPassword struct {
	Email string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
}

type Token struct {
	Token string `json:"token" validate:"required,max=255"`
}
