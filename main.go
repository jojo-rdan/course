package main

import (
	"fmt"

	"github.com/jojo-rdan/course/course"
)

func main() {
	Go := course.New("Go desde cero", 12.34, false)

	Go.UserIDs = []uint{12, 56, 898}
	Go.Classes = map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	}

	// Go := &course.Course{
	// 	Name:    "Go desde cero",
	// 	Price:   12.34,
	// 	IsFree:  false,
	// 	UserIDs: []uint{12, 56, 898},
	// 	Classes: map[uint]string{
	// 		1: "Introducción",
	// 		2: "Estructuras",
	// 		3: "Maps",
	// 	},
	// }

	//Go.PrintClasses()
	fmt.Printf("%+v", Go)
}
