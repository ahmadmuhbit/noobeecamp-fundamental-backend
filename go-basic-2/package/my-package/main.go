package main

import (
	"my-package/controllers"
	"my-package/models"
)

func main() {
	// define dulu modelsnya
	contactModel := models.NewContact("0812", "@noobeeid")
	userModel := models.NewUser("NooBeeID")

	// lalu di masukkan kedalam controller
	user := controllers.NewUserController(userModel)
	contact := controllers.NewContactController(contactModel)

	// dan siap digunakan
	user.SayHello()
	user.SetContact(contactModel)

	contact.DisplayMyContact()
	user.DisplayContact()
}
