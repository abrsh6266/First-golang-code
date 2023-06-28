package main
import ("fmt")
const (
	first = iota
	cat
	dog
	horse
)
func main(){
	const specialisttype int = cat
	fmt.Println(specialisttype==cat)
}