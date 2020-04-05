package main

import (
	"fmt"
	"learn-go/interfaces"
)

func main() {
	u := interfaces.NewUser()
	fmt.Println(u.GetUsername());
	fmt.Println(u.GetSurname());
}
