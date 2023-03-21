package main

import "fmt"

func main() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	saiyanList := []Saiyan{}

	saiyanList = append(saiyanList, goku)

	for index, value := range saiyanList {
		Super(&value)
		saiyanList[index] = value
	}
	for _, value := range saiyanList {
		fmt.Println(value)

	}

}
func Super(s *Saiyan) {
	s.Power += 10000
}

func (s *Saiyan) NameChange(name string) {
	s.Name = name
}

type Saiyan struct {
	Name  string
	Power int
}
