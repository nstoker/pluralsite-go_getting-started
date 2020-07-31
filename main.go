package main

import (
	"fmt"

	"github.com/pluralsite-go_getting-started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
