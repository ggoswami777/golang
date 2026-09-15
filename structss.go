package main

import "fmt"

type Vertex struct{
	X int
	Y int
}

func structsExample(){
	v:=Vertex{1,2}
	p:=&v
	p.X=1e9
	fmt.Println(v)
}

var(
	v1=Vertex{1,2}
	v2=Vertex{X:1}
	v3=Vertex{}
	p=&Vertex{1,2}
)