package main

import(
	"fmt"
	"math"
)

type vi struct{
	X,Y float64
}

// methods
// func (v Verex) Abs() float64{
// 	return math.Sqrt(v.X*v.X+v.Y*v.Y)
// }

func abs(v vi
	) float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}
func methodExample(){
	v:=vi{3,4}
	fmt.Println(abs(v))
	// fmt.Println(v.Abs())
}