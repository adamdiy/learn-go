package viewmodels

import (

)

type Home struct {
	Title string
	Active string
}

func GetHome() Home {
	result := Home {
		Title: "Lemonade Stand Society - dk",
		Active: "Home",
	}
	return result
}
