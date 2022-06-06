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

	// type interface
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

	/* data type
	=> non-decimal numeric */
	var positiveNumber uint8 = 24
	var negativeNumber int16 = -32768
	// overflow value
	/* var positiveNumber2 uint8 = 256 // message error 
	positiveNumber2 = 255 */

	fmt.Printf("%d is a positive number\n", positiveNumber)
	fmt.Printf("%d is a negative number\n", negativeNumber)

	// => decimal num``eric
	var rad = 2.34
	var decimalNum = 3.14
	fmt.Printf("area of circle where the radius is %.2f = %.2f", rad, decimalNum*rad*rad)
}