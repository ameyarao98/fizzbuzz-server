package fizzbuzz

func gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
