package main 
import "fmt"


// Unbuffered Channels (size=1)

/*func main() {
	// To declare channels: <name>  chan <datatype> 
	//var c chan int 

	// It is better to make the channel, so it is not empty 
	c := make(chan int)
	 
	// Send: using a go routine 
	go func(){
		c <- 1 // What we want to send needs to be of the same datatype of the channel
	}()

	// Sniff
	val := <-c
	fmt.Println(val)

	fmt.Println(c)

	go func(){
		c <- 2
	}()

	val = <-c
	fmt.Println(val)
}*/


// Buffered Channels (size>1)

/*func main(){
	c := make(chan int, 3)

	go func(){
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}*/

type Car struct{
	Model string
}

func main(){
	c := make(chan *Car, 3)

	go func(){
		c <- &Car{"1"}
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"}
		close(c)
	}()

	for i := range c{
		fmt.Println(i.Model)
	}
}