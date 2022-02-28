package main

import "fmt"

func main() {
	a := 10
	b := 20
	if c := 30; c >= a+b {
		fmt.Println("True")
	}
}
