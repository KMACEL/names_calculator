package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

var getNameTxt string
var a int
var total int
var getName string

func main() {
	start := time.Now()
	b, _ := ioutil.ReadFile("names.txt")

	getNameTxt = strings.Replace((strings.TrimSpace(string(b))), "\"", "", -1)

	var s []string = strings.Split(string(getNameTxt), ",")
	sort.Strings(s)

	for i := 0; i < len(s); i++ {
		getName = s[i]
		a = 0
		for j := 0; j < len(getName); j++ {
			a = a + int(getName[j]-'@')
		}
		total = total + (a * (i + 1))

	}
	fmt.Println(total)

	elapsed := time.Since(start)
	fmt.Printf("Total Time : %s \n", elapsed)
}
