package bluebird

import "testing"

func TestFizzBuzz(t *testing.T) {
	fizz := "fizz"
	buzz := "buzz"
	fizzbuzz := "fizzbuzz"

	if fizzbuzz1, _ := FizzBuzz(15); fizzbuzz1 != fizzbuzz {
		t.Fatalf("FizzBuzz() = %v, want %v", fizzbuzz1, fizzbuzz)
	}

	if fizz1, _ := FizzBuzz(3); fizz1 != fizz {
		t.Fatalf("FizzBuzz() = %v, want %v", fizz1, fizz)
	}

	if buzz1, _ := FizzBuzz(5); buzz1 != buzz {
		t.Fatalf("FizzBuzz() = %v, want %v", buzz1, buzz)
	}
}
