package domain

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	Password string `json:"-"`
}

type UserUsecases interface{}
type UserRepository interface{}
