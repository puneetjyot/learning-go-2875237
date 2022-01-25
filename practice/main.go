package main

import (
	"fmt"
)

const anotherStirng = "hey"
// fmt.Println(anotherStirng) 

func main() {

	var aString string = "This is go!"
	fmt.Println(aString)
	fmt.Printf("the variable type is %T\n", aString)

	var anInteger int = 6
	fmt.Println(anInteger)
	fmt.Printf("the variable type is %T\n", anInteger)

	myString := "this is another string"
	fmt.Println(anotherStirng)
	fmt.Printf("the variable type is %T\n",myString) 

}
