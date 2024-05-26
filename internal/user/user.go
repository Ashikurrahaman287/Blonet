package user

type User struct {
	Address    string
	PrivateKey string
}

func NewUser(address, privateKey string) *User {
	return &User{Address: address, PrivateKey: privateKey}
}
