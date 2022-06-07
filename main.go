package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// manifest typing
	var lang1 string = "Go"
	var lang2 string
	lang2 = "Javascript"

	fmt.Printf("Hi, My Name is Angga, I'm learning %s and %s programming language!\n", lang1, lang2)
	fmt.Println("My current skills :", lang1, lang2 + "!")

	// type inference
	goFramework := "echo"
	var jsFramework = "react.js"
	goFramework = "gin gonic"

	fmt.Printf("And also I study %s and %s!\n", goFramework, jsFramework)

	/* multi variable declaration
	variable declarations can also use type interfaces */
	var frontend, backend *string
	frontend, backend = &jsFramework, &goFramework

	// reserved variable ( _ )
	_ = "Belajar Bisnis"
	_ = "Belajar Marketing"
	platform, _ := "web", "mobile"

	fmt.Println("I've been working on a project using", *frontend, " and ", *backend, "for a", platform)

	// variable declaration using keyword 'new'
	database := new(string)
	*database = "MySQL"

	fmt.Printf("Variable declaration with new (addres) : %p\n", database)
	fmt.Printf("Variable declaration with new (value) : %s\n", *database)

	/* data type */
	// => non-decimal numeric
	var positiveNumber uint8 = 24
	var negativeNumber int16 = -32768
	// overflow value
	/* var positiveNumber2 uint8 = 256 // message error 
	positiveNumber2 = 255 */

	fmt.Printf("%d is a positive number\n", positiveNumber)
	fmt.Printf("%d is a negative number\n", negativeNumber)

	// => decimal numeric
	var rad = 2.34
	var decimalNum = 3.14
	fmt.Printf("area of circle where the radius is %.2f = %.2f\n", rad, decimalNum*rad*rad)

	// => boolean
	var loading bool = true
	fmt.Printf("loading... %t \n", loading)

	// => string another way to declaration string data type
	var message = `Package fmt implements formatted I/O with functions analogous to C's printf and scanf. 
	The format 'verbs' are derived from C's but are simpler.`
	fmt.Println(message)

	// constant
	const pi = 3.14
	fmt.Println(pi)

	// arithmetics operators 
	var circleCircumference = 3.14*(7*2)
	fmt.Println(circleCircumference)

	// comparison operators
	var val = (((45 - 20) + 5) * 20 - 50) / 20
	var isEqual = val == 30
	fmt.Printf("value of %d (%t) \n", val, isEqual)

	// logical operators
	var isLoading = false
	var length = 4
	var isEmpty = !isLoading && length > 0

	/* conditional statement */
	// if else
	if isEmpty {
		fmt.Printf("done, length = %d\n", length)
	} else {
		fmt.Printf("Please wait... loading is %t, length = %d", isLoading, length)
	}

	// variable temporary if else
	var point = 8239.5

	if grade := point / 100 ; grade >= 100 {
		fmt.Printf("%.1f%s perfect!\n", grade, "%")
	} else if grade >= 70 {
		fmt.Printf("%.1f%s good!\n", grade, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", grade, "%")
	}

	// switch case
	var value = 6

	switch {
	case value >= 8:
		fmt.Println("Perfect")
	case ( value < 8 ) && ( value >= 6 ):
		fmt.Println("Awesome")
	default:
		fmt.Println("You need to learn more")
	}
}