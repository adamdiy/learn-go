package greeting

import "fmt"

type Salutation struct {
	Name string
	Greeting string
}

type Renameable interface {
	Rename(newName string)
}

//writer interface
func(salutation *Salutation)Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}


func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

//new type that is a slice of salutations
type Salutations []Salutation
type Printer func(string) ()

//This is so weird.
//turn function into a method
func (salutations Salutations)Greet(do Printer, isFormal bool, times int) {
	for _, s := range salutations {
		message, alternate := CreateMessage(s.Name, s.Greeting)

		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func (salutations Salutations) ChannelGreeter(c chan Salutation) {
	for _, s := range salutations {
		c <- s
	}
	close(c)
}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string {
		"Bob" : "Mr. ",
		"Joe" : "Dr. ",
		"Amy" : "Dr. ",
		"Mary" : "Mrs. ",
	}

	prefixMap["Joe"] = "Jr. "
	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
}

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
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
