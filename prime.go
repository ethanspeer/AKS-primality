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



// Given two integers a and b, GCD(a,b) returns g,u,v where 
// g is the gcd(a,b) and au+bv = g 
func GCD(a, b int) (int,int,int) {
	// if b = 0, then the gcd is a 
	if b == 0 {
		return a,1,0
	}

	// Keeps tracks of the sign of a and b and makes sure 
	// a and b are non-negative 
	var na , nb   bool 
	if a < 0 {
		na = true 
		a *= -1 
	}
	if b < 0 {
		nb = true
		b *= -1 
	}

	// Variables we will return 
	var g, u, v int 
	u = 1 
	g = a
	x := 0  // keeps track of the number a's used in the Euclidean algorithm 
	y := b  // keeps track of the denominator in the Euclidean algorithm
	for y != 0 {
		t := ModN(uint(y),g) // find t,q with g = qy + t 
		q := g / y 
		s := u - q*x  
		u = x 
		g = y 
		x = s 
		y = t 
	}
	v = (g-a*u)/b  
	
	if !na && !nb {
		return g, u, v 
	} else if !na && nb {
		return g, u, -v 
	} else if na && !nb {
		return g, -u, v 
	} else {
		return g , -u, -v 
	}
}
