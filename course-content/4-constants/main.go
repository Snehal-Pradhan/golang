package main

import "fmt"

func main(){
	const name string = "Albert"
	fmt.Println(name)

	const (
		a = 12
		b = 34
		c = 65
		d = 90
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}