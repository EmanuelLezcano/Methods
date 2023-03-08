package methods

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	classes map[uint]string
}

func (c *course) CourseSetClasses(classes map[uint]string) {
	c.classes = classes
}

func NewCourse(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}
func PrintText() {
	fmt.Println("Hi methods.")
}

//Method of struct Course
func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.classes {
		text += class + ", "
	}
	fmt.Println(text)
}

//Cambia la propiedad Price de un struct Course
func (c *course) changePrice() {
	c.Price = 12.50
}
