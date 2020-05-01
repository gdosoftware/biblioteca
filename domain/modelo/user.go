package modelo

type User struct {
	username    string
	token       string
	permissions []string
}

func NewUser(username string, token string, permissions []string) *User {
	return &User{
		username:    username,
		token:       token,
		permissions: permissions,
	}
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetToken() string {
	return u.token
}

func (u *User) GetPermissions() []string {
	return u.permissions
}
