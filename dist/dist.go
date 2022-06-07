package dist

import (
	"fmt"
	"go-init/dist/conditional"
	forloop "go-init/dist/for-loop"
	"go-init/dist/operator"
	"go-init/dist/variable"
)

func Dist() {
	variable.Langs()

	fmt.Println("==> Type Inference & Multi variable declaration <==")
	goFw, jsFw := variable.Framework()
	fmt.Printf("And also I study %s and %s!\n", goFw, jsFw)
	var frontend, backend *string
	frontend, backend = &jsFw, &goFw

	var learn string = variable.LearnPath()

	fmt.Println("I've been working on a project using", *frontend, " and ", *backend, "for", learn, "App")

	fmt.Println("==> Variable declaration using keyword 'new' <==")
	db := variable.Database("MySQL")
	fmt.Printf("Variable declaration with new (addres) : %p\n", db)
	fmt.Printf("Variable declaration with new (value) : %s\n", *db)

	fmt.Println("Data Types")
	variable.NonDecimal()
	variable.Decimal()
	variable.BooleanType()

	fmt.Println("Arithmatic Operators")
	fmt.Println(operator.AreaOfCircle(30))
	fmt.Println("Comparison Operators")
	fmt.Println(operator.IsEvenNumber(1))
	fmt.Println("Logical Operators")
	fmt.Println(operator.OnProcess(false, 10))

	fmt.Println("Variable temporary if-else")
	fmt.Println(conditional.CheckGrade(8239.5))
	fmt.Println("Switch case")
	fmt.Println(conditional.CheckGrade(7))

	fmt.Println("For loop")
	forloop.LetsCount(10)
	forloop.WhichEven(15)
	forloop.LabelLoop(2)

	// switch case
	// var value = 6

	// for {
	// 	value++
	// 	if value == 10 {
	// 		if value % 2 == 0 && value < 10 {
	// 			fmt.Println("even", value)
	// 		} else {
	// 			fmt.Println("odd", value)
	// 		}
	// 		break
	// 	}
	// }
}