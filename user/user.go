package user

type User struct {
	Headers map[string]string
}

func NewUser(h map[string]string) *User {
	return &User{
		Headers: h,
	}
}
