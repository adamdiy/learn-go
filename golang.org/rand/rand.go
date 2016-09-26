package main

import ( 
	"fmt"
	"math/rand"
)


func main() {
	rand.Seed(42)
	answers := []string {
		"It's certain",
		"It's decidedly so",
		"Without doubt",
		"Definitely",
		"Rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])


}
