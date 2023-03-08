package methods

import "fmt"
func NewCourse(name string, price float64, isFree bool) *course{
	return &course{
		Name: name,
		Price: price,
		IsFree: isFree,
	}
}

type course struct {
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	classes map[uint]string
}
func printText(){
	fmt.Println("Hi methods.")
}
//Method of struct Course
func (c *Course) printClasses() {
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
