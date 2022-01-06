// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u user) changeEmail(email string) {
	u.email = email
}

func changeMail(user2 user, email string) {
	user2.email = email
}

// main is the entry point for the application.
func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	lisa := user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	lisa.changeEmail("lisa@newdomain.com")
	changeMail(lisa, "lisa@newdomain.com")
	fmt.Println(lisa)
	lisa.notify()


	str := string("aa")

	modifyString(str)

	fmt.Println(str)
}

func modifyString(str string) {
	str = "bb"
}
