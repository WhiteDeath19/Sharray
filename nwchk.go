package main 
import "fmt"

var x,y int32
func add(){
	fmt.Println("Enter first number")
	fmt.Scanln(x)
	fmt.Println("Enter second number")
	fmt.Scanln(y)
	z:= x+y
	fmt.Println("The sum is",z)
	return z

func main(){
	fmt.Println("Hello world????")
	add()
}
