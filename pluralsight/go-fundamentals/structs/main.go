package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//This gives us a pointer
	//DockerDeepDive := new(courseMeta)

	DockerDeepDive := courseMeta {
		Author: "dky",
		Level: "Beginner",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println(DockerDeepDive.Author)
	fmt.Println(DockerDeepDive.Level)
	fmt.Println(DockerDeepDive.Rating)
}

