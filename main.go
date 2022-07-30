package main

import (
	"fmt"

	"github.com/jojo-rdan/course/course"
)

func main() {
	Go := course.New("Go desde cero", 12.34, false)

	Go.SetUserIDs([]uint{12, 56, 898})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})
	Go.SetPrice(32.14)

	// Esto era cuando no teníamos un constructor
	// para las estructuras
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
	Go.SetName("POO con GO")
	fmt.Println(Go.Name())
	Go.PrintClasses()
	fmt.Printf("%+v", Go)
}
