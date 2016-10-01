package greeting

import "fmt"

type Salutation struct {
	Name string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {
		for _, s := range salutation {
			message, alternate := CreateMessage(s.Name, s.Greeting)

			if prefix := GetPrefix(s.Name); isFormal {
				do(prefix + message)
			} else {
				do(alternate)
			}
		}
}

func GetPrefix(name string) (prefix string) {

	//declare a map called prefixMap which is of type map which has a key and a value of type string.
	//var prefixMap map[string]string
	//initialize the map
	//prefixMap = make(map[string]string)

	//cleaning up map to do all
	prefixMap := map[string]string {
		"Bob" : "Mr. ",
		"Joe" : "Dr. ",
		"Amy" : "Dr. ",
		"Mary" : "Mrs. ",
	}

	//Assign new value for map value
	prefixMap["Joe"] = "Jr. "

	delete(prefixMap, "Mary")

	/*
	prefixMap["Bob"] = "Mr. "
	prefixMap["Joe"] = "Dr. "
	prefixMap["Amy"] = "Dr. "
	prefixMap["Mary"] = "Mrs. "
	*/

	return prefixMap[name]


//	switch name {
/*
	switch {
		case name == "Bob": 
			prefix = "Mr "
		case name == "Joe", name == "Amy", len(name) == 10:
			if name == "Amy" {
				prefix = "Boss woman "
			} else {
				prefix = "Boss man "
			}
		case name == "Mary": 
			prefix = "Mrs "
		default: 
			prefix = "Sup Dude "
	}
	return
	*/
}

//empty interface
func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	//switch t := x.(type) {
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float")
		case string:
			fmt.Println("string")
		case Salutation:
			fmt.Println("salutation")
		default:
			fmt.Println("unknown")
		}
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey!" + " " + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func PrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
