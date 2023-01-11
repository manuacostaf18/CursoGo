package main

import "fmt"

// Arrays 

/*func todo(){
	// Creating arrays
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}

	// Appending things to the array
	arr2 = append(arr2, "is", "Manuela")

	fmt.Println(arr, arr2)
}

func main(){
	todo()
}*/

// Pointers 

/*func main(){
	m1 := 2
	ptr := &m1 // Address of m1
	// fmt.Println(ptr)
	fmt.Println(*ptr) // De referencing the pointer 
}*/

/*func swap(m1, m2 *int){
	var temp int 
	temp = *m2
	*m2 = *m1
	*m1 = temp 
}

func main(){
	m1, m2 := 2, 3
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)
}*/

// Structures: datatype that groups together logically related data (data encapsulation)

type Car struct {
	Name 	string
	Age 	int
	ModelNo int
}

func (c Car) Print(){
	fmt.Println(c)
}

func (c Car) Drive(){
	fmt.Println("Driving...")
}

func (c Car) GetName() string{
	return c.Name
}

func main(){
	c := Car{
		Name: "chevy",
		Age: 1,
		ModelNo: 2,
	}
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}