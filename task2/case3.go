package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Persion struct {
	Name string
	Age  int
}

type Employee struct {
	Persion
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工ID: %d, 姓名: %s, 年龄: %d\n", e.EmployeeID, e.Name, e.Age)
}


func main_case3() {
	rect := Rectangle{Length: 5, Width: 3}
	circle := Circle{Radius: 2}

	shapes := []Shape{rect, circle}

	for _, shape := range shapes {
		fmt.Println("面积:", shape.Area())
		fmt.Println("周长:", shape.Perimeter())
	}

	employee := Employee{Persion: Persion{Name: "John", Age: 30}, EmployeeID: 123}
	employee.PrintInfo()
}
