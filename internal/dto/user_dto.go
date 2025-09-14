package dto

// Request ตอนสมัครสมาชิก
type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=admin staff customer"`
}

// Response ตอนส่งข้อมูล user กลับไป
type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type FindUserRequest struct {
	ID string `json:"id" validate:"numeric"`
}
