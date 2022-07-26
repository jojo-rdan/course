package main

import (
	"github.com/jojo-rdan/course/course"
)

func main() {
	Go := &course.Course{
		"Go desde cero",
		12.34,
		false,
		[]uint{12, 56, 898},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()
}
