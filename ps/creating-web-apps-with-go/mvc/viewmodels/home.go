package viewmodels

import (

)

type Home struct {
	Title string
	Active string
}

//helper fuction that returns Home viewmodel
func GetHome() Home {
	result := Home {
		Title: "Lemonade Stand Society - dk",
		Active: "Home",
	}
	return result
}
