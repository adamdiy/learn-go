package main

func main() {
	result := add(1, 2, 3)
	println(result)
}

func add(terms ...int)int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}
