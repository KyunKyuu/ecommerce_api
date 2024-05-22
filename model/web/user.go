package web

type RoleAllowed string

type UserUpdateServiceRequest struct {
	Name  string `json:"name"`
	Email string `validate:"email" json:"email"`
	Role  string `validate:"required,oneof=buyer seller" json:"role"`
}

type UserServiceRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	Role     string `validate:"required,oneof=buyer seller" json:"role"`
}

type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
