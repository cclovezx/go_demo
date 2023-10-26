package model

//import "fmt"

type student struct {
	name  string
	age   int
	score float64
}

func NewStudent(name string, age int, score float64) *student {
	return &student{
		name:  name,
		age:   age,
		score: score,
	}
}

func (s *student) SetName(name string) {
	s.name = name
}

func (s *student) GetName() string {
	return s.name
}
