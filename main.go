package main

import "fmt"

func main(){
	str1 := "睿"
	str2 := "135****2676666666666666睿"
	var sim float64
	SimilarText(str1,str2,&sim)
	fmt.Println(sim)
}

