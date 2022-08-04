package main

import(
	"fmt"
)

type AnimalIF interface{
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct{
	color string
	selftype string
}

func (this Cat) Sleep(){
	fmt.Println("Cat is sleep")
}

func(this Cat) GetColor()string{
	return this.color
}

func(this Cat) GetType()string{
	return "Cat"
}

type Dog struct{
	color string
}

func (this Dog) Sleep(){
	fmt.Println("Dog is sleep")
}

func(this Dog) GetColor()string{
	return this.color
}

func(this Dog) GetType()string{
	return "Dog"
}

func main(){
	var	iface AnimalIF

	iface = &Dog{"Red"}
	iface.Sleep()

}



