package main

import "fmt"

type car struct {
	Name string
	Year int
}

func Print(c car, i int) int{
	i_added := i + 1
	
	return i_added
}
func main(){
	c := car{
		Name:"Alto", 
		Year: 2011,
	}

	//c.Print() 
	i_added := Print(c, 100)
	fmt.Println(c, i_added)
}