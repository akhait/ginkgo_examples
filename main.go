package main 

import (
	"fmt"
	"log"
	//"ginkgo_examples/validation_testing"
	"ginkgo_examples/validation_ginkgo"
)

func main() {
	greeting, err := validation_ginkgo.RegisterNewUsername("Example_name")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(greeting)
	}
}