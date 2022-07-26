package main

import "fmt"

func main() {
	Go := Course{
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
	css := Course{
		Name:   "CSS desde cero",
		IsFree: true,
	}

	js := Course{}
	js.Name = "Curso JS"
	js.UserIDs = []uint{12, 67}

	Go.PrintClasses()
	Go.ChangePrice(67.12)
	fmt.Println(Go.Price)
	fmt.Println(css.Name)
}
