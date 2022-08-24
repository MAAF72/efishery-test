package models

// UserCreateRequest user create request struct
type UserCreateRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Password    string `json:"password" mapstructure:",omitempty"`
}

// UserCreateResponse user create response struct
type UserCreateResponse struct {
	Base
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}

// UserLoginRequest user login request struct
type UserLoginRequest struct {
	Base
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}
