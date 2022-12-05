package main

import "fmt"
// import "math"
// import "math/big"

//  TODO use math/big for integer exponentiation
//  TODO add tests



func main() {
  var n int
  fmt.Print("Enter an integer n: ")
  fmt.Scan(&n)
  fmt.Println("Your number is:", n)
  fmt.Println(aks(n))
}

func aks(n int) string {
  //1.
    return "prime"
}



func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}