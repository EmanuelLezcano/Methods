package main

import "fmt"

func main() {
	Go := Course{
		Name:    "Emanuel",
		Price:   15.34,
		IsFree:  true,
		UserIDs: []uint{32, 44, 55},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}
	/*
		Es lo mismo que lo de arriba
		NOTA: es obligatorio especificar los campos en el mismo orden que se declararon.
		Go := Course{
			"Emanuel",
			15.34,
			true,
			[]uint{32, 44, 55},
			map[uint]string{
				1: "Introducción",
				2: "Estructuras",
				3: "Maps",
			},
		}
	*/
	ObjVacio := Course{}
	Incompleto := Course{
		Name:   "Juan",
		IsFree: false,
	}
	fmt.Println(Go.Name)
	fmt.Printf("%v\n", ObjVacio)
	fmt.Printf("%v\n", Incompleto)

	Go.PrintClasses()
}
