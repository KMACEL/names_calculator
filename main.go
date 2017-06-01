package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	b, _ := ioutil.ReadFile("names.txt")
	var s []string = strings.Split(string(b), "\",\"")
	sort.Strings(s)
	var a int
	var total int
	for i := 0; i < len(s); i++ {
		getName := s[i]
		for j := 0; j < len(getName); j++ {
			a = a + int(getName[j]-'@')
		}
		total = total + (a * (i))
		a = 0
	}
	fmt.Println(total)

	elapsed := time.Since(start)

	fmt.Println("\nTotal Time : ", elapsed)
}
