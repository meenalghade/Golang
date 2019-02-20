package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "pa$$word123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPasswword := "pa$$word1234"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPasswword))
	if err != nil {
		fmt.Println("Your can't login")
		return
	}

	fmt.Println("You r logged in ")
}
