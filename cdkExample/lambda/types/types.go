package types

// Where we can define types and other misc. functions

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
