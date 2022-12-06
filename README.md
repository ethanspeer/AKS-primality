# aks-primality
AKS Primality Test in Golang

# Algorithm

Input: Integer n
1. Check if n is a perfect power, i.e. if n = a^b for integers a > 1 and b > 1. If it is, output composite.
2. Find the smallest r such that ord_r(n) > (log_2(n))^2.
3. For all 2 ≤ a ≤ min (r, n−1), check that a does not divide n: If a|n, output composite.
4. If n ≤ r, output prime.
5. For a = 1 to math.floor(sqrt(eulertotient(r)) * log_2(n))
if (X+a)n≠ Xn+a (mod Xr − 1,n), output composite.
6. Output prime.

https://en.wikipedia.org/wiki/AKS_primality_test#The_algorithm
