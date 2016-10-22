package main

func main() {
	/**
	for i := 0; i < 5; i++ {
		println(i)
	}
	**/
	//i := 0

	//infinite loop **
	/*/*

	for {
		i++
		println(i)

		if i > 100 {
			break
		}
	}
	**/

	//looping through slice.
	/**
	s := []string{"foo", "bar", "buz"}

	for idx, v := range s {
		println(idx, v)
	}
	**/
	//looping through a map
	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "buz"

	for k, v := range m {
		println(k, v)
	}

}
