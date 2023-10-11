package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
}

type UserUpdateRequest struct {
	Name  string `json:"name" binding:"omitempty,min=4,max=100"`
	Email string `json:"email" binding:"required,email"`
}
