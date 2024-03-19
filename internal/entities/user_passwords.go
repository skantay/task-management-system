package entities

type UserPasswords struct {
	UserID   int64  `db:"user_id"`
	Password []byte `db:"password"`
}
