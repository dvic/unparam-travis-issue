package main

import "fmt"

func main() {
	Check(
		func() (string, error) {
			fmt.Println("this is a test")
			return "", nil
		},
		"",
	)
	Check(
		func(v int) (string, error) {
			fmt.Println("this is another test")
			return "", nil
		},
		nil,
	)
}

func Check(f interface{}, g interface{}) {
	fmt.Println(f) // do nothing
	fmt.Println(g) // do nothing
}
