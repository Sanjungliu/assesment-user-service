package user

type RegisterUserInput struct {
	Fullname string `json:"fullname" validate:"required,min=2,max=100"`
	Role     string `json:"role" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
