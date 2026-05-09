package domain

import "strings"

type UUID string
type HashedPassword string

type User struct {
	ID       UUID           `json:"id"`
	Username string         `json:"username"`
	Password HashedPassword `json:"password"`
	Salt     string         `json:"salt"`
}

func NewUser(username, password string) (*User, error) {
	user := &User{
		Username: username,
		Password: HashedPassword(password),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func ClearSensitive(user *User) *User {
	if user == nil {
		return nil
	}

	return &User{
		ID:       user.ID,
		Username: user.Username,
	}
}

func (user *User) Validate() error {
	if user == nil {
		return ValidationErr
	}
	trimedUsername := strings.Trim(user.Username, " ")
	trimedPassword := strings.Trim(string(user.Password), " ")

	if len(trimedUsername) < 4 {
		return ValidationErr
	}

	if len(trimedPassword) < 8 {
		return ValidationErr
	}

	return nil
}
