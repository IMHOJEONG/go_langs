package main

import "fmt"

func main() {

	m := make(map[string]int)

	v, ok := m["Answer"]
	fmt.Println(v, ok)
}
