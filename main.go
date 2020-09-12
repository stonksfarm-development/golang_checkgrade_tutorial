package main

//golang language
import (
	"fmt"
)

//defining function checkgrade by using score as float and return as string
func checkGrade(score float64) string{
	var grade string
	if score > 100{
		grade = "error"
	}
	if score >=80{
		grade = "A"
	}
	if score >=70 && score <80{
		grade = "B"
	}
	if score >=60 && score <70{
		grade = "C"
	}
	if score >=50 && score<60{
		grade = "D"
	}
	if score<50{
		grade = "E"
	}
	return grade
}

//main function program gonna run this firstly
func main(){
	fmt.Println("hello world")
	grade := checkGrade(80.0)
	fmt.Println(grade) //output is A

}