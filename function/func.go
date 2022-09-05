package main

// simple function with parameter and return value (This parameter varArgs, named Variadic Function)
func sumNum(number ...int) int32 {
	total := 0
	for _, num := range number {
		total += num
	}

	return int32(total)
}

// simple function with parameter and two return value
func showFullName(fname string, lname string) (string, string){
	return fname, lname
}

// we can use name for return value of function
func no_type(type_identity string, no_identity string)(tipe, no_tipe string){
	tipe = type_identity
	no_tipe = no_identity

	return
}

// function save in variable
func bye(name string) string{
	return "bye "+name
}

func main() {
	pleaseSum := sumNum(3,2,3,4,5,1)
	println(pleaseSum)

	firstName, _ := showFullName("Moh", "Harianto")
	println(firstName)

	tipe_identitas, no_identitas := no_type("ktp", "6736497932430002")
	println(tipe_identitas, no_identitas)

	goodBye := bye
	println(goodBye("Hari"))
}