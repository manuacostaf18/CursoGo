package main 

import (
	"fmt"
	"strings"
)


// Hello word 
/*func main(){
	fmt.Println("hello word")
}*/

// Variables 
/*func main(){
	var m1 int
	m1 = 1
	fmt.Println(m1)
}*/
/*func main(){
	var (
		m1 = 1
		m2 = 2
	)
	fmt.Println(m1 + m2)
}*/
/*func main(){
	var m1 int32
	var m2 int64
	fmt.Println(int64(m1)+m2) // To avoid errors 
}*/
/*func main(){ // Different way of declaring variables 
	m1 := 1
	m2 := 2
	fmt.Println(m1+m2)
}*/
/*func main(){ // Normal inicialization of string variables 
	var m1 string 
	m1 = "my name"
	fmt.Println(m1)
}*/
/*func main(){ // Inicialization in one line of string variables 
	m1 := "my name" // := only used for declaration, not for assigment 
	fmt.Println(m1)
}*/
/*func main(){
	m1 := "my name"
	m1 = "hello"
	fmt.Println(m1)
}*/
/*func main(){ // Checking if one string contains the other 
	m1 := "my name"
	m2 := "name"
	fmt.Println(strings.Contains(m1, m2))
}*/
/*func main(){
	m1 := "my name"
	fmt.Println(strings.ReplaceAll(m1, "m", "NO"))
}*/
func main(){
	m1 := "my name"
	fmt.Println(strings.Split(m1, " "))
}

