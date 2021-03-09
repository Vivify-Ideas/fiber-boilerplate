package requests

// LoginRequest -> login input
type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=5,max=50"`
}

// Validate -> login input validation
func (input LoginRequest) Validate() ErrorResponse {
	return Validate(input)
}

// RegisterRequest -> login input
type RegisterRequest struct {
	FirstName string `validate:"required,min=2,max=32"`
	LastName  string `validate:"required,min=2,max=32"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=5,max=50"`
}

// Validate -> login input validation
func (input RegisterRequest) Validate() ErrorResponse {
	return Validate(input)
}

// ResetPasswordRequest -> reset password request
type ResetPasswordRequest struct {
	Email string `validate:"required,email"`
}

// Validate -> reset password validation
func (input ResetPasswordRequest) Validate() ErrorResponse {
	return Validate(input)
}

// ResetPasswordCompleteRequest -> complete reset password request
type ResetPasswordCompleteRequest struct {
	Token    string `validate:"required"`
	Password string `validate:"required,min=5,max=50"`
}

// Validate -> complete reset password validation
func (input ResetPasswordCompleteRequest) Validate() ErrorResponse {
	return Validate(input)
}
