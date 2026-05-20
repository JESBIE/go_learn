package user

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewUser(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func (u Person) DisplayInfo() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
}