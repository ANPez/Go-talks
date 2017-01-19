// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
type User struct {
	Name   string
	Height float32
}

func (u *User) Grow(height float32) { // HL
	u.Height += height
}

func main() {
	var user = User{Name: "name"}
	fmt.Println(user)
	user.Grow(1.5)
	fmt.Println(user)
}

// STOPMAIN OMIT
