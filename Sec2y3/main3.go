package main
import (
	"fmt"
	"time"
	"os"
)

// Select block

func Select(c chan int, quits chan struct{}){
	// Sintax equal to switch case 
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("received")
		case <-quits:
			fmt.Println("quit")
			os.Exit(0)
		}
	}
}

func main(){
	c := make(chan int)
	quits := make(chan struct{})
	go func(){
		c <- 1
		quits <- struct{}{}
	}()
	go Select(c, quits)
	select {} // Switch cases on the channel, as it waits indefenitely
}