package main

func main() {

	// conditional statement If - Else If - Else
	var name = string("Hari")

	if name == "hari" {
		println("Halo hari")
	} else {
		println("Hae guys")
	}

	if num := 15; num%2 == 0 {
		println("genap")
	} else if num%2 == 1 {
		println("ganjil")
	} else {
		println("Jangan input bilangan 0 dan negatif dulu ya")
	}

	// conditional statement Switch
	username := "Hari"

	switch {
	case len(username) < 3:
		println("username terlalu pendek")
	case len(username) > 6:
		println("nama terlalu panjanmg")
	default:
		println("sesuai")
	}

}
