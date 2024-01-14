// https://pkg.go.dev/strings#Fields

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	new_s := strings.Split(s, " ")

	new_map := make(map[string]int)

	for _, v := range new_s {

		_, ok := new_map[v]

		if ok == false {
			new_map[v] = 1
		} else {
			new_map[v] = new_map[v] + 1
		}

	}

	return new_map
}

func main() {
	wc.Test(WordCount)
}
