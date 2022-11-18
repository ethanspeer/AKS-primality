package main

import "fmt"
import "math"

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
    if(check_perfect_power(n)) {
        return "composite"
    }
  //2.
    var r int = find_smallest_r(n) //gives r
  //3.
    if(a_divides_n(n, r)) {
        return "composite"
    }
  //4.
    if(n <= r) {
        return "prime"
    }
  //5.
    polynomials(r, n)
  //6.
    return "prime"
}

func check_perfect_power(n int) bool {
    for a := 2; a < n; a++ {
        for b := 2; b < n; b++ {
            if(n == exponentiate(a, b)) {
                return true
            }
        }
    } 
    return false
}

func find_smallest_r(n int) int {
    bin := math.Log2((float64(n)))
    for r := 0; r < n; r++ {
        if(GCD(r, n) != 1) {
            continue
        }
        for k := 0; k < r; k++ {
           if(exponentiate(n, k) % r == 1) {
            if(float64(k) > math.Pow(bin, 2)) {
                return r
            }
           }
        }
    }
  return 0
}

func a_divides_n(n, r int) bool {
    min := 0
    if(n > r-1) {
        min = r - 1
    } else {
        min = n
    }
    for a := 2; a <= min; a++ {
        if(n % a == 0) {
            return true
        }
    }
    return false
}

func polynomials(r int, n int) {
    phi := math.Sqrt(float64(euler(r)))
    bin := math.Log2((float64(n)))
    a_bound := math.Floor(phi*bin)
    for a := 1; a < a_bound; a++ {

    }
}

func exponentiate(base int, exponent int) int {
  output := 1  
  for exponent != 0 {  
  output *= base  
  exponent -= 1  
 }  
 return output
}

func GCD(r, n int) int {
	for n != 0 {
		t := n
		n = r % n
		r = t
	}
	return r
}

func euler(n int) int {
    var result int = 1
    var i int
    for i = 2; i < n; i++ {
            if GCD(i, n) == 1 {
                    result++
            }
    }
    return result
}