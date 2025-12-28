package main

import "fmt"

func main(){
	// while loop

	i:=1
	for i <= 3 {
		fmt.Println(i)
		i = i +1
	}

	// for loop

	for i:=0;i<4;i++{
		fmt.Println(i-1)
	}

	for j:= range 12{
		fmt.Println((j))
	}
}