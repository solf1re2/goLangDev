package problem

import (
	"fmt"
	"math"
)

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

// broken/locks up laptop

var primeList []int
var n int

func Three() {
	n = 600851475143
	primeList = append(primeList, 1)
	for i := 2; i < n; i++ {
		//check if i is a factor of n
		if n%i != 0 {
			continue
		}
		// For each value of i - get the rounded up square root
		roundedUpSqrt := int(math.Ceil(math.Sqrt(float64(i))))
		// default bool to true, during checking it just needs to be disproved once to become false
		iisPrime := true
		// iterate over current slice of prime numbers
		for j := range primeList[0:] {
			fmt.Println(j)
			fmt.Println(roundedUpSqrt)
			if j != 0 {

				if (roundedUpSqrt % j) == 0 {
					iisPrime = false
				}
			}
		}
		if iisPrime {
			primeList = append(primeList, i)
		}
	}
	fmt.Println(primeList[len(primeList)])
}
