package main

import "fmt"

// Interfaces 

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

func NewModel(arg string) Car{
	return &Lambo{arg}
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Stop(){
	fmt.Println("Stopping Lambo")
}

func (l *Lambo) Drive(){
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive(){
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main(){
	l := Lambo{"Gallardo"}
	c := Chevy{"C369"}
	l.Drive()
	c.Drive()
}