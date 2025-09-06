package user

type UserRegisterReq struct {
	Username string `json:"username" binding:"required" extensions:"x-order=1"`

	// minimal length is 8, minimal 1 uppercase, minimal 1 lowercase, minimal 1 digit and minimal 1 punctuation
	Password string `json:"password" binding:"required" extensions:"x-order=2"`

	ConfirmPassword string `json:"confirmPassword" binding:"required" extensions:"x-order=3"`
}

type UserLoginReq struct {
	Username string `json:"username" binding:"required" extensions:"x-order=1"`

	Password string `json:"password" binding:"required" extensions:"x-order=2"`
}
