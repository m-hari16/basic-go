package main

// simple function with parameter and return value
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

func main() {
	pleaseSum := sumNum(3,2,3,4,5,1)
	println(pleaseSum)

	firstName, _ := showFullName("Moh", "Harianto")
	println(firstName)
}