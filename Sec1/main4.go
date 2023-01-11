package main
import "fmt"

// Interfaces 

// When not knowing the type of a structure we can define it as an interface (works like a black box)
func Anything(anything interface{}){
	fmt.Println(anything)
}

func main(){
	fmt.Println("vim.go")
	Anything(2.44)
	Anything("Manuela")
	Anything(struct{}{})

	//struct{} {} //empty structure

	mymap := make(map[string]interface{}) // the key of the map will be a string, and the value can be anything
	mymap["name"] = "Manuela"
	mymap["age"] = 10
	fmt.Println(mymap)
}