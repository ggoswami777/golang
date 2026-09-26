package main
import "fmt"

func slicesExample(){
	names:=[4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a:=names[0:2]
	b:=names[1:3]
	fmt.Println(a,b)
	b[0]="XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}
  
func slicesExample2(){
	q:=[]int{2,3,5,7,11,13}
	fmt.Println(q)

	r:=[]bool{true,false,true,true,false,true}

	s:=[]struct{
		i int
		b bool
	}{
		{2,true},
		{3,false},
		{5,true},
		{7,true},
		{11,false},
		{13,true},
	}
	fmt.Println(r,s)
	
}