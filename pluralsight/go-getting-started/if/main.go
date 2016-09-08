package main


func main() {
	//if foo == 1 {
	//Moving initializer into if block
	if foo := 2; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

}
