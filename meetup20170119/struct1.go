// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
type User struct { // HL
	Name     string
	Height   float32
	password string
}

func main() {
	u1 := new(User) // HL
	u2 := User{
		Name:     "user",
		Height:   1.75,
		password: "123456",
	}
	fmt.Println(u1, u2)
}

// STOPMAIN OMIT
