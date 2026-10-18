package main

import "fmt"

func IndexInt(s []int,x int)int{
	for i,v:=range s{
		if(v==x){
			return i
		}
	}
	return -1
}


func IndexString(s []string,x string)int{
	for i,v:=range s{
		if(v==x){
			return i
		}
	}
	return -1
}

func genericsEG(){
	 si:=[]int{10,20,15,-10}
	 fmt.Println(IndexInt(si,15))

	 ss:=[]string{"foo","bar","baz"}
	 fmt.Println(IndexString(ss,"hello"))  
	//  output : -1
}
//  taking type as parameter
func Index[T comparable](s []T,x T)int{
	for i,v:=range s{
		if v==x{
			return i
		}
	}
	return -1
}