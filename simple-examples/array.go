package main

import "fmt"

func simple(){
 // var arr []int

	int_arr := []int{1,2,3}
	str_arr := []string{"a", "b"}

	//append_arr  := []string{"appended", "something"}
	final_arr := append(str_arr, "99")
	fmt.Println(int_arr)
	fmt.Println(final_arr)
}

func main(){
	simple()
}