package session

// UserSession authenticated and do other stuff we will define for the user session
type UserSession struct {
}

// Authenticate is the implementation of crypt.Authenticator interface,
// it indicates the ability to determine whether a given username and password combination are valid or not.
func (ia *UserSession) Authenticate(username, password string) bool {
	return username == password
}
