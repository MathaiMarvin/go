// Go syntax

// It constitutes of the following:
// Package declaration
// Import packages
// function
//Statements and expressions

// In go any code that is executable belongs to the main package

// Package declaration - Every program in go is part of a package
// The program below belongs to the main package
package main

// Import packages
//The import below allows us to import files that are included in the fmt package
import ("fmt")

func practice() {
	// Data type specifies the size and type of variable values.
	// Given that Go is statitaclly typed it means that once a variable type is defined, it can only store data of that type.

	// Go has the following data types:
	// bool - either truo or false
	//Numeric - rep either integer, floats or complex types
	// String rep  string value.

	var a bool = true
	var b string= "David"
	var c float32 = 3.14
	var d int = 9990

	fmt.Println("Boolean: ", a)
	fmt.Println("String: ", b)
	fmt.Println("Float: ", c)
	fmt.Println("Integer: ", d)


	var x int = 600
	var y int = -80000

	fmt.Printf("Type: %T, Value: %v", x, x)
	fmt.Printf("Type: %T, Value: %v", y, y)

	// unisgned integers declared with the keyword uint store only non-negative values
	var n uint = 800
	fmt.Printf("Type: %T, Value: %v", n, n)

}

// function - Any code inside curly braces is executed
func main() {

	practice()

	//statements and expressions
	//fmt.Println("Hello World!") - A function that is made available from the fmt package.
	//It allows for output or printing out text. 
  fmt.Println("Hello World!")




  //Declaration of variables we can use the var keyword or :=

  var student1 string = "Joseph" //type is string
  var student2 = "Mary" //type is inferred
  x := 1 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)



  // While declaring a avariable without an initial value the type is declared to set the intial values as the type
  var a string
  var b bool
  var c int
  var d float32


  a = "Jeff"
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)

  // Its not possible to declare a variable using := without assigning a value to it and is only used inside functions the var can also be used outside functions


  var i, j, z, s int = 1, 2, 3, 4

  //By including the type keyword, then we are able to declare one type of variable per line.
  fmt.Println(i)
  fmt.Println(j)
  fmt.Println(z)
  fmt.Println(s)

  var m, n = "quincy", 5
  t, u :=  9, "world!"


  fmt.Println(m)
  fmt.Println(n)
  fmt.Println(t)
  fmt.Println(u)

  //DEclaration of variables as a block is preffered fro readability

  var (
	kim int 
	marv int = 100
	gaucho string = "name"
  )
  
  fmt.Println(kim)
  fmt.Println(marv)
  fmt.Println(gaucho)


  // For constants in go the names are usually written in uppercase letters for easy identification and differentiation from variables.
  // typed constants contain a type

  const K int = 70
  const L = "Cancun"


  fmt.Print(K)
  fmt.Print(L, "\n")

  // There are 3 functions to output text 
  //Print() - prints its arguments with their default format: incase you want to print arguments in a new linewe need to use \n
  //Println() - similar to Print() with the difference is that the whitespace is added between the arguments and a new line added at the end
  //Printf() - This function first formats its argument based on the given formatting verb and then prints them

  // For instance we can use the following formatting verbs:
  //%v is used to print the value of the arguments
  //%T is used to print the type of the arguments


  var kick string = "hello"
  var num int = 800

  fmt.Printf("kick has value: %v and type: %T \n", kick, kick)
  fmt.Printf("num has value: %v and type: %T \n", num, kick)
}



