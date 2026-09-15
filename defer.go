package main
import "fmt"
func deferExample(){
	defer fmt.Println("world")
	fmt.Println("hello")
}

