package variable

import "fmt"

// manifest typing
func Langs() {
	var lang1 string = "Go"
	var lang2 string
	lang2 = "Javascript"

	fmt.Println("==> Manifest Typing Example <==")
	fmt.Printf("Hi, My Name is Angga \nI'm learning %s and %s programming language!\n", lang1, lang2)
	fmt.Printf("My current skills : %s, %s!\n", lang1, lang2)
}

// type inference
func Framework() (string, string) {
	goFramework := "echo"
	var jsFramework = "react.js"
	goFramework = "gin gonic"

	// fmt.Printf("And also I study %s and %s!\n", goFramework, jsFramework)
	return goFramework, jsFramework
}

// reserved variable
func LearnPath() string {
	_ = "Advertising"
	_ = "Misthic"
	learn, _ := "Science & Techs", "Politics"
	return learn
}

func Database(database string) *string {
	db := new(string)
	*db = database
	return db
}