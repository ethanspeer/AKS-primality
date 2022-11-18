package main

import "fmt"
import "math"



func main() {
  var n int
  fmt.Print("Enter an integer n: ")
  fmt.Scan(&n)
  fmt.Println("Your number is:", n)
  fmt.Println(aks(n))

}

func aks(n int) string {
    if(check_perfect_power(n)) {
        return "composite"
    }
    return "prime"
}

func check_perfect_power(n int) bool {
    for a := 2; a < n; a++ {
        for b := 2; b < n; b++ {
            if(n == (int(math.Pow(float64(a), float64(b))))) {
                return true
            }
        }
    } 
    return false
}