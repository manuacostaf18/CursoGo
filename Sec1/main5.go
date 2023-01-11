package main
import "fmt"

// Control structures 

func main(){
	fmt.Println("Hello world")
	// Control structures: if else, for, switch case, break continue

	// If - else conditions
	f := true
	flag := &f // It is better to have everything in terms of pointers 

	if flag == nil{ // For two conditions we use cond1 && cond2 
		fmt.Println("Value is nil")
	} else if *flag{
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// For loops
	/*for i := 0; i < 10; i ++{
		fmt.Println(i)	
	}*/

	arr := []string{"Manuela", "Natalia", "Sara"}
	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "Manuela"
	mymap["age"] = 22
	for k, v := range mymap{
		fmt.Printf("Key: %s and Value: %v", k, v) // format printing
	}
}