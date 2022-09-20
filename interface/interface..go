package main

import (
	"errors"
	"fmt"
)

// declare interface and have one func
type HasName interface{
	GetName() string
}

// declare and use empty interface
func Ups(i int16) interface{}{
	if i == 1{
		return 1
	} else if i == 2{
		return false
	} else{
		return "Ups..."
	}
}

func sayHalo(name HasName){
	fmt.Println("Halo, ", name.GetName()) 
}

type Animal struct{
	Name string
}
func (animal Animal) GetName() string{
	return animal.Name
}

type Person struct{
	Name string
}
func (person Person) GetName() string{
	return person.Name
}

// error interface (built in)
func devide(val1 int, val2 int) (int, error){
	if val2 == 0{
		return 0, errors.New("Ga boleh dibagi dengan angka 0!")
	} else{
		return val1/val2, nil
	}
}

func main(){
	var kucing Animal
	kucing.Name = "Tomo"
	sayHalo(kucing)
	fmt.Println(kucing.GetName())

	uiDesigner := Person{Name: "Mario"}
	sayHalo(uiDesigner)

	// using empty interface
	var cekUps interface{}
	cekUps = Ups(3)
	fmt.Println(cekUps)

	// implement errors interface
	result, err := devide(2,0)
	if  err == nil{
		fmt.Println(result)
	} else{
		fmt.Println(err)
	}

	// Ini perubahan



}