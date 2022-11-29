package bluebird

func FizzBuzz(n int) (string, int) {
	fizz := "fizz"
	buzz := "buzz"

	if n%3 == 0 && n%5 == 0 {
		return fizz + buzz, 0
	} else if n%3 == 0 {
		return fizz, 0
	} else if n%5 == 0 {
		return buzz, 0
	}

	return "", n
}
