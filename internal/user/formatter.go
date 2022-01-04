package user

type UserFormatter struct {
	UserID       string `json:"user_id"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func FormatUser(user User, token, refreshToken string) UserFormatter {
	return UserFormatter{
		UserID:       user.UserID,
		Fullname:     user.Fullname,
		Email:        user.Email,
		Role:         user.Role,
		Token:        token,
		RefreshToken: refreshToken,
	}
}
