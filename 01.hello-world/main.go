package main

import "fmt"

func main() {
	// manifest typing
	var lang1 string = "Go"
	var lang2 string
	lang2 = "Javascript"

	// type interface
	goFramework := "echo"
	var jsFramework = "react.js"
	goFramework = "gin gonic"

	/* multi variable declaration
	variable declarations can also use type interfaces */
	var frontend, backend *string
	frontend, backend = &jsFramework, &goFramework

	fmt.Println("Hello World")
	fmt.Printf("Hi, My Name is Angga, I'm learning %s and %s programming language!\n", lang1, lang2)
	fmt.Println("My current skills :", lang1, lang2 + "!")
	fmt.Printf("And also I study %s and %s!\n", goFramework, jsFramework)
	fmt.Println("I've been working on a project using", *frontend, " and ", *backend)
}