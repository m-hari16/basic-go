package main

// simple function with parameter and return value (This parameter varArgs, named Variadic Function)
func sumNum(number ...int) int32 {
	defer logging()
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

// function as parameter
func showComment(word string, filter func(string)string) string{
	filteredWord := filter
	return filteredWord(word)
}
func filterFuck(word string) string{
	if word == "fuck"{
		return "..."
	} else {
		return word
	}
}

// recursive function 
func faktorial(val int)int{
	if val == 1 {
		return 1
	} else{
		return val*faktorial(val - 1)
	}
}

// function for defer and recover
func logging(){
	message := recover()
	if message != nil{
		println("Terjadi Error: ", message)
	}

	println("Ini adalah log aplikasi yang telah di eksekusi")
}

func devide(val1 int, val2 int) float32{
	defer logging()
	if val2 == 0{
		panic("Can Not devide by zero")
	}
	return float32(val1/val2)
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

	myWord1 := showComment("fuck", filterFuck)
	myWord2 := showComment("Hi", filterFuck)
	println(myWord1, myWord2)

	// anonymous function
	filterAsu := func (word string) string {
		if word == "asu"{
			return "..."
		} else{
			return word
		}
	}
	myWord3 := showComment("asu", filterAsu)
	println(myWord3)

	faktorialBilangan := faktorial
	println(faktorialBilangan(3))

	// implement defer, panic and recover
	result:=devide(2,0)
	println(result)
	println("Huh")
}