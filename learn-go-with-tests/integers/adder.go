package integers

import "fmt"

func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}

func main() {
	Add(2, 2)
}
