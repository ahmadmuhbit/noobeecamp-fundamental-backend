// models/User.go

package models // define name package

type User struct {
	Name    string
	Contact Contact // artinya, kita melakukan embeded struct. Silahkan lihat materi struct untuk lebih memahaminya.
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
