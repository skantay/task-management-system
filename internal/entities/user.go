package entities

type User struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	IsAdmin  bool   `json:"is_admin" db:"is_admin"`
}

type BannedUser struct {
	UserID      int64  `json:"user_id" db:"user_id"`
	Description string `json:"description" db:"description"`
}

const MaxBanDescriptionLen = 100
