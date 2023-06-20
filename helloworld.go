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

func array() {
	// It allows for the storage of multiple values of same type in a single variable. This is in place of declaring separate variable for each value.

  // The syntax for declaring an array is as follows:
  // var array_name [size] data_type
  // While using the := sign it is as follows

  // array_name := [length] datatype{values} - The length specifies the number of values to store in the array. In GO arrays have a fixed length

  var arr1 = [3]int{1,2,3}
  
  //for inferred lengths
  // var arr1 = [...]int{1,2,3}


  arr2 := [5]int{4,5,6,7,8}

  //arr2 := [...]int{4,5,6,7,8} - Length of the array will be inferred during compilation.

  var cars = [4]string{"ford", "BMW", "audi", "benz"}

  fmt.Println(cars)

  // Accessing array elements
  // The elements of an array are accessed using the index of the element. The index of the first element is 0 and the last element is length - 1
  fmt.Print(cars[1])
  
  fmt.Println(arr1)
  fmt.Println(arr2)

  // Array Initialization

  // If an array or one of its elements has not been initialized in the code. IT is assigned the default value of its type - int =0 and string = ""

  arr3 := [6]int{} // not initialized
  arr4 := [4]int{9, 10} //partially initialized
  arr5 := [3]int{113,45,67} //fully initialized

  fmt.Println(arr3)
  fmt.Println(arr4)
  fmt.Println(arr5)

  //Initializing a specific element in teh array
  arr6 := [5]int{1: 10, 2: 20} //arr6[1] = 10, arr6[2] = 20
  //Above means assign 10 to index 1 and 20 to index 2
  fmt.Print(arr6)

  
}

func slice(){
  //slices similar to arrays but more flexible and powerful
  //The length of slices grows and shrinks as one sees fit
  // Ways to make a slice

  // using []datatype{values} format
  // using make() function
  // create slice from an array.

  myslice1 := []int{}

  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1)) //- capacity of the slice the number of elements the slice can grow or shrink to
  fmt.Println(myslice1)

  myslice2 := []int{1,2,3,4,5}
  fmt.Println(len(myslice2))
  fmt.Println(len(myslice2))
  fmt.Println(myslice2)

  // using make() function
  // make([]datatype, length, capacity) - if capacity is not defined it will be equal the length

  myslice9 := make([]int, 5, 10)

  fmt.Printf("myslice9 = %v\n", myslice9)
  fmt.Printf("length = %d\n", len(myslice9))
  fmt.Printf("capacity = %d\n", cap(myslice9))

  //Appending to slices
  myslice2 = append(myslice2, 80, 90)
  fmt.Println(myslice2)

  //Appending a slice to another slice
  // slice3 = append(slice1, slice2...) - THe "..." is necessary when appending elements of one slice to another

  slice3 := append(myslice2, myslice9...)

  fmt.Printf("my slice: %v", slice3)

  //MEMORY EFFICIENCY IN GOLANG
  //when using slices, go will load all the underlying elements into memory
  // If the array is large and you only need a few elements. It is better to copy those elements using copy() function
  // THe copy function will create a new underlying array with only the required elements for the slice reducing memory used for the program.
  //syntax - copy(dest, src)

  // The copy function will take in two slices the dest and src, and will copy data from the src to the destination and will return the number of elements copied.
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  //Original slice

  fmt.Printf("numbers = %v \n", numbers)
  fmt.Printf("length = %d \n", len(numbers))
  fmt.Printf("Capacity = %d \n", cap(numbers))

  //Creating a new slice using the copy function
  neededNumbers := numbers[: len(numbers) - 10]
  numbersCopy := make([]int, len(neededNumbers))

  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v \n", numbersCopy)
  fmt.Printf("length = %d \n", len(numbersCopy))
  fmt.Printf("numbersCopy = %d \n", cap(numbersCopy))
}

func condition() {
  // if statement must be in lowercase letters else will generate errors
  if (20 > 10) {
    fmt.Println("20 is greater than 10")
  }

  time :=30

  if (time < 20) {
    fmt.Println("Good Morining")
  }else {
    fmt.Println("Good Evening")
  }

  // Nested if statement
  num := 20

  if (num >= 10){
    fmt.Println("Number is more than 10")

    if (num > 15){
      fmt.Println("Number is more than 15")
    }
  }else{
    fmt.Println("Number is less than 10")
  }


  // switch case - selects one of the many code blocks to be executed.
  // The switch statement executes case by case until a match is found.

  // How switch case works
  // - The expression is evaluated once
  // - The value of the switch expression is compared with the values of each case.
  // - If there is a match the associated block of code is executed
  // - The default keyword is optional and it specifies the code to be run if there is no case match

  day := 4

  switch day {
    case 1:
      fmt.Println("Monday")
    case 2:
      fmt.Println("Tuesday")
    case 3:
      fmt.Println("Wednesday")
    case 4:
      fmt.Println("Thursday")
    case 5:
      fmt.Println("Friday")
    case 6:
      fmt.Println("Saturday")
    case 7:
      fmt.Println("Sunday")
    default:
      fmt.Println("Invalid day")
  }

  // switch case with multiple expressions
  // The switch case can have multiple expressions separated by comma

}

func loop(){
  //For loop is the only available loop in go.
  // It loops through a block of code a specified number of times.
  // Syntax for statement1; statement2; statement3 {}
  // statement1 is executed before the loop starts - initialized the loop counter value
  // statement2 defines the condition for executing the loop - Evaluated for each loop iteratiob. If it evaluates true the loop continues if it is false the loop ends.
  // statement3 is executed at the end of each iteration - increments the loop counter value

  for i := 0; i<=5; i++{

    // The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
    if i == 3{
      //This code will skip 3 and continue with the iteration
      continue
    }

    // The break statement is used to terminate the loop or switch statement and transfer execution to the statement immediately following the loop or switch.
    if i == 4{
      break
    }

    fmt.Println(i)
  }

  // Range key word is used to iterate over an array, slice or map. It returns both the index and value

  // syntax for index, value :=array|slice|map{}

  fruits := [3]string{"banana", "apple", "orange"}
  for idx, val:= range fruits{
    fmt.Printf("fruits[%d] = %s \n", idx, val)
  }

  // for ommission of either index or value you replace the idx or val with an underscore.
}

// function - Any code inside curly braces is executed
func main() {

	practice()
  array()
  slice()
  condition()
  loop()

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



