package main

import (
	abstractfactory "dsa-golang/Chapter1/CreationPatter/factory/abstractFactory"
	"fmt"
)

/*
"A factory defined by a function that helps us to create instances of a certain structure with determinate values or values that could be provided in the function arguments". SO from SOLID can be achieved with this pattern
*/

type DataParser interface {
	ToString() string
}
type Robot struct {
	Name     string
	Kind     string
	Autonomy float32
}

func (r *Robot) ToString() string {
	return fmt.Sprintf("Name:%s\nKind: %s,\nAutonomy: %.2f", r.Name, r.Kind, r.Autonomy)
}

//factory
func newRobot(name string, kind string, autonomy float32) *Robot {
	return &Robot{
		Name:     name,
		Kind:     kind,
		Autonomy: autonomy,
	}
}

func main() {

	teachRobot := newRobot("Teacher Robot One", "Teacher Bot", 90)
	studentRobot := newRobot("Student Robot One", "Student Bot", 70)
	fmt.Println(teachRobot.ToString())
	fmt.Println(studentRobot.ToString())

	fmt.Println()
	slimShady2_0, _ := abstractfactory.CreateUser("admin")
	ali2_0, _ := abstractfactory.CreateUser("regular")

	fmt.Println(slimShady2_0.WhoAmI())
	fmt.Println(ali2_0.WhoAmI())

}
