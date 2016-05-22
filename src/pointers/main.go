// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 43

// 	fmt.Println(a)
// 	fmt.Println(&a)

// 	var b *int = &a
// 	fmt.Println(b)
// 	fmt.Println(*b)

// 	*b = 42

// 	fmt.Println(a)
// }

package main

import (
	"fmt"
)

func zero(x int) {
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	x = 0
}
func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
}
