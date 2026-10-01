package main
import "fmt"

type Vert struct{
	Lat,
	Long float64  
}

var m map[string]Vert

func mapExample(){
	m=make(map[string]Vert)
	m["Bell Labs"]=Vert{
		40.68433,-74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
