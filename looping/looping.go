package main

func main() {
	// Looping
	for i := 1; i <= 5; i++ {
		println("perulangan ke :", i)
	}

	slice := make([]string, 2, 3)
	slice[0] = "Halo"
	slice[1] = "Gaes"

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	// Using for range to access collaction slice
	for i, value := range slice {
		println("index ke: ", i, "adalah", value)
	}

	for _, val := range slice {
		println(val)
	}

	// Using for range to access collection map
	number := make(map[string]int)
	number["no_1"] = 1
	number["no_2"] = 2

	for key, val := range number {
		println(key, "=", val)
	}

	for _, val := range number {
		println(val)
	}

	// Using break and/or continue in looping
	for i := 0; i < 8; i++ {
		if i < 5 {
			continue
		}
		println("perulangan ke:", i)
	}

	for i := 2; i <= 11; i++ {
		if i == 7 {
			break
		}
		println("perulangan ke: ", i)
	}
}
