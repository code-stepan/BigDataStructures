package euler

func ModPow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}

func HasPrimitiveRoot(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 4 {
		return true
	}
	factors := Factorize(n)
	if _, hasTwo := factors[2]; hasTwo {
		if factors[2] > 1 {
			return false
		}
		if len(factors) != 2 {
			return false
		}
	}
	oddPrimes := 0
	for p := range factors {
		if p != 2 {
			oddPrimes++
		}
	}
	return oddPrimes == 1
}
