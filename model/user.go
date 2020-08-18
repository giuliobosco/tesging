package model

// User of the application
type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}
