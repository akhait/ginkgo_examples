package main 

import (
	"fmt"
	//"log"
	"ginkgo_examples/validation_testing"
	//"ginkgo_examples/validation_ginkgo"
	"ginkgo_examples/benchmarking"
)

func main() {
	
	greeting := validation_testing.RegisterNewUsername("Example_name")
	fmt.Println(greeting)
	/*
	greeting, err := validation_ginkgo.RegisterNewUsername("Example_name")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(greeting)
	}
	
	fmt.Println(benchmarking.ReverseOrder("test"))
	*/
}