package main 

import (
	"fmt"
	"log"
	"ginkgo_examples/validation"
)

func main() {
	greeting, err := validation.RegisterNewUsername("Example_name")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(greeting)
	}
}