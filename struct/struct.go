package main

import "fmt"

type Customer struct {
	name, address string
	age           int8
}

func (customer Customer) order(item string){
	fmt.Println(customer.name, " telah order", item)
}

func main() {
	var customer1 Customer
	customer1.name = "Har"
	customer1.address = "address"
	customer1.age = 23
	println(customer1.name)

	customer2 := Customer{
		name: "Ron",
		address: "addr",
		age: 32,
	}

	customer3 := Customer{
		name: "Za",
		address: "home",
		age: 21,
	}
	println(customer2.address)
	fmt.Println(customer2)
	fmt.Println(customer3)

	customer1.order("Komputer")
	customer2.order("Mouse")
	customer3.order("Monitor")
}