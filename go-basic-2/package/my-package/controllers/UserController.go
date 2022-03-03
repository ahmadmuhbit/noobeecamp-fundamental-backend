// controllers/UserController.go

package controllers

// kita melakukan import kepada package models, dengan cara
// nama-project/nama-package
import (
	"fmt"
	"my-package/models"
)

type UserController interface {
	SayHello()

	// tambahkan 2 method ini
	SetContact(contact *models.Contact)
	DisplayContact()
}

// Setelah itu, kita membuat sebuah struct baru untuk menerima
// tipe data User yang telah kita buat sebelumnya di package models.
// caranya ialah namaPackage.namaType
type User struct {
	User models.User
}

func NewUserController(user *models.User) UserController {
	return &User{
		User: *user,
	}
}

// melakukan implementasi method
func (u *User) SayHello() {
	fmt.Println("Hello", u.User.Name)
}

func (u *User) SetContact(contact *models.Contact) {
	u.User.Contact = *contact
}

func (u *User) DisplayContact() {
	fmt.Println("Called from user controller ...")
	fmt.Println("My Phone:", u.User.Contact.Phone)
	fmt.Println("My Instagram:", u.User.Contact.Instagram)
}
