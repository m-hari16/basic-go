package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello go lang")
	println("cek")

	// numeris type
	var duit int16
	duit = 9999
	println(duit)

	// change type
	duitgede := int8(duit)
	fmt.Println(duitgede)

	// boolean type
	var is_active bool
	is_active = true
	println(is_active)

	// string type
	var snack string
	snack = "kriptos"
	println(snack)

	var name = "Hari"
	fmt.Println(name)
	println(len(name))
	print(name[0])
	print(" ", string(name[0]), " ")

	// constant, can not be reassign
	const gender string = "male"
	print(gender)
	println("")

	// declaration alias tipe data
	type dataKTP string
	type status bool
	var no_type dataKTP
	no_type = "3525032707990001"
	is_merried := false
	println(no_type, is_merried)

	// array type
	var employee [3]dataKTP
	employee[0] = "35348719477420183"
	employee[1] = "23487974981748701"
	employee[2] = "09802173987308209"
	println(employee[1])

	var position = [...]string{
		"mobile apps developer",
		"web developer",
		"Backend developer",
		"UI",
		"UX",
	}
	fmt.Println(position)
	fmt.Println(len(position))

	// slice
	var slice1 = position[1:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1 = append(slice1, "PM")
	slice1[1] = "BE"
	println(slice1[1])
	println(slice1[3])
	fmt.Println(position)

	// make slice
	makeSlice := make([]string, 2, 4)
	makeSlice[0] = "Titin"
	makeSlice[1] = "Totti"
	fmt.Println("capacity: ", cap(makeSlice), "length: ", len(makeSlice))
	fmt.Println(makeSlice)

	// copy slice
	copySlice := make([]string, len(makeSlice), cap(makeSlice))
	copy(copySlice, makeSlice)
	fmt.Println(copySlice)

	// map
	person := map[string]int{
		"no_1": 1,
		"no_2": 2,
	}
	person["no_3"] = 3
	println(person["no_3"])

	manusia := make(map[int16]string)
	manusia[1] = "ini 1"
	manusia[2] = "ini 2"
	println(manusia[2], "panjang map adalah", len(manusia))

	delete(manusia, 2)
	println(manusia[2], "panjang map adalah", len(manusia))

}
