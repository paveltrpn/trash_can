package main

import (
	"fmt"
	"io"
	"os"
)

var (
	db map[string]*Student
)

type Student struct {
	name  string
	age   int
	grade int
}

func newStudent() *Student {
	return new(Student)
}

func (std *Student) getName() string {
	return std.name
}

func (std *Student) getAge() int {
	return std.age
}

func (std *Student) getGrade() int {
	return std.grade
}

func (std *Student) putName(name string) {
	std.name = name
}

func (std *Student) putAge(age int) {
	std.age = age
}

func (std *Student) putGrade(grade int) {
	std.grade = grade
}

func main() {
	var (
		inName  string
		inAge   int
		inGrade int
	)

	db = make(map[string]*Student)

	for {
		_, err := fmt.Scanf("%s %d %d", &inName, &inAge, &inGrade)
		if err == io.EOF {
			fmt.Println("students from storage:")
			for _, s := range db {
				fmt.Println(s.getName(), s.getAge(), s.getGrade())
			}
			os.Exit(0)
		}

		ns := newStudent()
		ns.putName(inName)
		ns.putAge(inAge)
		ns.putGrade(inGrade)

		db[ns.name] = ns
	}
}
