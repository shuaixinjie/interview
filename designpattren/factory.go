package designpattren

import "fmt"

// 工厂模式

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method")
}

type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("Inside Square::draw() method")
}

type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Inside Circle::draw() method")
}

type ShapeFactory struct {
}

func (f *ShapeFactory) NewShape(shapeType string) Shape {
	if shapeType == "CIRCLE" {
		return &Circle{}
	} else if shapeType == "RECTANGLE" {
		return &Rectangle{}
	} else if shapeType == "SQUARE" {
		return &Square{}
	}
	return nil
}

// 抽象工厂模式

type Worker interface {
	Work(task string)
}

type WorkerCreator interface {
	Create() Worker
}

type Programmer struct {
}

func (p *Programmer) Work(task string) {
	fmt.Println("Programmer process", task)
}

type ProgrammerCreator struct {
}

func (c *ProgrammerCreator) Create() Worker {
	s := new(Programmer)
	return s
}

type Farmer struct {
}

func (f *Farmer) Work(task string) {
	fmt.Println("Farmer process", task)
}

type FarmerCreator struct {
}

func (c *FarmerCreator) Create() Worker {
	s := new(Farmer)
	return s
}
