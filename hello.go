//FUNCTIONS IN GO


// package main

// import "fmt"

// func main() {

// 	fmt.Println("main function invoked.")

// 	funFunction()
	
// }   

// func funFunction (){
// 	fmt.Println("fun function invoked :)")// main function invoked.
// 										  //fun function invoked :)
// }
//-------
package main

import "fmt"
func main(){
    fmt.Println("Main Function")
    funFunction()
 
    // addNums()
    // addNums(10)
    addNums(10,20)
    // addNums(10,20,30)
 
}
 
 
func funFunction(){
    fmt.Println("fun function invoked")
}
 
func addNums(num1 int,num2 int){
    fmt.Println(num1+num2)
}



// //TYPE CONVERSION IN GO

// package main

// import "fmt"

// func main() {
// 	bonus:=10.25
// 	salary := 10000
// 	//pay:= salary+bonus//type prob float and int so wont work 'invalid operation' and no implicit type convertion, gotta explicitly do it
	
	
	
// 	pay:= salary + int(bonus)
// 	pay1:= float64(salary)+ bonus
	
// 	fmt.Println(pay)
// 	fmt.Println(pay1)
	

// }   







//DATATYPES IN GO

// package main

// import "fmt"

// func main() {
// 	eid:= 101
// 	firstName :="Vibha"  //no commands ok
// 	salary := 10000
// 	isMarried:= false

// 	fmt.Println(eid)
// 	fmt.Println(firstName)
// 	fmt.Println(salary)
// 	fmt.Println(isMarried)

// }   








//VARIABLES HOW TO TYPE AND STUFF 


// package main

// import "fmt"

// func main(){
// 	var a int = 10 //variable dec TYPE 1
// 	//OR
// 	b:=- 20//TYPE 2


// 	fmt.Println(b)
// 	fmt.Println(a)
// 	fmt.Println("Hello world")

// }   // usually how the go code looks haha hehe eee no using emicolons anywhere, case sens also.