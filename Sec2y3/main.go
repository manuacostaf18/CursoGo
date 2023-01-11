package main 
import (
	"fmt"
	//"time"
	"sync"
)

// Go routines - it is not a function, it is the way we call a function

/*func heavy(){
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("heavy")
	}
}

func superHeavy(){
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("super heavy")
	}
}

func main(){
	go heavy() // Calling the function as a Go routine (función asincrónica). It keeps it in the background
	go superHeavy()
	fmt.Println("Fin")
	//time.Sleep(time.Second * 5) // The go routine runs for 5 additional seconds 
	select {} // The function main waits indefenitely 
}*/

// Anonymous or Lambda functions 

/*func main(){
	go func (){
		fmt.Println("hello")
	}() // () makes this function to be excecuted, these functions are defined because we do not need them after their excecution
	go func(){
		fmt.Println("world")
	}()
	fmt.Println("Fin")
}*/
// If we excecute the lambda functions as go routines, it does not excecute.

// Wait groups

func main(){
	wg := &sync.WaitGroup{}

	// Add

	wg.Add(2) // Adds to functions in the main group 
	go func (){
		fmt.Println("hello")
		wg.Done()
	}() 
	go func(){
		fmt.Println("world")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Fin")
	select {}
}