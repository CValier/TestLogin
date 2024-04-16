package entity

// Represents a user
type User struct {
	User        string `json:"user" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// Represents the credentials to access
type Credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
