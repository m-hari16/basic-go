package main

// simple function with parameter and return value
func sumNum(number ...int) int32 {
	total := 0
	for _, num := range number {
		total += num
	}

	return int32(total)
}

func main() {
	pleaseSum := sumNum(3,2,3,4,5,1)
	println(pleaseSum)
}