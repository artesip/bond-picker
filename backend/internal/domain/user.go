package domain

type UUID string
type HashedPassword string

type User struct {
	ID       UUID
	Username string         `json:"username"`
	Password HashedPassword `json:"password"`
	Salt     string
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

func (user *User) Validate() error {
	return nil
}
