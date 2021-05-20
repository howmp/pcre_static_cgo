package main

import (
	"fmt"
	"regexp"
)

func pcre() {
	m := MustCompile(`\xFF\xFF`, 0).Matcher([]byte("123\xFF\xFF123"), 0)
	i := m.Index()
	fmt.Println("pcre", i)

}
func gore() {
	m := regexp.MustCompile(`\xFF\xFF`)
	i := m.FindIndex([]byte("123\xFF\xFF123"))
	fmt.Println("gpre", i)
}

func main() {
	pcre()
	gore()
}
