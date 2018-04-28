package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName string
	id int
}

func (p *Person) person(first_name string, last_name string, identification int) {
	p.firstName = first_name
	p.lastName = last_name
	p.id = identification
}

func (p *Person) printPerson() {
	fmt.Printf("Name: %s, %s\n", p.lastName, p.firstName)
	fmt.Printf("ID: %d\n", p.id)
}

type Student struct {
	Person

	testScores []int
} 

func (s *Student) calculate() {
	var sum int = 0
	var grade string
    for i := 0; i < len(s.testScores); i++ {
    	sum += s.testScores[i]
    }
        
    average_score := sum / len(s.testScores);
    if average_score >= 90 && average_score <= 100 {
        grade = "O"
    }else if average_score >= 80 && average_score < 90 {
        grade = "E"
    }else if average_score >= 70 && average_score < 80 {
        grade = "A"
    }else if average_score >= 55 && average_score < 70 {
        grade = "P"
    }else if average_score >= 40 && average_score < 55 {
        grade = "D"
    }else if average_score < 40 {
        grade = "T"
    }

    fmt.Println("Grade: ", grade)
}

func main() {
	var first_name, last_name string
	var id, num int

	fmt.Scan(&first_name)
	fmt.Scan(&last_name)
	fmt.Scan(&id)
    fmt.Scan(&num)

    scores := make([]int, num)
    for i := 0; i < num; i++ {
    	fmt.Scan(&scores[i])
    }

    s := Student{Person{first_name, last_name, id}, scores}
	s.printPerson()
	s.calculate()
}