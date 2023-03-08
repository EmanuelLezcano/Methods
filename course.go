package methods

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func NewCourse(name string, price float64, isFree bool) *Course {
	return &Course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}
func PrintText() {
	fmt.Println("Hi methods.")
}

//Method of struct Course
func (c *Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text)
}

//Cambia la propiedad Price de un struct Course
func (c *Course) changePrice() {
	c.Price = 12.50
}
