package domain

type User struct {
	Username     string
	PasswordHash string // password hash to simulate secure database storage
	Salt         []byte
	Balance      uint32
}
