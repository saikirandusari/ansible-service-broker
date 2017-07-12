package auth

type AuthProvider interface {
	getPrincipal(*Request) string // Principal
}

type User struct {
	Username string
	Password string
}

type UserServiceAdapter interface {
	findByLogin(string) User
}

type BasicAuth struct {
}

func (b BasicAuth) getPrincipal(r *Request) string { // Principal
}
