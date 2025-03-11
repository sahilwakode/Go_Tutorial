package main
import "fmt"
// package variables
// here L of Logintoken is capital to denote it is public:
const Logintoken string = "kwdjw2d"
func main() {
	var username string = "sahil"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	//default values and some alias
	var anothervar int
	fmt.Println(anothervar)
	fmt.Printf("Variable is of type: %T\n", anothervar)

	var anothervar2 string
	fmt.Println(anothervar2)
	fmt.Printf("Variable is of type: %T\n", anothervar2)

	//implicit type directly put it
	var take = "sahil"
	fmt.Println(take)

	takeme := "bit"
	fmt.Println(takeme)

	fmt.Println(Logintoken)
}
