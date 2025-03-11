package main
import "fmt"

func main() {
	// no iheritance in golang
	// no super no parent
	sahil := User{"sahil", "shqidj@gmail.com", true, 21}
	fmt.Println(sahil.Age)
	// fmt.Println("%+v\n",sahil)
	fmt.Println(sahil)
	sahil.GetStatus()
	sahil.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// all are capital because we want to access outside
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}
func (u User) NewMail() {
	u.Email = "hello@go"
	fmt.Println(" email:", u.Email)

}
