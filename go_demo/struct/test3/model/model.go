package model

import "fmt"

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
func (s *student) SetAge(age int) {
	if age > 10 && age < 30 {
		s.age = age
	} else {
		fmt.Print("学生太老了")
	}

}

func (s *student) GetAge() int {
	return s.age
}

func (s *student) SetScore(score float64) {
	if score > 0 && score < 100 {
		s.score = score
	} else {
		fmt.Print("你的垃圾分数填错了")
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
