package entities

type User struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
	Gender    string
}

type Role int

const (
	Admin Role = iota
	Guess
)
