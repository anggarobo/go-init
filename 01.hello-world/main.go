package main

import "fmt"

func main() {
	// manifest typing
	var lang1 string = "Go"
	var lang2 string
	lang2 = "Javascript"

	// type interface
	goFramework := "gin gonic"
	var jsFramework = "react.js"

	fmt.Println("Hello World")
	fmt.Printf("Hi, My Name is Angga, I'm learning %s and %s programming language!\n", lang1, lang2)
	fmt.Println("My current skills:", lang1, lang2 + "!")
	fmt.Printf("And also I study %s and %s!\n", goFramework, jsFramework)
}