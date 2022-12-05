package main

import (
  "fmt"
  "math"
)
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
  if(check_perfect_power(n)) {
    return "composite"
  }
  //2.
  r := find_smallest_r(n)
  //3.
  for a := 2; a <= Math.min(r, n-1); a++ {
    if n % a == 0 {
      return "composite"
    }
  }
  //4.
  if n <= r {
    return "prime"
  }
  //5.

  //6.
  return "prime"
}

func check_perfect_power(n int) bool {
    max := math.Sqrt(float64(n)) + 1
    for i := 2; i < int(max); i++ {
      a := n;
      b := 0;
      for a % i == 0 {
        a /= i
        b += 1
        if a == 1 {
          return true
        }
      }
    }
    return false
}

func find_smallest_r(n int) int {
  bound := int(math.Ceil(math.Log2(float64(n)) * math.Log2(float64(n))))

  for r := 2; r < math.MaxInt64; r++ {
    if multiplicativeOrder(n, r) > bound {
      return r
    }
  }
  return -1
}



func multiplicativeOrder(a, n int) int {
  if(GCD(a, n) != 1) {
    return -1
  }
  result := 1
  k := 1
  for k < n {
    result = (result * a) % n
    if(result == 1) {
      return k
    }
    k++
  }
  return -1
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}