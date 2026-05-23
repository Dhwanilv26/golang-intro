// package main

// import "fmt"

// type triangle struct {
// 	height float64
// 	base   float64
// }
// type square struct {
// 	sideLength float64
// }

// type shape interface {
// 	getArea() float64
// }

// func main() {

// 	sq := square{45}
// 	t := triangle{4, 5}

// 	printArea(sq)
// 	printArea(t)
// }

// func printArea(s shape) {
// 	fmt.Println("area is", s.getArea())
// }

// func (t triangle) getArea() float64 {
// 	return 0.5 * t.height * t.base
// }

// func (s square) getArea() float64 {
// 	return s.sideLength * s.sideLength
// }

package main

import (
	"fmt"
	"io"
	"os"
)

// first time temp / var binary file banegi, next time cache se lega go sab data

func main() {
	f, err := os.Open(os.Args[1]) // even though f is a file pointer (*File), go auto dereferences the pointers for us, the same case occured in *http.Response too

	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
