package problem

import "fmt"

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

var n int

func Three() {
	n = 600851475143
	// n = 13195

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n = n / i
			}
		}
	}
	fmt.Println(n)
}
