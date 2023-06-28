package main
import ("fmt")
const (
	first = iota
	cat = 1 << (10 * iota)
	dog
	horse
)
func main(){
	const specialisttype int = cat
	fmt.Println(horse)
}