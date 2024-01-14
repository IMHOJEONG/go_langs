package main

import (
	"fmt"
)

var (
	name1 = "TEST"
	name2 = "TEST2"
)

func main() {

	main := "Anna"
	fmt.Println(main, name1, name2)

	names := make([]string, 0, 2)

	fmt.Println(len(names))

	names = append(names, "TEST1")
	names = append(names, "TEST2")

	fmt.Println(len(names))
	fmt.Println(names)

	testing := []string{"a", "b", "c"}
	for i, v := range testing {
		fmt.Println(i, v)
	}

	for _, v := range testing {
		print(v)
	}

}
