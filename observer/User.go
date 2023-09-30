package main

import "fmt"

type User struct {
	id int
}

func (u *User) update(msg string) {
	fmt.Printf("%s for %d id\n", msg, u.id)
}

func (u *User) getID() int {
	return u.id
}
