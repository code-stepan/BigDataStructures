package euler

import (
	"math"
)

type Complex struct {
	Re float64
	Im float64
}

func NewComplex(re, im float64) Complex {
	return Complex{Re: re, Im: im}
}

func (c Complex) Add(other Complex) Complex {
	return Complex{Re: c.Re + other.Re, Im: c.Im + other.Im}
}

func (c Complex) Sub(other Complex) Complex {
	return Complex{Re: c.Re - other.Re, Im: c.Im - other.Im}
}

func (c Complex) Mul(other Complex) Complex {
	return Complex{
		Re: c.Re*other.Re - c.Im*other.Im,
		Im: c.Re*other.Im + c.Im*other.Re,
	}
}

func (c Complex) Scale(s float64) Complex {
	return Complex{Re: c.Re * s, Im: c.Im * s}
}

func (c Complex) Conjugate() Complex {
	return Complex{Re: c.Re, Im: -c.Im}
}

func (c Complex) Abs() float64 {
	return math.Sqrt(c.Re*c.Re + c.Im*c.Im)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func EulerPhiByDefinition(n int) int {
	if n <= 0 {
		return 0
	}
	count := 0
	for k := 1; k <= n; k++ {
		if gcd(k, n) == 1 {
			count++
		}
	}
	return count
}

func Factorize(n int) map[int]int {
	factors := make(map[int]int)
	temp := n
	for i := 2; i*i <= temp; i++ {
		for temp%i == 0 {
			factors[i]++
			temp /= i
		}
	}
	if temp > 1 {
		factors[temp]++
	}
	return factors
}

func EulerPhiByFactorization(n int) int {
	if n <= 0 {
		return 0
	}
	result := n
	for p := range Factorize(n) {
		result -= result / p
	}
	return result
}

func EulerPhiByDFT(n int) float64 {
	if n <= 0 {
		return 0
	}
	sum := 0.0
	for k := 1; k <= n; k++ {
		g := gcd(k, n)
		angle := 2.0 * math.Pi * float64(k) / float64(n)
		sum += float64(g) * math.Cos(angle)
	}
	return math.Round(sum)
}

func AllRootsOfUnity(n int) []Complex {
	roots := make([]Complex, n)
	for k := 0; k < n; k++ {
		angle := 2.0 * math.Pi * float64(k) / float64(n)
		roots[k] = NewComplex(math.Cos(angle), math.Sin(angle))
	}
	return roots
}

func PrimitiveRootsOfUnity(n int) []Complex {
	allRoots := AllRootsOfUnity(n)
	primitive := make([]Complex, 0)
	for k := 0; k < n; k++ {
		if gcd(k, n) == 1 {
			primitive = append(primitive, allRoots[k])
		}
	}
	return primitive
}

func complexPower(base Complex, exp int) Complex {
	if exp < 0 {
		posResult := complexPower(base, -exp)
		return posResult.Conjugate()
	}
	result := NewComplex(1, 0)
	for exp > 0 {
		if exp%2 == 1 {
			result = result.Mul(base)
		}
		base = base.Mul(base)
		exp /= 2
	}
	return result
}

func VandermondeMatrix(n int) [][]Complex {
	w := NewComplex(math.Cos(2*math.Pi/float64(n)), math.Sin(2*math.Pi/float64(n)))
	mat := make([][]Complex, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]Complex, n)
		for j := 0; j < n; j++ {
			mat[i][j] = complexPower(w, i*j)
		}
	}
	return mat
}

func InverseVandermondeMatrix(n int) [][]Complex {
	w := NewComplex(math.Cos(2*math.Pi/float64(n)), math.Sin(2*math.Pi/float64(n)))
	inv := make([][]Complex, n)
	for i := 0; i < n; i++ {
		inv[i] = make([]Complex, n)
		for j := 0; j < n; j++ {
			inv[i][j] = complexPower(w, -i*j).Scale(1.0 / float64(n))
		}
	}
	return inv
}

func multiplyMatrixVector(mat [][]Complex, vec []Complex) []Complex {
	n := len(mat)
	result := make([]Complex, n)
	for i := 0; i < n; i++ {
		sum := NewComplex(0, 0)
		for j := 0; j < n; j++ {
			sum = sum.Add(mat[i][j].Mul(vec[j]))
		}
		result[i] = sum
	}
	return result
}

func DFT(x []Complex) []Complex {
	n := len(x)
	mat := VandermondeMatrix(n)
	return multiplyMatrixVector(mat, x)
}

func InverseDFT(x []Complex) []Complex {
	n := len(x)
	mat := InverseVandermondeMatrix(n)
	return multiplyMatrixVector(mat, x)
}

// func Divisors(n int ) []int {
// 	divs := []int{}
// 	for i := 1; i*i <= n; i++ {
// 		if n%i == 0 {
// 			divs = append(divs, i)
// 			if i != n/i {
// 				divs = append(divs, n/i)
// 			}
// 		}
// 	}
// 	sort.Ints(divs)
// 	return divs
// }
