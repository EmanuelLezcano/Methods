package methods

import "fmt"

type course struct {
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	classes map[uint]string
}

func (c *course) CourseSetName(name string) {
	c.name = name
}
func (c *course) CourseName() string {
	return c.name
}
func (c *course) CourseSetPrice(price float64) {
	c.price = price
}
func (c *course) CoursePrice() float64 {
	return c.price
}
func (c *course) CourseSetIsFree(isFree bool) {
	c.isFree = isFree
}
func (c *course) CourseIsFree() bool {
	return c.isFree
}
func (c *course) CourseSetUserIDs(userIDs []uint) {
	c.userIDs = userIDs
}
func (c *course) CourseUserIDs() []uint {
	return c.userIDs
}
func (c *course) CourseSetClasses(classes map[uint]string) {
	c.classes = classes
}
func (c *course) CourseClasses() map[uint]string {
	return c.classes
}

func NewCourse(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
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
